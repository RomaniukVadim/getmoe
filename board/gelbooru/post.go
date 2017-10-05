package gelbooru

// Post contains native Gelbooru data
type Post struct {
	Directory    string      `json:"directory"`
	Hash         string      `json:"hash"`
	Height       int         `json:"height"`
	ID           int         `json:"id"`
	Image        string      `json:"image"`
	Change       int         `json:"change"`
	Owner        string      `json:"owner"`
	ParentID     interface{} `json:"parent_id"`
	Rating       string      `json:"rating"`
	Sample       bool        `json:"sample"`
	SampleHeight int         `json:"sample_height"`
	SampleWidth  int         `json:"sample_width"`
	Score        int         `json:"score"`
	Tags         string      `json:"tags"`
	Width        int         `json:"width"`
	FileURL      string      `json:"file_url"`
}
