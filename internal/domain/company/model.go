package company

type Company struct {
	Id          int    `json:"id"`
	FullName    string `json:"fullName"`
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
}
