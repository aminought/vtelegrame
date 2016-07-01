package http

import "net/http"
import "github.com/op/go-logging"
import "io/ioutil"

var log = logging.MustGetLogger("http")

// Request does http request to the remote server and returns result
func Request(link string) []byte {
	resp, err := http.Get(link)
	if err != nil {
		log.Error("Unable to connect to the remote server")
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("Unable to read server answer")
	}

	return buf
}
