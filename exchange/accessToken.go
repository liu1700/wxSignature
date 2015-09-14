package exchange

import (
  "encoding/json"
  "fmt"
  "net/http"
)

const (
  AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&"
)

type AccessToken struct {
  Token     string `json:"access_token,omitempty"`
  ExpireIn  int    `json:"expires_in,omitempty"`
  ErrorCode int    `json:"errcode,omitempty"`
}

func GetAccessToken(appid string, secret string) (*AccessToken, error) {
  token := new(AccessToken)
  reqUrl := fmt.Sprintf("%sappid=%s&secret=%s", AccessTokenUrl, appid, secret)
  res, err := http.Get(reqUrl)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  err = json.NewDecoder(res.Body).Decode(token)

  return token, err
}
