package converters

import (
	"encoding/json"
	"encoding/xml"

	structs "go-xsd/internal/structs"
)

func XMLToJSON(xmlData []byte) ([]byte, error) {
	var doc structs.ACCCDOC
	if err := xml.Unmarshal(xmlData, &doc); err != nil {
		return nil, err
	}
	return json.MarshalIndent(doc, "", "  ")
}

func JSONToXML(jsonData []byte) ([]byte, error) {
	var doc structs.ACCCDOCComplexType
	if err := json.Unmarshal(jsonData, &doc); err != nil {
		return nil, err
	}
	return xml.MarshalIndent(doc, "", "  ")
}
