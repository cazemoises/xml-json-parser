package schemas

import (
	"encoding/xml"
	schemas "go-xsd/internal/structs/common"
)

type ACCC471ComplexType struct {
	GrupoACCC471GarSCR []*GrupoACCC471GarSCRComplexType `xml:"Grupo_ACCC471_GarSCR"`
}

type GrupoACCC471GarSCRComplexType struct {
	XMLName                     xml.Name                                `xml:"Grupo_ACCC471_GarSCRComplexType"`
	CodErroAttr                 string                                  `xml:"CodErro,attr,omitempty"`
	IdentdPartAdmdo             string                                  `xml:"IdentdPartAdmdo"`
	NumCtrlIFGar                string                                  `xml:"NumCtrlIFGar"`
	TpGar                       schemas.TpGarCodErr                     `xml:"TpGar"`
	TpGarSCR                    schemas.TpGarSCRCodErr                  `xml:"TpGarSCR"`
	SubTpGarSCR                 schemas.SubTpGarSCRCodErr               `xml:"SubTpGarSCR"`
	SitGar                      string                                  `xml:"SitGar"`
	IndrBemFincd                schemas.IndrCodErr                      `xml:"IndrBemFincd"`
	PercGar                     schemas.PercentualCodErr                `xml:"PercGar"`
	VlrOrGar                    schemas.ValorCodErr                     `xml:"VlrOrGar"`
	VlrGarDtReaval              schemas.ValorCodErr                     `xml:"VlrGarDtReaval"`
	DtReaval                    schemas.DataCodErr                      `xml:"DtReaval"`
	ClassRscCesta               string                                  `xml:"ClassRscCesta"`
	GrupoACCC471Veic            *GrupoACCC471VeicComplexType            `xml:"Grupo_ACCC471_Veic"`
	GrupoACCC471Imovel          *GrupoACCC471ImovelComplexType          `xml:"Grupo_ACCC471_Imovel"`
	GrupoACCC471GarFidejussoria *GrupoACCC471GarFidejussoriaComplexType `xml:"Grupo_ACCC471_GarFidejussoria"`
	GrupoACCC471Equipmnt        schemas.GrupoEquipmntComplexType        `xml:"Grupo_ACCC471_Equipmnt"`
	GrupoACCC471InstntoFinanc   *GrupoInstntoFinancComplexType          `xml:"Grupo_ACCC471_InstntoFinanc"`
	GrupoACCC471Mercdria        *GrupoMercdriaComplexType               `xml:"Grupo_ACCC471_Mercdria"`
}

type GrupoACCC471ImovelComplexType struct {
	XMLName               xml.Name                          `xml:"Grupo_ACCC471_ImovelComplexType"`
	TpImovl               string                            `xml:"TpImovl"`
	TpUsoImovl            string                            `xml:"TpUsoImovl"`
	NumInscrMuncplImovl   string                            `xml:"NumInscrMuncplImovl"`
	NumMatriclImovl       string                            `xml:"NumMatriclImovl"`
	IdCartrio             string                            `xml:"IdCartrio"`
	GrupoACCC471EndImovel schemas.GrupoEndImovelComplexType `xml:"Grupo_ACCC471_EndImovel"`
}

type GrupoRETImovelComplexType struct {
	CodErroAttr              string                               `xml:"CodErro,attr,omitempty"`
	TpImovl                  schemas.TpImovlCodErr                `xml:"TpImovl"`
	TpUsoImovl               schemas.TpUsoImovlCodErr             `xml:"TpUsoImovl"`
	NumInscrMuncplImovl      schemas.NumInscrMuncplImovlCodErr    `xml:"NumInscrMuncplImovl"`
	NumMatriclImovl          schemas.NumMatriclImovlCodErr        `xml:"NumMatriclImovl"`
	IdCartrio                schemas.IdCartrioInstntoFinancCodErr `xml:"IdCartrio"`
	GrupoACCC471RETEndImovel schemas.GrupoEndImovelComplexType    `xml:"Grupo_ACCC471RET_EndImovel"`
}

