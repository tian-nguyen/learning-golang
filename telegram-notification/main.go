package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

var (
	ChatID   = "-4035316275"
	BotToken = "7852293021:AAF3Bx118TvSPAI7QzcGYvlU03VsNk-6IfQ"
	Text     = "Tung Nguyen demo telegram notification"
)

// curl https://api.telegram.org/bot7852293021:AAF3Bx118TvSPAI7QzcGYvlU03VsNk-6IfQ/getUpdates

func main() {
	fmt.Println("main")
	jsonStr := []byte(fmt.Sprintf(`{"chat_id": "%s", "text": "%s"}`, ChatID, Text))

	resp, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BotToken), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Default().Println("err:", err)
		return
	}
	defer resp.Body.Close()
	log.Default().Println("successfully")

}
