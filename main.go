package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	cookie := ""
	day := 2
	filename := "0" + strconv.Itoa(day) + "/input.txt"
	url := "https://adventofcode.com/2022/day/" + strconv.Itoa(day) + "/input/" + cookie
	log.Println(url)
	/*err := DownloadFile(url, filename)
	if err != nil {
		log.Fatal(err)
	}
	*/
	b, err := HTTPwithCookies(url, "", filename)
	if err != nil {
		panic(err)
	}
	println(string(b))
}

func HTTPwithCookies(url, sessionid, filepath string) (b []byte, err error) {
	out, err := os.Create(filepath)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionid})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}

func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)

	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
