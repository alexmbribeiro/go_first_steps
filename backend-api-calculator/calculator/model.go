package calculator

type TwoNumbersRequest struct {
	Number1 float32 `json:"number1"`
	Number2 float32 `json:"number2"`
}

type DivideRequest struct {
	Dividend float32 `json:"dividend"`
	Divisor  float32 `json:"divisor"`
}

type ResultResponse struct {
	Result float32 `json:"result"`
}
