package moddle

type UserMenu struct {
	Role int `json:"role"`
	Path string	`json:"path"`
	Name string	`json:"name"`
	Icon string `json:"icon"`
	Label string `json:"label"`
	Url	string	`json:"url"`
}
