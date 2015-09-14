package config

type Config struct {
  Appid  string
  Secret string
  Url    string
}

var conf Config

func init() {
  conf.Appid = "wxd98888751036c960"
  conf.Secret = "6aa0925e117874335068d95c37088029"
  conf.Url = "http://test.weixin.bigertech.com"
}

func Get() *Config {
  return &conf
}
