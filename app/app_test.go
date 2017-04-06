package main

import (
	"io/ioutil"
	"testing"

	"net/http/httptest"

	"google.golang.org/appengine/aetest"
)

func TestHello(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	helloWorld(w, r)
	data, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	if "Hello World!" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}
