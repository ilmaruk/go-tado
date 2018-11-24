package v1

import (
	"errors"
	"testing"
)

func TestGetMeFailure(t *testing.T) {
	old := requestRunner
	defer func() {
		requestRunner = old
	}()

	requestRunner = func(method, path string, data interface{}) error {
		return errors.New("boom")
	}

	_, err := GetMe()
	if err == nil {
		t.Error("no error returned")
	}
	if err.Error() != "boom" {
		t.Errorf("incorrect error; expected 'boom' received '%s'", err.Error())
	}
}

func TestGetMeSuccess(t *testing.T) {
	old := requestRunner
	defer func() {
		requestRunner = old
	}()

	meIn := Me{
		Email: "ruggero.ferretti@gmail.com",
	}

	requestRunner = func(method, path string, data interface{}) error {
		data = meIn
		return nil
	}

	_, err := GetMe()
	if err != nil {
		t.Error("unexpected error")
	}
	//if meOut != meIn {
	//	t.Errorf("invalid data; expected %v received %v", meIn, meOut)
	//}
}
