package domain

// Base 通用的 response 结构
type Base struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data, omitempty"`
}

// 分页结构中的数据部分
type PageData struct {
	TotalNum  int64       `json:"total_item"`
	TotalPage int64       `json:"total_page"`
	ItemList  interface{} `json:"item_list,omitempty"`
}

// Page 分页结构
type Page struct {
	Base
	Data PageData `json:"data"`
}
