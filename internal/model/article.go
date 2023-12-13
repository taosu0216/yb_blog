package model

type Article struct {
	*GblueModel
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Content  string `json:"content"`
	CoverImg string `json:"cover_img"`
	State    uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}
