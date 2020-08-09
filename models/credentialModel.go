package models

import (
	"errors"

	"github.com/culturadevops/cindi/libs"
	"github.com/jinzhu/gorm"
)

var VarCredential *Credential

type Credential struct {
	gorm.Model

	Owner    string `gorm:"type:varchar(50);not null;"`
	Name     string `gorm:"type:varchar(100);not null;"`
	Account  string `gorm:"type:varchar(200);not null;"`
	Password string `gorm:"type:varchar(320);not null;"`
}

func (this *Credential) Add(owner string, name string, account string, password string) error {

	this.Owner = owner
	this.Name = name
	this.Account = account
	this.Password = password
	if !libs.DB.Where("owner = ? AND  name = ? ", owner, name).First(&Credential{}).RecordNotFound() {
		return errors.New("Ya existe un item con el nombre " + name)
	}
	if err := libs.DB.Create(&this).Error; err != nil {
		return err
	}
	return nil
}

func (this *Credential) GetForId(owner string, id int64) (Credential, error) {
	var data = Credential{}

	if libs.DB.Where("owner = ? AND  id = ? ", owner, id).Find(&data).RecordNotFound() {
		return Credential{}, errors.New("no se encontro las Credentiales ")
	}
	return data, nil
}
func (this *Credential) Get(owner string, name string) (Credential, error) {
	var data = Credential{}

	if libs.DB.Where("owner = ? AND  name = ? ", owner, name).Find(&data).RecordNotFound() {
		return Credential{}, errors.New("no se encontro las Credentiales " + name)
	}
	return data, nil
}
func (this *Credential) List(owner string) []Credential {
	var data = []Credential{}

	err := libs.DB.Where("owner = ?", owner).Find(&data).Error
	if err != nil {
		//log.Fatalln(err)
	}
	return data
}
func (this *Credential) Del(owner string, name string) error {
	var data Credential

	if err := libs.DB.Where("owner = ? AND  name = ? ", owner, name).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Credential) DelForId(owner string, id int64) error {
	var data Credential

	if err := libs.DB.Where("owner = ? AND  id = ? ", owner, id).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Credential) Update(owner string, name string, account string, password string) error {
	var varCredential Credential

	varCredential.Name = name

	if libs.DB.Where("owner = ? AND  name = ? ", owner, name).Find(&varCredential).RecordNotFound() {
		return errors.New("no existe la Credential " + name)
	}
	varCredential.Account = account
	varCredential.Password = password
	if err := libs.DB.Save(&varCredential).Error; err != nil {
		return err
	}
	return nil
}
