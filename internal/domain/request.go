package domain

type PageSearch struct {
	PageNum  int `json:"page_num,omitempty" binding:"required" form:"page_num"`
	PageSize int `json:"page_size,omitempty" binding:"required" form:"page_size"`
}

func (ps *PageSearch) GetPageNum() int {
	return ps.PageNum
}

func (ps *PageSearch) GetPageSize() int {
	return ps.PageSize
}

type TimeParam struct {
	DateType  string `json:"date_type" form:"date_type"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

func (t *TimeParam) GetDateType() string {
	return t.DateType
}

func (t *TimeParam) GetStartTime() int64 {
	return t.StartTime
}

func (t *TimeParam) GetEndTime() int64 {
	return t.EndTime
}

type TimeAndPageSearch struct {
	PageSearch
	TimeParam
}
