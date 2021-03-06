package models

// ResponseJSONRow 这里重点讲一下，proto自动生成的每一个字段都会加上omitempty，当ParentID和Sorts为0时，会不显示该字段。所以重新转义
type ResponseJSONRow struct {
	ID             int64  `json:"ID"`
	Name           string `json:"Name"`
	URL            string `json:"URL"`
	ParentID       int64  `json:"ParentID"`
	Sorts          int32  `json:"Sorts"`
	IsShowNav      int32  `json:"IsShowNav"`
	CssIcon        string `json:"CssIcon"`
	CreateTime     int64  `json:"CreateTime"`
	LastUpdateTime int64  `json:"LastUpdateTime"`
}
