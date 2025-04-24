package mailgun

import (
	"context"
	"errors"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

// class 成員變數
type MailConfig struct {
	Domain     string `json:"domain"`
	PrivateKey string `json:"private_key" `
	Sender     string `json:"sender"`
}

// constructor 用來初始化
func NewMailConfig() *MailConfig {
	return &MailConfig{
		Domain:     viper.GetString("mailgun.domain"),
		PrivateKey: viper.GetString("mailgun.privateAPIKey"),
		Sender:     viper.GetString("mailgun.sender"),
	}
}

type Mail interface {
	SendMail(subject string, emails []string, body string) (err error)
}

// Method
func (m *MailConfig) CheckConfig() (err error) {
	if m.Domain == "" || m.PrivateKey == "" || m.Sender == "" {
		err = errors.New("Mailgun config error!")
		return
	}
	return
}

func (m *MailConfig) SendMail(subject string, emails []string, body string) (err error) {
	err = m.CheckConfig()
	if err != nil {
		return
	}
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(m.Domain, m.PrivateKey)
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
		go func(size int) {
			defer wg.Done()
			for job := range mailsCh {
				message := mg.NewMessage(m.Sender, subject, body, job)
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
