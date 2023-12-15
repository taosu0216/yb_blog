package service

type ArticleRequest struct {
	ID    uint `form:"id" binding:"required,gte=1"`
	State uint `form:"state,default=1" binding:"oneof=0 1"`
}
type ArticleListRequest struct {
	TagID uint `form:"tag_id" binding:"required,gte=1"`
	State uint `form:"state,default=1" binding:"oneof=0 1"`
}
type CreateArticleRequest struct {
	TagID     uint   `form:"tag_id" binding:"required,gte=1"`
	Title     string `form:"title" binding:"max=100"`
	State     uint   `form:"state,default=1" binding:"oneof=0 1"`
	Desc      string `form:"desc" binding:"max=100"`
	Content   string `form:"content"`
	CoverImg  string `form:"cover_img"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}
type UpdateArticleRequest struct {
	ID         uint   `form:"id" binding:"required,gte=1"`
	TagID      uint   `form:"tag_id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"max=100"`
	State      uint   `form:"state,default=1" binding:"oneof=0 1"`
	Desc       string `form:"desc" binding:"max=100"`
	Content    string `form:"content"`
	CoverImg   string `form:"cover_img"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}
type DeleteArticleRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}
