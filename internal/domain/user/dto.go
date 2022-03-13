package user

type GetUsersDTO struct {
	Query struct {
		Id          string `json:"id"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		CountryCode int    `json:"countryCode"`
		RegionCode  int    `json:"regionCode"`
		CreatedAt   int    `json:"createdAt"`
		UpdatedAt   int    `json:"updatedAt"`
	} `json:"query"`
	Sort struct {
		Id          string `json:"id"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		CountryCode string `json:"countryCode"`
		RegionCode  string `json:"regionCode"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
	}
}

type CreateUserDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CountryCode int    `json:"countryCode"`
	RegionCode  int    `json:"regionCode"`
}
