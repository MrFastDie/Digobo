package osu

type Achievement struct {
	IconUrl      string      `json:"icon_url"`
	Id           int         `json:"id"`
	Name         string      `json:"name"`
	Grouping     string      `json:"grouping"`
	Ordering     int         `json:"ordering"`
	Slug         string      `json:"slug"`
	Description  string      `json:"description"`
	Mode         interface{} `json:"mode"`
	Instructions string      `json:"instructions"`
}