type GrupoACCC471GarFidejussoriaComplexType struct {
	XMLName                     xml.Name                                  `xml:"Grupo_ACCC471_GarFidejussoriaComplexType"`
	SeqGarFidjssria             int                                       `xml:"SeqGarFidjssria"`
	TpPessoaGarFidjssria        string                                    `xml:"TpPessoaGarFidjssria"`
	CNPJCPFPessoaGarFidjssria   string                                    `xml:"CNPJ_CPFPessoaGarFidjssria"`
	GrupoACCC471GarFidjssriaPF  *GrupoACCC471GarFidjssriaPFComplexType    `xml:"Grupo_ACCC471_GarFidjssriaPF"`
	GrupoACCC471GarFidjssriaPJ  *GrupoACCC471GarFidjssriaPJComplexType    `xml:"Grupo_ACCC471_GarFidjssriaPJ"`
	GrupoACCC471EndGarFidjssria []*GrupoACCC471EndGarFidjssriaComplexType `xml:"Grupo_ACCC471_EndGarFidjssria"`
}

type GrupoRETGarFidejussoriaComplexType struct {
	CodErroAttr                    string                                    `xml:"CodErro,attr,omitempty"`
	SeqGarFidjssria                schemas.SeqGarFidjssriaCodErr             `xml:"SeqGarFidjssria"`
	TpPessoaGarFidjssria           schemas.TpPessoaCodErr                    `xml:"TpPessoaGarFidjssria"`
	CNPJCPFPessoaGarFidjssria      schemas.CNPJCPFCodErr                     `xml:"CNPJ_CPFPessoaGarFidjssria"`
	GrupoACCC471RETGarFidjssriaPF  *GrupoACCC471GarFidjssriaPFComplexType    `xml:"Grupo_ACCC471RET_GarFidjssriaPF"`
	GrupoACCC471RETGarFidjssriaPJ  *GrupoACCC471RETGarFidjssriaPJComplexType `xml:"Grupo_ACCC471RET_GarFidjssriaPJ"`
	GrupoACCC471RETEndGarFidjssria []*GrupoACCC471EndGarFidjssriaComplexType `xml:"Grupo_ACCC471RET_EndGarFidjssria"`
}

type GrupoACCC471GarFidjssriaPJComplexType struct {
	XMLName                   xml.Name                            `xml:"Grupo_ACCC471_GarFidjssriaPJComplexType"`
	NomEmpGarFidjssria        string                              `xml:"NomEmpGarFidjssria"`
	DtAbert                   string                              `xml:"DtAbert"`
	TelComl                   string                              `xml:"TelComl"`
	NomContato                string                              `xml:"NomContato"`
	TpCtrl                    int                                 `xml:"TpCtrl"`
	Fatrmnt                   float64                             `xml:"Fatrmnt"`
	PortePessoaGarFidjssriaPJ int                                 `xml:"PortePessoaGarFidjssriaPJ"`
	GrupoACCC471ReprLegal     []schemas.GrupoReprLegalComplexType `xml:"Grupo_ACCC471_ReprLegal"`
}

type GrupoACCC471RETGarFidjssriaPJComplexType struct {
	XMLName                   xml.Name                            `xml:"Grupo_ACCC471RET_GarFidjssriaPJComplexType"`
	CodErroAttr               string                              `xml:"CodErro,attr,omitempty"`
	NomEmpGarFidjssria        schemas.NomeCodErr                  `xml:"NomEmpGarFidjssria"`
	DtAbert                   schemas.DataCodErr                  `xml:"DtAbert"`
	TelComl                   schemas.NumTelCodErr                `xml:"TelComl"`
	NomContato                schemas.NomeCodErr                  `xml:"NomContato"`
	TpCtrl                    schemas.TpCtrlCodErr                `xml:"TpCtrl"`
	Fatrmnt                   schemas.ValorCodErr                 `xml:"Fatrmnt"`
	PortePessoaGarFidjssriaPJ schemas.PorteCliCodErr              `xml:"PortePessoaGarFidjssriaPJ"`
	GrupoACCC471RETReprLegal  []schemas.GrupoReprLegalComplexType `xml:"Grupo_ACCC471RET_ReprLegal"`
}

type GrupoACCC471GarFidjssriaPFComplexType struct {
	XMLName                   xml.Name                   `xml:"Grupo_ACCC471_GarFidjssriaPFComplexType"`
	CodErroAttr               string                     `xml:"CodErro,attr,omitempty"`
	NomPessoaGarFidjssria     schemas.NomeCodErr         `xml:"NomPessoaGarFidjssria"`
	NomMae                    schemas.NomeCodErr         `xml:"NomMae"`
	EstadoCivil               schemas.EstadoCivilCodErr  `xml:"EstadoCivil"`
	CPFConjuge                schemas.CPFCodErr          `xml:"CPFConjuge"`
	NomConjuge                schemas.NomeCodErr         `xml:"NomConjuge"`
	PortePessoaGarFidjssriaPF schemas.PorteCliCodErr     `xml:"PortePessoaGarFidjssriaPF"`
	TpIdentc                  schemas.TpIdentcCodErr     `xml:"TpIdentc"`
	NumIdentc                 schemas.NumDocIdentcCodErr `xml:"NumIdentc"`
}

