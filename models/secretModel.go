package models

import (

	"errors"

	"github.com/culturadevops/cindi/libs"
	"github.com/jinzhu/gorm"
)
type Secret struct {
	gorm.Model
	Name     string `gorm:"type:varchar(200);not null;"`
	Secret     string `gorm:"type:text;not null;"`
}

func (this *Secret) Add(name string,text string) error {
	var varSecret Secret
	db := libs.DB
	varSecret.Name=name
	varSecret.Secret=text

	if !db.Where("name = ? ", varSecret.Name).First(&Secret{}).RecordNotFound() {
		return errors.New("Ya existe un item con el nombre "+varSecret.Name )
	}
	if err := db.Create(&varSecret).Error; err != nil {
		return err
	}
	return nil
}
func (this *Secret) Get(name string) (Secret,error){
	var data = Secret{}
	db := libs.DB

	if db.Where("name in (?)", name).Find(&data).RecordNotFound() {
		return Secret{}, errors.New("no se encontro las credenciales "+ name)
	}
	return data, nil
}
func (this *Secret) List() []Secret {
	var data = []Secret{}
	db := libs.DB

	err := db.Find(&data).Error
	if err != nil {
		//log.Fatalln(err)
	}
	return data
}
func (this *Secret) Del(name string) error {
	var data Secret
	db := libs.DB
	if err := db.Where("name = ?", name).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Secret) DelForId(id int64) error {
	var data Secret
	db := libs.DB


	if err := db.Where("id = ?", id).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Secret) Update(name string,secret string) error {
	var varSecret Secret
	db := libs.DB
	varSecret.Name=name
	
	if db.Where("name = ?", varSecret.Name).Find(&varSecret).RecordNotFound() {
		return errors.New("no existe la credencial " + name)
	}
	varSecret.Secret=secret

	if err := db.Save(&varSecret).Error; err != nil {
		return err
	}
	return nil
}