package function

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// Slacker exported
func Slacker() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3897246883060-3895225666610-kRd4FPaN789IYq9g4rzWWnmH")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03SDJTKE68-3897641955188-0780ac99fc1d2f9c1e4e661c265997a3c3e88195b79ca33fff7165438208ca01")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	
	bot.Command("my yob is year <year>", &slacker.CommandDefinition{
		Description : "yob calculator",
		Examples : []string{"my yob is 2020"},
		Handler : func ( botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil { fmt.Println("Error") }
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})
	

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err :=  bot.Listen(ctx)
	if err != nil { log.Fatal(err) }
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel {
		fmt.Println("command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}