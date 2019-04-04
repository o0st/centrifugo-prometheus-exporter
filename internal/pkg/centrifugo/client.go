package centrifugo

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ErrStatusCode struct {
	Code int
}

func (e ErrStatusCode) Error() string {
	return fmt.Sprintf("wrong status code: %d", e.Code)
}

var DefaultHTTPClient = &http.Client{Transport: &http.Transport{
	MaxIdleConnsPerHost: 100,
}, Timeout: time.Second * 5}

type Client struct {
	endpoint   string
	apiKey     []byte
	httpClient *http.Client
}

func NewClient(endpoint, apiKey string) *Client {
	endpoint = strings.TrimRight(endpoint, "/")
	if !strings.HasSuffix(endpoint, "/api") {
		endpoint = endpoint + "/api/"
	}

	return &Client{
		httpClient: DefaultHTTPClient,
		endpoint:   endpoint,
		apiKey:     []byte(apiKey),
	}
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) generateApiSign(data []byte) string {
	sign := hmac.New(sha256.New, c.apiKey)
	sign.Write(data)
	return hex.EncodeToString(sign.Sum(nil))
}

func (c *Client) GetStats(ctx context.Context) (*StatsResponse, error) {
	body, err := json.Marshal(&StatsCommand{
		Method: "stats",
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Sign", c.generateApiSign(body))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrStatusCode{resp.StatusCode}
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responses := make([]*StatsResponse, 0)

	if err := json.Unmarshal(body, &responses); err != nil {
		return nil, err
	}

	if len(responses) == 0 {
		return nil, errors.New("empty server response")
	}

	return responses[0], nil
}
