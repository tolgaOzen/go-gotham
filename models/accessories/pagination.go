package accessories

type Paginator struct {
	TotalRecord int         `json:"total_record"`
	Records     interface{} `json:"records"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
}
