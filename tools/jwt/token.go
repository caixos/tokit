package jwt

import (
	"caixin.app/caixos/tokit/config"
	"caixin.app/caixos/tokit/constant"
	"caixin.app/caixos/tokit/tools/util"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"errors"
)

type Token struct {
	Claims *Claims
	config *config.TokenConfig
}

func New() *Token {
	token := &Token{
		Claims: &Claims{},
	}
	token.config = config.LoadTokenConfig()
	return token
}

func (s *Token) SetId(id string) *Token {
	s.Claims.Id = id
	return s
}

func (s *Token) SetName(name string) *Token {
	s.Claims.Name = name
	return s
}

func (s *Token) SetRole(role string) *Token {
	s.Claims.Role = role
	return s
}

func (s *Token) GetToken() string {
	s.Claims.Iat = time.Now().Unix()
	s.Claims.Exp = s.getExpTime()
	jsonClaim, err := json.Marshal(s.Claims)
	if err != nil {
		panic(err)
	}
	payload := base64.StdEncoding.EncodeToString(jsonClaim)
	sign := s.getSign(s.Claims)
	ret := string(payload) + "." + sign
	return ret
}

func (s *Token) VerifyToken(sign string) (*Claims, error) {
	m := strings.Split(sign, ".")
	if len(m) < 1 {
		return nil, errors.New(constant.ErrTokenFmt)
	}
	jsonClaim, decodeErr := base64.StdEncoding.DecodeString(m[0])
	if decodeErr != nil {
		return nil, decodeErr
	}
	claims := &Claims{}
	jsonErr := json.Unmarshal(jsonClaim, claims)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if claims.Exp < time.Now().Unix() {
		return nil, errors.New(constant.ErrTokenExp)
	}

	if m[1] != s.getSign(claims) {
		return nil, errors.New(constant.ErrTokenSign)
	}
	return claims, nil
}
func (s *Token) getExpTime() int64 {
	period := s.config.Exp
	return s.Claims.Iat + period
}
func (s *Token) getSign(claims *Claims) string {
	key := s.config.Key
	keyPlain := claims.Id + strconv.Itoa(int(claims.Iat)) + key
	return util.Md5(keyPlain)

}

