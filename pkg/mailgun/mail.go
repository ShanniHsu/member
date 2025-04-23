package mailgun

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

func SendMail(subject string, emails []string, body string) (err error) {
	var yourDomain string = viper.GetString("mailgun.domain")
	var privateAPIKey string = viper.GetString("mailgun.privateAPIKey")
	var sender string = viper.GetString("mailgun.sender")
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
