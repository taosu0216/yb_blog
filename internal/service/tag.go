package service

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint   `form:"state,default=1" binding:"oneof=0 1"`
}
type TafListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint   `form:"state,default=1" binding:"oneof=0 1"`
}
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint   `form:"state,default=1" binding:"oneof=0 1"`
}
type UpdateTagRequest struct {
	ID         uint   `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,min=3,max=100"`
	State      uint   `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}
type DeleteTagRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}
