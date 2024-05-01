package dto

type CatReq struct {
	Name string `json:"name" validate:"required,min=1,max=30"`
	Race string `json:"race" validate:"required,min=1,max=30"`
}
