package controllers

import (
	"beego-erp/helpers"
	"beego-erp/models"
	requestStruct "beego-erp/requstStruct"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type CarsControllers struct {
	beego.Controller
}

func (u *CarsControllers) RegisterCar() {
	var cars requestStruct.CarsInsert
	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)
	result, _ := models.RegisterCar(cars)
	if result != nil {
		helpers.ApiSuccessResponse(&u.Controller, "", "Car Register Successfully")
	}
	helpers.ApiFailedResponse(&u.Controller, "Please Try Again")

}

func (u *CarsControllers) UpdateCar() {
	var cars requestStruct.UpdateCars
	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)
	result, _ := models.UpdateCars(cars)
	if result != nil {
		helpers.ApiSuccessResponse(&u.Controller, "", "Car Updated Successfully")
	}
	helpers.ApiFailedResponse(&u.Controller, "Please Try Again")
}

func (u *CarsControllers) DeleteCar() {
	var cars requestStruct.DeleteCar
	json.Unmarshal(u.Ctx.Input.RequestBody, &cars)
	result := models.DeleteCar(cars)
	if result != 0 {
		helpers.ApiSuccessResponse(&u.Controller, "", "Car Deleted Successfully")
	}
	helpers.ApiFailedResponse(&u.Controller, "Please Try Again")
}

func (u *CarsControllers) FetchCar() {

	// json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	result, _ := models.FetchCars()
	if result != nil {
		helpers.ApiSuccessResponse(&u.Controller, result, "Cars Found Successfully")
	}
	helpers.ApiFailedResponse(&u.Controller, "Something Wrong Please Try Again")
}
