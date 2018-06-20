package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	type args struct {
		reaction SlackRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr error
	}{
		{
			"token verification fail",
			args{
				SlackRequest{
					Token:     "HELLO",
					Challenge: "challenge",
				}},
			events.APIGatewayProxyResponse{
				StatusCode: http.StatusUnauthorized,
				Headers:    jsonHeader,
			},
			errTokenVerification,
		},
		{
			"slack event success",
			args{
				SlackRequest{
					Token:     "VALUE",
					EventType: eventCallback,
				}},
			events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Headers:    jsonHeader,
			},
			nil,
		},
		{
			"unknown event",
			args{
				SlackRequest{
					Token:     "VALUE",
					EventType: "unknown event",
				}},
			events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers:    jsonHeader,
			},
			errUnknownEventType,
		},
		{
			"challenge success",
			args{
				SlackRequest{
					Token:     "VALUE",
					Challenge: "challenge",
					EventType: urlVerification,
				}},
			events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Headers:    jsonHeader,
				Body:       (`{"challenge":"challenge"}`),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler(tt.args.reaction)
			if (err != nil) && (err != tt.wantErr) {
				t.Errorf("handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
