package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/culturadevops/GORM/libs"
	"github.com/jinzhu/gorm"
)

var VarSecret *Secret

type Items struct {
	Type  string            `json:"type"`
	Items map[string]string `json:"item"`
}

func (this *Secret) CreateItem(secret string, types string) Items {
	i := make(map[string]string)
	i["secret"] = secret
	return Items{Type: types, Items: i}

}
func (this *Secret) itemCommand(secret string) Items {
	return this.CreateItem(secret, "command")
}
func (this *Secret) itemFile(secret string) Items {
	return this.CreateItem(secret, "file")
}

func (this *Secret) itemToken(secret string) Items {
	return this.CreateItem(secret, "token")
}

func (this *Secret) itemCredencial(user string, pass string) Items {
	res := this.CreateItem(pass, "credential")
	res.Items["user"] = user
	return res
}
func (this *Secret) itemAmazonCredencial(accountid string, user string, pass string) Items {
	res := this.itemCredencial(user, pass)
	res.Items["account"] = accountid
	res.Type = "amazon"

	return res
}
func (this *Secret) jsoncode(item Items) string {
	bs1, _ := json.Marshal(item)
	return string(bs1)
}

type Secret struct {
	gorm.Model
	Owner  string `gorm:"type:varchar(50);not null;"`
	Name   string `gorm:"type:varchar(100);not null;"`
	Secret string `gorm:"type:blob;not null;"`
}

func (t *Secret) AdditemFile(owner string, name string, token string) error {
	text := t.jsoncode(t.itemFile(token))
	return t.Add(owner, name, text)
}
func (t *Secret) AdditemCommand(owner string, name string, token string) error {
	text := t.jsoncode(t.itemCommand(token))
	return t.Add(owner, name, text)
}
func (t *Secret) Additem(owner string, name string, token string) error {
	text := t.jsoncode(t.itemToken(token))
	return t.Add(owner, name, text)
}
func (t *Secret) Additem1(owner string, name string, user string, pass string) error {
	text := t.jsoncode(t.itemCredencial(user, pass))
	return t.Add(owner, name, text)
}

func (t *Secret) Additem2(owner string, name string, account string, user string, pass string) error {
	text := t.jsoncode(t.itemAmazonCredencial(account, user, pass))
	return t.Add(owner, name, text)
}

func (this *Secret) AddNormal(owner string, name string, text string) error {
	this.Owner = owner
	this.Name = name
	this.Secret = text

	if !libs.DB.Where("owner = ? AND  name = ? ", owner, name).First(&Secret{}).RecordNotFound() {
		return errors.New("Ya existe un item con el nombre " + name)
	}

	libs.DB.Create(&this)

	if err := libs.DB.Create(&this).Error; err != nil {
		return err
	}
	return nil
}
func (this *Secret) Add(owner string, name string, text string) error {
	result := fmt.Sprintf("INSERT INTO `cindi`.`secret`(`owner`,`name`,`secret`)VALUES('%s','%s',AES_ENCRYPT('%s','password'));", owner, name, text)

	if err := libs.DB.Exec(result).Error; err != nil {
		fmt.Println("error en scan Add")
		fmt.Println(err)
		return err
	}

	return nil
}

func (this *Secret) GetForId(owner string, id int64) (Secret, error) {
	var data = Secret{}

	result := fmt.Sprintf("select `owner`,`name`,CAST(AES_DECRYPT(secret,'password')AS CHAR) AS `secret` FROM  `cindi`.`secret` WHERE owner = '%s' AND  id = '%v';", owner, id)
	if err := libs.DB.Raw(result).Scan(&data).Error; err != nil {
		fmt.Println(err)
		return Secret{}, errors.New("no se encontro las credenciales ")
	}

	return data, nil
}
func (this *Secret) Get(owner string, name string) (Secret, error) {
	var data = Secret{}

	result := fmt.Sprintf("select `owner`,`name`,CAST(AES_DECRYPT(secret,'password')AS CHAR) AS `secret` FROM  `cindi`.`secret` WHERE owner = '%s' AND  name = '%v';", owner, name)
	if err := libs.DB.Raw(result).Scan(&data).Error; err != nil {
		fmt.Println(err)
		return Secret{}, errors.New("no se encontro las credenciales ")
	}

	return data, nil
}
func (this *Secret) List(owner string) []Secret {
	var data = []Secret{}
	result := fmt.Sprintf("select `id`, `owner`,`name`,CAST(AES_DECRYPT(secret,'password')AS CHAR) AS `secret` FROM  `cindi`.`secret` WHERE owner = '%s';", owner)
	if err := libs.DB.Raw(result).Scan(&data).Error; err != nil {
		fmt.Println(err)
		return []Secret{}
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
	if libs.DB.Where("owner = ? AND  name = ? ", owner, name).Find(&varSecret).RecordNotFound() {
		return errors.New("no existe la credencial " + name)
	}
	varSecret.Name = name
	varSecret.Secret = secret

	if err := libs.DB.Save(&varSecret).Error; err != nil {
		return err
	}
	return nil
}