type GrupoACCC471EndGarFidjssriaComplexType struct {
	XMLName             xml.Name             `xml:"Grupo_ACCC471_EndGarFidjssriaComplexType"`
	CodErroAttr         string               `xml:"CodErro,attr,omitempty"`
	TpEndGarFidjssria   schemas.TpEndCodErr  `xml:"TpEndGarFidjssria"`
	EndGarFidjssria     schemas.EndCodErr    `xml:"EndGarFidjssria"`
	CEPEndGarFidjssria  schemas.CEPCodErr    `xml:"CEPEndGarFidjssria"`
	CidEndGarFidjssria  schemas.CidadeCodErr `xml:"CidEndGarFidjssria"`
	UFEndGarFidjssria   schemas.UFCodErr     `xml:"UFEndGarFidjssria"`
	PaisEndGarFidjssria schemas.PaisCodErr   `xml:"PaisEndGarFidjssria"`
}

type GrupoInstntoFinancComplexType struct {
	TpInstntoFinanc         string                             `xml:"TpInstntoFinanc"`
	CodInstntoFinanc        string                             `xml:"CodInstntoFinanc"`
	DtEmissInstntoFinanc    string                             `xml:"DtEmissInstntoFinanc"`
	DtinialInstntoFinanc    string                             `xml:"DtinialInstntoFinanc"`
	DtVencInstntoFinanc     string                             `xml:"DtVencInstntoFinanc"`
	QtdEmitd                int                                `xml:"QtdEmitd"`
	VlrUnit                 float64                            `xml:"VlrUnit"`
	NumMatriclInstntoFinanc string                             `xml:"NumMatriclInstntoFinanc"`
	IdCartrioInstntoFinanc  string                             `xml:"IdCartrioInstntoFinanc"`
	NumFatr                 string                             `xml:"NumFatr"`
	TpDuplicata             int                                `xml:"TpDuplicata"`
	Obs                     string                             `xml:"Obs"`
	ConfcActe               string                             `xml:"ConfcActe"`
	DtConfcActe             string                             `xml:"DtConfcActe"`
	VlrDesct                float64                            `xml:"VlrDesct"`
	DtIniDesct              string                             `xml:"DtIniDesct"`
	DtFimDesct              string                             `xml:"DtFimDesct"`
	GrupoACCC471Pessoa      []*GrupoPessoa                     `xml:"Grupo_ACCC471_Pessoa"`
	GrupoACCC471NotaFiscal  []schemas.NotaFiscalDupComplexType `xml:"Grupo_ACCC471_NotaFiscal"`
	GrupoACCC471VenctoOrdem []schemas.VenctoOrdemComplexType   `xml:"Grupo_ACCC471_VenctoOrdem"`
	GrupoACCC471Boleto      []schemas.BoletoDupComplexType     `xml:"Grupo_ACCC471_Boleto"`
}

