package models

type Passenger struct {
	ID       string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name     string `sql:"type:varchar(30)"`
	Survived uint8
	PClass   uint8
	Sex      string
	Age      uint8
	SibSp    uint8
	Parch    uint8
	Ticket   string
	Fare     float32
	Cabin    string
	Embarked string
}
