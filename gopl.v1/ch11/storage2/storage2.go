package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	sender   = "ik@ikaroskun.me"
	password = "password"
	hostname = "smtp.ikaroskun.me"
	template = `Warning: you are using %d bytes if storage, %d%% of your quota.`
)

func bytesInUse(username string) int64 {
	return 0
}

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendEmail(%s) failed: %s", username, err)
	}
}

func checkQuota(username string) {
	used := bytesInUse(username)

	const quota = 1000 * 1000 * 1000
	percent := 100 * used / quota

	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
