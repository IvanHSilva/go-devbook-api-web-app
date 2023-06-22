package models

type Post struct {
	Id       uint64 `json:"id,omitempty"`
	Title    uint64 `json:"title,omitempty"`
	Content  uint64 `json:"content,omitempty"`
	AuthorId uint64 `json:"authorid,omitempty"`
	Likes    uint64 `json:"likes"`
	RegDate  string `json:"regdate,omitempty"`
}
