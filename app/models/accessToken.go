package models

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/revel/revel"
  "net/http"
)

type AccessToken struct {
  Token     string `json:"access_token,omitempty"`
  ExpireIn  int    `json:"expires_in,omitempty"`
  ErrorCode int    `json:"errcode,omitempty"`
}

func GetAccessToken(appid string, secret string) (*AccessToken, error) {

  api, found := revel.Config.String("wxAccessTokenAPI")
  if !found {
    return nil, errors.New("accesstoken api not found")
  }

  token := new(AccessToken)
  reqUrl := fmt.Sprintf("%sappid=%s&secret=%s", api, appid, secret)
  res, err := http.Get(reqUrl)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  err = json.NewDecoder(res.Body).Decode(token)

  return token, err
}
