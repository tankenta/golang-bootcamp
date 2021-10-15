package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"ch04/ex10/issues/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var within1Month, within1Year, over1Year []*github.Issue
	now := time.Now()
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(now.AddDate(0, -1, 0)):
			within1Month = append(within1Month, item)
		case item.CreatedAt.After(now.AddDate(-1, 0, 0)):
			within1Year = append(within1Year, item)
		case item.CreatedAt.Before(now.AddDate(-1, 0, 0)):
			over1Year = append(over1Year, item)
		}
	}

	const dateFormat = "2006-01-02"
	fmt.Printf("%d issues within a month:\n", len(within1Month))
	for _, item := range within1Month {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format(dateFormat), item.User.Login, item.Title)
	}
	fmt.Println()
	fmt.Printf("%d issues within a year:\n", len(within1Year))
	for _, item := range within1Year {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format(dateFormat), item.User.Login, item.Title)
	}
	fmt.Println()
	fmt.Printf("%d issues over a year ago:\n", len(over1Year))
	for _, item := range over1Year {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format(dateFormat), item.User.Login, item.Title)
	}
}
