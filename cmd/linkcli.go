package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func GetProperties(url string, num string) error {
	var resp *http.Response
	var err error
	if num != "" {
		resp, err = http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		fmt.Println("Response status:", resp.Status)
	}
	posNum, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	if posNum > 0 {
		for i := 0; i < posNum; i++ {
			start := time.Now()
			resp, err = http.Get(url)
			if err != nil {
				return err
			}
			end := time.Now()
			fmt.Println("Request took:", end.Sub(start))
			_ = fmt.Sprintf("Response %d status: %s", i, resp.Status)
		}
		defer resp.Body.Close()

	}
	return nil
}
