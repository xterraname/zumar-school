package dto

type ClassResponse struct {
	ID     uint   `json:"id"`
	Degree uint8  `json:"degree"`
	Group  string `json:"group"`
}
