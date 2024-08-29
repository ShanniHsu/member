package simpleFactory

import "testing"

func TestType1(t *testing.T) {
	// 這個對簡單工廠模式來說是客戶端(Client) shanni
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	// 這個對簡單工廠模式來說是客戶端(Client) shanni
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Type2 test fail")
	}
}
