package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)
	
	go loadRecipient("./assets/email.csv", recipientChannel)

	var wg sync.WaitGroup
	workerCount := 5

	for i := 0; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()
}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("assets/email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	
	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}