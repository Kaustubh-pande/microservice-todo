package model

import (
	// test "github.com/docker/docker/api/types/time"
	// "time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	// ID       int32     `gorm:"column:Id" json:"Id"`
	Name     string    `gorm:"column:Name" json:"Name"`
	Email    string    `gorm:"column:Email" json:"Email"`
	Password string    `gorm:"column:Password" json:"Password"`
	Token    string    `gorm:"column:Token" json:"Token"`
	Todos []Todo `gorm:"foreignkey:uid; association_foreignkey:	id" `

	// CreatedAt time.Time `gorm:"column:CreatedAt" sql:"DEFAULT:current_timestamp" json:"created_at"`
	
	//ModifiedDate time.Time `gorm:"column:ModifiedDate" json:"ModifiedDate"`
}

type Todo struct {
	gorm.Model
	// ID       int32     `gorm:"column:Id" json:"Id"`
	Task  string `gorm:"column:Task" json:"Task"`		
	Uid int32 `gorm:"column:uid; DEFAULT:null " json:"uid"`
	//User []User `gorm:"foreignkey:id" gorm:"column:Task"json:"Task"`
	
}

// TableName sets the insert table name for this struct type
func (an *User) TableName() string {
	return "users"
}

func (an *User) CreateUser(db *gorm.DB) error {
	return db.Create(an).Error
}
func (an *Todo) TableName() string {
	return "todos"
}

func (an *Todo) CreateTodo(db *gorm.DB) error {
	return db.Create(an).Error
}