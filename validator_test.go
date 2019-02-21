package gocfgvalidtor

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type CfgMockScalarWithValidator string

func (m CfgMockScalarWithValidator) Validate() error {
	mStr := string(m)
	if strings.HasPrefix(mStr, "error:") {
		return errors.New(strings.TrimPrefix(mStr, "error:"))
	}
	return nil
}

type CfgMockValidatorContainsValidator struct {
	Validator Validator
	err       error
}

func (m CfgMockValidatorContainsValidator) Validate() error {
	return m.err
}

type CfgMockValidatorContainsScalar struct {
	Scalar string
	err    error
}

func (m CfgMockValidatorContainsScalar) Validate() error {
	return m.err
}

type CfgMockWrapper struct {
	err error
}

func (m CfgMockWrapper) Validate() error {
	return m.err
}

type CfgMocNotimplementedValidator struct {
}

func TestRecursiveValidate_StrictModeCorrectCfg_WithoutError(t *testing.T) {
	testStruct := struct {
		CfgMockWrapper
		First  CfgMockValidatorContainsScalar
		Second CfgMockValidatorContainsValidator
		Third  CfgMockScalarWithValidator
		Scalar string
	}{
		Second: CfgMockValidatorContainsValidator{
			Validator: CfgMockValidatorContainsValidator{
				Validator: CfgMockScalarWithValidator("sdfsf"),
			},
		},
	}
	component := New()
	if err := component.RecursiveValidate(testStruct); err != nil {
		fmt.Printf("Unexpected err: '%s'\n", err.Error())
		t.FailNow()
	}
}

func TestRecursiveValidate_StrictModeTooDeepRecursion_Error(t *testing.T) {
	testStruct := struct {
		CfgMockWrapper
		First CfgMockValidatorContainsValidator
	}{
		First: CfgMockValidatorContainsValidator{
			Validator: CfgMockValidatorContainsValidator{
				Validator: CfgMockValidatorContainsValidator{
					Validator: CfgMockScalarWithValidator("sdfsf"),
				},
			},
		},
	}
	component := New(MustWithDeepOfRecursion(4))
	err := component.RecursiveValidate(testStruct)
	if err == nil || err != errCfgToDeep {
		fmt.Printf("Expeceted err: '%#v'\nActual err: %#v\n", errCfgToDeep, err)
		t.FailNow()
	}
}

func TestRecursiveValidate_StrictModeStructNotImplementedValidator_Error(t *testing.T) {
	testStruct := struct {
		First CfgMocNotimplementedValidator
		CfgMockWrapper
	}{
		First: CfgMocNotimplementedValidator{},
	}
	component := New()
	err := component.RecursiveValidate(testStruct)
	if err == nil || !strings.HasSuffix(err.Error(), "must implement Validator interface in strictMode") {
		fmt.Printf(
			`In strict mode expected error when we try Validate struct not iplemented Validator.
Actual err: '%#v'
`,
			err,
		)
		t.FailNow()
	}
}

func TestRecursiveValidate_SoftModeStructNotImplementedValidator_NotError(t *testing.T) {
	testStruct := struct {
		First CfgMocNotimplementedValidator
		CfgMockWrapper
	}{
		First: CfgMocNotimplementedValidator{},
	}
	component := New(WithStrictMode(false))
	if err := component.RecursiveValidate(testStruct); err != nil {
		fmt.Printf("Unexpected err: %#v", err)
		t.FailNow()
	}
}
