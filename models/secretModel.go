package models

import (
	"errors"

	"github.com/culturadevops/cindi/libs"
	"github.com/jinzhu/gorm"
)

var VarSecret *Secret

type Secret struct {
	gorm.Model
	Owner  string `gorm:"type:varchar(50);not null;"`
	Name   string `gorm:"type:varchar(100);not null;"`
	Secret string `gorm:"type:text;not null;"`
}

func (this *Secret) Add(owner string, name string, text string) error {
	this.Owner = owner
	this.Name = name
	this.Secret = text

	if !libs.DB.Where("owner = ? AND  name = ? ", owner, name).First(&Secret{}).RecordNotFound() {
		return errors.New("Ya existe un item con el nombre " + name)
	}
	if err := libs.DB.Create(&this).Error; err != nil {
		return err
	}
	return nil
}

func (this *Secret) GetForId(owner string, id int64) (Secret, error) {
	var data = Secret{}

	if libs.DB.Where("owner = ? AND  id = ? ", owner, id).Find(&data).RecordNotFound() {
		return Secret{}, errors.New("no se encontro las credenciales ")
	}
	return data, nil
}
func (this *Secret) Get(owner string, name string) (Secret, error) {
	var data = Secret{}

	if libs.DB.Where("owner = ? AND  name = ? ", owner, name).Find(&data).RecordNotFound() {
		return Secret{}, errors.New("no se encontro las credenciales " + name)
	}
	return data, nil
}
func (this *Secret) List(owner string) []Secret {
	var data = []Secret{}

	err := libs.DB.Where("owner = ? ", owner).Find(&data).Error
	if err != nil {
		//log.Fatalln(err)
	}
	return data
}
func (this *Secret) Del(owner string, name string) error {
	var data Secret

	if err := libs.DB.Where("owner = ? AND  name = ? ", owner, name).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Secret) DelForId(owner string, id int64) error {
	var data Secret

	if err := libs.DB.Where("owner = ? AND  id = ? ", owner, id).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Secret) Update(owner string, name string, secret string) error {
	var varSecret Secret

	varSecret.Name = name

	if libs.DB.Where("owner = ? AND  name = ? ", owner, name).Find(&varSecret).RecordNotFound() {
		return errors.New("no existe la credencial " + name)
	}
	varSecret.Secret = secret

	if err := libs.DB.Save(&varSecret).Error; err != nil {
		return err
	}
	return nil
}
