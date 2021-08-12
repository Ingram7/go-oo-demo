package controller

import (
	"fmt"
	app "go-oo-demo/internal/pkg/app"
)

type Captcha struct {

}

func NewCaptcha() *Captcha {
	captcha := new(Captcha)
	return captcha
}

func (ctr *Captcha) Send(c *app.Context) (Data, error) {

	fmt.Println(c.GetQuery("mobile"))
	fmt.Println(c.Query("mobile"))

	//captchaQuery := new(query.CaptchaSend)
	//if err := c.Bind(captchaQuery); err != nil {
	//	return nil, err
	//}
	//
	//captcha := new(domain.Captcha)
	//captcha.Generate(captchaQuery.GetMobile())
	//
	//captcha.Invalid()
	//if err := ctr.repo.Save(captcha); err != nil {
	//	return nil, err
	//}
	//
	//errSend := ctr.send(captcha)
	//if err := ctr.repo.Save(captcha); err != nil {
	//	return nil, err
	//}
	//
	//if errSend != nil {
	//	return nil, errSend
	//}

	return nil, nil
}