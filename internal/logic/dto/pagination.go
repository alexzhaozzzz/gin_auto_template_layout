package dto

type Pagination struct {
	PageIndex int `json:"page_index" form:"page_index" `
	PageSize  int `json:"page_size" form:"page_size" `
	TotalNum  int `json:"total_num" form:"total_num" `
}

func (s *Pagination) GetPageIndex() int {
	if s.PageIndex <= 0 {
		s.PageIndex = 1
	}
	return s.PageIndex
}

func (s *Pagination) GetPageSize() int {
	if s.PageSize <= 0 {
		s.PageSize = 10
	}
	return s.PageSize
}

func (s *Pagination) GetOffset() int {
	return (s.GetPageIndex() - 1) * s.GetPageSize()
}
