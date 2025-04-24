package mailgun

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sync"
	"time"
)

// class 成員變數
type mailConfig struct {
	Domain     string `json:"domain"`
	PrivateKey string `json:"private_key" `
	Sender     string `json:"sender"`
}

// constructor 用來初始化
func NewMailConfig(domain string, privateKey string, sender string) *mailConfig {
	return &mailConfig{
		Domain:     domain,
		PrivateKey: privateKey,
		Sender:     sender,
	}
}

// Method
func (m *mailConfig) CheckConfig() (config *mailConfig) {
	var ctx *gin.Context
	var yourDomain string = viper.GetString("mailgun.domain")
	var privateAPIKey string = viper.GetString("mailgun.privateAPIKey")
	var sender string = viper.GetString("mailgun.sender")
	if yourDomain == "" || privateAPIKey == "" || sender == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Mailgun config error!",
		})
		return
	}
	config = &mailConfig{
		Domain:     yourDomain,
		PrivateKey: privateAPIKey,
		Sender:     sender,
	}
	return
}

func (m *mailConfig) SendMail(subject string, emails []string, body string) {
	var yourDomain string
	var privateAPIKey string
	var sender string
	config := m.CheckConfig()
	if config != nil {
		yourDomain = config.Domain
		privateAPIKey = config.PrivateKey
		sender = config.Sender
	}
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)
	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	mailsCh := make(chan string, len(emails))
	workerSize := 2
	wg := &sync.WaitGroup{}
	for _, email := range emails {
		mailsCh <- email
	}
	close(mailsCh)
	//recipient := emails
	// The message object allows you to add attachments and Bcc recipients

	for i := 0; i < workerSize; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range mailsCh {
				message := mg.NewMessage(sender, subject, body, job)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				defer cancel()

				// Send the message with a 10 second timeout
				resp, id, err := mg.Send(ctx, message)
				if err != nil {
					log.Printf("Error sending email #%d:%s", id, err)
				}

				fmt.Printf("ID: %s Resp: %s\n", id, resp)
			}

		}(i)
	}
	wg.Wait()
	return
}
