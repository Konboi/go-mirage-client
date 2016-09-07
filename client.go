package mirage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (cli *Client) Launch(subdomain string, image string, params map[string]string) error {
	values := url.Values{}
	values.Add("subdomain", subdomain)
	values.Add("image", image)

	for k, v := range params {
		values.Add(k, v)
	}

	res, err := http.PostForm(fmt.Sprintf("%s/api/launch", cli.Host), values)
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
	values := url.Values{}
	values.Add("subdomain", subdomain)

	res, err := http.PostForm(fmt.Sprintf("%s/api/terminate", cli.Host), values)
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
