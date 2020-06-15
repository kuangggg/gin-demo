package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error

	if err == gorm.ErrRecordNotFound {
		return false
	}

	if err != nil {
		log.Printf("db find error %v\n", err)
		return false
	}

	return true
}

func ExistTagById(id int) bool {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	fmt.Println(id)
	fmt.Println(err)

	if err == gorm.ErrRecordNotFound {
		return false
	}

	if err != nil {
		log.Printf("db find error %v\n", err)
		return false
	}

	return true
}

func AddTag(name string, state int, createdBy string) error {
	return db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}


