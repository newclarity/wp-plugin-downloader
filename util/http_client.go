package util

import (
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
	response *http.Response
	body     []byte
	url      string
	status   status.Status
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (me *HttpClient) Download(url string) (b []byte, sts status.Status) {
	for range only.Once {
		c := NewHttpClient()
		sts = c.GET(url)
		if is.Error(sts) {
			break
		}
		b, sts = c.GetBody()
		if is.Error(sts) {
			break
		}
		_ = c.response.Body.Close()
		if c.response.StatusCode == http.StatusNotFound {
			sts = status.Fail().
				SetHttpStatus(http.StatusNotFound).
				SetMessage("not found: %s", url)
			break
		}
	}
	return b, sts
}

func (me *HttpClient) GET(url string) (sts status.Status) {
	me.url = url
	r, err := http.Get(url)
	if err != nil {
		sts = status.Wrap(err, &status.Args{
			Message: fmt.Sprintf("error while requesting '%s'", url),
		})
	}
	me.response = r
	me.status = sts
	return sts
}

func (me *HttpClient) GetBody() (body []byte, sts status.Status) {
	for range only.Once {
		if me.response == nil {
			log.Fatal("cannot call HttpClient.GetBody() before response is set")
		}
		var err error
		body, err = ioutil.ReadAll(me.response.Body)
		if err != nil {
			sts = status.Wrap(err, &status.Args{
				Message: fmt.Sprintf("error reading response body for '%s'", me.url),
				Data:    me.response,
			})
		}
		_ = me.response.Body.Close()
	}
	me.body = body
	return body, sts
}
