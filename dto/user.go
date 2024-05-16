package dto

import (
	"SCIProj/model"
	"github.com/go-playground/validator/v10"
)

type UserLoginDTO struct {
	SevenID  string `json:"seven_id" form:"seven_id" validate:"required,len=7"`
	Password string `form:"password" validate:"required,min=6,max=20"`
}

func (m *UserLoginDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type UserRegisterDTO struct {
	Username   string `json:"username" form:"username" validate:"required,min=2,max=20"`
	SevenID    string `json:"seven_id" form:"seven_id" validate:"required,len=7"`
	Password   string `form:"password" validate:"required,min=6,max=20"`
	RePassword string `form:"re_password" validate:"required,eqfield=Password"`
	StudentID  string `json:"student_id" form:"student_id" validate:"required,len=10"`
}

func (m *UserRegisterDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

func (m *UserRegisterDTO) ConvertToModel(student *model.Student) {
	student.Username = m.Username
	student.Password = m.Password
	student.SevenID = m.SevenID
	student.StudentID = m.StudentID
}
