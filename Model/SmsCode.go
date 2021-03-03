package Model

type Code struct {
	Id         int64  `xorm:"pk autoincr" json:"id"`
	Name       string `xorm:"varchar(11)" json:"name"`
	Code       string `xorm:"varchar(6)" json:"code"`
	CreateTime int64  `xorm:"bigint" json:"create_time"`
}
