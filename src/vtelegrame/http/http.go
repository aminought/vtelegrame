package http

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("http")

// GetRequest does http get request to the remote server and returns result
func GetRequest(link string, data map[string]string) []byte {
	req, _ := http.NewRequest("GET", link, nil)

	if data != nil {
		q := req.URL.Query()
		for k, v := range data {
			if v != "" {
				q.Add(k, v)
			}
		}
		req.URL.RawQuery = q.Encode()
	}

	client := http.Client{}
	resp, err := client.Do(req)
	logIfConnectionError(err)
	return readData(resp)
}

// PostRequest does http post request to the remote server and return result
func PostRequest(link string, data map[string]string) []byte {
	values := url.Values{}
	if data != nil {
		for k, v := range data {
			if v != "" {
				values.Add(k, v)
			}
		}
	}
	log.Info(values)
	log.Info(values.Encode())
	resp, err := http.PostForm(link, values)
	logIfConnectionError(err)
	return readData(resp)
}

func readData(resp *http.Response) []byte {
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	logIfAnswerReadingError(err)
	return buf
}

func logIfConnectionError(err error) {
	if err != nil {
		log.Error("Unable to connect to the remote server")
	}
}

func logIfAnswerReadingError(err error) {
	if err != nil {
		log.Error("Unable to read server answer")
	}
}
