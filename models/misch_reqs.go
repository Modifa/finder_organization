package models

type GetOrganizationProfile struct {
	EmailAddress string `json:"email_address"`
}

type OrganizationRequest1 struct {
	Email    string `json:"email_address"`
	Password string `json:"password"`
}

//orgization_id bigint
type GetOrganizationProfileID struct {
	OrganizationID int64 `json:"orgization_id"`
}
