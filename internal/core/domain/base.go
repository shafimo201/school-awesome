package domain

import "time"

type AuditFields struct {
	ID        string     `json:"id" db:"id"`
	SchoolID  string     `json:"school_id" db:"school_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	DeletedBy string     `json:"deleted_by,omitempty" db:"deleted_by"`
}
