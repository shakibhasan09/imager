package models

type Project struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
