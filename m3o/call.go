package m3o

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const m3oV1BaseUrl = "https://api.m3o.com/v1"

type Data map[string]interface{}

type micro struct {
	token    string
	response *http.Response
	error    error
}

func NewMicro(token string) *micro {
	return &micro{
		token: token,
	}
}

func (m *micro) Post(servicePath string, data Data) *micro {
	url := fmt.Sprintf("%s%s", m3oV1BaseUrl, servicePath)
	method := "POST"

	bs, err := json.Marshal(data)
	if err != nil {
		m.error = err
		return m
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(bs))

	if err != nil {
		m.error = err
		return m
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		m.error = err
		return m
	}

	m.response = resp

	return m
}

func (m *micro) Bytes() ([]byte, error) {
	if m.error != nil {
		return nil, m.error
	}

	if m.response.StatusCode != http.StatusOK {
		return nil, errors.New("micro response status is not 200")
	}

	defer m.response.Body.Close()

	data, err := ioutil.ReadAll(m.response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *micro) JSON(v interface{}) error {
	data, err := m.Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (m *micro) String() (string, error) {
	data, err := m.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}
