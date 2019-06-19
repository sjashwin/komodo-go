package komodo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Auth struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

type Client struct {
	*Auth
	Request *http.Request
}

type dataBinary struct {
	JSON   string `json:"jsonrpc"`
	ID     string `json:"id"`
	Method string `json:"method"`
	Params Params `json:"params,omitempty"`
}

type Params map[string]interface{}

// KomodoClient returns a new komodo client
// filePath of the json file that contains the authentication
// details
func NewClient(filePath string) (*Client, error) {
	var auth = &Auth{}
	var data []byte
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	if _, err = file.Read(data); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, auth); err != nil {
		return nil, err
	}
	return &Client{Auth: auth}, nil
}

func (c *Client) send(method string, params Params) ([]byte, error) {
	body, err := c.generateCurl(method, params)
	if err != nil {
		return nil, err
	}
	host := fmt.Sprintf("http://%s:%s/", c.Host, strconv.Itoa(c.Port))
	request, err := http.NewRequest("POST", host, body)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(c.User, c.Password)
	request.Header.Set("Content-Type", "text/plain;")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	d, err := ioutil.ReadAll(resp.Body)
	return d, err
}

func (c *Client) generateCurl(method string, params Params) (*strings.Reader, error) {
	data := &dataBinary{
		JSON:   "1.0",
		ID:     "curltest",
		Method: method,
		Params: params,
	}
	d, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal method and paramerters")
	}
	return strings.NewReader(string(d)), nil
}
