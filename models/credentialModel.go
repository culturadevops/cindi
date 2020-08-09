package models

import (

	"errors"

	"github.com/culturadevops/cindi/libs"
	"github.com/jinzhu/gorm"
)
type Credencial struct {
	gorm.Model
	Name     string `gorm:"type:varchar(200);not null;"`
	Account     string `gorm:"type:varchar(200);not null;"`
	Password string    `gorm:"type:varchar(320);not null;"`

}

func (this *Credencial) Add(name string,account string, password string) error {
	var varCredencial Credencial
	db := libs.DB
	varCredencial.Name=name
	varCredencial.Account=account
	varCredencial.Password=password
	if !db.Where("name = ? ", varCredencial.Name).First(&Credencial{}).RecordNotFound() {
		return errors.New("Ya existe un item con el nombre "+varCredencial.Name )
	}
	if err := db.Create(&varCredencial).Error; err != nil {
		return err
	}
	return nil
}
func (this *Credencial) Get(name string) (Credencial,error){
	var data = Credencial{}
	db := libs.DB

	if db.Where("name in (?)", name).Find(&data).RecordNotFound() {
		return Credencial{}, errors.New("no se encontro las credenciales "+ name)
	}
	return data, nil
}
func (this *Credencial) List() []Credencial {
	var data = []Credencial{}
	db := libs.DB

	err := db.Find(&data).Error
	if err != nil {
		//log.Fatalln(err)
	}
	return data
}
func (this *Credencial) Del(name string) error {
	var data Credencial
	db := libs.DB
	if err := db.Where("name = ?", name).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Credencial) DelForId(id int64) error {
	var data Credencial
	db := libs.DB


	if err := db.Where("id = ?", id).Unscoped().Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
func (this *Credencial) Update(name string,account string, password string) error {
	var varCredencial Credencial
	db := libs.DB
	varCredencial.Name=name
	
	if db.Where("name = ?", varCredencial.Name).Find(&varCredencial).RecordNotFound() {
		return errors.New("no existe la credencial " + name)
	}
	varCredencial.Account=account
	varCredencial.Password=password
	if err := db.Save(&varCredencial).Error; err != nil {
		return err
	}
	return nil
}