package models

// ColumnRow 解析栏目的struct
type ColumnRow struct {
	ID             int64  `json:"ID"`
	Name           string `json:"Name"`
	URL            string `json:"URL"`
	ParentID       int64  `json:"ParentID,omitempty"`
	Sorts          int64  `json:"Sorts,omitempty"`
	IsShowNav      int32  `json:"IsShowNav,omitempty"`
	CssIcon        string `json:"CssIcon,omitempty"`
	CreateTime     int64  `json:"CreateTime"`
	LastUpdateTime int64  `json:"LastUpdateTime"`
	Nodes          []*ColumnRow
}
