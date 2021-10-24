package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	CorpId string
	AppToken string
	EncodingAesKey string
}
