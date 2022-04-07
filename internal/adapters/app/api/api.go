package api

import "hexagon/internal/ports"

type Adapter struct {
	arithmetic ports.Arithmetic
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Addition(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Division(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
