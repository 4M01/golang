package main

func TestHello(t *testing.T) {
	got := Hello("Chris", "es")
	want := "Hola, Chris"

	if got != want{
		t.Error
	}
}
