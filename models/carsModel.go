package models

import (
	requestStruct "crudDemo/requstStruct"
	"time"

	"github.com/beego/beego/orm"
)

func RegisterCar(c requestStruct.CarsInsert) (interface{}, error) {
	db := orm.NewOrm()
	res := CarsMasterTable{
		Name:        c.Name,
		CreatedBy:   1,
		UpdatedBy:   0,
		UpdatedDate: time.Now(),
		Description: c.Description,
		CreatedDate: time.Now(),
	}

	_, err := db.Insert(&res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func UpdateCars(c requestStruct.UpdateCars) (interface{}, error) {
	db := orm.NewOrm()

	cars := CarsMasterTable{CarsId: c.CarsId}
	if db.Read(&cars) == nil {
		cars.Name = c.Name
		cars.Description = c.Description
		if num, err := db.Update(&cars); err == nil {
			return num, nil
		}

	}
	return nil, orm.ErrArgs

}

func DeleteCar(c requestStruct.DeleteCar) int {
	db := orm.NewOrm()
	cars := CarsMasterTable{CarsId: c.CarsId}
	if _, err := db.Delete(&cars); err == nil {
		return 1
	}
	return 0

}
func FetchCars() (interface{}, error) {
	db := orm.NewOrm()
	var cars []struct {
		Name        string    `json:"car_name"`
		Description string    `json:"description"`
		CreatedDate time.Time `json:"created_date"`
	}
	_, err := db.Raw("SELECT name,description , created_date FROM cars_master_table").QueryRows(&cars)

	if err != nil {
		return nil, err
	}

	if len(cars) == 0 {
		return "Not Found Cars", nil
	}

	return cars, nil
}
