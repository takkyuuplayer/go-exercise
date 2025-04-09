package main

import (
	"encoding/csv"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("MY_SLACK_TOKEN"))

	msgs := make(map[string]slack.SearchMessage, 0)
	reactions := [...]string{"絶好調", "cdo", "jikicdo次期cdo"}
	for _, reaction := range reactions {
		for page := 1; true; page++ {
			res, err := api.SearchMessages(
				fmt.Sprintf(`has::%s: before: 2024-10-01 after:2023-09-30`, reaction),
				slack.SearchParameters{
					Count: 100,
					Sort:  "timestamp",
					Page:  page,
				},
			)
			if err != nil {
				panic(err)
			}

			for _, match := range res.Matches {
				msgs[match.Permalink] = match
			}

			if res.Pages <= page {
				break
			}
		}
	}

	headers := []string{
		"username",
		"text",
	}
	headers = append(headers, reactions[:]...)
	headers = append(headers, "url", "root_url")

	w := csv.NewWriter(os.Stdout)
	w.Comma = '\t'
	_ = w.Write(headers)

	for _, message := range msgs {
		if message.Channel.IsPrivate {
			continue
		}

		replies, _, _, err := api.GetConversationReplies(&slack.GetConversationRepliesParameters{
			ChannelID:          message.Channel.ID,
			Timestamp:          message.Timestamp,
			IncludeAllMetadata: true,
		})
		if err != nil {
			panic(err)
		}

		r := replies[0]
		reacted := make(map[string]int, len(reactions))
		for _, reaction := range r.Reactions {
			reacted[reaction.Name] = reaction.Count
		}
		values := []string{
			message.Username,
			message.Text,
		}
		for _, reaction := range reactions {
			values = append(values, fmt.Sprintf("%d", reacted[reaction]))
		}
		values = append(values, message.Permalink)
		if r.ThreadTimestamp != "" {
			u, _ := url.Parse(message.Permalink)
			u.Path = fmt.Sprintf("/archives/%s/p%s", message.Channel.ID, strings.ReplaceAll(r.ThreadTimestamp, ".", ""))
			u.RawQuery = ""
			values = append(values, u.String())
		} else {
			values = append(values, message.Permalink)
		}

		_ = w.Write(values)
	}
	w.Flush()
}
