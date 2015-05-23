package main

import (
	"fmt"
	"goseapi"
	"os"
)

func main() {
	questionid := os.Args[1]
	var answers []goseapi.Answer
	_, err := goseapi.Do("/questions/"+questionid+"/answers", &answers, &goseapi.Params{
		Site:  goseapi.StackOverflow,
		Sort:  goseapi.SortScore,
		Order: "desc",
	})
	if err != nil {
		fmt.Printf("Unable to fetch answer for %s: '%v'", questionid, err)
		os.Exit(1)
	}
	for _, answer := range answers {
		fmt.Println(answer.ID)
		var comments []goseapi.Comment
		_, err := goseapi.Do(fmt.Sprintf("/answers/%d/comments", answer.ID), &comments, &goseapi.Params{
			Site:   goseapi.StackOverflow,
			Sort:   goseapi.SortCreationDate,
			Order:  "asc",
			Filter: "!40nvjILjlVtsS1Mrk",
		})
		if err != nil {
			fmt.Printf("Unable to fetch comments for answer %s: '%v'", answer.ID, err)
			os.Exit(1)
		}
		for _, comment := range comments {
			fmt.Printf("As mentioned by [%s](%s) in [the comments](%s)  \n", comment.Owner.DisplayName, comment.Owner.Link, comment.Link)
			fmt.Printf("> %s\n\n", comment.Body)
		}
	}
}
