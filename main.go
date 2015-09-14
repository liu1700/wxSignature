package main

import (
  "./config"
  "./exchange"
  "fmt"
  "strconv"
)

func main() {
  wx := config.Get()
  getAccessToken(wx.Appid, wx.Secret)
  accessToken := exchange.GetGlobalTokens().G_AccessToken
  getTicket(accessToken.Token)
  fmt.Println(exchange.GetGlobalTokens().G_JSTicket.Ticket)
}

func getAccessToken(appid string, secret string) {
  token, err := exchange.GetAccessToken(appid, secret)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  if token.ErrorCode != 0 {
    fmt.Println(fmt.Sprintf("Error Code: %s", strconv.Itoa(token.ErrorCode)))
  }
  exchange.Set(token)
}

func getTicket(accessToken string) {
  ticket, err := exchange.GetJSTicket(accessToken)
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  if ticket.ErrorCode != 0 {
    fmt.Println(fmt.Sprintf("Error Code: %s", strconv.Itoa(ticket.ErrorCode)))
  }
  exchange.Set(ticket)
}
