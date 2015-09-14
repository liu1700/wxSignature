package models

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/revel/revel"
  "net/http"
)

type JSTicket struct {
  Ticket    string `json:"ticket,omitempty"`
  ExpireIn  int    `json:"expires_in,omitempty"`
  ErrorCode int    `json:"errcode,omitempty"`
}

func GetJSTicket(token string) (*JSTicket, error) {

  api, found := revel.Config.String("wxTicketAPI")
  if !found {
    return nil, errors.New("ticket api not found")
  }

  ticket := new(JSTicket)
  reqUrl := fmt.Sprintf("%saccess_token=%s&type=jsapi", api, token)
  res, err := http.Get(reqUrl)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  err = json.NewDecoder(res.Body).Decode(ticket)

  return ticket, err
}
