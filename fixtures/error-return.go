// Test of error-return.

package fixtures

func f1() {
}

func f2() error {
	return nil
}

func f3() int {
	return 0
}

func f4() (int, error) {
	return 0, nil
}

func f5() (x int, err error) {
	return 0, nil
}

func f6() (error, int) { // MATCH /error should be the last type when returning multiple items/
	return nil, 0
}

func f7() (int, error, int) { // MATCH /error should be the last type when returning multiple items/
	return 0, nil, 0
}

func f8() (x int, err error, y int) { // MATCH /error should be the last type when returning multiple items/
	return 0, nil, 0
}

func f9() (int, error, error) {
	return 0, nil, nil
}

func f10() (int, error, int, error) {
	return 0, nil, 0, nil
}
