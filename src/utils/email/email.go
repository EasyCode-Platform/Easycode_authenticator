package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmail(emailAddress, content, subject string) error {
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", "2531693734@qq.com")
	//接收人
	m.SetHeader("To", emailAddress)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", subject)
	//内容

	m.SetBody("text/html", content)

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "2531693734@qq.com", "xdybbpdixtccechb")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		return err
	}
	fmt.Printf("send mail to %s success\n", emailAddress)
	return nil
}
