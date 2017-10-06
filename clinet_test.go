package mirage

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockTopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dummy Top")
}

func MockListHandler(w http.ResponseWriter, r *http.Request) {
	jsonStr := `{"result":[{"id":"asdfg12345","short_id":"as12","subdomain":"do1","branch":"bra1","image":"img1","ipaddress":"127.0.0.1"},{"id":"hjkl7890","short_id":"hj78","subdomain":"sub2","branch":"bra2","image":"img2","ipaddress":"127.0.0.2"}]}`

	w.Header().Set("Content-Tyep", "application/json")
	fmt.Fprintf(w, jsonStr)
}

func MockLaunchHandler(w http.ResponseWriter, r *http.Request) {
	jsonStr := `{"result": "ok"}`
	r.ParseForm()

	if r.Form.Get("subdomain") == "" || r.Form.Get("image") == "" || r.Form.Get("branch") == "" {
		jsonStr = `{"result": "false"}`
	}

	w.Header().Set("Content-Tyep", "application/json")
	fmt.Fprintf(w, jsonStr)
}

func MockTerminateHandler(w http.ResponseWriter, r *http.Request) {
	jsonStr := `{"result": "ok"}`

	r.ParseForm()

	if r.Form.Get("subdomain") == "" {
		jsonStr = `{"result": "false"}`
	}

	w.Header().Set("Content-Tyep", "application/json")
	fmt.Fprintf(w, jsonStr)
}

func TestNewClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(MockTopHandler))
	defer ts.Close()

	_, err := NewClient("http://example.mirage.fuga")
	if err == nil {
		t.Fatal("dummy is not working mirage")
	}

	if _, err := NewClient("http://example.mirage.fuga", NoInitPing()); err != nil {
		t.Fatal("check is false")
	}

	_, err = NewClient(ts.URL)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(MockListHandler))
	defer ts.Close()

	cli, err := NewClient(ts.URL)
	if err != nil {
		t.Fatal(err.Error())
	}

	list, err := cli.List()
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(list.Result) != 2 {
		t.Fatalf("[error] /api/list response: result num %d", len(list.Result))
	}

	if list.Result[0].ID != "asdfg12345" || list.Result[1].ID != "hjkl7890" {
		t.Fatal("[error] /api/list response")
	}
}

func TestLaunch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(MockLaunchHandler))
	defer ts.Close()

	cli, err := NewClient(ts.URL)
	if err != nil {
		t.Fatal(err.Error())
	}

	params := make(map[string]string)
	params["branch"] = "dummy-branch"
	err = cli.Launch("dummy-sub", "dummy-image", params)
	if err != nil {
		t.Fatalf("[error] Launch: %s", err.Error())
	}
}

func TestTerminate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(MockTerminateHandler))
	defer ts.Close()

	cli, err := NewClient(ts.URL)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = cli.Terminate("dummy-sub")
	if err != nil {
		t.Fatalf("[error] Launch: %s", err.Error())
	}
}
