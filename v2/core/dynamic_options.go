package core

type OffsetPaginationMeta struct {
	Offset     int  `json:"offset"`
	Limit      int  `json:"limit"`
	TotalItems int  `json:"totalItems"`
	HasMore    bool `json:"hasMore"`
}

type DynamicOptionsResponse struct {
	Metadata OffsetPaginationMeta `json:"metadata"`
	Items    any                  `json:"items"`
}

type DynamicOptionsFilterParams struct {
	Offset     int    `json:"offset"` // The offset of the first item to return (default: 0)
	Limit      int    `json:"limit"`  // The maximum number of items to return (default: 10)
	FilterTerm string `json:"filterTerm"`
}
