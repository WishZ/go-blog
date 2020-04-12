package dao

import (
	"go-blog/entity"
)



func GetTags(pageNum int, pageSize int, maps interface{}) (tags []entity.Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&entity.Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByID(tagId int) bool {
	var tag entity.Tag
	db.Select("id").Where("id = ?", tagId).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByName(name string) bool {
	var tag entity.Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&entity.Tag{
		Name:    name,
		State:   state,
		Creator: createdBy,
	})

	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&entity.Tag{})

	return true
}

func EditTag(id int, data interface {}) bool {
	db.Model(&entity.Tag{}).Where("id = ?", id).Updates(data)

	return true
}