type GrupoRETInstntoFinancComplexType struct {
	CodErroAttr                string                                `xml:"CodErro,attr,omitempty"`
	TpInstntoFinanc            schemas.TpInstntoFinancCodErr         `xml:"TpInstntoFinanc"`
	CodInstntoFinanc           schemas.CodInstntoFinancCodErr        `xml:"CodInstntoFinanc"`
	DtEmissInstntoFinanc       schemas.DataCodErr                    `xml:"DtEmissInstntoFinanc"`
	DtinialInstntoFinanc       schemas.DataCodErr                    `xml:"DtinialInstntoFinanc"`
	DtVencInstntoFinanc        schemas.DataCodErr                    `xml:"DtVencInstntoFinanc"`
	QtdEmitd                   schemas.QtdCodErr                     `xml:"QtdEmitd"`
	VlrUnit                    schemas.ValorCodErr                   `xml:"VlrUnit"`
	NumMatriclInstntoFinanc    schemas.NumMatriclInstntoFinancCodErr `xml:"NumMatriclInstntoFinanc"`
	IdCartrioInstntoFinanc     schemas.IdCartrioInstntoFinancCodErr  `xml:"IdCartrioInstntoFinanc"`
	NumFatr                    schemas.NumFatrCodErr                 `xml:"NumFatr"`
	TpDuplicata                schemas.TpDuplicataCodErr             `xml:"TpDuplicata"`
	Obs                        schemas.ObsCodErr                     `xml:"Obs"`
	ConfcActe                  schemas.ConfcActeCodErr               `xml:"ConfcActe"`
	DtConfcActe                schemas.DataCodErr                    `xml:"DtConfcActe"`
	VlrDesct                   schemas.ValorCodErr                   `xml:"VlrDesct"`
	DtIniDesct                 schemas.DataCodErr                    `xml:"DtIniDesct"`
	DtFimDesct                 schemas.DataCodErr                    `xml:"DtFimDesct"`
	GrupoACCC471RETPessoa      []*GrupoPessoaRET                     `xml:"Grupo_ACCC471RET_Pessoa"`
	GrupoACCC471RETNotaFiscal  []schemas.NotaFiscalDupComplexType    `xml:"Grupo_ACCC471RET_NotaFiscal"`
	GrupoACCC471RETVenctoOrdem []schemas.VenctoOrdemComplexType      `xml:"Grupo_ACCC471RET_VenctoOrdem"`
	GrupoACCC471RETBoleto      []schemas.BoletoDupComplexType        `xml:"Grupo_ACCC471RET_Boleto"`
}

type GrupoPessoa struct {
	TpPessoaInstntoFinanc      string                                 `xml:"TpPessoaInstntoFinanc"`
	CNPJCPFPessoaInstntoFinanc string                                 `xml:"CNPJ_CPFPessoaInstntoFinanc"`
	NomPessoaInstntoFinanc     string                                 `xml:"NomPessoaInstntoFinanc"`
	TpPessoaVincdInstntoFinanc schemas.TpPessoaFisicaJuridicaCodErr   `xml:"TpPessoaVincdInstntoFinanc"`
	InscrEstadl                string                                 `xml:"InscrEstadl"`
	PcaPgto                    string                                 `xml:"PcaPgto"`
	GrupoACCC471EndPessoa      []schemas.GrupoEndPessoaDupComplexType `xml:"Grupo_ACCC471_EndPessoa"`
}

type GrupoPessoaRET struct {
	CodErroAttr                string                                 `xml:"CodErro,attr,omitempty"`
	TpPessoaInstntoFinanc      schemas.TpPessoaInstntoFinancCodErr    `xml:"TpPessoaInstntoFinanc"`
	CNPJCPFPessoaInstntoFinanc schemas.CNPJCPFCodErr                  `xml:"CNPJ_CPFPessoaInstntoFinanc"`
	NomPessoaInstntoFinanc     schemas.NomeCodErr                     `xml:"NomPessoaInstntoFinanc"`
	TpPessoaVincdInstntoFinanc schemas.TpPessoaCodErr                 `xml:"TpPessoaVincdInstntoFinanc"`
	InscrEstadl                schemas.InscrEstadlCodErr              `xml:"InscrEstadl"`
	PcaPgto                    schemas.PcaPgtoCodErr                  `xml:"PcaPgto"`
	GrupoACCC471RETEndPessoa   []schemas.GrupoEndPessoaDupComplexType `xml:"Grupo_ACCC471RET_EndPessoa"`
}

type GrupoMercdriaComplexType struct {
	TpProdt         string                                `xml:"TpProdt"`
	CodProdt        string                                `xml:"CodProdt"`
	DescProdt       string                                `xml:"DescProdt"`
	PesoLiqdProdt   float64                               `xml:"PesoLiqdProdt"`
	AnoSaf          string                                `xml:"AnoSaf"`
	QtdProdt        int                                   `xml:"QtdProdt"`
	UniddProdt      int                                   `xml:"UniddProdt"`
	GrupoACCC471End []schemas.GrupoMercdriaEndComplexType `xml:"Grupo_ACCC471_End"`
}

type GrupoRETMercdriaComplexType struct {
	CodErroAttr        string                                `xml:"CodErro,attr,omitempty"`
	TpProdt            schemas.TpProdtCodErr                 `xml:"TpProdt"`
	CodProdt           schemas.CodProdtCodErr                `xml:"CodProdt"`
	DescProdt          schemas.DescProdtCodErr               `xml:"DescProdt"`
	PesoLiqdProdt      schemas.PesoCodErr                    `xml:"PesoLiqdProdt"`
	AnoSaf             schemas.AnoCodErr                     `xml:"AnoSaf"`
	QtdProdt           schemas.QtdCodErr                     `xml:"QtdProdt"`
	UniddProdt         schemas.QtdCodErr                     `xml:"UniddProdt"`
	GrupoACCC471RETEnd []schemas.GrupoMercdriaEndComplexType `xml:"Grupo_ACCC471RET_End"`
}

