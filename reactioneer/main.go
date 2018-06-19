package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// expected slack event types
	urlVerification = "url_verification"
	eventCallback   = "event_callback"

	errTokenVerification = errors.New("failed to verify slack token")
	errUnknownEventType  = errors.New("unknown event type received")

	jsonHeader = map[string]string{"Content-Type": "application/json"}
)

// ItemStruct keeps slack event item details
type ItemStruct struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

// SlackEvent details in SlackRequest
type SlackEvent struct {
	Type     string     `json:"type"`
	User     string     `json:"user"`
	Item     ItemStruct `json:"item"`
	Reaction string     `json:"reaction"`
	EventTs  string     `json:"event_ts"`
}

// SlackRequest is parsed from incoming requests
type SlackRequest struct {
	Token       string     `json:"token"`
	Challenge   string     `json:"challenge"`
	TeamID      string     `json:"team_id"`
	APIAppID    string     `json:"api_app_id"`
	Event       SlackEvent `json:"event"`
	EventType   string     `json:"type"`
	EventID     string     `json:"event_id"`
	EventTime   int        `json:"event_time"`
	AuthedUsers []string   `json:"authed_users"`
}

func handler(reaction SlackRequest) (events.APIGatewayProxyResponse, error) {
	// verify slack token
	if reaction.Token != os.Getenv("SLACK_TOKEN") {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
		}, errTokenVerification
	}

	if reaction.EventType == urlVerification {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       fmt.Sprintf(`{"challenge":"%s"}`, reaction.Challenge),
			Headers:    jsonHeader,
		}, nil
	} else if reaction.EventType != eventCallback {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, errUnknownEventType
	}

	// TODO record dynamodb event
	fmt.Println("new event received for reaction", reaction)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    jsonHeader,
	}, nil
}

func main() {
	lambda.Start(handler)
}
