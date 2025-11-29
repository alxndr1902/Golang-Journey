package main

import "fmt"

func main() {
	err := SaveDate("1", "a")
	if err != nil {
		//	if validationErr, ok := err.(*validationError); ok {
		//		fmt.Println("Validation Error:", validationErr.Error())
		//	} else if notFoundErr, ok := err.(*notFoundError); ok {
		//		fmt.Println("NotFound Error:", notFoundErr.Error())
		//	} else {
		//		fmt.Println("Unknown Error:", err.Error())
		//	}
		//} else {
		//	fmt.Println("OK")
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found Error:", finalError.Error())
		default:
			fmt.Println("Unknown Error:", finalError.Error())
		}
	}
}

func SaveDate(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}
	if data != "Alex" {
		return &notFoundError{"Data not found"}
	}
	return nil
}

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}
