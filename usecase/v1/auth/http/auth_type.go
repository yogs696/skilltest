package uv1authroute

type register__Reqp struct {
	Username string `json:"username" form:"username" xml:"username" validate:"required"`
	Email    string `json:"email" form:"email" xml:"email" validate:"required"`
	Password string `json:"password" form:"password" xml:"password" validate:"required"`
}

type Login__Reqp struct {
	Email    string `json:"email" form:"email" xml:"email" validate:"required"`
	Password string `json:"password" form:"password" xml:"password" validate:"required"`
}
