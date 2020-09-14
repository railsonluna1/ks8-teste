package main

import "testing"

func TestGreeting(t *testing.T) {
	var message = "Code.education Rocks!"

	var messageResponse = greeting(message)

	if "<b>Code.education Rocks!</b>" != messageResponse {
		t.Errorf("greeting did not work, %s", "<b>"+message+"</b>")
	}
}
