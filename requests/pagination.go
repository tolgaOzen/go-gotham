package requests

type Pagination struct {
	/**
    * PATH
    */
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
	OrderBy string `query:"order_by"`
	SortBy  string `query:"sort_by"`
}
