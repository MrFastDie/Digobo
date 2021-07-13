package osu

type User struct {
	Username         string `json:"username"`
	PreviousUsername string `json:"previousUsername,omitempty"`
	Url              string `json:"url"`
}
