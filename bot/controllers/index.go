package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"

	"simplechat/setup"
)

const (
	stockAPIURL    = "https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv"
	errorMsgPrefix = "bot-api : %s - %s"
	botName        = "simpleChatBot"
)

type Msg struct {
	From    string `json:"from"`
	Room    string `json:"room"`
	Content string `json:"content"`
}

func StockHandler(cfg *setup.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		go func() {
			pushMsg(cfg, generateMsg(c))
		}()

		c.JSON(http.StatusOK, gin.H{
			"msg": "request added successfully",
		})
	}
}

func pushMsg(cfg *setup.Config, msg *Msg) {
	// 1. Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@host.docker.internal:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// 2. Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// 3. Declare a queue
	q, err := ch.QueueDeclare(
		msg.Room, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// 4. Publish a message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := fmt.Sprintf("<strong>%s - %s:</strong><br>&nbsp;%s", time.Now().Format("15:04:05"), msg.From, msg.Content)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key (queue name)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Println(" [x] Sent %s\n", body)
}

func generateMsg(c *gin.Context) *Msg {
	outMsg := &Msg{
		From:    botName,
		Room:    c.Query("room"),
		Content: "",
	}

	/*
		apiURL := fmt.Sprintf(stockAPIURL, url.QueryEscape(c.Query("code")))
		log.Println("calling GET ", apiURL)

		res, err := http.Get(apiURL)
		if err != nil {
			outMsg.Content = fmt.Sprintf(errorMsgPrefix, "calling stock api error", err)
		}

		if res.StatusCode == http.StatusOK {
			content, err := csv.NewReader(res.Body).ReadAll()
			if err != nil {
				outMsg.Content = fmt.Sprintf(errorMsgPrefix, "csv parse error", err)
			}

			name := strings.ToUpper(content[1][0])
			closeVal := content[1][6]
			if closeVal == "N/D" {
				outMsg.Content = fmt.Sprintf(errorMsgPrefix, "quote not found", name)
			}
			outMsg.Content = fmt.Sprintf("%s quote is $%s per share", name, closeVal)
		}
	*/
	outMsg.Content = fmt.Sprintf("%s quote is $%s per share", "massimo", "5000$")

	return outMsg
}
