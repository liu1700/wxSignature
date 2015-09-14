package exchange

import (
  "encoding/json"
  "fmt"
  "net/http"
)

const (
  TicketUrl = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?"
)

type JSTicket struct {
  Ticket    string `json:"ticket,omitempty"`
  ExpireIn  int    `json:"expires_in,omitempty"`
  ErrorCode int    `json:"errcode,omitempty"`
}

func GetJSTicket(token string) (*JSTicket, error) {
  ticket := new(JSTicket)
  reqUrl := fmt.Sprintf("%saccess_token=%s&type=jsapi", TicketUrl, token)
  res, err := http.Get(reqUrl)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  err = json.NewDecoder(res.Body).Decode(ticket)

  return ticket, err
}
