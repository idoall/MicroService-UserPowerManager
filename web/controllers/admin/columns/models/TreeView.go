package models

// TreeView struct
type TreeView struct {
	ID    int64          `json:"id"`
	Text  string         `json:"text"`
	Href  string         `json:"href"`
	Tags  []string       `json:"tags"`
	Nodes []*TreeView    `json:"nodes"`
	State *TreeViewState `json:"state"`
}

type TreeViewState struct {
	Checked  bool `json:"checked"`
	Disabled bool `json:"disabled"`
	Expanded bool `json:"expanded"`
	Selected bool `json:"selected"`
}
