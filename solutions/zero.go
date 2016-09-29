package solutions

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)

// Zero solves the first challenge by fetching the website and extracting the hidden content
func Zero() {

	// fetch data
	resp, err := http.Get(URL);
	if (err != nil) {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// convert to string
	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		log.Fatal(err)
	}
	html := string(body);

	// extract comment
	startPos := strings.Index(html, "<!-") + len("<!-")
	endPos := strings.Index(html, "->")

	if (startPos == -1 || endPos == -1) {
		log.Fatal("could not extract comments")
	}

	// print it
	fmt.Println(html[startPos:endPos])
}