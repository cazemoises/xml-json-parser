package schemas

import (
	"encoding/xml"
	schemas "go-xsd/internal/structs/common"
)

type ACCCDOCComplexType struct {
	XMLName xml.Name                 `xml:"ACCCDOC"`
	Xmlns   string                   `xml:"xmlns,attr"`
	BCARQ   schemas.BCARQComplexType `xml:"BCARQ"`
	SISARQ  SISARQComplexType        `xml:"SISARQ"`
	ESTARQ  string                   `xml:"ESTARQ,omitempty"`
}

type SISARQComplexType struct {
	ACCC470    *ACCC470ComplexType    `xml:"ACCC470"`
	ACCC470RET *ACCC470RETComplexType `xml:"ACCC470RET"`
}

type ACCCDOC *ACCCDOCComplexType

type ACCC470ComplexType struct {
	IdentdPartAdmdo            string                                   `xml:"IdentdPartAdmdo"`
	GrupoACCC470NuGarantia     *GrupoACCC470NuGarantiaComplexType       `xml:"Grupo_ACCC470_NuGarantia_ComplexType"`
	GrupoACCC470ChaveOrdem     []*GrupoACCC470ChaveOrdemComplexType     `xml:"Grupo_ACCC470_ChaveOrdem"`
	GrupoACCC470ChaveDuplicata []*GrupoACCC470ChaveDuplicataComplexType `xml:"Grupo_ACCC470_ChaveDuplicata"`
}

type GrupoACCC470NuGarantiaComplexType struct {
	XMLName xml.Name `xml:"Grupo_ACCC470_NuGarantia_ComplexType"`
	NUGar   []string `xml:"NUGar"`
}

type GrupoACCC470ChaveOrdemComplexType struct {
	XMLName            xml.Name                       `xml:"Grupo_ACCC470_ChaveOrdemComplexType"`
	NumOrdem           string                         `xml:"NumOrdem"`
	CNPJEmit           string                         `xml:"CNPJEmit"`
	CNPJCPFSacd        string                         `xml:"CNPJ_CPFSacd"`
	DtEmiss            string                         `xml:"DtEmiss"`
	Valor              float64                        `xml:"Valor"`
	GrupoACCC470ChNota *GrupoACCC470ChNotaComplexType `xml:"Grupo_ACCC470_ChNota"`
}

type GrupoACCC470ChaveDuplicataComplexType struct {
	XMLName             xml.Name `xml:"Grupo_ACCC470_ChaveDuplicataComplexType"`
	CodInstntoFinanc    string   `xml:"CodInstntoFinanc"`
	CNPJEmit            string   `xml:"CNPJEmit"`
	CNPJCPFSacd         string   `xml:"CNPJ_CPFSacd"`
	DtVencInstntoFinanc string   `xml:"DtVencInstntoFinanc"`
	VlrUnit             float64  `xml:"VlrUnit"`
}

type GrupoACCC470ChNotaComplexType struct {
	XMLName xml.Name `xml:"Grupo_ACCC470_ChNotaComplexType"`
	ChNota  string   `xml:"ChNota"`
}

type ACCC470RETComplexType struct {
	IdentdPartAdmdo   schemas.CNPJBaseCodErr        `xml:"IdentdPartAdmdo"`
	GrupoGarExistte   *GrupoGarExistteComplexType   `xml:"Grupo_Gar_Existte"`
	GrupoGarInexistte *GrupoGarInexistteComplexType `xml:"Grupo_Gar_Inexistte"`
	GrupoGarErro      *GrupoGarExistteComplexType   `xml:"Grupo_Gar_Erro"`
}

type GrupoACCC470RETNuGarantiaComplexType struct {
	XMLName xml.Name `xml:"Grupo_ACCC470RET_NuGarantiaComplexType"`
	NUGar   []string `xml:"NUGar"`
}

type GrupoGarExistteComplexType struct {
	XMLName                       xml.Name                                    `xml:"Grupo_Gar_ExistteComplexType"`
	GrupoACCC470RETNuGarantia     *GrupoACCC470RETNuGarantiaComplexType       `xml:"Grupo_ACCC470RET_NuGarantia"`
	GrupoACCC470RETChaveOrdem     []*GrupoACCC470RETChaveOrdemComplexType     `xml:"Grupo_ACCC470RET_ChaveOrdem"`
	GrupoACCC470RETChaveDuplicata []*GrupoACCC470RETChaveDuplicataComplexType `xml:"Grupo_ACCC470RET_ChaveDuplicata"`
}

type GrupoACCC470RETChaveOrdemComplexType struct {
	XMLName     xml.Name `xml:"Grupo_ACCC470RET_ChaveOrdemComplexType"`
	NumOrdem    string   `xml:"NumOrdem"`
	CNPJEmit    string   `xml:"CNPJEmit"`
	CNPJCPFSacd string   `xml:"CNPJ_CPFSacd"`
	DtEmiss     string   `xml:"DtEmiss"`
	Valor       float64  `xml:"Valor"`
	SitRegistro string   `xml:"SitRegistro"`
}

type GrupoACCC470RETChaveDuplicataComplexType struct {
	XMLName             xml.Name                       `xml:"Grupo_ACCC470RET_ChaveDuplicataComplexType"`
	CodErroAttr         string                         `xml:"CodErro,attr,omitempty"`
	CodInstntoFinanc    schemas.CodInstntoFinancCodErr `xml:"CodInstntoFinanc"`
	CNPJEmit            schemas.CNPJCodErr             `xml:"CNPJEmit"`
	CNPJCPFSacd         schemas.CNPJCPFCodErr          `xml:"CNPJ_CPFSacd"`
	DtVencInstntoFinanc schemas.DataCodErr             `xml:"DtVencInstntoFinanc"`
	VlrUnit             schemas.ValorCodErr            `xml:"VlrUnit"`
	GrupoACCC470ChNota  *GrupoACCC470ChNotaComplexType `xml:"Grupo_ACCC470_ChNota"`
	SitRegistro         string                         `xml:"SitRegistro"`
	SituacaoCorda       string                         `xml:"SituacaoCorda"`
	LocalRegistro       string                         `xml:"LocalRegistro"`
}

type GrupoGarInexistteComplexType struct {
	XMLName                       xml.Name                                    `xml:"Grupo_Gar_InexistteComplexType"`
	GrupoACCC470RETNuGarantia     *GrupoACCC470RETNuGarantiaComplexType       `xml:"Grupo_ACCC470RET_NuGarantia"`
	GrupoACCC470RETChaveOrdem     []*GrupoACCC470RETChaveOrdemComplexType     `xml:"Grupo_ACCC470RET_ChaveOrdem"`
	GrupoACCC470RETChaveDuplicata []*GrupoACCC470RETChaveDuplicataComplexType `xml:"Grupo_ACCC470RET_ChaveDuplicata"`
}
