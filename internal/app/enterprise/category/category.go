package category

import (
	"yk/internal/pkg/infra"
)

func FindAll() (*[]Category, error) {
	var cs []Category

	db := infra.Mysql
	if err := db.Model(&cs).Where("cat_fid = ?", 0).Find(&cs).Error; err != nil {
		return nil, db.Error
	}

	return &cs, nil
}

func FindOne(id int) (*[]Category, error) {
	var cs []Category

	db := infra.Mysql
	if err := db.Model(&cs).Where("cat_fid = ?", id).Find(&cs).Error; err != nil {
		return nil, db.Error
	}

	return &cs, nil
}

func Create() {

}

func Update() {

}

func Delete() {

}
