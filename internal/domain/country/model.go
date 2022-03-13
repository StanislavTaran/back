package country

type Country struct {
	Code      int    `json:"code"`
	ShortName string `json:"shortName"`
	FullName  string `json:"fullName"`
}
