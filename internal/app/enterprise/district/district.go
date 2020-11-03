package district

import (
	"yk/internal/pkg/infra"

	"github.com/fatih/structs"
)

func FindAll(params WhereParams) (*[]District, error) {
	var ds []District

	// 判断params是否为空
	st := structs.New(params)

	db := infra.Mysql

	if !st.IsZero() {
		db = db.Model(&ds).Where("area_type = ?", params.Level).Find(&ds)
	} else {
		db = db.Model(&ds).Where("area_pid = ?", 0).Find(&ds)
	}

	if err := db.Error; err != nil {
		return nil, db.Error
	}

	return &ds, nil
}

func FindOne(id int) (*[]District, error) {
	var ds []District

	db := infra.Mysql
	if err := db.Model(&ds).Where("area_pid = ?", id).Find(&ds).Error; err != nil {
		return nil, db.Error
	}

	return &ds, nil
}

func Create() {

}

func Update() {

}

func Delete() {

}
