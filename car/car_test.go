package car

import (
	"testing"
)

func TestNewFail(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		t.Fatal("got errors: ", err)
	}

	if c == nil {
		t.Error("car should be nil")
	}
}

func TestNewPass(t *testing.T) {
	c, err := New("Shanni", 200)
	if err != nil {
		t.Fatal("got errors: ", err)
	}

	if c == nil {
		t.Error("car should not be nil")
	}

	t.Log("car is pass:", c)
}

/*
可用 go test -v ./car
-v 代表會顯示測試的詳細過程(不然只會有PASS、ok的那兩行)

go test -v -run {testFuncName} {filePath}
go test -v -run TestNew ./car
代表指定這檔案底下的TestNew這個Function

t.Fatal出現錯誤後會終止
*/
