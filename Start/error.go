package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func f1(arg int) (int, error) {

	if arg == 42 {
		return -1, errors.New("can not work with 42")
	}
	return arg + 3, nil
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can not work with 42"}
	}

	return arg + 3, nil
}

func errorSample() {
	err1 := argError{arg: 1, prob: "some error"}
	fmt.Println(err1.Error())

	// arg1, err1 := f1(42)
	// fmt.Println(arg1, err1)

	// arg2, err2 := f1(45)
	// fmt.Println(arg2, err2)

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println("ok: ", ok)
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
