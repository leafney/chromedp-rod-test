/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     chromedp-rod-test
 * @Date:        2023-05-30 09:39
 * @Description:
 */

package main

import "github.com/go-resty/resty/v2"

type MyResty struct {
	*resty.Client
}

func NewMyResty() *MyResty {
	return &MyResty{
		Client: resty.New(),
	}
}

func (r *MyResty) SetHeader(header map[string]string) *MyResty {
	r.Client.SetHeaders(header)
	return r
}

func (r *MyResty) SetQueryParam(param map[string]string) *MyResty {
	r.Client.SetQueryParams(param)
	return r
}

func (r *MyResty) Get(url string, params map[string]string) (*resty.Response, error) {
	resp, err := r.Client.R().
		SetQueryParams(params).
		Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *MyResty) Post(url string, body interface{}) (*resty.Response, error) {
	resp, err := r.Client.R().
		SetBody(body).
		Post(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
