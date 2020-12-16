package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/chap4/github"
)

func main() {
	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	thisMonth := make(map[int]*github.Issue)
	thisYear := make(map[int]*github.Issue)
	moreThanYear := make(map[int]*github.Issue)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) {
			thisMonth[item.Number] = item
		} else if item.CreatedAt.After(oneYearAgo) {
			thisYear[item.Number] = item
		} else {
			moreThanYear[item.Number] = item
		}
	}
	fmt.Println("1ヶ月以内")
	show(thisMonth)
	fmt.Println("1年以内")
	show(thisYear)
	fmt.Println("1年以上前")
	show(moreThanYear)
}

func show(issues map[int]*github.Issue) {
	for _, issue := range issues {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			issue.Number, issue.User.Login, issue.Title, issue.CreatedAt.Format("2006/01/02"))
	}
}
