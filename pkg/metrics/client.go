package metrics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/dcos/dcos-cli/pkg/httpclient"
)

// Client is a metrics client for DC/OS.
type Client struct {
	http *httpclient.Client
}

// NewClient creates a new metrics client.
func NewClient(baseClient *httpclient.Client) *Client {
	return &Client{
		http: baseClient,
	}
}

// Node returns the units of a certain node.
func (c *Client) Node(mesosID string) (*Node, error) {
	resp, err := c.http.Get(fmt.Sprintf("/system/v1/agent/%s/metrics/v0/node", mesosID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		var node Node
		err = json.NewDecoder(resp.Body).Decode(&node)
		if err != nil {
			return nil, err
		}
		return &node, nil
	default:
		return nil, fmt.Errorf("HTTP %d error", resp.StatusCode)
	}
}

// Values returns the units of a certain node.
func (c *Client) Values() (*Values, error) {
	// https://soak113s.testing.mesosphe.re/service/monitoring/prometheus/api/v1/label/__name__/values
	resp, err := c.http.Get("/service/monitoring/prometheus/api/v1/label/__name__/values")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		var values Values
		err = json.NewDecoder(resp.Body).Decode(&values)
		if err != nil {
			return nil, err
		}
		return &values, nil
	default:
		return nil, fmt.Errorf("HTTP %d error", resp.StatusCode)
	}
}

// Query returns the units of a certain node.
func (c *Client) Query(query string) error {
	resp, err := c.http.Get("/service/monitoring/prometheus/api/v1/query?query=" + url.QueryEscape(query))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	return nil
}

// Series returns the units of a certain node.
func (c *Client) Series() error {
	resp, err := c.http.Get("/service/monitoring/prometheus/api/v1/series")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	return nil
}
