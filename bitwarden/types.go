package bitwarden

type Folder struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Notes string `json:"notes"`
}
