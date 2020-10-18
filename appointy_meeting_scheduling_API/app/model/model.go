package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/MongoDB"
)

type meeting struct {
	gorm.Model
	Title    string `gorm:"unique" json:"title"`
	Archived bool   `json:"archived"`
	Tasks    []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
}

func (p *meeting) Archive() {
	p.Archived = true
}

func (p *meeting) Restore() {
	p.Archived = false
}

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Priority  string     `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Done      bool       `json:"done"`
	MeetingID uint       `json:"meeting_id"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = false
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Meeting{}, &Task{})
	db.Model(&Task{}).AddForeignKey("meeting_id", "meeting(id)", "CASCADE", "CASCADE")
	return db
}

