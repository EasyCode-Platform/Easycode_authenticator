package model

import (
	"fmt"

	"github.com/EasyCode/Easycode-authenticator-backend/src/utils/email"
)

const BASE_HTML = `<head>
<base target="_blank" />
<style type="text/css">::-webkit-scrollbar{ display: none; }</style>
<style id="cloudAttachStyle" type="text/css">#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>
<style id="blockquoteStyle" type="text/css">blockquote{display:none;}</style>
<style type="text/css">
	body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}
	td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}
	pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%%}
	th,td{font-family:arial,verdana,sans-serif;line-height:1.666}
	img{ border:0}
	header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}
	blockquote{margin-right:0px}
</style>
</head>
<body tabindex="0" role="listitem">
<table width="700" border="0" align="center" cellspacing="0" style="width:700px;">
<tbody>
<tr>
	<td>
		<div style="width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;">
			<table border="0" cellpadding="0" cellspacing="0" width="700" height="39" style="font:12px Tahoma, Arial, 宋体;">
				<tbody><tr><td width="210"></td></tr></tbody>
			</table>
		</div>
		<div style="width:680px;padding:0 10px;margin:0 auto;">
			<div style="line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;">
				<strong style="display:block;margin-bottom:15px;">尊敬的用户：<span style="color:#f60;font-size: 16px;"></span>您好！</strong>
				<strong style="display:block;margin-bottom:15px;">
					您正在进行<span style="color: red">%s</span>操作，请在验证码输入框中输入：<span style="color:#f60;font-size: 24px">%s</span>，以完成操作。
				</strong>
			</div>
			<div style="margin-bottom:30px;">
				<small style="display:block;margin-bottom:20px;font-size:12px;">
					<p style="color:#747474;">
						注意：此操作可能会修改您的密码、登录邮箱或绑定手机。如非本人操作，请及时登录并修改密码以保证帐户安全
						<br>（工作人员不会向你索取此验证码，请勿泄漏！)
					</p>
				</small>
			</div>
		</div>
		<div style="width:700px;margin:0 auto;">
			<div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
				<p>此为系统邮件，请勿回复<br>
					请保管好您的邮箱，避免账号被他人盗用
				</p>
				<p>Easycode Platform</p>
			</div>
		</div>
	</td>
</tr>
</tbody>
</table>
</body>
`

func SendSubscriptionEmail(emailAddress string) error {

	// client := resty.New()
	// resp, err := client.R().
	// 	SetBody(map[string]string{"emailAddress": emailAddress}).
	// 	Post(EMAIL_DELIVER_BASEURL + EMAIL_DELIVER_USAGE_SUBSCRIBE)
	// if resp.StatusCode() != http.StatusOK || err != nil {
	// 	return errors.New("failed to send subscription emailAddress")
	// }
	// fmt.Printf("response: %+v, err: %+v", resp, err)
	return nil
}

func SendVerificationEmail(emailAddress, code, usage string) error {
	content := fmt.Sprintf(BASE_HTML, usage, code)
	return email.SendEmail(emailAddress, content, "验证您的邮箱")
}

func SendInviteEmail(m *EmailInviteMessage) error {
	fmt.Printf("%v\n", m)

	return email.SendEmail(m.Email, m.JoinLink, "邀请！")
	// client := resty.New()
	// resp, err := client.R().
	// 	SetBody(payload).
	// 	Post(EMAIL_DELIVER_BASEURL + EMAIL_DELIVER_INVITE_EMAIL)
	// if resp.StatusCode() != http.StatusOK || err != nil {
	// 	fmt.Printf("response: %+v, err: %+v", resp, err)
	// 	return errors.New("failed to send invite emailAddress")
	// }
	// fmt.Printf("response: %+v, err: %+v", resp, err)
}

func SendShareAppEmail(m *EmailShareAppMessage) error {
	fmt.Printf("%v\n", m)

	return email.SendEmail(m.Email, m.AppLink, "共享App邀请！")
	// client := resty.New()
	// resp, err := client.R().
	// 	SetBody(payload).
	// 	Post(EMAIL_DELIVER_BASEURL + EMAIL_DELIVER_INVITE_EMAIL)
	// if resp.StatusCode() != http.StatusOK || err != nil {
	// 	fmt.Printf("response: %+v, err: %+v", resp, err)
	// 	return errors.New("failed to send invite emailAddress")
	// }
	// fmt.Printf("response: %+v, err: %+v", resp, err)
	// fmt.Printf("%v\n", m)
	// payload := m.Export()
	// client := resty.New()
	// resp, err := client.R().
	// 	SetBody(payload).
	// 	Post(EMAIL_DELIVER_BASEURL + EMAIL_DELIVER_SHARE_APP_EMAIL)
	// if resp.StatusCode() != http.StatusOK || err != nil {
	// 	fmt.Printf("response: %+v, err: %+v", resp, err)
	// 	return errors.New("failed to send share app emailAddress")
	// }
	// fmt.Printf("response: %+v, err: %+v", resp, err)
	// return nil
}
