package monta

// PageMeta holds page metadata.
type PageMeta struct {
	// ItemCount is the number of items in the page.
	ItemCount int32 `json:"itemCount"`
	// CurrentPage is the index of the current page.
	CurrentPage int32 `json:"currentPage"`
	// PerPage is the requested number of items per page.
	PerPage int32 `json:"perPage"`
	// TotalPageCount is the total number of pages.
	TotalPageCount int32 `json:"totalPageCount"`
	// TotalItemCount is the total number of items.
	TotalItemCount int64 `json:"totalItemCount"`
}
