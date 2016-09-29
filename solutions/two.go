package solutions;

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"net/url"
	"bufio"
	"os"
)

// Two posts you name to the url
func Two() {

	// read user name
	fmt.Print("\n Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	if (len(name) < 1) {
		log.Fatal("You have to give me your name!")
	}

	data := url.Values{}
	data.Add("name", name)

	// make the post
	resp, err := http.PostForm(URL, data)
	if (err != nil) {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		log.Fatal(err)
	}

	// print the solution
	html := string(body);
	fmt.Println(html);
}