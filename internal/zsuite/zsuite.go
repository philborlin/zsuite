package zsuite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Zsuite struct {
	token      string
	id         string
	year       int
	httpPrefix string
	client     *http.Client
}

func New(token, id, httpPrefix string, year int) *Zsuite {
	return &Zsuite{
		token:      token,
		id:         id,
		httpPrefix: httpPrefix,
		year:       year,
		client:     &http.Client{},
	}
}

func (z *Zsuite) get(url string, respBody any) error {
	req, err := http.NewRequest("GET", z.fqURL(url), nil)
	if err != nil {
		return err
	}

	return z.do(req, respBody)
}

func (z *Zsuite) put(url string, reqBody any) error {
	bs, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(bs))

	req, err := http.NewRequest("PUT", z.fqURL(url), bytes.NewReader(bs))
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}

	return z.do(req, nil)
}

func (z *Zsuite) fqURL(partialURL string) string {
	return z.httpPrefix + partialURL
}

func (z *Zsuite) do(req *http.Request, respBody any) error {
	req.AddCookie(&http.Cookie{Name: "SESSION", Value: z.token})

	resp, err := z.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		fmt.Println(resp.Status)
	}

	if respBody != nil {
		bsr, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(bsr, respBody)
		if err != nil {
			fmt.Printf("zsuite.do() %s - %v\n", string(bsr), respBody)
		}

		return err
	}

	return nil
}
