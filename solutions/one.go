package solutions;

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)


// One makes a post to the url and examines the response
func One() {
	resp, err := http.PostForm(URL, nil)
	if (err != nil) {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		log.Fatal(err)
	}
	html := string(body);

	fmt.Println(html);
}