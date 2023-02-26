package models

type User struct {
	ID       uint   `gorm:"primary key"`
	Name     string `gorm:"type:varchar(50);not null"`
	Login    string `gorm:"type:varchar(20); not null"`
	Password string `gorm:"not null"`
}

type Item struct {
	ID    uint    `gorm:"primary key"`
	Name  string  `gorm:"type:varchar(50); not null"`
	Price float32 `gorm:"not null"`
	Rate  float32 `gorm:"not null"`
}
