package models

type DataTemplate struct {
	DirIn  string   `json:"dir_in"`
	DirOut string   `json:"dir_out"`
	Action string   `json:"action"`
	Exts   []string `json:"exts"`
	Stack  *Stack
}

type Stack struct {
	Ffound  map[string]int // file found
	Err     string         // error handled
	Pattern []string       // urls path
	Fcount  int            // all files count
}
