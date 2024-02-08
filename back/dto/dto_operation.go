package dto

type OperationDto struct {
	Id      int
	Number1 float32 `json:"number1"`
	Number2 float32 `json:"number2"`
	Result  float32 `json:"result"`
}
