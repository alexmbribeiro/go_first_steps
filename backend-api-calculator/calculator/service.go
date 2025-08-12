package calculator

import "errors"

func Add(req TwoNumbersRequest) ResultResponse {
	return ResultResponse{Result: req.Number1 + req.Number2}
}

func Subtract(req TwoNumbersRequest) ResultResponse {
	return ResultResponse{Result: req.Number1 - req.Number2}
}

func Multiply(req TwoNumbersRequest) ResultResponse {
	return ResultResponse{Result: req.Number1 * req.Number2}
}

func Divide(req DivideRequest) (ResultResponse, error) {
	if req.Divisor == 0 {
		return ResultResponse{}, errors.New("division by zero")
	}
	return ResultResponse{Result: req.Dividend / req.Divisor}, nil
}

func Sum(numbers []int) ResultResponse {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return ResultResponse{Result: sum}
}
