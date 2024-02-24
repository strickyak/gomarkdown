package main

import "github.com/microcosm-cc/bluemonday"
import "github.com/russross/blackfriday"
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Markdown(s string) string {
	t := blackfriday.MarkdownCommon([]byte(s))
	html := bluemonday.UGCPolicy().SanitizeBytes(t)
	return string(html)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(Markdown(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
