package restkit

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

type Struct1 struct {
	Name string `json:"name"`
	Foo  int
}

type Struct1List []Struct1

func TestJSON(t *testing.T) {
	tests := []struct {
		ServerData   interface{}
		ClientResult interface{}
	}{
		{
			ServerData:   &Struct1{Name: "hello", Foo: 543},
			ClientResult: &Struct1{},
		},
		{
			ServerData:   &Struct1List{Struct1{Name: "entry1", Foo: 10}, Struct1{Name: "entry2", Foo: 20}},
			ClientResult: &Struct1List{},
		},
		{
			ServerData:   nil,
			ClientResult: nil,
		},
	}

	for _, test := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			JSON(w, test.ServerData, http.StatusOK)
		}))
		defer ts.Close()

		url, err := url.Parse(ts.URL)
		if err != nil {
			t.Fatalf("Failed to parse '%s': %#v", ts.URL, err)
		}
		rc := NewRestClient(url)
		err = rc.Request("GET", "/", nil, nil, test.ClientResult)
		if !reflect.DeepEqual(test.ServerData, test.ClientResult) {
			t.Fatalf("Comparison failed: expected '%#v', got '%#v' ", test.ServerData, test.ClientResult)
		}
	}
}
