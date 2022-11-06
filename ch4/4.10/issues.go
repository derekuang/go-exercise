// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"go-exercise/ch4/4.10/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issues created in one month:\n")
	oneMonthBefore := time.Now().UTC().AddDate(0, -1, 0)
	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthBefore) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
