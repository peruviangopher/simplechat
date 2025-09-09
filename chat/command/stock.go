package command

import (
	"fmt"
	"io"
	"net/http"
	"simplechat/setup"
	"strings"
)

const (
	commandPrefix  = "/stock="
	botAPIURL      = "http://host.docker.internal:8081/stock?code=%s&room=%s"
	errorMsgPrefix = "server-api : %s - %s"
)

func EvaluateMsgCmd(cfg *setup.Config, msg string, room string) {
	if !strings.HasPrefix(msg, commandPrefix) {
		fmt.Println("no '/stock=' command")
		return
	}

	client := &http.Client{}

	stockCode := msg[len(commandPrefix):]

	apiURL := fmt.Sprintf(botAPIURL, stockCode, room)
	fmt.Println("calling GET ", apiURL)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add(cfg.BotAPIKeyName(), cfg.BotAPIKey())

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf(errorMsgPrefix, "calling bot api error", err)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode == http.StatusOK {
		fmt.Printf("success call to bot api for stock %s and room %s - %s", stockCode, room, string(bodyBytes))
	}
}
