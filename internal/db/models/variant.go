package models

type Fit string

const (
	FitCover   Fit = "cover"
	FitContain Fit = "contain"
)

type Variant struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Metadata    bool   `json:"metadata"`
	ProjectUuid string `json:"project_uuid"`
	Fit         Fit    `json:"fit"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