type GrupoACCC471VeicComplexType struct {
	XMLName                xml.Name                           `xml:"Grupo_ACCC471_VeicComplexType"`
	VlrEntdVeic            schemas.ValorCodErr                `xml:"VlrEntdVeic"`
	IdentdChassi           schemas.IdChassiCodErr             `xml:"IdentdChassi"`
	TpVeic                 int                                `xml:"TpVeic"`
	TpTabVeicl             string                             `xml:"TpTabVeicl"`
	CodTabVeicl            string                             `xml:"CodTabVeicl"`
	UFCodTabVeicl          string                             `xml:"UFCodTabVeicl"`
	IndrVeicUsado          string                             `xml:"IndrVeicUsado"`
	NumPlaca               string                             `xml:"NumPlaca"`
	UFNumPlaca             string                             `xml:"UFNumPlaca"`
	KM                     int                                `xml:"KM"`
	RENAVAM                string                             `xml:"RENAVAM"`
	NumNota                string                             `xml:"NumNota"`
	NumSerNota             string                             `xml:"NumSerNota"`
	VlrNota                float64                            `xml:"VlrNota"`
	DtEmiss                string                             `xml:"DtEmiss"`
	GrupoACCC471IdentcVeic schemas.GrupoIdentcVeicComplexType `xml:"Grupo_ACCC471_IdentcVeic"`
}

type GrupoACCC471AvalstComplexType struct {
	XMLName       xml.Name               `xml:"Grupo_ACCC471_AvalstComplexType"`
	CodErroAttr   string                 `xml:"CodErro,attr,omitempty"`
	TpCliAvalst   schemas.TpPessoaCodErr `xml:"TpCli_Avalst"`
	CNPJCPFAvalst schemas.CNPJCPFCodErr  `xml:"CNPJ_CPF_Avalst"`
}

type GrupoACCC471CliPFComplexType struct {
	XMLName              xml.Name                      `xml:"Grupo_ACCC471_CliPFComplexType"`
	NomCli               string                        `xml:"NomCli"`
	DtNasc               string                        `xml:"DtNasc"`
	Sexo                 string                        `xml:"Sexo"`
	Nacnl                string                        `xml:"Nacnl"`
	TpIdentc             string                        `xml:"TpIdentc"`
	NumIdentc            string                        `xml:"NumIdentc"`
	DtIdentc             string                        `xml:"DtIdentc"`
	OrgaoExpddr          string                        `xml:"OrgaoExpddr"`
	UFOrgaoExpddr        string                        `xml:"UFOrgaoExpddr"`
	NomMae               string                        `xml:"NomMae"`
	EstadoCivil          string                        `xml:"EstadoCivil"`
	NomConjuge           string                        `xml:"NomConjuge"`
	CPFConjuge           string                        `xml:"CPFConjuge"`
	GrupoACCC471EndCliPF []schemas.GrupoEndComplexType `xml:"Grupo_ACCC471_EndCliPF"`
	TelRes               string                        `xml:"TelRes"`
	TelComl              string                        `xml:"TelComl"`
	TelCel               string                        `xml:"TelCel"`
	EndEletrnc           string                        `xml:"EndEletrnc"`
	PorteCliPF           int                           `xml:"PorteCli_PF"`
}

type ACCC471RETComplexType struct {
	GrupoGarActo   *GrupoGarActoComplexType   `xml:"Grupo_Gar_Acto"`
	GrupoGarRecsdo *GrupoGarRecsdoComplexType `xml:"Grupo_Gar_Recsdo"`
}

type GrupoGarActoComplexType struct {
	XMLName                xml.Name                             `xml:"Grupo_Gar_ActoComplexType"`
	GrupoACCC471RETGarActo []*GrupoACCC471RETGarActoComplexType `xml:"Grupo_ACCC471RET_GarActo"`
}

type GrupoACCC471RETGarActoComplexType struct {
	XMLName         xml.Name `xml:"Grupo_ACCC471RET_GarActoComplexType"`
	IdentdPartAdmdo string   `xml:"IdentdPartAdmdo"`
	NumCtrlIFGar    string   `xml:"NumCtrlIFGar"`
	SitRegistro     string   `xml:"SitRegistro"`
}

