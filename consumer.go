package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient:= range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		formatedMsg := fmt.Sprintf("To: %s\r\nSubject:Hi %s, Test Email\r\n\r\n%s\r\n", recipient.Email, recipient.Name, "Hello World")
		msg:= []byte(formatedMsg)

		err := smtp.SendMail(smtpHost + ":" + smtpPort, nil, "dipendramule@gmail.com", []string{recipient.Email}, msg)
		if err != nil {
			log.Fatal(err)
		}

		// Delay for bypass rate limiting

		time.Sleep(50*time.Microsecond)

		fmt.Printf("Worker %d: Sending email to %s \n", id, recipient.Email)
	}
}