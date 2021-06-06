package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	// TODO: create a script that runs this zoomRem code and adds the created reminder script to crontab

	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)

	/* list, err := srv.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("can't access calendar list")
	}
	if len(list.Items) == 0 {
		fmt.Println("No calendar lists")
	} else {
		for _, item := range list.Items {
			fmt.Printf("summary: %v, id: %v, desc: %v\n", item.Summary, item.Id, item.Description)
		}
	} */

	// t0doist calendar id: 6uud686gk210hlq5qm0p8if0t8@group.calendar.google.com

	// FIX: maxResults to be a non-arbitrary/reasonable value
	events, err := srv.Events.List("primary").
		SingleEvents(true).TimeMin(t).OrderBy("startTime").MaxResults(5).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			if strings.Contains(item.Location, "zoom") {
				// fmt.Printf("%v (%v), location: %v\n", item.Summary, date, item.Location)

				fmt.Printf("desc: %v, sum: %v, id: %v", item.Description, item.Summary, item.Id)
				f := item.Summary
				// eventually want this to write Reminder scripts in a user specified directory
				// create a crontab to have this run, say every 12 hours. (can change later)

				// item.Summary is title/name of event
				// item.Location is location of event
				// item.Start.DateTime is start of of event

				// TODO: determine location of terminal-notifier on the users system
				terminalNotifierLocation := "'/Applications/terminal-notifier.app/Contents/MacOS/terminal-notifier'"
				quotedItemSummary := fmt.Sprintf("'%s'", item.Summary)
				if err != nil {
					log.Fatalf("cannot convert item.Location to quotedItem")
				}
				titleOfRem := " -title " + quotedItemSummary
				// TODO: get remMessage from user or set as environmental variable so they don't have to keep interacting with
				// this script over and over again
				remMessage := " -message 'hello'"
				quotedItemLocation := fmt.Sprintf("'%s'", item.Location)
				meetingLink := " -open " + quotedItemLocation

				// "'/Applications/terminal-notifier.app/Contents/MacOS/terminal-notifier' -title 'CS 111 Lec' -message 'operate' -open https://ucla.zoom.us/j/98811081960?pwd=T1M5UWJTRTd2K3Ntd3Z5bUxrVExjZz09"
				var sb strings.Builder
				sb.WriteString(terminalNotifierLocation)
				sb.WriteString(titleOfRem)
				sb.WriteString(remMessage)
				sb.WriteString(meetingLink)

				if f != "" {
					err := ioutil.WriteFile(f, []byte(sb.String()), 0744)
					if err != nil {
						log.Fatalf("cannot write to %v, err: %v", f, err)
					}
				}
			}
		}
	}
}