type GrupoGarRecsdoComplexType struct {
	XMLName                  xml.Name                            `xml:"Grupo_Gar_RecsdoComplexType"`
	GrupoACCC471RETGarRecsdo []*GrupoACCC471RETGarSCRComplexType `xml:"Grupo_ACCC471RET_GarRecsdo"`
}

type GrupoACCC471RETGarSCRComplexType struct {
	XMLName                        xml.Name                            `xml:"Grupo_ACCC471RET_GarSCRComplexType"`
	CodErroAttr                    string                              `xml:"CodErro,attr,omitempty"`
	IdentdPartAdmdo                schemas.CNPJBaseCodErr              `xml:"IdentdPartAdmdo"`
	NumCtrlIFGar                   schemas.ControleIFCCCCodErr         `xml:"NumCtrlIFGar"`
	TpGar                          schemas.TpGarCodErr                 `xml:"TpGar"`
	TpGarSCR                       schemas.TpGarSCRCodErr              `xml:"TpGarSCR"`
	SubTpGarSCR                    schemas.SubTpGarSCRCodErr           `xml:"SubTpGarSCR"`
	SitGar                         schemas.SitGarCodErr                `xml:"SitGar"`
	IndrBemFincd                   schemas.IndrCodErr                  `xml:"IndrBemFincd"`
	PercGar                        schemas.PercentualCodErr            `xml:"PercGar"`
	VlrOrGar                       schemas.ValorCodErr                 `xml:"VlrOrGar"`
	VlrGarDtReaval                 schemas.ValorCodErr                 `xml:"VlrGarDtReaval"`
	DtReaval                       schemas.DataCodErr                  `xml:"DtReaval"`
	ClassRscCesta                  schemas.ClassRscCestaCodErr         `xml:"ClassRscCesta"`
	GrupoACCC471RETVeic            *GrupoRETVeicComplexType            `xml:"Grupo_ACCC471RET_Veic"`
	GrupoACCC471RETImovel          *GrupoRETImovelComplexType          `xml:"Grupo_ACCC471RET_Imovel"`
	GrupoACCC471RETGarFidejussoria *GrupoRETGarFidejussoriaComplexType `xml:"Grupo_ACCC471RET_GarFidejussoria"`
	GrupoACCC471RETEquipmnt        schemas.GrupoEquipmntComplexType    `xml:"Grupo_ACCC471RET_Equipmnt"`
	GrupoACCC471RETInstntoFinanc   *GrupoRETInstntoFinancComplexType   `xml:"Grupo_ACCC471RET_InstntoFinanc"`
	GrupoACCC471RETMercdria        *GrupoRETMercdriaComplexType        `xml:"Grupo_ACCC471RET_Mercdria"`
}

type GrupoRETVeicComplexType struct {
	CodErroAttr               string                             `xml:"CodErro,attr,omitempty"`
	VlrEntdVeic               schemas.ValorCodErr                `xml:"VlrEntdVeic"`
	IdentdChassi              schemas.IdChassiCodErr             `xml:"IdentdChassi"`
	TpVeic                    schemas.TpVeiclCodErr              `xml:"TpVeic"`
	TpTabVeicl                schemas.TpTabVeiclCodErr           `xml:"TpTabVeicl"`
	CodTabVeicl               schemas.CodTabVeiclCodErr          `xml:"CodTabVeicl"`
	UFCodTabVeicl             schemas.UFCodErr                   `xml:"UFCodTabVeicl"`
	IndrVeicUsado             schemas.IndrCodErr                 `xml:"IndrVeicUsado"`
	NumPlaca                  schemas.NumPlacaVeicCodErr         `xml:"NumPlaca"`
	UFNumPlaca                schemas.UFCodErr                   `xml:"UFNumPlaca"`
	KM                        schemas.QtdCodErr                  `xml:"KM"`
	RENAVAM                   schemas.RenavamCodErr              `xml:"RENAVAM"`
	NumNota                   schemas.NumNotaFiscalCodErr        `xml:"NumNota"`
	NumSerNota                string                             `xml:"NumSerNota"`
	VlrNota                   schemas.ValorCodErr                `xml:"VlrNota"`
	DtEmiss                   schemas.DataCodErr                 `xml:"DtEmiss"`
	GrupoACCC471RETIdentcVeic schemas.GrupoIdentcVeicComplexType `xml:"Grupo_ACCC471RET_IdentcVeic"`
}
