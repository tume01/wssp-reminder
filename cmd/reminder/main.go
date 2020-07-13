package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/sfreiberg/gotwilio"
)

type TwilioConfig struct {
	AccountID string `env:"TWILIO_ACCOUNT,required"`
	SecretKey string `env:"TWILIO_KEY,required"`
	FromNumber string `env:"TWILIO_FROM,required"`
}

type Notifier struct {
	Client *gotwilio.Twilio
	FromNumber string
}

type AppConfig struct {
	MediaURL string `env:"MEDIA_URL,required"`
	ToNumber string `env:"TO_NUMBER,required"`
}

func (n Notifier) notify(message string, to string) error {
	_, _, err := n.Client.SendWhatsApp(n.FromNumber, to, message, "", n.Client.AccountSid)
	if err != nil {
		return err
	}
	return nil
}

func (n Notifier) notifyMedia(message string, mediaURLs []string, to string) error {
	_, _, err := n.Client.SendWhatsAppMedia(n.FromNumber, to, message, mediaURLs, "", n.Client.AccountSid)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	tConfig := TwilioConfig{}
	err := env.Parse(&tConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	app := AppConfig{}
	err = env.Parse(&app)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tClient := gotwilio.NewTwilioClient(tConfig.AccountID, tConfig.SecretKey)
	n := Notifier{
		Client: tClient,
		FromNumber: tConfig.FromNumber,
	}
	mediaURLs := []string{app.MediaURL}
	err = n.notifyMedia("", mediaURLs, app.ToNumber)
	if err != nil {
		fmt.Println(err.Error())
	}
}
