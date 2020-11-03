package merchant

import (
	"errors"
	"yk/internal/pkg/infra"
	"yk/internal/pkg/lib"

	"github.com/fatih/structs"
)

func FindAll(page, size int, params WhereParams) (*infra.CountRows, error) {
	var ms []Merchant

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
		countResult := db.Model(&ms).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Find(&ms)
	} else {
		countResult := db.Model(&ms).Where(m).Count(&count)
		db = countResult.Limit(size).Offset((page - 1) * size).Find(&ms)
	}

	if db.Error != nil {
		return nil, db.Error
	}

	return &infra.CountRows{
		Count: count,
		Rows:  ms,
	}, nil
}

func FindOne(id int) (*Merchant, error) {
	m := Merchant{
		DB: infra.Mysql,
		ID: id,
	}

	if err := m.DB.First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func Create(params CreateParams) (*Merchant, error) {
	body := params.Body

	// 暂时不用图片列表
	//marshal, err := json.Marshal(body.Images)
	//if err != nil {
	//	return nil, err
	//}

	long, lat, err := lib.GetFloat64LongAndLat(body.Long, body.Lat)
	if err != nil {
		return nil, err
	}

	m := Merchant{
		DB:           infra.Mysql,
		Name:         body.Name,
		Introduction: body.Introduction,
		Logo:         body.Logo,
		Images:       body.Images,
		//Contact:      body.Contact,
		Tel:        body.Tel,
		Email:      body.Email,
		CategoryID: body.CategoryID,
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
		Account:  body.Account,
		Password: &body.Password,
		Status:   body.Status,
	}

	if err := m.DB.Create(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func Update(id int, params UpdateParams) (*Merchant, error) {
	body := params.Body
	m := Merchant{
		DB: infra.Mysql,
		ID: id,
	}
	st := structs.New(body)
	if st.IsZero() {
		return nil, nil
	}

	if err := m.DB.Model(&m).First(&m).Error; err != nil {
		return nil, err
	}

	if err := m.DB.Model(&m).Updates(st.Map()).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func Delete(id int) error {
	m := Merchant{
		DB: infra.Mysql,
		ID: id,
	}

	if err := m.DB.Model(&m).First(&m).Error; err != nil {
		return err
	}

	//暂不支持删除
	//if err := m.DB.Model(&m).Update("status", 0).Error; err != nil {
	//	return err
	//}

	return errors.New("暂不支持删除商户")
}
