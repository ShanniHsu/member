package mailgun

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"
	"log"
	"time"
)

func SendMail(subject string, email string, body string) (err error) {
	var yourDomain string = viper.GetString("mailgun.domain")
	var privateAPIKey string = viper.GetString("mailgun.privateAPIKey")
	var sender string = viper.GetString("mailgun.sender")
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)
	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	recipient := email
	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return
}
