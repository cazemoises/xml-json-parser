package schemas

import (
	"encoding/xml"
	"time"
)

type ACCCDOCPROComplexType struct {
	BCARQ *BCARQComplexType `xml:"BCARQ"`
}

type ACCCDOCERRComplexType struct {
	BCARQ *BCARQERRComplexType `xml:"BCARQ"`
}

type BCARQComplexType struct {
	NomArq           string               `xml:"NomArq"`
	NumCtrlEmis      string               `xml:"NumCtrlEmis,omitempty"`
	NumCtrlDestOr    string               `xml:"NumCtrlDestOr,omitempty"`
	ISPBEmissor      string               `xml:"ISPBEmissor"`
	ISPBDestinatario string               `xml:"ISPBDestinatario"`
	DtHrArq          string               `xml:"DtHrArq"`
	SitReq           int                  `xml:"SitReq"`
	GrupoSeq         *GrupoSeqComplexType `xml:"Grupo_Seq"`
	DtRef            string               `xml:"DtRef"`
}

type GrupoSeqComplexType struct {
	XMLName  xml.Name `xml:"Grupo_Seq"`
	NumSeq   int      `xml:"NumSeq"`
	IndrCont string   `xml:"IndrCont"`
}

type BCARQERRComplexType struct {
	NomArq *NomArqCodErrComplexType `xml:"NomArq"`
}

type ESTARQSimpleType string

type Ano string

type AnoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type AnoMes string

type AnoMesCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CNPJ string

type CNPJCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CNPJBase string

type CNPJBaseCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpPessoaINTCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*TpPessoaINT
}

type TpPessoaFisicaJuridica string

type TpPessoaFisicaJuridicaCodErr struct {
	XMLName     xml.Name `xml:"TpPessoaFisica_JuridicaCodErr"`
	CodErroAttr string   `xml:"CodErro,attr,omitempty"`
	Value       string   `xml:",chardata"`
}

type TpPessoaSacdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*TpPessoaSacd
}

type NumBcoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*NumBco
}

type DigtVerificdrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*DigtVerificdr
}

type LinhaDigitvlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*LinhaDigitvl
}

type CNPJCPF string

type CNPJCPFCodErr struct {
	XMLName     xml.Name `xml:"CNPJ_CPFCodErr"`
	CodErroAttr string   `xml:"CodErro,attr,omitempty"`
	Value       string   `xml:",chardata"`
}

type CPF string

type CPFCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodCatg string

type CodCatgCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodContrtoOr string

type CodContrtoOrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodContrtoSCR string

type CodContrtoSCRCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type RetenSubtlRsc int

type RetenSubtlRscErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type CodErro string

type CodInstntoFinanc string

type CodInstntoFinancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodMarca string

type CodMarcaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodModl string

type CodModlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodMoedaCNAB string

type CodMoedaCNABCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodProdt string

type CodProdtCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodTabVeicl string

type CodTabVeiclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ControleIF string

type ControleIFCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ControleIFCCC string

type NumCtrlIFGar string

type ControleIFCCCCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Data string

type DataCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type DataHoraCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type DataHora string

type Hora time.Time

type DescProdt string

type DescProdtCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Justificativa string

type JustificativaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type DescCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*Desc
}

type NumConsProtErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*NumConsProt
}

type DtHrConsProtErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*DtHrConsProt
}

type CodConsProtErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*CodConsProt
}

type NumOrdemCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*NumOrdem
}

type EndEletrnc string

type EndEletrncCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type GrupoDadObjtoContrtoConsigdComplexType struct {
	CodErroAttr             string                    `xml:"CodErro,attr,omitempty"`
	CNPJBaseEnteCons        *CNPJBaseCodErr           `xml:"CNPJBase_EnteCons"`
	TpEnteCons              *TpEnteConsSCRCodErr      `xml:"TpEnteCons"`
	CNPJBaseOrgaoPagdr      *CNPJBaseCodErr           `xml:"CNPJBaseOrgaoPagdr"`
	CNPJOrgaoPagdrServdr    *CNPJCodErr               `xml:"CNPJOrgaoPagdrServdr"`
	IdentcOrgaoPagdr        *IdentcOrgaoPagdrCodErr   `xml:"IdentcOrgaoPagdr"`
	NUAvebcSCC              *NUCodErr                 `xml:"NUAvebcSCC"`
	NumContrtoEnteCons      *NumContrtoEnteConsCodErr `xml:"NumContrtoEnteCons"`
	NumBenfcEnteCons        *NumBenfcEnteConsCodErr   `xml:"NumBenfcEnteCons"`
	NumMatriclEnteCons      *NumMatriclEnteConsCodErr `xml:"NumMatriclEnteCons"`
	AnoMesComptcPrimroDesct *AnoMesCodErr             `xml:"AnoMesComptcPrimroDesct"`
	VlrMargConsigd          *ValorCodErr              `xml:"VlrMargConsigd"`
	TpVincEmprtcio          *TpVincEmprtcioCodErr     `xml:"TpVincEmprtcio"`
	DtAdmiss                *DataCodErr               `xml:"DtAdmiss"`
}

