package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tld = flag.String("d", "com", "Set tld (defualt:com)")

const allowedChares = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChares, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + *tld)
	}
}
