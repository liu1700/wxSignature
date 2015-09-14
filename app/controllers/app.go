package controllers

import (
  "errors"
  "fmt"
  "github.com/revel/revel"
  "strconv"
  "wxSignature/app/models"
)

type App struct {
  *revel.Controller
}

func (c App) Index() revel.Result {
  return c.Render()
}

func (c App) Generate() revel.Result {

  appid, found := revel.Config.String("wx.appid")
  if !found {
    return c.RenderJson(errors.New("appid not found"))
  }

  secret, found := revel.Config.String("wx.secret")
  if !found {
    return c.RenderJson(errors.New("secret not found"))
  }

  getAccessToken(appid, secret)
  accessToken := models.GetGlobalTokens().G_AccessToken
  getTicket(accessToken.Token)
  revel.INFO.Println(models.GetGlobalTokens().G_JSTicket.Ticket)

  return c.RenderJson(nil)
}

// get Access token and store it global
func getAccessToken(appid string, secret string) {
  token, err := models.GetAccessToken(appid, secret)
  if err != nil {
    revel.ERROR.Println(err.Error())
    return
  }
  if token.ErrorCode != 0 {
    revel.ERROR.Println(fmt.Sprintf("Error Code: %s", strconv.Itoa(token.ErrorCode)))
  }
  models.Set(token)
}

// get jsTicktet with access token and store it in global
func getTicket(accessToken string) {
  ticket, err := models.GetJSTicket(accessToken)
  if err != nil {
    revel.ERROR.Println(err.Error())
    return
  }
  if ticket.ErrorCode != 0 {
    revel.ERROR.Println(fmt.Sprintf("Error Code: %s", strconv.Itoa(ticket.ErrorCode)))
  }
  models.Set(ticket)
}