type GrupoEndImovelComplexType struct {
	CodErroAttr string        `xml:"CodErro,attr,omitempty"`
	TpEndImovl  *TpEndCodErr  `xml:"TpEndImovl"`
	EndImovl    *EndCodErr    `xml:"EndImovl"`
	CEPImovl    *CEPCodErr    `xml:"CEPImovl"`
	CidImovl    *CidadeCodErr `xml:"CidImovl"`
	UFImovl     *UFCodErr     `xml:"UFImovl"`
	PaisImovl   *PaisCodErr   `xml:"PaisImovl"`
}

type GrupoEndCliComplexType struct {
	CodErroAttr string        `xml:"CodErro,attr,omitempty"`
	TpEnd       *TpEndCodErr  `xml:"TpEnd"`
	End         *EndCodErr    `xml:"End"`
	CEP         *CEPCodErr    `xml:"CEP"`
	Cid         *CidadeCodErr `xml:"Cid"`
	UF          *UFCodErr     `xml:"UF"`
	Pais        *PaisCodErr   `xml:"Pais"`
}

type GrupoEndGarFidjssriaComplexType struct {
	CodErroAttr         string        `xml:"CodErro,attr,omitempty"`
	TpEndGarFidjssria   *TpEndCodErr  `xml:"TpEndGarFidjssria"`
	EndGarFidjssria     *EndCodErr    `xml:"EndGarFidjssria"`
	CEPEndGarFidjssria  *CEPCodErr    `xml:"CEPEndGarFidjssria"`
	CidEndGarFidjssria  *CidadeCodErr `xml:"CidEndGarFidjssria"`
	UFEndGarFidjssria   *UFCodErr     `xml:"UFEndGarFidjssria"`
	PaisEndGarFidjssria *PaisCodErr   `xml:"PaisEndGarFidjssria"`
}

type GrupoReprLegalComplexType struct {
	CodErroAttr      string         `xml:"CodErro,attr,omitempty"`
	CNPJCPFReprLegal *CNPJCPFCodErr `xml:"CNPJ_CPF_ReprLegal"`
}

type GrupoEquipmntComplexType struct {
	CodErroAttr     string                  `xml:"CodErro,attr,omitempty"`
	TpEquipmnt      *TpEquipmntCodErr       `xml:"TpEquipmnt"`
	VlrEntdEquipmnt *ValorCodErr            `xml:"VlrEntdEquipmnt"`
	CNPJFabrcte     *CNPJCodErr             `xml:"CNPJFabrcte"`
	CNPJForndr      *CNPJCodErr             `xml:"CNPJForndr"`
	NumNota         *NumNotaFiscalCodErr    `xml:"NumNota"`
	NumSerNota      *NumSerNotaFiscalCodErr `xml:"NumSerNota"`
	VlrNota         *ValorCodErr            `xml:"VlrNota"`
	DtEmiss         *DataCodErr             `xml:"DtEmiss"`
	PaisFabrcc      *PaisCodErr             `xml:"PaisFabrcc"`
	NumSerEquipmnt  *NumSerEquipmntCodErr   `xml:"NumSerEquipmnt"`
}

type GrupoEndPessoaComplexType struct {
	CodErroAttr string        `xml:"CodErro,attr,omitempty"`
	TpEnd       *TpEndCodErr  `xml:"TpEnd"`
	End         *EndCodErr    `xml:"End"`
	CEP         *CEPCodErr    `xml:"CEP"`
	Cid         *CidadeCodErr `xml:"Cid"`
	UF          *UFCodErr     `xml:"UF"`
	Pais        *PaisCodErr   `xml:"Pais"`
}

type GrupoEndPessoaDupComplexType struct {
	CodErroAttr string            `xml:"CodErro,attr,omitempty"`
	TpEnd       *TpEndCodErr      `xml:"TpEnd"`
	End         *EndCodErr        `xml:"End"`
	CEP         *CEPCodErr        `xml:"CEP"`
	Cid         *CidadeCodErr     `xml:"Cid"`
	CodMunIBGE  *CodMunIBGECodErr `xml:"CodMunIBGE"`
	UF          *UFCodErr         `xml:"UF"`
	Pais        *PaisCodErr       `xml:"Pais"`
}

