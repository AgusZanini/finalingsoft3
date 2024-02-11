package dto

type OperationDto struct {
	Id      int
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
	Result  float64 `json:"result"`
}
