package main

import (
	"fmt"
	"goseapi"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`<a href="([^"]+)".*?>(.*?)</a>`)

func main() {
	questionid := os.Args[1]
	goseapi.Verbose = true

	var questions []goseapi.Question
	_, err := goseapi.Do("/questions/"+questionid, &questions, &goseapi.Params{
		Site:  goseapi.StackOverflow,
		Sort:  goseapi.SortScore,
		Order: "desc",
	})
	if err != nil {
		fmt.Printf("Unable to fetch question for %s: '%v'", questionid, err)
		os.Exit(1)
	}
	question := questions[0]
	op := question.Owner
	fmt.Printf("The [OP %s](%s)\n\n\n", op.DisplayName, op.Link)

	var answers []goseapi.Answer
	_, err = goseapi.Do("/questions/"+questionid+"/answers", &answers, &goseapi.Params{
		Site:  goseapi.StackOverflow,
		Sort:  goseapi.SortScore,
		Order: "desc",
	})
	if err != nil {
		fmt.Printf("Unable to fetch answer for %s: '%v'", questionid, err)
		os.Exit(1)
	}
	for _, answer := range answers {
		fmt.Printf("--------------\nAnswer '%d':\n\n", answer.ID)
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
			body := comment.Body
			body = strings.Replace(body, "<i>", "*", -1)
			body = strings.Replace(body, "</i>", "*", -1)
			body = strings.Replace(body, "<code>", "`", -1)
			body = strings.Replace(body, "</code>", "`", -1)
			body = strings.Replace(body, "&quot;", "\"", -1)
			body = strings.Replace(body, "&#47;", "/", -1)
			body = strings.Replace(body, "&#39;", "'", -1)
			body = re.ReplaceAllString(body, "[$2]($1)")
			fmt.Printf("> %s\n\n", body)
		}
	}
}
