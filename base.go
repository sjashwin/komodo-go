package komodo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Auth struct {
	Password string `json:"password"`
	User     string `json:"user"`
	Port     int    `json:"port"`
}

type Client struct {
	*Auth
	Request *http.Request
}

type dataBinary struct {
	Json   string `json:"json"`
	ID     string `json:"id"`
	Method string `json:"method"`
	Params Params `json:"params,omitempty"`
}

type Params map[string]interface{}

// KomodoClient returns a new komodo client
// filePath of the json file that contains the authentication
// details
func KomodoClient(filePath string) *Client {
	var auth = &Auth{}
	var data []byte
	file, err := os.Open(filePath)
	if err != nil {

	}
	file.Read(data)
	if err := json.Unmarshal(data, auth); err != nil {

	}
	return &Client{Auth: auth}
}

func (c *Client) send(body *strings.Reader) (*http.Response, error) {
	request, err := http.NewRequest("POST", "http://127.0.0.1:"+string(c.Port), body)
	if err != nil {
		return nil, nil
	}
	request.SetBasicAuth(c.User, c.Password)
	request.Header.Set("Content-Type", "text/plain;")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, err
}

func (c *Client) generateCurl(method string, params Params) (*strings.Reader, error) {
	data := &dataBinary{
		Json:   "1.0",
		ID:     "curltest",
		Method: method,
		Params: params,
	}
	d, err := json.Marshal(data)
	fmt.Println(string(d))
	if err != nil {
		return nil, fmt.Errorf("could not marshal method and paramerters")
	}
	return strings.NewReader(string(d)), nil
}
