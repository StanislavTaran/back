package education_institution

type EducationInstitution struct {
	Id          int    `json:"id"`
	FullName    string `json:"fullName"`
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
}
