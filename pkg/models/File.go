package models

type File struct {
	DirIn  string   `json:"dir_in"`
	DirOut string   `json:"dir_out"`
	Action string   `json:"action"`
	Exts   []string `json:"exts"`
}
