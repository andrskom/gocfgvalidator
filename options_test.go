package gocfgvalidtor

import (
	"fmt"
	"testing"
)

func TestWithStrictMode(t *testing.T) {
	{
		c := New()
		if !c.strictMode {
			fmt.Println("expected that strict mode enabled by default")
			t.FailNow()
		}
	}
	{
		c := New(WithStrictMode(false))
		if c.strictMode {
			fmt.Println("expected that strict mode disabled")
			t.FailNow()
		}
	}
	{
		c := New(WithStrictMode(true))
		if !c.strictMode {
			fmt.Println("expected that strict mode enabled")
			t.FailNow()
		}
	}
}

func TestMustWithDeepOfRecursion(t *testing.T) {
	{
		c := New()
		if c.deepOfRecursion != 7 {
			fmt.Println("expected that deep of recursion is 7 by default")
			t.FailNow()
		}
	}
	{
		c := New(MustWithDeepOfRecursion(10))
		if c.deepOfRecursion != 10 {
			fmt.Println("expected that deep of recursion is 10")
			t.FailNow()
		}
	}
	{
		c := New(MustWithDeepOfRecursion(1))
		if c.deepOfRecursion != 1 {
			fmt.Println("expected that deep of recursion is 1")
			t.FailNow()
		}
	}
	{
		func() {
			defer func() {
				if rec := recover(); rec == nil {
					fmt.Println("expected panic for not valid arg")
					t.FailNow()
				}
			}()
			MustWithDeepOfRecursion(0)
		}()
	}
}
