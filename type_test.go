package mirage

import (
	"encoding/json"
	"testing"
)

func TestListUnmarshal(t *testing.T) {
	jsonStr := `{"result":[{"id":"asdfg12345","short_id":"as12","subdomain":"do1","branch":"bra1","image":"img1","ipaddress":"127.0.0.1"},{"id":"hjkl7890","short_id":"hj78","subdomain":"sub2","branch":"bra2","image":"img2","ipaddress":"127.0.0.2"}]}`

	list := &List{}

	err := json.Unmarshal([]byte(jsonStr), list)
	if err != nil {
		t.Fatal("[error] type List json.Unmarshal:", err.Error())
	}

	if len(list.Result) != 2 {
		t.Fatal("[error] type List json.Unmarshal Result num:%d", len(list.Result))
	}

	if list.Result[0].ID != "asdfg12345" || list.Result[1].ID != "hjkl7890" {
		t.Fatal("[error] type List json.Unmarshal Result[0].ID:%s,Result[1].ID:%s", list.Result[0].ID, list.Result[1].ID)
	}
}

func TestStatusUnmarshal(t *testing.T) {
	jsonStr := `{"result": "ok"}`
	status := &Status{}

	err := json.Unmarshal([]byte(jsonStr), status)
	if err != nil {
		t.Fatal("[error] type List json.Unmarshal:", err.Error())
	}

	if status.Result != "ok" {
		t.Fatal("[error] type List json.Unmarshal Result:%s", status.Result)
	}

	jsonStr2 := `{"result": "fail"}`
	status2 := &Status{}

	err = json.Unmarshal([]byte(jsonStr2), status2)
	if err != nil {
		t.Fatal("[error] type List json.Unmarshal:", err.Error())
	}

	if status2.Result != "fail" {
		t.Fatal("[error] type List json.Unmarshal Result:%s", status2.Result)
	}
}
