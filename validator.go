package gocfgvalidtor

import (
	"errors"
	"fmt"
	"reflect"
)

var errCfgToDeep = errors.New("cfg is too deep")

type optionFunc func(c *Component)

//Component need for configure validation
type Component struct {
	deepOfRecursion        int
	strictMode             bool
	validatorInterfaceType reflect.Type
}

//New init with option configs
func New(options ...optionFunc) *Component {
	c := &Component{
		deepOfRecursion:        7,
		strictMode:             true,
		validatorInterfaceType: reflect.TypeOf((*Validator)(nil)).Elem(),
	}
	for _, op := range options {
		op(c)
	}
	return c
}

//RecursiveValidate is entrypoint for validation
func (c *Component) RecursiveValidate(v Validator) error {
	return c.recursiveValidate(v, c.deepOfRecursion)
}

func (c *Component) recursiveValidate(v Validator, deepOfRecursion int) error {
	if deepOfRecursion <= 0 {
		return errCfgToDeep
	}
	if reflect.TypeOf(v).Kind() == reflect.Struct {
		val := reflect.ValueOf(v)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Type().Implements(c.validatorInterfaceType) {
				if err := c.recursiveValidate(val.Field(i).Interface().(Validator), deepOfRecursion-1); err != nil {
					return err
				}
			} else if field.Type().Kind() == reflect.Struct && c.strictMode {
				return fmt.Errorf("value %#v must implement Validator interface in strictMode", field)
			}
		}
	}
	return v.Validate()
}
