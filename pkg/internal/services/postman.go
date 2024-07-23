package services

import (
	"context"
	"crypto/tls"
	"firebase.google.com/go/messaging"
	"fmt"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"github.com/jordan-wright/email"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"github.com/sideshow/apns2"
	payload2 "github.com/sideshow/apns2/payload"
	"github.com/spf13/viper"
	"net/smtp"
	"net/textproto"
	"time"
)

func DealDeliveryTask(task any) {
	switch tk := task.(type) {
	case *proto.DeliverEmailRequest:
		if tk.GetEmail().HtmlBody != nil {
			_ = SendMailHTML(tk.GetTo(), tk.GetEmail().GetSubject(), tk.GetEmail().GetHtmlBody())
		} else {
			_ = SendMail(tk.GetTo(), tk.GetEmail().GetSubject(), tk.GetEmail().GetTextBody())
		}
	case *proto.DeliverNotificationRequest:
		switch tk.GetProvider() {
		case "firebase":
			_ = PushFirebaseNotify(tk.GetDeviceToken(), tk.GetNotify())
		case "apple":
			_ = PushAppleNotify(tk.GetDeviceToken(), tk.GetNotify())
		}
	}
}

func PushFirebaseNotify(token string, in *proto.NotifyRequest) error {
	if ExtFire == nil {
		return fmt.Errorf("firebase push notification is unavailable")
	}

	start := time.Now()

	ctx := context.Background()
	client, err := ExtFire.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("failed to create firebase client")
	}

	var image string
	if in.Picture != nil {
		image = *in.Picture
	}
	var subtitle string
	if in.Subtitle != nil {
		subtitle = "\n" + *in.Subtitle
	}
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title:    in.Title,
			Body:     subtitle + in.Body,
			ImageURL: image,
		},
		Token: token,
	}

	if response, err := client.Send(ctx, message); err != nil {
		log.Warn().Err(err).Msg("An error occurred when notify subscriber via FCM...")
	} else {
		log.Debug().
			Dur("elapsed", time.Since(start)).
			Str("response", response).
			Msg("Push a notify via firebase")
	}

	return nil
}

func PushAppleNotify(token string, in *proto.NotifyRequest) error {
	if ExtAPNS == nil {
		return fmt.Errorf("apple push notification is unavailable")
	}

	start := time.Now()

	var metadata map[string]any
	_ = jsoniter.Unmarshal(in.GetMetadata(), &metadata)

	data := payload2.
		NewPayload().
		AlertTitle(in.GetTitle()).
		AlertBody(in.GetBody()).
		Category(in.GetTopic()).
		Custom("metadata", metadata).
		Sound("default").
		MutableContent()
	if in.Avatar != nil {
		data = data.Custom("avatar", *in.Avatar)
	}
	if in.Picture != nil {
		data = data.Custom("picture", *in.Picture)
	}
	rawData, err := data.MarshalJSON()
	if err != nil {
		log.Warn().Err(err).Msg("An error occurred when preparing to notify subscriber via APNs...")
	}
	payload := &apns2.Notification{
		DeviceToken: token,
		Topic:       viper.GetString("apns_topic"),
		Payload:     rawData,
	}

	if resp, err := ExtAPNS.Push(payload); err != nil {
		log.Warn().Err(err).Msg("An error occurred when notify subscriber via APNs...")
	} else {
		log.Debug().
			Dur("elapsed", time.Since(start)).
			Str("reason", resp.Reason).
			Int("status", resp.StatusCode).
			Msg("Push a notify via firebase")
	}

	return nil
}

func SendMail(target string, subject string, content string) error {
	mail := &email.Email{
		To:      []string{target},
		From:    viper.GetString("mailer.name"),
		Subject: subject,
		Text:    []byte(content),
		Headers: textproto.MIMEHeader{},
	}
	return mail.SendWithTLS(
		fmt.Sprintf("%s:%d", viper.GetString("mailer.smtp_host"), viper.GetInt("mailer.smtp_port")),
		smtp.PlainAuth(
			"",
			viper.GetString("mailer.username"),
			viper.GetString("mailer.password"),
			viper.GetString("mailer.smtp_host"),
		),
		&tls.Config{ServerName: viper.GetString("mailer.smtp_host")},
	)
}

func SendMailHTML(target string, subject string, content string) error {
	mail := &email.Email{
		To:      []string{target},
		From:    viper.GetString("mailer.name"),
		Subject: subject,
		HTML:    []byte(content),
		Headers: textproto.MIMEHeader{},
	}
	return mail.SendWithTLS(
		fmt.Sprintf("%s:%d", viper.GetString("mailer.smtp_host"), viper.GetInt("mailer.smtp_port")),
		smtp.PlainAuth(
			"",
			viper.GetString("mailer.username"),
			viper.GetString("mailer.password"),
			viper.GetString("mailer.smtp_host"),
		),
		&tls.Config{ServerName: viper.GetString("mailer.smtp_host")},
	)
}
