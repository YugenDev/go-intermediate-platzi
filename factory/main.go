package main

import "fmt"

// SMS EMAIL

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

type SMSNotificationSender struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification SMS")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending via Email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func GetNotificationFactory(notificationType string ) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification type")
}

func SendNotification (f INotificationFactory) {
	f.SendNotification()
}

func GetMethod (f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main(){
	smsFactory, _ := GetNotificationFactory("SMS")
	emailFactory, _ := GetNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	GetMethod(smsFactory)
	GetMethod(emailFactory)
}
