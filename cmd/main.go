package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jasonlvhit/gocron"
	"github.com/nafisfaysal/awsautomation/iam"
	"github.com/nafisfaysal/awsautomation/ses"
	"strings"
)

func main() {
	lambda.Start(SendAlertHandler)
}

// cron jobs to send alert to the user.
func SendAlertHandler() {
	gocron.Every(1).Day().At("06:30").Do(SendAlertEmail)
	gocron.Every(1).Day().At("17:30").Do(SendAlertEmail)
}

func SendAlertEmail() {
	emailSubject := "Credential Expiration Notice From AWS Account"
	CharSet := "UTF-8"
	region := "us-east-1"

	csvDts := iam.ReadCsvFile("file_path_of_your_data")
	for i := 0; i < len(csvDts); i++ {

		userName := csvDts[i][0]
		if strings.TrimSpace(userName) == "userName" {
			continue
		}

		accesKeyID := csvDts[i][1]
		if strings.TrimSpace(accesKeyID) == "accessKeyId" {
			continue
		}

		userEmail := csvDts[i][3]
		if strings.TrimSpace(userName) == "email" {
			continue
		}

		timeLeft := iam.TimeLeftOfUserAgeExpedited(userName)
		switch {
		// 20 days to minute 28800
		case timeLeft < 28800:
			ses.SendEmail("awsauthorsender@sender.com", userEmail, emailSubject, "\n\tYour AccessKey ID left 20 days post expiration. It will be deactivated very soon", CharSet, region)
			// 15 days to minute
		case timeLeft < 21600:
			ses.SendEmail("awsauthorsender@sender.com", userEmail, emailSubject, "\n\tYour AccessKey ID left 15 days post expiration. It will be deactivated very soon", CharSet, region)
		// 10 days to minute 14400
		case timeLeft < 14400:
			ses.SendEmail("awsauthorsender@sender.com", userEmail, emailSubject, "\n\tYour AccessKey ID left 10 days post expiration. It will be deactivated very soon", CharSet, region)

		case timeLeft <= 0:
			ses.SendEmail("awsauthorsender@sender.com", userEmail, emailSubject, "\n\tYour AccessKey ID post expiration. It has been deactivated", CharSet, region)
			iam.DisableUsersKey(userName, accesKeyID)
		}
	}
}
