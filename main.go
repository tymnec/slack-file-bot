package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6177862970372-6174772949781-TuXtOH2n8xzGxUuDmVkw6MkJ")
	os.Setenv("CHANNEL_ID", "C065J1CGV33")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArray := []string{os.Getenv("CHANNEL_ID")}

	fileArray := []string{"Latency_Issues_in_Internet_of_Things_A_Review_of_L.pdf"}

	for i := 0; i < len(fileArray); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArray,
			File:     fileArray[i],
		}

		file, err := api.UploadFile(params)

		// Check for the errors
		if err != nil {
			// Print the error
			fmt.Println(err)
			return
		}

		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
