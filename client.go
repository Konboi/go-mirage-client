package mirage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Host     string
	initPing bool
}

type Option func(*Client)

func NoInitPing() Option {
	return func(c *Client) {
		c.initPing = false
	}
}

func NewClient(host string, options ...Option) (*Client, error) {
	cli := &Client{
		Host:     host,
		initPing: true,
	}

	for _, option := range options {
		option(cli)
	}

	if cli.initPing {
		res, err := http.Get(host)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("please check mirage host. %s return %d", res.StatusCode)
		}
	}

	return cli, nil
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
