package customerimporter

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestEmailString(t *testing.T) {
	e := email("test@example.com")
	expected := "test@example.com"

	if e.String() != expected {
		t.Errorf("Expected: %s, Got: %s", expected, e.String())
	}
}

func TestEmailIsValid(t *testing.T) {
	e := email("test@example.com")

	if !e.isValid() {
		t.Error("The email is supposed to be valid")
	}
}

func TestEmailIsInvalid(t *testing.T) {
	e := email("invalid.email")

	if e.isValid() {
		t.Error("The email is supposed to be invalid")
	}
}

func TestEmailDomainWithValidEmail(t *testing.T) {
	e := email("test@example.com")
	expected := "example.com"

	if e.domain() != expected {
		t.Errorf("Expected domain: %s, Got: %s", expected, e.domain())
	}
}

func TestEmailDomainWithInvalidEmail(t *testing.T) {
	e := email("invalid.email")
	expected := "*invalid"

	if e.domain() != expected {
		t.Errorf("Expected domain: %s, Got: %s", expected, e.domain())
	}
}

func TestFromCsv(t *testing.T) {
	content := []byte(`first_name,last_name,email,gender,ip_address
	Mildred,Hernandez,test@example.com,Female,38.194.51.128`)
	tmpfile, err := os.CreateTemp("", "example.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	FromCsv(tmpfile.Name())

	w.Close()
	os.Stdout = oldStdout

	var output strings.Builder
	io.Copy(&output, r)

	expectedOutput := "example.com: 1\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}
