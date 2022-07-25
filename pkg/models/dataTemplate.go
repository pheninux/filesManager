package models

type DataTemplate struct {
	DirIn  string   `json:"dir_in"`
	DirOut string   `json:"dir_out"`
	Action string   `json:"action"`
	Exts   []string `json:"exts"`
	Count  int      `json:"ccount"`
}
