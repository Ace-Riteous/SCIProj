package model

type Student struct {
	Username  string    `gorm:"type:varchar(20);not null;index:studentname;comment:学生姓名" json:"username"`
	Password  string    `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	Email     string    `gorm:"type:varchar(100);default:null;comment:邮箱" json:"email"`
	Phone     string    `gorm:"type:varchar(20);default:null;comment:电话号码" json:"phone"`
	Role      string    `gorm:"type:varchar(50);default:null;comment:职位" json:"role"`
	Avatar    string    `gorm:"type:varchar(100);default:null;comment:头像" json:"avatar"`
	Age       int       `gorm:"type:int;comment:年龄" json:"age"`
	AuthID    string    `gorm:"type:varchar(10);primary_key;not null;unique;comment:统一认证码" json:"authid"`
	StudentID string    `gorm:"type:varchar(20);unique;index:studentid;comment:学号" json:"studentid"`
	MyTeacher []Teacher `gorm:"type:varchar(20);many2many:student_teacher;comment:指导老师" json:"my_teacher"`
}

type Teacher struct {
	Username      string        `gorm:"type:varchar(20);not null;index:teachername;comment:老师姓名" json:"username"`
	Password      string        `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	Email         string        `gorm:"type:varchar(100);default:null;comment:邮箱" json:"email"`
	Phone         string        `gorm:"type:varchar(20);default:null;comment:电话号码" json:"phone"`
	Role          string        `gorm:"type:varchar(50);default:null;comment:职位" json:"role"`
	Avatar        string        `gorm:"type:varchar(100);default:null;comment:头像" json:"avatar"`
	Age           int           `gorm:"type:int;comment:年龄" json:"age"`
	TeacherID     string        `gorm:"type:varchar(10);not null;primary_key;comment:统一认证码" json:"teacherid"`
	MyStudent     []Student     `gorm:"type:varchar(20);many2many:student_teacher;comment:我的学生" json:"mystudent"`
	MyTeam        []Team        `gorm:"type:varchar(20);foreignkey:TeamID;comment:我的团队" json:"myteam"`
	MyCompetition []Competition `gorm:"type:varchar(10);foreignkey:CID;我组织的竞赛" json:"mycompetition"`
}

type Team struct {
	Student []Student `gorm:"type:varchar(20);many2many:team_student;comment:团队成员" json:"teamstudent"`
	Teacher Teacher   `gorm:"type:varchar(20);not null;comment:指导老师" json:"teamteacher"`
	TeamId  string    `gorm:"type:varchar(20);primary_key;unique_index;comment:团队号" json:"teamid"`
	Name    string    `gorm:"type:varchar(20);not null;index:teamname;comment:团队姓名" json:"teamname"`
}
