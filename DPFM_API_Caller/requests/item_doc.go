package requests

type ItemDoc struct {
	FreightAgreement         int    `json:"FreightAgreement"`
	FreightAgreementItem     int    `json:"FreightAgreementItem"`
	DocType                  string `json:"DocType"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	FileExtension            string `json:"FileExtension"`
	FileName                 string `json:"FileName"`
	FilePath                 string `json:"FilePath"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