type GrupoGarFidjssriaPFComplexType struct {
	CodErroAttr               string              `xml:"CodErro,attr,omitempty"`
	NomPessoaGarFidjssria     string              `xml:"NomPessoaGarFidjssria"`
	NomMae                    *NomeCodErr         `xml:"NomMae"`
	EstadoCivil               *EstadoCivilCodErr  `xml:"EstadoCivil"`
	CPFConjuge                *CPFCodErr          `xml:"CPFConjuge"`
	NomConjuge                *NomeCodErr         `xml:"NomConjuge"`
	PortePessoaGarFidjssriaPF *PorteCliCodErr     `xml:"PortePessoaGarFidjssriaPF"`
	TpIdentc                  *TpIdentcCodErr     `xml:"TpIdentc"`
	NumIdentc                 *NumDocIdentcCodErr `xml:"NumIdentc"`
}

type GrupoMercdriaEndComplexType struct {
	CodErroAttr string           `xml:"CodErro,attr,omitempty"`
	IdentcEnd   *IdentcEndCodErr `xml:"IdentcEnd"`
	TpEnd       *TpEndCodErr     `xml:"TpEnd"`
	End         *EndCodErr       `xml:"End"`
	CEP         *CEPCodErr       `xml:"CEP"`
	Cid         *CidadeCodErr    `xml:"Cid"`
	UF          *UFCodErr        `xml:"UF"`
	Pais        *PaisCodErr      `xml:"Pais"`
}

type GrupoPagtoParclComplexType struct {
	CodErroAttr            string               `xml:"CodErro,attr,omitempty"`
	NumCtrlIFEvt           *ControleIFCCCCodErr `xml:"NumCtrlIFEvt"`
	NUEvt                  *NUCodErr            `xml:"NUEvt"`
	DtPagtoParcl           *DataCodErr          `xml:"DtPagtoParcl"`
	DtContb                *DataCodErr          `xml:"DtContb"`
	VlrPago                *ValorCodErr         `xml:"VlrPago"`
	VlrPrincipalPagtoParcl *ValorCodErr         `xml:"VlrPrincipalPagtoParcl"`
	VlrJurosPagtoParcl     *ValorCodErr         `xml:"VlrJurosPagtoParcl"`
	VlrJurosMoraPagtoParcl *ValorCodErr         `xml:"VlrJurosMoraPagtoParcl"`
	VlrMultaPagtoParcl     *ValorCodErr         `xml:"VlrMultaPagtoParcl"`
	VlrIOFAtrPagtoParcl    *ValorCodErr         `xml:"VlrIOFAtrPagtoParcl"`
	VlrDesctPagtoParcl     *ValorCodErr         `xml:"VlrDesctPagtoParcl"`
	VlrAbattPagtoParcl     *ValorCodErr         `xml:"VlrAbattPagtoParcl"`
	VlrDespPagtoParcl      *ValorCodErr         `xml:"VlrDespPagtoParcl"`
	IndrLiquidExtraJud     *IndrCodErr          `xml:"IndrLiquidExtraJud"`
	IndrLei                *IndrCodErr          `xml:"IndrLei"`
	VlrLei                 *ValorCodErr         `xml:"VlrLei"`
	IndrSegrPago           *IndrCodErr          `xml:"IndrSegrPago"`
}

type GrupoPagtoParclActoComplexType struct {
	XMLName      xml.Name             `xml:"GrupoPagtoParclActo_ComplexType"`
	NumCtrlIFEvt *ControleIFCCCCodErr `xml:"NumCtrlIFEvt"`
	NUEvt        *NUCodErr            `xml:"NUEvt"`
}

type GrupoParclPagasComplexType struct {
	XMLName              xml.Name             `xml:"Grupo_ParclPagasComplexType"`
	CodErroAttr          string               `xml:"CodErro,attr,omitempty"`
	NumCtrlIFParclIni    *ControleIFCCCCodErr `xml:"NumCtrlIFParclIni"`
	NumCtrlIFParclFim    *ControleIFCCCCodErr `xml:"NumCtrlIFParclFim"`
	TpParclParclPaga     *TpParclCodErr       `xml:"TpParclParclPaga"`
	VlrParclPaga         *ValorCodErr         `xml:"VlrParclPaga"`
	VlrPagoParclPaga     *ValorCodErr         `xml:"VlrPagoParclPaga"`
	IndrParclPagaDt      *IndrCodErr          `xml:"IndrParclPagaDt"`
	IndrParclPagaCedd    *IndrCodErr          `xml:"IndrParclPagaCedd"`
	CNPJBaseCesParclPaga *CNPJBaseCodErr      `xml:"CNPJBase_CesParclPaga"`
	SitRepParclPaga      *SitRepCodErr        `xml:"SitRepParclPaga"`
}

type SitConclcao string

type SitConclcaoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitNotaFiscal string

type SitNotaFiscalCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type GrupoAltParclActoComplexType struct {
	XMLName              xml.Name     `xml:"Grupo_AltParclActoComplexType"`
	CodErroAttr          string       `xml:"CodErro,attr,omitempty"`
	NovVlrParcl          *ValorCodErr `xml:"NovVlrParcl"`
	NovVlrPrincipalParcl *ValorCodErr `xml:"NovVlrPrincipalParcl"`
	NovVlrJurosParcl     *ValorCodErr `xml:"NovVlrJurosParcl"`
	NovVlrSldDevdrParcl  *ValorCodErr `xml:"NovVlrSldDevdrParcl"`
	NovVlrSldPrejzParcl  *ValorCodErr `xml:"NovVlrSldPrejzParcl"`
}

type GrupoParcelaComplexType struct {
	CodErroAttr string      `xml:"CodErro,attr,omitempty"`
	NUParcl     []*NUCodErr `xml:"NUParcl"`
}

type GrupoIdentcVeicComplexType struct {
	CodErroAttr    string           `xml:"CodErro,attr,omitempty"`
	VlrMercVeic    *ValorCodErr     `xml:"VlrMercVeic"`
	CodCatg        *CodCatgCodErr   `xml:"CodCatg"`
	CodMarca       *CodMarcaCodErr  `xml:"CodMarca"`
	CodModl        *CodModlCodErr   `xml:"CodModl"`
	AnoModlVeicl   *AnoCodErr       `xml:"AnoModlVeicl"`
	AnoFabrccVeicl *AnoCodErr       `xml:"AnoFabrccVeicl"`
	TpCombtvl      *TpCombtvlCodErr `xml:"TpCombtvl"`
	PaisFabrcc     *PaisCodErr      `xml:"PaisFabrcc"`
}

type GrupoImpEncargoComplexType struct {
	CodErroAttr        string                 `xml:"CodErro,attr,omitempty"`
	TpEncargo          *TpEncargoFinancCodErr `xml:"TpEncargo"`
	VlrEncargo         *ValorCodErr           `xml:"VlrEncargo"`
	IndrFincmntEncargo *IndrCodErr            `xml:"IndrFincmntEncargo"`
}

type GrupoIdentdContrtoComplexType struct {
	XMLName             xml.Name            `xml:"Grupo_Identd_ContrtoComplexType"`
	CodErroAttr         string              `xml:"CodErro,attr,omitempty"`
	CodContrtoOr        *CodContrtoOrCodErr `xml:"CodContrtoOr"`
	CNPJBaseIFOrContrto *CNPJBaseCodErr     `xml:"CNPJBase_IFOrContrto"`
	TpCli               *TpPessoaCodErr     `xml:"TpCli"`
	CNPJCPFCli          *CNPJCPFCodErr      `xml:"CNPJ_CPFCli"`
	TpContrto           *TpContrtoCodErr    `xml:"TpContrto"`
}

type GrupoJurosMoraContrtoComplexType struct {
	CodErroAttr        string                    `xml:"CodErro,attr,omitempty"`
	QtdDiasJurosMora   *QtdCodErr                `xml:"QtdDiasJurosMora"`
	TpCobrJurosMora    *TpCobrCodErr             `xml:"TpCobrJurosMora"`
	VlrJurosMora       *ValorCodErr              `xml:"VlrJurosMora"`
	TpCalcJurosMora    *TpCalcCodErr             `xml:"TpCalcJurosMora"`
	FormaCalcJurosMora *FormaCalcJurosMoraCodErr `xml:"FormaCalcJurosMora"`
	TaxJurosMora       *PercentualCodErr         `xml:"TaxJurosMora"`
	PeriodTaxJurosMora *PeriodicidadeCodErr      `xml:"PeriodTaxJurosMora"`
	IndxJurosMora      *IndxCodErr               `xml:"IndxJurosMora"`
	PercIndxJurosMora  *PercentualCodErr         `xml:"PercIndxJurosMora"`
}

type GrupoMultaContrtoComplexType struct {
	CodErroAttr   string               `xml:"CodErro,attr,omitempty"`
	QtdDiasMulta  *QtdCodErr           `xml:"QtdDiasMulta"`
	TpCobrMulta   *TpCobrCodErr        `xml:"TpCobrMulta"`
	VlrMulta      *ValorCodErr         `xml:"VlrMulta"`
	BaseCalcMulta *BaseCalcMultaCodErr `xml:"BaseCalcMulta"`
	PercMulta     *PercentualCodErr    `xml:"PercMulta"`
	IndxMulta     *IndxCodErr          `xml:"IndxMulta"`
}

type GrupoIdentdContrtoOrigdrComplexType struct {
	XMLName                   xml.Name            `xml:"Grupo_Identd_ContrtoOrigdrComplexType"`
	CodErroAttr               string              `xml:"CodErro,attr,omitempty"`
	CodContrtoOrigdr          *CodContrtoOrCodErr `xml:"CodContrtoOrigdr"`
	CNPJBaseIFOrContrtoOrigdr *CNPJBaseCodErr     `xml:"CNPJBase_IFOrContrtoOrigdr"`
	TpCliContrtoOrigdr        *TpPessoaCodErr     `xml:"TpCliContrtoOrigdr"`
	CNPJCPFCliContrtoOrigdr   *CNPJCPFCodErr      `xml:"CNPJ_CPFCliContrtoOrigdr"`
	ModOpContrtoOr            *ModOpCodErr        `xml:"ModOpContrtoOr"`
}

