package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"go-xsd/config"
	"go-xsd/internal/converters"
	schemas "go-xsd/internal/structs"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
	"github.com/xuri/xgen"
)

type API struct {
	Config *config.Config
}

func NewRouter(cfg *config.Config) *mux.Router {
	api := &API{Config: cfg}

	r := mux.NewRouter()
	r.HandleFunc("/convert-xml-to-json", api.ConvertXMLToJSON).Methods("POST")
	r.HandleFunc("/convert-json-to-xml", api.ConvertJSONToXML).Methods("POST")
	r.HandleFunc("/validate-xml", api.ValidateXML).Methods("POST")
	r.HandleFunc("/generate-structs", api.GenerateStructs).Methods("POST")
	return r
}

func (api *API) ConvertXMLToJSON(w http.ResponseWriter, r *http.Request) {
	xmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	jsonData, err := converters.XMLToJSON(xmlData)
	if err != nil {
		http.Error(w, "Erro ao converter XML para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (api *API) ConvertJSONToXML(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	var doc schemas.ACCCDOCComplexType
	err = json.Unmarshal(jsonData, &doc)
	if err != nil {
		http.Error(w, "Erro ao fazer unmarshal do JSON", http.StatusInternalServerError)
		return
	}

	doc.Xmlns = "http://www.cip-bancos.org.br/ARQ/ACCC470.xsd"

	xmlData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		http.Error(w, "Erro ao converter JSON para XML", http.StatusInternalServerError)
		return
	}

	xmlHeader := `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	finalXML := xmlHeader + string(xmlData)

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(finalXML))
}

func (api *API) ValidateXML(w http.ResponseWriter, r *http.Request) {
	schemaName := r.URL.Query().Get("schema")
	if schemaName == "" {
		http.Error(w, "Parâmetro 'schema' é obrigatório", http.StatusBadRequest)
		return
	}

	xsdPath := fmt.Sprintf("./schemas/ACCC%s.xsd", schemaName)

	schemaPath, err := filepath.Abs(xsdPath)
	if err != nil {
		fmt.Println("Erro ao obter o caminho absoluto do XSD", err)
		return
	}

	absPath, err := filepath.Abs("./schemas/ACCCTIPOS.xsd")

	fmt.Println(absPath)

	if err != nil {
		fmt.Println("Erro ao obter o caminho absoluto do XSD dependente", err)
		return
	}

	schemaContent, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Arquivo XSD não encontrado: %s", schemaPath), http.StatusNotFound)
		return
	}

	xmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	xsdSchema, err := xsd.Parse(schemaContent)
	if err != nil {
		http.Error(w, "Erro ao compilar o XSD", http.StatusInternalServerError)
		return
	}
	defer xsdSchema.Free()

	doc, err := libxml2.Parse(xmlData)
	if err != nil {
		http.Error(w, "Erro ao fazer parse do XML", http.StatusBadRequest)
		return
	}
	defer doc.Free()

	if err := xsdSchema.Validate(doc); err != nil {
		http.Error(w, "Erro na validação do XML: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("XML validado com sucesso!"))
}

func (api *API) GenerateStructs(w http.ResponseWriter, r *http.Request) {
	// Pegar o parâmetro 'schema' da query string para definir o XSD
	schemaName := r.URL.Query().Get("schema")
	if schemaName == "" {
		http.Error(w, "Parâmetro 'schema' é obrigatório", http.StatusBadRequest)
		return
	}

	// Formatar o caminho do XSD com base no parâmetro schemaName
	xsdPath := fmt.Sprintf("./schemas/ACCC%s.xsd", schemaName)

	fmt.Println(xsdPath + " - " + schemaName)

	outputDir := fmt.Sprintf("./internal/structs/ACCC%s", schemaName)
	schemaName = fmt.Sprintf("ACCC%s", schemaName)

	// Definir as opções para xgen
	options := &xgen.Options{
		FilePath:            xsdPath,     // Utiliza o caminho absoluto do XSD
		OutputDir:           outputDir,   // Diretório de saída para os structs gerados
		InputDir:            "./schemas", // Diretório de entrada com os schemas XSD
		Lang:                "go",        // Linguagem de saída (Go)
		Package:             schemaName,  // Nome do pacote gerado
		Extract:             false,
		IncludeMap:          make(map[string]bool),
		LocalNameNSMap:      make(map[string]string),
		NSSchemaLocationMap: make(map[string]string),
		ParseFileList:       make(map[string]bool),
		ParseFileMap:        make(map[string][]interface{}),
		RemoteSchema:        make(map[string][]byte),
	}

	// Gerar os structs a partir do XSD
	err := options.Parse()
	if err != nil {
		http.Error(w, "Erro ao gerar structs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Structs gerados a partir de ACCC%s.xsd com sucesso!", schemaName)))
}
