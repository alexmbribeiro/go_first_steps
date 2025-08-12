package calculator

type TwoNumbersRequest struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type DivideRequest struct {
	Dividend int `json:"dividend"`
	Divisor  int `json:"divisor"`
}

type ResultResponse struct {
	Result int `json:"result"`
}
