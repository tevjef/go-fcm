package fcm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"net/http/httputil"
)

const (
	endpointFormat = "https://fcm.googleapis.com/v1/projects/%s/messages:send"
)

// Client abstracts the interaction between the application server and the
// FCM server via HTTP protocol. The developer must obtain a service account
// private key in JSON and the Firebase project id from the Firebase console and pass it to the `Client`
// so that it can perform authorized requests on the application server's behalf.
type Client struct {
	// https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages/send
	projectID     string
	endpoint      string
	client        *http.Client
	tokenProvider *tokenProvider
}

// NewClient creates new Firebase Cloud Messaging Client based on a json service account file credentials file.
func NewClient(projectID string, credentialsLocation string, opts ...Option) (*Client, error) {
	c := &Client{
		endpoint: fmt.Sprintf(endpointFormat, projectID),
		client:   http.DefaultClient,
		//tokenProvider: tp,
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	customerCli := &http.Client{
		Transport: c.client.Transport,
	}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, customerCli)
	tp, err := newTokenProvider(ctx, credentialsLocation)
	if err != nil {
		return nil, err
	}
	c.tokenProvider = tp
	return c, nil
}

// NewClientFromBytes creates new Firebase Cloud Messaging Client based on a json service account file credentials file.
func NewClientFromBytes(projectID string, jsonKey []byte, opts ...Option) (*Client, error) {
	c := &Client{
		endpoint: fmt.Sprintf(endpointFormat, projectID),
		client:   http.DefaultClient,
		//tokenProvider: tp,
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	customerCli := &http.Client{
		Transport: c.client.Transport,
	}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, customerCli)
	tp, err := newTokenProviderFromBytes(ctx, jsonKey)
	if err != nil {
		return nil, err
	}
	c.tokenProvider = tp
	return c, nil
}

// Send sends a message to the FCM server.
func (c *Client) Send(req *SendRequest) (*Message, error) {
	// validate
	if err := req.Message.Validate(); err != nil {
		return nil, err
	}

	// marshal message
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return c.send(data)
}

// send sends a request.
func (c *Client) send(data []byte) (*Message, error) {
	// create request
	req, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// get bearer token
	token, err := c.tokenProvider.token()
	if err != nil {
		return nil, err
	}

	// add headers
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	// execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		requestBytes, _ := httputil.DumpRequest(req, true)
		responseBytes, _ := httputil.DumpResponse(resp, true)

		if resp.StatusCode >= http.StatusInternalServerError {
			return nil, HttpError{
				RequestDump:  string(requestBytes),
				ResponseDump: string(responseBytes),
				Err:          fmt.Errorf(fmt.Sprintf("%d error: %s", resp.StatusCode, resp.Status)),
			}
		}
		return nil, HttpError{
			RequestDump:  string(requestBytes),
			ResponseDump: string(responseBytes),
			Err:          fmt.Errorf("%d error: %s", resp.StatusCode, resp.Status),
		}
	}

	// build return
	response := new(Message)
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

// HttpError contains the dump of the request and response for debugging purposes.
type HttpError struct {
	RequestDump  string
	ResponseDump string
	Err          error
}

func (fcmError HttpError) Error() string {
	return fcmError.Err.Error()
}
