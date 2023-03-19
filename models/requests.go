package models

type OrganizationRequest struct {
	Name          string `json:"org_name"`
	Description   string `json:"org_desc"`
	Logo_url      string `json:"org_logo"`
	Email_address string `json:"org_email_address"`
	Mobile_Number string `json:"org_mobile_number"`
	Type          int64  `json:"organizatio_type"`
	Password      string `json:"org_password"`
}
