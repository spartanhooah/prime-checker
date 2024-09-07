package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		input    int
		expected bool
		msg      string
	}{
		{"negative number", -1, false, "Negative numbers are not prime, by definition"},
		{"zero", 0, false, "0 is not prime, by definition"},
		{"one", 1, false, "1 is not prime, by definition"},
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
	}

	for _, test := range primeTests {
		result, msg := isPrime(test.input)

		if test.expected && !result {
			t.Errorf("%s expected true but got false", test.name)
		}

		if !test.expected && result {
			t.Errorf("%s expected false but got true", test.name)
		}

		if test.msg != msg {
			t.Errorf("%s expected %s but got %s", test.name, test.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of the standard output
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// reassign standard out so we can capture it
	os.Stdout = w

	prompt()

	// close writer
	_ = w.Close()

	// reset standard output
	os.Stdout = oldOut

	// read output of prompt from the read pipe
	out, _ := io.ReadAll(r)

	// make assertions
	if string(out) != "-> " {
		t.Errorf("Incorrect prompt. Expected '-> ' but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("Intro text not correct. Got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Please enter a whole number."},
		{"zero", "0", "0 is not prime, by definition"},
		{"one", "1", "1 is not prime, by definition"},
		{"two", "2", "2 is a prime number"},
		{"three", "3", "3 is a prime number"},
		//{"four", "4", "4 is not a prime number because it is divisible by 2"},
		{"negative", "-1", "Negative numbers are not prime, by definition"},
		{"typed", "three", "Please enter a whole number."},
		{"decimal", "1.1", "Please enter a whole number."},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, test.expected) {
			t.Errorf("%s: expected %s but got %s", test.name, test.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
