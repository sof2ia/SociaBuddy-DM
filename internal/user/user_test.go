package user

import (
	"testing"
)

type argsStr struct {
	input    string
	hasError bool
}

type argsInt struct {
	input    int
	hasError bool
}

func TestValidationName(t *testing.T) {
	tests := []argsStr{
		{
			input:    "User First",
			hasError: false,
		},
		{
			input:    "User2 Second",
			hasError: true,
		},
		{
			input:    "User",
			hasError: true,
		},
		{
			input:    "User 4",
			hasError: true,
		},
		{
			input:    "Fifth User Test",
			hasError: false,
		},
		{
			input:    "User second Test",
			hasError: true,
		},
		{
			input:    "",
			hasError: true,
		},
	}
	var i int
	for i = 0; i < len(tests); i++ {
		actualError := nameValidation(tests[i].input)
		if (actualError == nil && tests[i].hasError == true) || (actualError != nil && tests[i].hasError == false) {
			t.Fatalf("the test is failed %d: %v", i, actualError)
		}
		t.Logf("the test is passed %d", i)
	}
}

func TestValidationAge(t *testing.T) {
	tests := []argsInt{
		{
			input:    25,
			hasError: false,
		},
		{
			input:    16,
			hasError: true,
		},
		{
			input:    200,
			hasError: true,
		},
	}
	var i int
	for i = 0; i < len(tests); i++ {
		actualError := ageValidation(tests[i].input)
		if (actualError == nil && tests[i].hasError == true) || (actualError != nil && tests[i].hasError == false) {
			t.Fatalf("the test is failed %d: %v", i, actualError)
		}
		t.Logf("the test is passed %d", i)
	}
}

func TestValidationDocument(t *testing.T) {
	tests := []argsStr{
		{
			input:    "777.666.555-44",
			hasError: false,
		},
		{
			input:    "555.888.100",
			hasError: true,
		},
		{
			input:    "1234.123.123-12",
			hasError: true,
		},
		{
			input:    "123.12.123-12",
			hasError: true,
		},
		{
			input:    "123.123.123-1",
			hasError: true,
		},
		{
			input:    "1.1.1-1",
			hasError: true,
		},
		{
			input:    "i22.1oo.222-11",
			hasError: true,
		},
		{
			input:    "",
			hasError: true,
		},
	}
	var i int
	for i = 0; i < len(tests); i++ {
		actualError := documentValidation(tests[i].input)
		if (actualError == nil && tests[i].hasError == true) || (actualError != nil && tests[i].hasError == false) {
			t.Fatalf("the test is failed %d: %v", i, actualError)
		}
		t.Logf("the test is passed %d", i)
	}
}

func TestValidationEmail(t *testing.T) {
	tests := []argsStr{
		{
			input:    "user1@gmail.com",
			hasError: false,
		},
		{
			input:    "user2@",
			hasError: true,
		},
		{

			input:    "@gmail.com",
			hasError: true,
		},
		{
			input:    "user4@gmai",
			hasError: true,
		},
		{
			input:    "user5@gmai.",
			hasError: true,
		},
		{
			input:    "user6@outlook.com",
			hasError: false,
		},
	}
	var i int
	for i = 0; i < len(tests); i++ {
		actualError := emailValidation(tests[i].input)
		if (actualError == nil && tests[i].hasError == true) || (actualError != nil && tests[i].hasError == false) {
			t.Fatalf("the test is failed %d: %v", i, actualError)
		}
		t.Logf("the test is passed %d", i)
	}
}

func TestValidationPhone(t *testing.T) {
	tests := []argsStr{
		{
			input:    "+55 12 94321 1257",
			hasError: false,
		},
		{
			input:    "+55 12 77781 3456",
			hasError: true,
		},
		{
			input:    "+5512 97781 3456",
			hasError: true,
		},
		{
			input:    "+55 11 94321 123",
			hasError: true,
		},
		{
			input:    "+55 11 961234 1234",
			hasError: true,
		},
		{
			input:    "55 11 94321 1234",
			hasError: true,
		},
		{
			input:    "1",
			hasError: true,
		},
	}
	var i int
	for i = 0; i < len(tests); i++ {
		actualError := phoneValidation(tests[i].input)
		if (actualError == nil && tests[i].hasError == true) || (actualError != nil && tests[i].hasError == false) {
			t.Fatalf("the test is failed %d: %v", i, actualError)
		}
		t.Logf("the test is passed %d", i)
	}
}