type GrupoEndComplexType struct {
	XMLName     xml.Name      `xml:"Grupo_EndComplexType"`
	CodErroAttr string        `xml:"CodErro,attr,omitempty"`
	TpEnd       *TpEndCodErr  `xml:"TpEnd"`
	End         *EndCodErr    `xml:"End"`
	CEP         *CEPCodErr    `xml:"CEP"`
	Cid         *CidadeCodErr `xml:"Cid"`
	UF          *UFCodErr     `xml:"UF"`
	Pais        *PaisCodErr   `xml:"Pais"`
}

type GrupoIdentdResComplexType struct {
	XMLName      xml.Name            `xml:"Grupo_Identd_ResComplexType"`
	CodErroAttr  string              `xml:"CodErro,attr,omitempty"`
	DtOpRes      *DataCodErr         `xml:"DtOpRes"`
	CNPJBaseCed  *CNPJBaseCodErr     `xml:"CNPJBase_Ced"`
	CNPJBaseCes  *CNPJBaseCodErr     `xml:"CNPJBase_Ces"`
	NumIdentdRes *NumIdentdResCodErr `xml:"NumIdentdRes"`
}

type IdChassi string

type IdChassiCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Indr string

type IndrSemDominio string

type IdCartrioInstntoFinanc string

type NumFatr string

type NumFatrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpDuplicata int

type TpDuplicataCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type Obs string

type ObsCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ConfcActe string

type ConfcActeCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type InscrEstadl string

type PcaPgto string

type PcaPgtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type InscrEstadlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ChNota string

type ChNotaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumNotaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*NumNota
}

type SerieNFeErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*SerieNFe
}

type IdCartrioInstntoFinancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type IndrCess string

type IndrLiquidFinancCamr string

type IndrLiquidFinancCamrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type IndrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type IndrSemDominioCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type IdentcEnd string

type IdentcEndCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type IdentdLiqdFinanc string

type IdentcOrgaoPagdr int

type IdentcOrgaoPagdrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type LitrlTpVeicl string

type LitrlTpVeiclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Nome string

type NomeCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NomArqCodErrComplexType struct {
	XMLName     xml.Name `xml:"NomArqCodErr_ComplexType"`
	CodErroAttr string   `xml:"CodErro,attr"`
	Value       string   `xml:",chardata"`
}

type NU string

type NossoNumeroBOL string

type NUCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NossoNumeroBOLErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type MotvCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	*Motv
}

type NumBenfcEnteCons int

type NumBenfcEnteConsCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type NumContrtoEnteCons string

type NumContrtoEnteConsCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumIdentdRes string

type NumIdentdResCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumInscrMuncplImovl string

type NumInscrMuncplImovlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumMatriclImovl string

type NumMatriclImovlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumMatriclInstntoFinanc string

type NumMatriclInstntoFinancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumMatriclEnteCons string

type GrupoAtlzParcCedCodContrtoSCRComplexType struct {
	CodErroAttr      string               `xml:"CodErro,attr,omitempty"`
	NUCess           *NUCodErr            `xml:"NUCess"`
	CodContrtoSCRCes *CodContrtoSCRCodErr `xml:"CodContrtoSCRCes"`
}

type NumMatriclEnteConsCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumNotaFiscal string

type NumNotaFiscalCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumSerEquipmnt string

type NumSerEquipmntCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumSerNotaFiscal string

type NumSerNotaFiscalCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumSeq int

type NumSeqCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type NumSeqInf int

type NumSeqInfCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type NumPlacaVeic string

type NumPlacaVeicCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumTel string

type NumTelCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Pais string

type PaisCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Percentual float64

type Percentual5 float64

type PercentualCodErr struct {
	CodErroAttr string  `xml:"CodErro,attr,omitempty"`
	Value       float64 `xml:",chardata"`
}

type Percentual5CodErr struct {
	CodErroAttr string  `xml:"CodErro,attr,omitempty"`
	Value       float64 `xml:",chardata"`
}

type PercentualCTC float64

type Qtd int

type Peso float64

type QualifcCessSemFinanc string

type QualifcCessSemFinancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type RegmAmtzc string

type RegmAmtzcCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Renavam string

type RenavamCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ResFase int

type QtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type SeqGarFidjssria int

type SeqGarFidjssriaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type Sit int

type SitCons int

type SitContrto string

type SitContrtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitContrtoGar string

type SitEvt string

type SitEvtCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitGar string

type SitGarCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitGarAgtValiddr string

type SitGarAgtValiddrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitMenslddRecup int

type SitMenslddRecupCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type SitParcl int

type SitParclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type SitPortbdd string

type SitRep string

type SitRepCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitReq int

type SitRes string

type SitRetContrto int

type SitRegistro string

type SituacaoCorda string

type SitRetContrtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type SitProc int

type LocalRegistro string

type SitSolicitacao int

type SitSolicitacaoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type Taxa float64

type TpCalc string

type TpBaixContrto string

type TpBaixContrtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpBaixaParcl string

type TpBaixaParclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpAcao string

type TpAcaoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpCalcCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpCobr string

type TpCoobr string

type TpCoobrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpCombtvl string

type TpCombtvlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpCobrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpContrto string

type TpContrtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpDivgte int

type InfDivgte int

type TpEnteCons int

type TpEnteConsCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpEnteConsSCR string

type TpEnteConsSCRCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpEquipmnt string

type TpEquipmntCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpEvt string

type TpInst string

type TpInstPart string

type TpInstntoFinanc string

type TpInstntoFinancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpImovl string

type TpImovlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpLeiaute string

type TpLeiauteCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Motv string

type TpMotv string

type TpMotvCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpMotvAltDadCad string

type TpMotvAltDadCadCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpPart string

type TpParcl string

type TpParclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpPartResponsTarCtrapart int

type TpPartResponsTarCtrapartCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpPessoa string

type TpPessoaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpPessoaInstntoFinanc string

type TpPessoaInstntoFinancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpProdt string

type TpProdtCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpRecompra string

type TpRecompraCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpRep string

type TpRepCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpRes string

type TpResCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpSolicitacao int

type TpSolicitacaoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpSituacao int

type TpSituacaoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpTabVeicl string

type TpTabVeiclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpTaxa string

type TpTaxaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpUsoImovl string

type TpUsoImovlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpVincEmprtcio int

type TpVincEmprtcioCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpVarcVlrFinancOp string

type TpVarcVlrFinancOpCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpVeicl int

type TpVeiclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpVencParcl string

type TpVencParclCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpManut string

type TpManutCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumIPOCSCR string

type NumIPOCSCRCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type FormaCalcJurosMora string

type FormaCalcJurosMoraCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type FormaPgto int

type FormaPgtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type Banco int

type BaseCalcMulta string

type BaseCalcMultaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type BancoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type Agencia int

type AgenciaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type CtBancaria string

type CtBancariaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Cidade string

type CodMunIBGE int

type CidadeCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodMunIBGECodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type ClassRscCesta string

type ClassRscCestaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CMC7 string

type CMC7CodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type LinhaDigtvl string

type LinhaDigtvlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpEncargoFinanc int

type TpEncargoFinancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpGar int

type TpGarCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpGarSCR int

type TpGarSCRCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpGarBacen int

type TpGarBacenCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type SubTpGarSCR int

type SubTpGarSCRCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpSexo string

type TpSexoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumDocIdentc string

type NumDocIdentcCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type EstadoCivil string

type EstadoCivilCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type End string

type EndCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CaracEspecial string

type CaracEspecialCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CaracEspecialBacen string

type CaracEspecialBacenCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CEP string

type CEPCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Periodicidade string

type PeriodicidadeCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type PorteCli int

type PorteCliCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type Indx string

type IndxCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type MesAno string

type MesAnoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ModOp string

type ModOpCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type MtvCanceltPortldd string

type MtvRetenContrto int

type MtvDevLiquidPortldd int

type Nacnl string

type NacnlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NatuOp int

type NatuOpCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type OrgaoExpddr string

type OrgaoExpddrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitEftcPortldd string

type TpEnd string

type TpEndCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpIdentc string

type TpIdentcCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpCtrl int

type TpCtrlCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type TpSolctcRel string

type TpSolctcRelCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type UF string

type UFCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type Valor float64

type ValorCodErr struct {
	CodErroAttr string  `xml:"CodErro,attr,omitempty"`
	Value       float64 `xml:",chardata"`
}

type PesoCodErr struct {
	CodErroAttr string  `xml:"CodErro,attr,omitempty"`
	Value       float64 `xml:",chardata"`
}

type NotaFiscalComplexType struct {
	CodErroAttr     string              `xml:"CodErro,attr,omitempty"`
	ChNota          *ChNotaCodErr       `xml:"ChNota"`
	NumNota         *NumNotaCodErr      `xml:"NumNota"`
	DtHrEmiss       *DataHoraCodErr     `xml:"DtHrEmiss"`
	DtVenc          *DataCodErr         `xml:"DtVenc"`
	CNPJEmit        *CNPJCodErr         `xml:"CNPJEmit"`
	InscrEstadlEmit *InscrEstadlCodErr  `xml:"InscrEstadlEmit"`
	CNPJCPFSacd     *CNPJCPFCodErr      `xml:"CNPJ_CPFSacd"`
	TpPessoaSacd    *TpPessoaSacdCodErr `xml:"TpPessoaSacd"`
	CNPJCPFINT      *CNPJCPFCodErr      `xml:"CNPJ_CPFINT"`
	TpPessoaINT     *TpPessoaINTCodErr  `xml:"TpPessoaINT"`
	Valor           *ValorCodErr        `xml:"Valor"`
	Desc            *DescCodErr         `xml:"Desc"`
}

