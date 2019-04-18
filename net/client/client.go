package client

import (
	"crypto/tls"
	"encoding/json"
	"github.com/dulumao/Guten-utils/net/httplib"
	"math/rand"
	"net/http"
	"time"
)

type Client struct {
	host     string
	method   string
	params   interface{}
	username string
	password string
	Headers  map[string]string
	Timeout  int // Second
	Debug    bool
	IsSSL    bool
	Request  *httplib.HTTPRequest
}

func New(client *Client) *Client {
	return client
}

func Default() *Client {
	return New(&Client{
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36",
		},
		Timeout: 30,
		Debug:   true,
	})
}

func (self *Client) newRequest(url, method string) *httplib.HTTPRequest {
	self.Request = httplib.NewRequest(url, method)

	return self.Request
}

func (self *Client) NewEndPoint(host string, method ...string) *Client {
	self.host = host

	if len(method) > 0 {
		self.method = method[0]
	}

	return self
}

func (self *Client) SetBasicAuth(username, password string) *Client {
	self.username = username
	self.password = password

	return self
}

func (self *Client) SetMethod(name string, params ...interface{}) *Client {
	if len(params) > 0 {
		self.params = params[0]
	} else {
		self.params = nil
	}

	self.method = name

	return self
}

func (self *Client) Execute() *httplib.HTTPRequest {
	if self.method == "" {
		self.method = "GET"
	}

	self.Request = self.newRequest(self.host, self.method)

	self.Request.Debug(self.Debug)
	self.Request.SetTimeout(time.Duration(self.Timeout)*time.Second, time.Duration(self.Timeout)*time.Second)

	if self.IsSSL {
		self.Request.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	for k, v := range self.Headers {
		self.Request.Header(k, v)
	}

	playLoad := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  self.method,
		"id":      rand.Intn(10000),
	}

	if self.params != nil {
		playLoad["params"] = self.params
	}

	self.Request.JSONBody(playLoad)
	self.Request.SetBasicAuth(self.username, self.password)

	return self.Request
}

func (self *Client) GetResponse() (*http.Response, error) {
	return self.Execute().Response()
}

func (self *Client) GetBody() (string, error) {
	return self.Execute().String()
}

func (self *Client) GetMustBody() string {
	if body, err := self.Execute().String(); err != nil {
		return ""
	} else {
		return body
	}
}

func (self *Client) GetJson(v interface{}) error {
	if s, err := self.GetBody(); err != nil {
		return err
	} else {
		err := json.Unmarshal([]byte(s), v)

		return err
	}
}
