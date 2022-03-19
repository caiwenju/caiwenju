package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string) (map[string]interface{}, error) {

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("qing qiu guan bi shi bai")
		}
	}(resp.Body)
	defer func() {
		err := recover()
		if err != nil {
			errResult := fmt.Sprintf("%sclose req fail", err)
			fmt.Println(errResult)
		}
	}()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	data,JsonRrr := StringToJson(result.String())
	return data,JsonRrr
}

func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("close fail")
		}
	}(resp.Body)

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func StringToJson(targetSting string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	JsonRrr := json.Unmarshal([]byte(targetSting), &data)
	return data,JsonRrr
}



