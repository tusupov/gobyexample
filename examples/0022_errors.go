package main

import (
	"errors"
	"fmt"
)

func f1(i int) (int, error)  {
	if i == 45 {
		return -1, errors.New("can't work with 45")
	}
	return i + 3, nil
}

type argError struct {
	arg int
	prob string
}

func (e argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(i int) (int, error)  {
	if i == 45 {
		return -1, argError{i, "can't work with 45"}
	}
	return i + 3, nil
}

func main() {

	for _, i := range []int{7, 45} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 45} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(45)
	if ae, ok := e.(argError); ok {
		fmt.Println(ok)
		fmt.Println(ae.arg)
	}

}
