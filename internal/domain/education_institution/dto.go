package education_institution

type CreateEducationInstitutionInputDTO struct {
	FullName    string `json:"fullName"`
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
}