type NotaFiscalDupComplexType struct {
	CodErroAttr     string                        `xml:"CodErro,attr,omitempty"`
	ChNota          *ChNotaCodErr                 `xml:"ChNota"`
	NumNota         *NumNotaFiscalCodErr          `xml:"NumNota"`
	SerieNFe        *SerieNFeErr                  `xml:"SerieNFe"`
	DtHrEmiss       *DataHoraCodErr               `xml:"DtHrEmiss"`
	DtVenc          *DataCodErr                   `xml:"DtVenc"`
	CNPJEmit        *CNPJCodErr                   `xml:"CNPJEmit"`
	InscrEstadlEmit *InscrEstadlCodErr            `xml:"InscrEstadlEmit"`
	CNPJCPFSacd     *CNPJCPFCodErr                `xml:"CNPJ_CPFSacd"`
	TpPessoaSacd    *TpPessoaFisicaJuridicaCodErr `xml:"TpPessoaSacd"`
	CNPJCPFINT      *CNPJCPFCodErr                `xml:"CNPJ_CPFINT"`
	TpPessoaINT     *TpPessoaFisicaJuridicaCodErr `xml:"TpPessoaINT"`
	Valor           *ValorCodErr                  `xml:"Valor"`
	Desc            *DescCodErr                   `xml:"Desc"`
	NumConsProt     *NumConsProtErr               `xml:"NumConsProt"`
	DtHrConsProt    *DtHrConsProtErr              `xml:"DtHrConsProt"`
	CodConsProt     *CodConsProtErr               `xml:"CodConsProt"`
}

type Desc string

type NumConsProt int

type DtHrConsProt string

type CodConsProt string

type TpPessoaSacd string

type TpPessoaINT string

type NumNota string

type SerieNFe string

type NumOrdem string

type VenctoOrdemComplexType struct {
	CodErroAttr string          `xml:"CodErro,attr,omitempty"`
	NumOrdem    *NumOrdemCodErr `xml:"NumOrdem"`
	DtVencto    *DataCodErr     `xml:"DtVencto"`
	VlrOrdem    *ValorCodErr    `xml:"VlrOrdem"`
}

type BoletoComplexType struct {
	CodErroAttr      string               `xml:"CodErro,attr,omitempty"`
	IDCIPBoleto      *NUCodErr            `xml:"IDCIPBoleto"`
	IDINSTIFGeradora *NUCodErr            `xml:"IDINSTIFGeradora"`
	NumBco           *NumBcoCodErr        `xml:"NumBco"`
	DigtVerificdr    *DigtVerificdrCodErr `xml:"DigtVerificdr"`
	LinhaDigitvl     *LinhaDigitvlCodErr  `xml:"LinhaDigitvl"`
	DtVenct          *DataCodErr          `xml:"DtVenct"`
	Valor            *ValorCodErr         `xml:"Valor"`
	AgCod            *AgCodCodErr         `xml:"AgCod"`
}

type BoletoDupComplexType struct {
	CodErroAttr      string               `xml:"CodErro,attr,omitempty"`
	IDCIPBoleto      *NUCodErr            `xml:"IDCIPBoleto"`
	NossoNumeroBOL   *NossoNumeroBOLErr   `xml:"NossoNumeroBOL"`
	NumOrdem         *NumOrdemCodErr      `xml:"NumOrdem"`
	IDINSTIFGeradora *NUCodErr            `xml:"IDINSTIFGeradora"`
	NumBco           *NumBcoCodErr        `xml:"NumBco"`
	DigtVerificdr    *DigtVerificdrCodErr `xml:"DigtVerificdr"`
	LinhaDigitvl     *LinhaDigitvlCodErr  `xml:"LinhaDigitvl"`
	DtVenct          *DataCodErr          `xml:"DtVenct"`
	Valor            *ValorCodErr         `xml:"Valor"`
	AgCod            *AgCodCodErr         `xml:"AgCod"`
}

type AgCod string

type AgCodCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type LinhaDigitvl string

type DigtVerificdr string

type NumBco string

type NUSolctcLiquidCIP string

type NUSolctcLiquidCIPCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitSolctcLiquid string

type NUSolctcLiquidOpCip string

type NUSolctcLiquidOpCipCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitSolctcLiquidOp string

type TpAtivoLiquid string

