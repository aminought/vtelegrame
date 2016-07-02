package http

import "net/http"
import "github.com/op/go-logging"
import "io/ioutil"
import "net/url"
import "bytes"
import "strconv"

var log = logging.MustGetLogger("http")

// GetRequest does http get request to the remote server and returns result
func GetRequest(link string) []byte {
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

// PostRequest does http post request to the remote server and return result
func PostRequest(link string, data map[string]string) []byte {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, v)
	}
	client := &http.Client{}
	r, _ := http.NewRequest("POST", link, bytes.NewBufferString(values.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Lenght", strconv.Itoa(len(values.Encode())))

	resp, err := client.Do(r)
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
