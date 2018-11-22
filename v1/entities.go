package v1

type MeStruct struct {
	Email    string `json:"email"`
	Enabled  bool   `json:"enabled"`
	HomeId   int    `json:"homeId"`
	Id       string `json:"id"`
	Locale   string `json:"locale"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Username string `json:"username"`
}
