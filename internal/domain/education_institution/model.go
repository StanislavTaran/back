package education_institution

type EductionInstitution struct {
	Id          int    `json:"id"`
	FullName    string `json:"fullName"`
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
}
