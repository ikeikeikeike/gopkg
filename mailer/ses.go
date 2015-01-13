package mailer

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/exp/ses"
)

type SesMailer struct {
	M      *ses.SES
	E      *ses.Email
	auth   aws.Auth
	region aws.Region
}

func NewSesMailer() *SesMailer {
	a := aws.Auth{
		AccessKey: beego.AppConfig.String("AwsAccessKey"),
		SecretKey: beego.AppConfig.String("AwsSecretKey"),
	}
	r := aws.Region{
		Name:        "us-east-1",
		SESEndpoint: "https://email.us-east-1.amazonaws.com",
	}

	mailer := &SesMailer{M: ses.NewSES(a, r), auth: a, region: r}
	mailer.SetDefaultInfo()
	return mailer
}

func (s *SesMailer) SetAuth(auth aws.Auth) *SesMailer {
	s.auth = auth
	s.M = ses.NewSES(auth, s.region)
	return s
}

func (s *SesMailer) SetRegion(region aws.Region) *SesMailer {
	s.region = region
	s.M = ses.NewSES(s.auth, region)
	return s
}

func (s *SesMailer) SetDefaultInfo() *SesMailer {
	s.E = ses.NewEmail()
	s.E.SetSource(fmt.Sprintf("no-reply@%s", beego.AppConfig.String("httpaddr")))
	return s
}

func (s *SesMailer) SendMail(to, subject, htmlbody string) error {
	s.E.AddTo(to)
	s.E.SetSubject(subject)
	s.E.SetBodyHtml(htmlbody)

	return s.M.SendEmail(s.E)
}

// Async Send mail message
func (s *SesMailer) SendAsync(to, subject, htmlbody string) {
	go func() {
		if err := s.SendMail(to, subject, htmlbody); err != nil {
			beego.Error(fmt.Sprintf("Async send email not send emails: %s", err))
		}
	}()
}
