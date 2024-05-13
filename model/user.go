package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null;comment:学生姓名" json:"username"`
	Password  string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	Email     string `gorm:"type:varchar(100);default:null;comment:邮箱" json:"email"`
	Phone     string `gorm:"type:varchar(20);default:null;comment:电话号码" json:"phone"`
	Role      string `gorm:"type:varchar(50);default:null;comment:职位" json:"role"`
	Avatar    string `gorm:"type:varchar(100);default:null;comment:头像" json:"avatar"`
	Age       int    `gorm:"type:int;comment:年龄" json:"age"`
	SevenID   string `gorm:"type:varchar(10);default:null;comment:统一认证码" json:"sevenid"`
	StudentID string `gorm:"type:varchar(20);comment:学号" json:"studentid"`
	MyTeacher string `gorm:"type:longtext;comment:指导老师" json:"my_teacher"`
}

type Teacher struct {
	gorm.Model
	TeacherID     string `gorm:"type:varchar(10);comment:统一认证码" json:"teacherid"`
	Username      string `gorm:"type:varchar(20);not null;comment:老师姓名" json:"username"`
	Password      string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	Email         string `gorm:"type:varchar(100);default:null;comment:邮箱" json:"email"`
	Phone         string `gorm:"type:varchar(20);default:null;comment:电话号码" json:"phone"`
	Role          string `gorm:"type:varchar(50);default:null;comment:职位" json:"role"`
	Avatar        string `gorm:"type:varchar(100);default:null;comment:头像" json:"avatar"`
	Age           int    `gorm:"type:int;comment:年龄" json:"age"`
	MyStudent     string `gorm:"type:longtext;comment:我的学生" json:"mystudent"`
	MyTeam        string `gorm:"type:longtext;comment:我的团队" json:"myteam"`
	MyCompetition string `gorm:"type:longtext;comment:我组织的竞赛" json:"mycompetition"`
}

type Team struct {
	gorm.Model
	StudentIds string `gorm:"type:varchar(100);comment:团队成员" json:"teamstudentids"`
	TeacherId  string `gorm:"type:varchar(20);comment:指导老师" json:"teamteacherid"`
	CID        int    `gorm:"type:int;comment:竞赛号" json:"cid"`
	Name       string `gorm:"type:varchar(20);not null;comment:团队姓名" json:"teamname"`
	IsFull     bool   `gorm:"type:bool;default:false;comment:是否满员" json:"isfull"`
}
