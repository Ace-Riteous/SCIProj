package dto

import (
	"SCIProj/model"
	"github.com/go-playground/validator/v10"
)

type UserLoginDTO struct {
	SevenID  int64  `json:"seven_id" form:"seven_id" validate:"required,min=1000000,max=9999999"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
}

func (m *UserLoginDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type UserRegisterDTO struct {
	Username   string `json:"username" form:"username" validate:"required,min=2,max=20"`
	SevenID    int64  `json:"seven_id" form:"seven_id" validate:"required,min=1000000,max=9999999"`
	Password   string `json:"password" form:"password" validate:"required,min=6,max=20"`
	RePassword string `json:"re_password" form:"re_password" validate:"required,eqfield=Password"`
	StudentID  int64  `json:"student_id" form:"student_id" validate:"required,min=1000000000,max=9999999999"`
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
