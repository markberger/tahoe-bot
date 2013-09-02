package main

import (
	"encoding/json"
	"fmt"
	"github.com/markberger/tally"
	"io/ioutil"
)

type Settings struct {
	Password string
}

func main() {
	var s Settings
	f, err := ioutil.ReadFile("./password.json")
	if err != nil {
		fmt.Printf("Error loading password.json: %v\n", err)
		return
	}
	err = json.Unmarshal(f, &s)
	if err != nil {
		fmt.Printf("Error unmarshalling password.json: %v\n", err)
		return
	}

	tally.InitLogging()
	bot := tally.NewBot()
	bot.AddAction(`:\x01ACTION (\w*) tahoe-bot\x01`, ParseAction, RespondToAction)
	bot.Connect()
	bot.PrivateMsg("NickServ", "identify "+s.Password)
	bot.Run()
}
