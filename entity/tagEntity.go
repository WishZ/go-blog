package entity

type Tag struct {
	Model

	Name      string `json:"name"`
	Creator   string `json:"creator"`
	State     int    `json:"state"`
}