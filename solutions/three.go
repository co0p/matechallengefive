package solutions

import (
	"crypto/sha1"
	"fmt"
	"log"
	"math"
	"net/url"
	"strings"
)

func sha1encoder(candidate string) [20]byte {
	return sha1.Sum([]byte(candidate))
}

// Three solved the proof-of-work system by generating
func Three() {

	// read user name
	username, err := getAnswer("username please:")
	if err != nil {
		log.Fatal(err)
	}

	// read email domain
	domain, err := getAnswer("domain please:")
	if err != nil {
		log.Fatal(err)
	}

	// read target prefix
	fmt.Printf("I am attempting to brute force a sha1 hash using %s+<rand>%s. What should the hash start with?\n", username, domain)
	targetPrefix, err := getAnswer("target prefix (try a51dea5):")
	if err != nil {
		log.Fatal(err)
	}

	// generate string which sha1 starts with 'a51deas' brute force style
	// example+19546430@example.com -> a51dea52388b4acd501e296382873057560c4a0c
	var candidate string
	var sha1candidate string
	var match string
	var i uint64
	for i = 0; i < math.MaxUint64; i++ {
		candidate = fmt.Sprintf("%s+%d%s", username, i, domain)
		sha1candidate = fmt.Sprintf("%x", sha1encoder(candidate))

		if strings.HasPrefix(sha1candidate, targetPrefix) {
			fmt.Printf("found: %s --> %s", candidate, sha1candidate)
			match = candidate
			break
		}
	}

	if match == "" {
		log.Fatal("sry. I was not able to create a hash with prefix '%s'", targetPrefix)
	}

	data := url.Values{}
	data.Add("name", username)
	data.Add("email", match)

	html, err := Postdata(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)
}
