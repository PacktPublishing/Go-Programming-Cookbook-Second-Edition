package mockgen

import (
	"errors"
	"testing"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter9/mockgen/internal"
	"github.com/golang/mock/gomock"
)

func TestExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetSetter := internal.NewMockGetSetter(ctrl)

	var k string
	mockGetSetter.EXPECT().Get("we can put anything here!").Do(func(key string) {
		k = key
	}).Return("", nil)

	customError := errors.New("failed this time")

	mockGetSetter.EXPECT().Get(gomock.Any()).Return("", customError)

	if _, err := mockGetSetter.Get("we can put anything here!"); err != nil {
		t.Errorf("got %#v; want %#v", err, nil)
	}
	if k != "we can put anything here!" {
		t.Errorf("bad key")
	}

	if _, err := mockGetSetter.Get("key"); err == nil {
		t.Errorf("got %#v; want %#v", err, customError)
	}
}
