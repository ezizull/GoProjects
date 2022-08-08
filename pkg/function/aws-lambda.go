package function

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name string `json:"what is yout name?"`
	Age 	int `json:"How old are you"`
}

type myResponse struct {
	Message string `json:"Answer:"`
}

// BasicLamda exported
func BasicLamda() {
	lambda.Start(handleLambdaEvent)
}

func handleLambdaEvent(event myEvent)(myResponse, error){
	return myResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}