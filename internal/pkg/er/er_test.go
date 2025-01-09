package er

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"testing"
)

var _test_unknown *BuiltInError
var myErr error

type _test_err_type struct {
	s string
}

func (t *_test_err_type) Error() string {
	return t.s
}

// 每个测试前执行
func setup() {
	_test_unknown = NewErr(2, "_test_unknown", "_test_unknown", Error, true) // 未知错误（服务端错误）
	myErr = WSEF(_test_unknown, zap.String("test", "test"))
}

// 每个测试后执行
func teardown() {
	fmt.Println("After all tests")
}

// 用 TestMain 整合 setup 和 teardown
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	//teardown()
	os.Exit(code)
}

func TestIs(t *testing.T) {
	if !errors.Is(myErr, _test_unknown) {
		t.Fatal("errors.Is False")
	}
}

func TestIsBaseErr(t *testing.T) {
	err := &_test_err_type{
		"_test_err_type",
	}
	mer := WSEF(_test_unknown)
	mer.WithErr(err)
	if !errors.Is(mer, err) {
		t.Fatal("errors.Is _test_err_type False")
	}
}

func TestAsBaseErr(t *testing.T) {
	err := &_test_err_type{
		"_test_err_type",
	}
	mer := WSEF(_test_unknown)
	mer.WithErr(err)
	teErr := new(_test_err_type)
	if !errors.As(mer, &teErr) {
		t.Fatal("errors.As ShellError False")
	}
}

func TestAsMyErr(t *testing.T) {
	myErr := new(ShellError)
	if !errors.As(myErr, &myErr) {
		t.Fatal("errors.As ShellError False")
	}
}

func TestAsBuiltInError(t *testing.T) {
	e := new(BuiltInError)
	if !errors.As(myErr, &e) {
		t.Fatal("errors.Is BuiltInError False")
	}
	if e.Code() != _test_unknown.Code() {
		t.Fatal("BuiltInError False")
	}
}
