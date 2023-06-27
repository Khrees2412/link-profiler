package app

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	Fastest              float64 `json:"fastest"`
	Slowest              float64 `json:"slowest"`
	Mean                 float64 `json:"mean"`
	Median               float64 `json:"median"`
	PercentageSuccessful string  `json:"percentageSuccessful"`
	ErrorCodes           []int   `json:"errorCodes"`
	BytesOfSmallest      int     `json:"bytesOfSmallest"`
	ByteOfLargest        int     `json:"byteOfLargest"`
}

func GetProperties(url string, num string) (*Result, error) {

	var resp *http.Response
	var err error
	var body []byte

	posNum, err := strconv.Atoi(num)
	if err != nil {
		return nil, err
	}
	if posNum < 1 {
		return nil, errors.New("profiling number must not be less than 1")
	}

	requests := make([]float64, 0)
	responseSizes := make([]int, 0)
	var successfulResponses int
	errorCodes := make([]int, 0)

	for i := 0; i < posNum; i++ {

		start := time.Now()

		resp, err = http.Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode >= 400 {
			errorCodes = append(errorCodes, resp.StatusCode)
		}
		if resp.StatusCode == http.StatusOK {
			successfulResponses += 1
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		size := len(body)

		end := time.Now()
		timeTaken := end.Sub(start)

		requests = append(requests, timeTaken.Seconds())
		responseSizes = append(responseSizes, size)

	}
	defer func(body io.ReadCloser) {
		err = body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	slowest := getHighestFloat(requests)
	fastest := getLowestFloat(requests)
	mean := getMeanOfRequests(requests)
	median := getMedianRequest(requests)
	largest := getHighestInt(responseSizes)
	smallest := getLowestInt(responseSizes)
	percentSuccessful := fmt.Sprintf("%d%%", successfulResponses/posNum*100)

	result := Result{
		Fastest:              fastest,
		Slowest:              slowest,
		Mean:                 mean,
		Median:               median,
		PercentageSuccessful: percentSuccessful,
		ErrorCodes:           errorCodes,
		BytesOfSmallest:      smallest,
		ByteOfLargest:        largest,
	}

	return &result, nil
}
