package models

var global Tokens

type Tokens struct {
  G_AccessToken *AccessToken
  G_JSTicket    *JSTicket
  ExpiredAt     int
}

func GetGlobalTokens() *Tokens {
  return &global
}

func Set(t interface{}) {
  switch t.(type) {
  case *AccessToken:
    global.G_AccessToken = t.(*AccessToken)
    break
  case *JSTicket:
    global.G_JSTicket = t.(*JSTicket)
    break
  default:
    break
  }
}
