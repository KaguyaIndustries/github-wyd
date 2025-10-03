package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		print("Usage: wyd <GITHUB USERNAME>")
		return
	}

	getEvent(args[0])
	return
}

func getEvent(user string) {
	client := &http.Client{}
	resp, err := client.Get("https://api.github.com/users/" + user + "/events")
	var eventData []map[string]interface{}
	// var events [...]string

	if err != nil {
		println("Some Error happened lol")
		return
	}

	if resp.StatusCode == 200 {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(resp.Body)
		body, _ := io.ReadAll(resp.Body)

		err = json.Unmarshal(body, &eventData)
		if err != nil {
			fmt.Println("Could not unmarshal event data: ", err)
			return
		}

		fmt.Println(user, " is doing the following:")

		for _, event := range eventData {
			if eventType, ok := event["type"].(string); ok {
				switch eventType {
				case "CommitCommentEvent":
					fmt.Println("CommitCommentEvent")
				case "CreateEvent":
					fmt.Println("Created event")
				case "DeleteEvent":
					fmt.Println("Deleted event")
				case "ForkEvent":
					fmt.Println("Forked event")
				case "GollumEvent":
					fmt.Println("Gollum event")
				case "IssueCommentEvent":
					fmt.Println("Issue comment event")
				case "IssuesEvent":
					fmt.Println("Issues event")
				case "MemberEvent":
					fmt.Println("Member event")
				case "PublicEvent":
					fmt.Println("Public event")
				case "PullRequestEvent":
					fmt.Println(" - Created a Pull Request at", event["repo"].(map[string]interface{})["name"].(string))
				case "PullRequestReviewEvent":
					fmt.Println("Pull request review event")
				case "PullRequestReviewCommentEvent":
					fmt.Println("Pull request review comment event")
				case "PullRequestReviewThreadEvent":
					fmt.Println("Pull request review thread event")
				case "PushEvent":
					fmt.Println(" - Pushed ", len(event["payload"].(map[string]interface{})["commits"].([]interface{})), "commits into", event["repo"].(map[string]interface{})["name"].(string))
				case "ReleaseEvent":
					fmt.Println("Release event")
				case "SponsorEvent":
					fmt.Println("Sponsor event")
				case "WatchEvent":
					fmt.Println("Watch event")
				}
			}
		}
	} else if resp.StatusCode == 404 {
		println("User not found")
	}
}
