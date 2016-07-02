package http

import "net/http"
import "github.com/op/go-logging"
import "io/ioutil"
import "net/url"

var log = logging.MustGetLogger("http")

// GetRequest does http get request to the remote server and returns result
func GetRequest(link string) []byte {
	resp, err := http.Get(link)
	logIfConnectionError(err)
	return readData(resp)
}

// PostRequest does http post request to the remote server and return result
func PostRequest(link string, data map[string]string) []byte {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, v)
	}

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
