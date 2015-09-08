package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Camera struct {
	Camera string `json:"camera"`
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
	Name   string `json:"name"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	DownloadImages("http://207.251.86.238/cctv601.jpg")
	// jsonBytes, err := ioutil.ReadFile("nyc-cameras.json")
	// if err != nil {
	// 	panic(err)
	// }
	// var cameras []Camera
	// err = json.Unmarshal(jsonBytes, &cameras)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, camera := range cameras {
	// 	client := &http.Client{}
	// 	req, err := http.NewRequest("GET", camera.Camera, nil)
	// 	req.Header.Add("User-Agent", RandStringBytes(10))
	// 	resp, err := client.Do(req)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	defer resp.Body.Close()
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	err = ioutil.WriteFile(camera.Camera[22:], body, 0644)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }
}

func DownloadImages(link string) {
	for {
		time.Sleep(time.Second * 3)
		go func() {
			resp, err := http.Get(link)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}
			filename := fmt.Sprintf("cctv261_%d.jpg", time.Now().Unix())
			ioutil.WriteFile(filename, body, 0644)
		}()
	}
}
