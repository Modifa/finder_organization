package models

type OrganizationResponse struct {
	ID           int64  `db:"organization_id"`
	Name         string `db:"org_name"`
	Description  string `db:"org_desc"`
	EmailAddress string `db:"emailaddress"`
	MobileNumber string `db:"mobile_number"`
	Status       string `db:"status"`
	Type_Name    string `db:"type_desc"`
	Type_ID      int64  `db:"type_id"`
	DateAdded    string `db:"dateadded"`
	Password     string `db:"org_password"`
}
