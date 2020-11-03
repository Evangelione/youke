package store

import (
	"encoding/json"
	"yk/internal/app/enterprise/merchant"
	"yk/internal/pkg/infra"
	"yk/internal/pkg/lib"

	"gorm.io/datatypes"

	"github.com/fatih/structs"
)

func FindAndCountAll(pid, page, size int, params WhereParams) (*infra.CountRows, error) {
	if _, err := merchant.FindOne(pid); err != nil {
		return nil, err
	}

	var ss []Store

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
		countResult := db.Model(&ss).Where("status <> ? AND mer_id = ?", 0, pid).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Order("sort desc").Find(&ss)
	} else {
		countResult := db.Model(&ss).Where(m).Where("status <> ? AND mer_id = ?", 0, pid).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Order("sort desc").Find(&ss)
	}

	if db.Error != nil {
		return nil, db.Error
	}

	return &infra.CountRows{
		Count: count,
		Rows:  ss,
	}, nil
}

func FindOne(pid, id int) (*Store, error) {
	if _, err := merchant.FindOne(pid); err != nil {
		return nil, err
	}

	s := Store{
		DB:    infra.Mysql,
		ID:    id,
		MerID: pid,
	}

	if err := s.DB.Where("status <> ?", 0).First(&s).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func Create(pid int, params CreateParams) (*Store, error) {
	if _, err := merchant.FindOne(pid); err != nil {
		return nil, err
	}

	body := params.Body
	marshal, err := json.Marshal(body.Images)
	if err != nil {
		return nil, err
	}

	long, lat, err := lib.GetFloat64LongAndLat(body.Long, body.Lat)
	if err != nil {
		return nil, err
	}

	s := Store{
		DB:           infra.Mysql,
		MerID:        pid,
		Name:         body.Name,
		Introduction: body.Introduction,
		OpeningHours: body.OpeningHours,
		Logo:         body.Logo,
		Images:       datatypes.JSON(marshal),
		Contact:      body.Contact,
		Tel:          body.Tel,
		Email:        body.Email,
		CategoryID:   body.CategoryID,
		Districts: infra.Districts{
			ProvinceID: body.ProvinceID,
			CityID:     body.CityID,
			AreaID:     body.AreaID,
		},
		CircleID: body.CircleID,
		MarketID: body.MarketID,
		Address:  body.Address,
		Long:     long,
		Lat:      lat,
		Status:   1,
	}

	if err := s.DB.Create(&s).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func Update(pid, id int, params UpdateParams) (*Store, error) {
	if _, err := merchant.FindOne(pid); err != nil {
		return nil, err
	}

	body := params.Body
	s := Store{
		DB:    infra.Mysql,
		ID:    id,
		MerID: pid,
	}
	st := structs.New(body)
	if st.IsZero() {
		return nil, nil
	}

	if err := s.DB.Model(&s).Where("status <> ?", 0).First(&s).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&s).Updates(st.Map()).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func Delete(pid, id int) error {
	if _, err := merchant.FindOne(pid); err != nil {
		return err
	}

	s := Store{
		DB:    infra.Mysql,
		ID:    id,
		MerID: pid,
	}

	if err := s.DB.Model(&s).Where("status <> ?", 0).First(&s).Error; err != nil {
		return err
	}

	if err := s.DB.Model(&s).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