type TpLanc string

type CodAtivo string

type Nom string

type NumCotas string

type VlrTotLiqdar float64

type NumCtrlIMFParcr string

type NumCtrlOpIMFParcr string

type NUCtrlCIP string

type QtdCotas string

type IdentdPartPrincipalCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type IdentdPartAdmdoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CNPJIMFParcrCodErr struct {
	XMLName     xml.Name `xml:"CNPJ_IMFParcrCodErr"`
	CodErroAttr string   `xml:"CodErro,attr,omitempty"`
	Value       string   `xml:",chardata"`
}

type DtMovtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumCtrlIMFParcrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpAtivoLiquidCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpLancCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CodAtivoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpPessoaEmissorTitCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CNPJCPFEmissorTitCodErr struct {
	XMLName     xml.Name `xml:"CNPJ_CPFEmissorTitCodErr"`
	CodErroAttr string   `xml:"CodErro,attr,omitempty"`
	Value       string   `xml:",chardata"`
}

type NomEmissorTitCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type NumCotasCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type VlrTotLiqdarCodErr struct {
	CodErroAttr string  `xml:"CodErro,attr,omitempty"`
	Value       float64 `xml:",chardata"`
}

type TpPessoaLancDebtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CNPJCPFLancDebtdCodErr struct {
	XMLName     xml.Name `xml:"CNPJ_CPFLancDebtdCodErr"`
	CodErroAttr string   `xml:"CodErro,attr,omitempty"`
	Value       string   `xml:",chardata"`
}

type NomLancDebtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ISPBIFDebtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type AgLancDebtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type CtLancDebtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpPessoaLancCredtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CNPJCPFLancCredtdCodErr struct {
	XMLName     xml.Name `xml:"CNPJ_CPFLancCredtdCodErr"`
	CodErroAttr string   `xml:"CodErro,attr,omitempty"`
	Value       string   `xml:",chardata"`
}

type NomLancCredtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type ISPBIFCredtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type AgLancCredtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       int    `xml:",chardata"`
}

type CtLancCredtdCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type QtdCotasCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type VlrTotLiqdarGrupoCodErr struct {
	CodErroAttr string  `xml:"CodErro,attr,omitempty"`
	Value       float64 `xml:",chardata"`
}

type NumCtrlOpIMFParcrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type SitLiquidOp string

type VerfcIdentdd string

type CanContr string

type Device string

type Ip string

type Porta string

type LatLong string

type HashDocOr string

type HashDocAssinado string

type TpFormalizContrto string

type Cosif string

type CdCli string

type VerfcIdentddCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type CanContrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type DeviceCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type IpCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type PortaCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type LatLongCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type HashDocOrCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type HashDocAssinadoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type TpFormalizContrtoCodErr struct {
	CodErroAttr string `xml:"CodErro,attr,omitempty"`
	Value       string `xml:",chardata"`
}

type GrupoGeoLoclcDispositComplexType struct {
	XMLName xml.Name `xml:"Grupo_GeoLoclc_DispositComplexType"`
	Device  string   `xml:"Device"`
	Ip      string   `xml:"Ip"`
	Porta   string   `xml:"Porta"`
	Lat     string   `xml:"Lat"`
	Long    string   `xml:"Long"`
}

type GrupoDossieprobatComplexType struct {
	XMLName         xml.Name `xml:"Grupo_Dossie_probatComplexType"`
	DtHrCriacAlt    string   `xml:"DtHrCriacAlt"`
	HashDocOr       string   `xml:"HashDocOr"`
	DtHrAssinatura  string   `xml:"DtHrAssinatura"`
	HashDocAssinado string   `xml:"HashDocAssinado"`
}

type GrupoRETGeoLoclcDispositComplexType struct {
	XMLName     xml.Name       `xml:"GrupoRET_GeoLoclc_DispositComplexType"`
	CodErroAttr string         `xml:"CodErro,attr,omitempty"`
	Device      *DeviceCodErr  `xml:"Device"`
	Ip          *IpCodErr      `xml:"Ip"`
	Porta       *PortaCodErr   `xml:"Porta"`
	Lat         *LatLongCodErr `xml:"Lat"`
	Long        *LatLongCodErr `xml:"Long"`
}

type GrupoRETDossieprobatComplexType struct {
	XMLName         xml.Name               `xml:"GrupoRET_Dossie_probatComplexType"`
	CodErroAttr     string                 `xml:"CodErro,attr,omitempty"`
	DtHrCriacAlt    *DataHoraCodErr        `xml:"DtHrCriacAlt"`
	HashDocOr       *HashDocOrCodErr       `xml:"HashDocOr"`
	DtHrAssinatura  *DataHoraCodErr        `xml:"DtHrAssinatura"`
	HashDocAssinado *HashDocAssinadoCodErr `xml:"HashDocAssinado"`
}
