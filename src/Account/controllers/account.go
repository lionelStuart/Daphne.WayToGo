package controllers

import (
	"Common/Email"
	"account/models"
	"fmt"
	"github.com/satori/go.uuid"
)

type AccountController struct {
	baseController
}

func (a *AccountController) Login() {
	if a.Ctx.Request.Method == "POST" {
		fmt.Println("post in methods ")

		account := a.GetString("account")
		password := a.GetString("password")
		remember := a.GetString("remember")
		fmt.Println("get param", account, password, remember)
		if account == "" || password == "" {
			a.Redirect("login", 302)
		}

		acc := models.Account{}
		acc.Uname = account
		acc.QueryByName()
		acc.Update()
	}
	a.TplName = "index.tpl"
}

func (a *AccountController) Logout() {

}

func (a *AccountController) Register() {
	fmt.Println("register call")
	if a.Ctx.Request.Method == "POST" {
		username := a.GetString("username")
		// password := a.GetString("password")
		email := a.GetString("email")

		verification, _ := uuid.NewV1()
		verificationUrl := fmt.Sprintf("http://localhost:8080/account/verification?username=%s&verification=%s", username, verification)

		body := fmt.Sprintf("<h1>auth your account </h1> <p>check your account at:  %s </p>", verificationUrl)

		m := Email.MailContent{
			Address: email,
			Name:    email,
			Title:   "Verify Your Account",
			Body:    body,
		}
		if err := Email.SendMail(&m); err != nil {
			fmt.Println("failure send mail ", err)
		}
		session, _ := CookieStore.Get(a.Ctx.Request, "ver")
		session.Values["username"] = username
		session.Values["verification"] = verification.String()

		if err := session.Save(a.Ctx.Request, a.Ctx.ResponseWriter); err != nil {
			fmt.Println("failure error", err)
		}
		a.Redirect("login", 302)
	}

	a.TplName = "register.html"
}

func (a *AccountController) Verification() {
	session, _ := CookieStore.Get(a.Ctx.Request, "ver")
	fmt.Println("123")
	fmt.Println("cookie ", session.Values["username"], session.Values["verification"])

	username := a.GetString("username")
	verification := a.GetString("verification")

	if username != session.Values["username"] || verification != session.Values["verification"] {
		fmt.Println("verification failure")
	} else {
		session.AddFlash("verification success")
		session.Save(a.Ctx.Request, a.Ctx.ResponseWriter)
		//a.Redirect("login",302)
		fmt.Println("verification success")
	}

	a.TplName = "index.tpl"

}
