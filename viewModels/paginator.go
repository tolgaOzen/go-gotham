package viewModels

type Paginator struct {
	TotalRecord int64       `json:"total_record"`
	Records     interface{} `json:"records"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
}
