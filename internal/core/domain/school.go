package domain

type School struct {
	AuditFields
	Name      string `json:"name" db:"name"`
	SubDomain string `json:"sub_domain" db:"sub_domain"`
	IsActive  bool   `json:"is_active" db:"is_active"`
}
