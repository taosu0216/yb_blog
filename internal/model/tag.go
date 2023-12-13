package model

type Tag struct {
	*GblueModel
	Name  string `json:"name"`
	State uint   `json:"state"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}
