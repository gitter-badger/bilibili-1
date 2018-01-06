package bilibili

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// Catalogue URL
const (
	VideoURL = "https://www.bilibili.com/video"
)

// Regexp pattern
const (

	// For a multi-p video, cid could be found with pattern `cid='\d+?'`,
	// but for a single p video, cid could only be found by `cid=\d+?&`.
	cidPattern = `cid='?(\d+?)[&']`
)

// ConnectionError is raised when the status code of a response is >= 400.
type ConnectionError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface.
func (e *ConnectionError) Error() string {
	return fmt.Sprintf("%s. Status code: %d", e.Message, e.StatusCode)
}

// FindCID parse the video page to get the cid slice.
// A video could contain several *P*, and each P has a unique cid.
func FindCID(av int) ([]int, error) {

	url, _ := url.ParseRequestURI(VideoURL)
	url.Path = fmt.Sprintf("/av%d", av)

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &ConnectionError{
			StatusCode: resp.StatusCode,
			Message:    "fail to get the video page",
		}
	}

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(bb)

	re := regexp.MustCompile(cidPattern)
	found := re.FindAllStringSubmatch(body, -1)

	var result []int
	for _, o := range found {
		cid, err := strconv.Atoi(o[1])
		if err != nil {
			return nil, err
		}
		result = append(result, cid)
	}

	// Remove the repeated cid, which should be the first p of this video.
	if len(result) > 1 && result[0] == result[len(result)-1] {
		result = result[:len(result)-1]
	}

	return result, nil

}
