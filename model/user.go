package model

type Student struct {
	Username  string `gorm:"type:varchar(20);not null;index:studentname;comment:学生姓名" json:"username"`
	Password  string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	Email     string `gorm:"type:varchar(100);default:null;comment:邮箱" json:"email"`
	Phone     string `gorm:"type:varchar(20);default:null;comment:电话号码" json:"phone"`
	Role      string `gorm:"type:varchar(50);default:null;comment:职位" json:"role"`
	Avatar    string `gorm:"type:varchar(100);default:null;comment:头像" json:"avatar"`
	Age       int    `gorm:"type:int;comment:年龄" json:"age"`
	SevenID   string `gorm:"type:varchar(10);default:null;index:sevenid;comment:统一认证码" json:"sevenid"`
	StudentID string `gorm:"type:varchar(20);index:studentid;primary_key;comment:学号" json:"studentid"`
	MyTeacher string `gorm:"type:longtext;comment:指导老师" json:"my_teacher"`
}

type Teacher struct {
	TeacherID     string `gorm:"type:varchar(10);index:teacherid;primary_key;comment:统一认证码" json:"teacherid"`
	Username      string `gorm:"type:varchar(20);not null;index:teachername;comment:老师姓名" json:"username"`
	Password      string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	Email         string `gorm:"type:varchar(100);default:null;comment:邮箱" json:"email"`
	Phone         string `gorm:"type:varchar(20);default:null;comment:电话号码" json:"phone"`
	Role          string `gorm:"type:varchar(50);default:null;comment:职位" json:"role"`
	Avatar        string `gorm:"type:varchar(100);default:null;comment:头像" json:"avatar"`
	Age           int    `gorm:"type:int;comment:年龄" json:"age"`
	MyStudent     string `gorm:"type:longtext;comment:我的学生" json:"mystudent"`
	MyTeam        string `gorm:"type:longtext;comment:我的团队" json:"myteam"`
	MyCompetition string `gorm:"type:longtext;我组织的竞赛" json:"mycompetition"`
}

type Team struct {
	Student string `gorm:"type:varchar(50);comment:团队成员" json:"teamstudent"`
	Teacher string `gorm:"type:varchar(20);comment:指导老师" json:"teamteacher"`
	TeamId  string `gorm:"type:varchar(20);index:teamid;primary_key;comment:团队号" json:"teamid"`
	Name    string `gorm:"type:varchar(20);not null;index:teamname;comment:团队姓名" json:"teamname"`
}
