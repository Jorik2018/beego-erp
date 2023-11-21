package requestStruct

type CarsInsert struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCars struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CarsId      int    `json:"car_id"`
}

type DeleteCar struct {
	CarsId int `json:"car_id"`
}
