package company

type CreateCompanyInputDTO struct {
	FullName    string `json:"fullName"`
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
}
