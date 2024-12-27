package main

import (
	"errors"
	"fmt"
)

func main() {
	testDeferFunc(true)
}

func testDeferFunc(isVcc bool) {
	var err error

	defer func() {
		if err != nil {
			fmt.Println("err=", err.Error())
		}
	}()

	num, err := err1()
	if err != nil {
		fmt.Println("test err1", err, "num==", num)
	}

	if isVcc {
		str, err := err2()
		if err != nil {
			fmt.Println("test err2", err, "str==", str)
		}

		if err = err3(); err == nil {
			fmt.Println("test err 33")
		}

	}
	return

}

func err1() (int, error) {
	return 0, nil
}

func err2() (string, error) {
	return "", nil
}

func err3() error {
	return errors.New("err333")
}
