package main

import (
	"net/http/httptest"
	"testing"
)

/*
func TestPing(t *testing.T) {
	fmt.Printf("Comecando o test\n")
	const expected_body = "<html><body>ping</body></html>"

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	Handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Status code = %d", resp.StatusCode)
	//fmt.Println(resp.Header.Get("Content-Type"))
	if string(body) != expected_body {
		t.Fatalf("expected body %q, got %q", expected_body, string(body))
	}
}
*/

func TestReturnCode(t *testing.T) {
	const expected_code = 200

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	Handler(w, req)

	resp := w.Result()

	if resp.StatusCode != expected_code {
		t.Fatalf("expected code %q, got %q", expected_code, resp.StatusCode)
	}
}
