package ilorest

import (
	"encoding/base64"
	"fmt"

	"github.com/go-resty/resty/v2" //Package resty provides simple HTTP and REST client for Go inspired by Ruby rest-client.
	//Reference for resty: https://pkg.go.dev/github.com/go-resty/resty
)

//creating a struct which contains client's details
type RFClient struct {
	BaseUrl       string `json:"baseUrl"`
	UserName      string `json:"userName"`
	Password      string `json:"password"`
	BiosPassword  string `json:"biosPassword"`
	DefaultPrefix string `json:"defaultPrefix"`
	Timeout       int    `json:"timeout"`
	// concurrent int `json:"concurrent"`
	Proxy            string        `json:"proxy"`
	SessionKey       string        `json:"sessionKey"`
	AuthorizationKey string        `json:"authorizationKey"` //for authorization
	SessionLocation  string        `json:"sessionLocation"`
	rest             *resty.Client //resty client object
}

func NewRFClient(config RFClient) *RFClient {
	return &RFClient{
		BaseUrl:       config.BaseUrl,
		UserName:      config.UserName,
		Password:      config.Password,
		BiosPassword:  config.BiosPassword,
		SessionKey:    config.SessionKey,
		DefaultPrefix: config.DefaultPrefix,
		Timeout:       config.Timeout,
		Proxy:         config.Proxy,
		rest:          resty.New(),
	}
}

type requestOptions struct {
	Method      string
	Path        string            //relative URI of RESTful/Redfish
	Body        interface{}       //body
	QueryParams map[string]string //query parameters
	Headers     map[string]string //Additional headers can be passed-in, default to null
	Timeout     int
}

// Interface are the custom type that is used to specify a set of one or more method signatures which are allowed to create a variable of an
// interface type and this variable can be assigned with a concrete type value that has the methods the interface requires.
type RFClientInterface interface {
	Get(path string, queryParams map[string]string, headers map[string]string, timeout int) (*resty.Response, error)
	Patch(path string, queryParams map[string]string, headers map[string]string, timeout int) (*resty.Response, error)

}

// Get makes a get request to the ilorest server
func (c *RFClient) Get(path string, queryParams map[string]string, headers map[string]string, timeout int) (map[string]interface{}, error) {
	opts := requestOptions{
		Method:      "GET",
		Path:        path,
		QueryParams: queryParams,
		Headers:     headers,
		Timeout:     timeout,
	}
	output := map[string]interface{}{}

	_, err := c.rest.R().SetQueryParams(opts.QueryParams).SetHeaders(opts.Headers).SetResult(&output).Get(c.BaseUrl + opts.Path)
	if err != nil {
		//if there is an error
		fmt.Println(err)
		return nil, err
	}
	return output, nil //if no error it returns the output
}

func (c *RFClient) Patch(path string, queryParams map[string]string, headers map[string]string, timeout int) (map[string]interface{}, error) {
	opts := requestOptions{
		Method:      "PATCH",
		Path:        path,
		QueryParams: queryParams,
		Headers:     headers,
		Timeout:     timeout,
	}
	output := map[string]interface{}{}

	_, err := c.rest.R().SetQueryParams(opts.QueryParams).SetHeaders(opts.Headers).SetResult(&output).Patch(c.BaseUrl + opts.Path)
	if err != nil {
		
		fmt.Println(err)
		return nil, err
	}
	return output, nil 
}

// Login logs a person into the olorest server. it supports basic authentication and auth key authentication
func (c *RFClient) Login(username, password string, authMethod string) (map[string]interface{}, error) {
	opts := requestOptions{
		Path: "/Login",

		Timeout: 10,
	}
	output := map[string]interface{}{}

	if authMethod == AuthBasic {
		authKey := base64.StdEncoding.EncodeToString([]byte("Basic " + username + ":" + password))
		opts.Headers = map[string]string{
			"Authorization": authKey,
		}
		opts.Method = "GET"

	} else if authMethod == AuthSession {
		opts.Body = map[string]string{
			"UserName": username,
			"Password": password,
		}
		opts.Method = "POST"
	} else {
		return nil, fmt.Errorf("Invalid auth method")
	}

	_, err := c.rest.R().SetBody(opts.Body).SetResult(&output).Post(c.BaseUrl + opts.Path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if authMethod == AuthSession {
		c.SessionKey = output["SessionToken"].(string)
	}
	return output, nil
}
