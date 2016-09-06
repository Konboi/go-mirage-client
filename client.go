package mirage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Host string
}

func NewClient(host string) (cli *Client, err error) {
	res, err := http.Get(host)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("please check mirage host. %s return %d", res.StatusCode)
	}

	return &Client{
		Host: host,
	}, nil
}

func (cli *Client) List() (list List, err error) {
	res, err := http.Get(fmt.Sprintf("%s/api/list", cli.Host))
	if err != nil {
		return list, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&list)

	return list, nil
}

func (cli *Client) Launch(subdomain string, image string, branch string) error {
	rp := &RequestParam{
		Subdomain: subdomain,
		Image:     image,
		Branch:    branch,
	}
	params, err := json.Marshal(rp)
	if err != nil {
		return err
	}
	res, err := http.Post(fmt.Sprintf("%s/api/launch", cli.Host), "application/json", bytes.NewBuffer(params))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	status := &Status{}
	json.NewDecoder(res.Body).Decode(status)

	if status.Result != "ok" {
		return fmt.Errorf("%s launch error: %s", subdomain, status.Result)
	}

	return nil
}

func (cli *Client) Terminate(subdomain string) error {
	rp := &RequestParam{
		Subdomain: subdomain,
	}
	params, err := json.Marshal(rp)
	if err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/terminate", cli.Host), "application/json", bytes.NewBuffer(params))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	status := &Status{}
	json.NewDecoder(res.Body).Decode(&status)
	if status.Result != "ok" {
		return fmt.Errorf("%s launch error: %s", subdomain, status.Result)
	}

	return nil
}
