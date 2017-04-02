package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	ts := httptest.NewServer(nil)
	defer ts.Close()

	// リクエストの送信先はテストサーバのURLへ。
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	if "Hello World!" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}
