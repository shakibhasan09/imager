package models

type Image struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Metadata    string `json:"metadata"`
	ProjectUuid string `json:"project_uuid"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
