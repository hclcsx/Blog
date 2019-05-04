package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model

	Name  string `json:"name"`
	State int    `json:"state"`
}

func GetTags(offset int, limit int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(offset).Limit(limit).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExitsTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int) bool {
	db.Create(&Tag{Name: name, State: state})
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
