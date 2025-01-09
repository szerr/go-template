package domain

import (
	"math"
)

func FillInPageResponseData(req PageSearch, count int64, data any) *PageData {
	return &PageData{
		TotalNum:  count,
		TotalPage: int64(math.Ceil(float64(count) / float64(req.PageSize))),
		ItemList:  data,
	}
}
