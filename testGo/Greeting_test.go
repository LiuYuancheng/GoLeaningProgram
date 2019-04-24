package testGo

import "testing"

func TestGreeting(t *testing.T){
	got := Greeting("George")
	want := "Hello George"
	if got != want{
		t.Fatalf("Expected %q, got %q", want , got)
	}
}

func TestTranslate(t *testing.T){
	got := translate("fr-Fr4")
	want := "Bonjour"
	if got != want{
		t.Fatalf("Expect %q, got %q", want, got)
	}
}