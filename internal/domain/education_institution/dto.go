package education_institution

type CreateEducationInstitutionDTO struct {
	FullName    string `json:"fullName"`
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
}
