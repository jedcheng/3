package script

import (
	"fmt"
)

type compileErr string

func (c *compileErr) Error() string {
	return string(*c)
}

func newCompileErr(msg ...interface{}) *compileErr {
	e := compileErr(fmt.Sprint(msg...))
	return &e
}