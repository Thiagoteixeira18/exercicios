package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	testcases := []struct {
		name     string
		expected string
	}{
		{name: "thiago", expected: "Hello thiago"},
		{name: "maria", expected: "Hello maria"},
		{name: "", expected: "Hello word"},
	}

	srv := httptest.NewServer(http.HandlerFunc(HandleHello))
	defer srv.Close()

	for _, tc := range testcases {
		path := fmt.Sprintf("%s/?name=%s", srv.URL, tc.name)
		req, erro := http.NewRequest("GET", path, nil)
		if erro != nil {
			t.Fatal(erro)
		}

		response, erro := http.DefaultClient.Do(req)
		if erro != nil {
			t.Fatal(erro)
		}

		if response.StatusCode != 200 {
			t.Fatalf("expected: HTTP 200, got: HTTP %d", response.StatusCode)
		}
		body, erro := ioutil.ReadAll(response.Body)
		if erro != nil {
			t.Fatal(erro)
		}

		var responsebody map[string]string
		json.Unmarshal(body, &responsebody)
		if erro != nil {
			t.Fatal(erro)
		}
		if responsebody["message"] != tc.expected {
			t.Fatalf("expected: %s, got: %s", tc.expected, responsebody["message"])
		}
	}
}
