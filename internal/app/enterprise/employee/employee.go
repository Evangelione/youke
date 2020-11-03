package employee

import (
	"yk/internal/app/enterprise/merchant"
	"yk/internal/app/enterprise/store"
	"yk/internal/pkg/infra"

	"github.com/fatih/structs"
)

func FindAndCountAll(merID, storeID, page, size int, params WhereParams) (*infra.CountRows, error) {
	if _, err := store.FindOne(merID, storeID); err != nil {
		return nil, err
	}

	var es []Employee

	// where条件创建
	var m map[string]interface{}

	// 判断params是否为空
	st := structs.New(params)
	if !st.IsZero() {
		// params不为空则转为 map[string]interface{}
		m = structs.Map(params)
	}

	db := infra.Mysql
	var count int64
	if m == nil {
		countResult := db.Model(&es).Where("status <> ? AND mer_id = ? AND store_id = ?", 0, merID, storeID).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Find(&es)
	} else {
		countResult := db.Model(&es).Where(m).Where("status <> ? AND mer_id = ? AND store_id = ?", 0, merID, storeID).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Find(&es)
	}

	if db.Error != nil {
		return nil, db.Error
	}

	return &infra.CountRows{
		Count: count,
		Rows:  es,
	}, nil
}

func FindOne(merID, storeID, id int) (*Employee, error) {
	if _, err := store.FindOne(merID, storeID); err != nil {
		return nil, err
	}

	e := Employee{
		DB:      infra.Mysql,
		ID:      id,
		MerID:   merID,
		StoreID: storeID,
	}

	if err := e.DB.Where("status <> ?", 0).First(&e).Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func Create(merID, storeID int, params CreateParams) (*Employee, error) {
	if _, err := store.FindOne(merID, storeID); err != nil {
		return nil, err
	}

	body := params.Body
	e := Employee{
		DB:      infra.Mysql,
		MerID:   merID,
		StoreID: storeID,
		Name:    body.Name,
		Tel:     body.Tel,
		Email:   body.Email,
		Status:  1,
	}

	if err := e.DB.Create(&e).Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func Update(merID, storeID, id int, params UpdateParams) (*Employee, error) {
	if _, err := store.FindOne(merID, storeID); err != nil {
		return nil, err
	}

	body := params.Body
	e := Employee{
		DB:      infra.Mysql,
		ID:      id,
		MerID:   merID,
		StoreID: storeID,
	}
	st := structs.New(body)
	if st.IsZero() {
		return nil, nil
	}

	if err := e.DB.Model(&e).Where("status <> ?", 0).First(&e).Error; err != nil {
		return nil, err
	}

	if err := e.DB.Model(&e).Updates(st.Map()).Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func Delete(merID, storeID, id int) error {
	if _, err := store.FindOne(merID, storeID); err != nil {
		return err
	}

	e := Employee{
		DB:      infra.Mysql,
		ID:      id,
		MerID:   merID,
		StoreID: storeID,
	}

	if err := e.DB.Model(&e).Where("status <> ?", 0).First(&e).Error; err != nil {
		return err
	}

	if err := e.DB.Model(&e).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}

func FindAndCountAllOfMerchant(id, page, size int, params WhereParamsOfMerchant) (*infra.CountRows, error) {
	if _, err := merchant.FindOne(id); err != nil {
		return nil, err
	}

	var es []Employee

	// where条件创建
	var m map[string]interface{}

	// 判断params是否为空
	st := structs.New(params)
	if !st.IsZero() {
		// params不为空则转为 map[string]interface{}
		m = structs.Map(params)
	}

	db := infra.Mysql
	var count int64
	if m == nil {
		countResult := db.Model(&es).Where("status <> ? AND mer_id = ?", 0, id).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Find(&es)
	} else {
		countResult := db.Model(&es).Where(m).Where("status <> ? AND mer_id", 0, id).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Find(&es)
	}

	if db.Error != nil {
		return nil, db.Error
	}

	return &infra.CountRows{
		Count: count,
		Rows:  es,
	}, nil
}
