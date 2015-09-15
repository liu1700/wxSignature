package controllers

import (
  "crypto/sha1"
  "encoding/hex"
  "errors"
  "fmt"
  "github.com/revel/revel"
  "github.com/satori/go.uuid"
  "strconv"
  "time"
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

  url, found := revel.Config.String("wx.url")
  if !found {
    return c.RenderJson(errors.New("url not found"))
  }

  getAccessToken(appid, secret)
  accessToken := models.GetGlobalTokens().G_AccessToken
  getTicket(accessToken.Token)

  uid := uuid.NewV4()
  noncestr := uid.String()
  timestamp := strconv.FormatInt(time.Now().Unix(), 10)
  ticket := models.GetGlobalTokens().G_JSTicket.Ticket

  string1 := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s",
    ticket, noncestr, timestamp, url)

  h := sha1.New()
  h.Write([]byte(string1))
  bs := h.Sum(nil)
  revel.INFO.Printf("%x", bs)
  signature := hex.EncodeToString(bs)

  return c.RenderJson(map[string]string{
    "noncestr":     noncestr,
    "jsapi_ticket": ticket,
    "timestamp":    timestamp,
    "url":          url,
    "signature":    signature})
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
