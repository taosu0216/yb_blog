package model

type ArticalTag struct {
	*GblueModel
	ArticalID uint `json:"artical_id"`
	TagID     uint `json:"tag_id"`
}

func (a *ArticalTag) TableName() string {
	return "blog_article_tag"
}
