package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name      string `json:"name"`
	Creator   string `json:"creator"`
	State     int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByID(tagId int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", tagId).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:    name,
		State:   state,
		Creator: createdBy,
	})

	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface {}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}


func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	fmt.Println(time.Now().UTC())
	_ = scope.SetColumn("CreatedAt", time.Now().UTC())

	return nil
}
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("UpdatedAt", time.Now().UTC())

	return nil
}