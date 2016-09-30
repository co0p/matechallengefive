package solutions

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getAnswer(question string) (string, error) {
	fmt.Print(question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	name := scanner.Text()
	if len(name) < 1 {
		return "", errors.New("No name given")
	}
	return name, nil
}

func Postdata(data url.Values) (string, error) {
	resp, err := http.PostForm(URL, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Two posts you name to the url
func Two() {

	name, err := getAnswer("your name please:")
	if err != nil {
		log.Fatal(err)
	}

	data := url.Values{}
	data.Add("name", name)

	// make the post
	html, err := Postdata(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

}
