package infra

import "errors"

func main() {
	errs := make([]error, 0, 2+2)
	errs = append(errs, errors.New("error 1"))
	e := errors.Join(errs...)

	println(e.Error())
}
