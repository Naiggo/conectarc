package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	invalidTokenMsg    string = "Invalid token"
	errorDecodeMsg     string = "Cant decode token part: "
	errorJsonDecodeMsg string = "Cant decode JSON struct: "
	expiredTokenMsg    string = "Expired token!"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_Use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func TokenValidation(token string) (bool, error, string) {
	tokenParts := strings.Split(token, ".")

	if len(tokenParts) != 3 {
		fmt.Println(invalidTokenMsg)
		return false, nil, invalidTokenMsg
	}

	userInfo, err := base64.StdEncoding.DecodeString(tokenParts[1]) //payload decode
	if err != nil {
		fmt.Println(errorDecodeMsg, err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println(errorJsonDecodeMsg, err.Error())
		return false, err, err.Error()
	}

	timeNow := time.Now()
	timeExpiration := time.Unix(int64(tkj.Exp), 0)

	if timeExpiration.Before(timeNow) {
		fmt.Println("Token expiration date = " + timeExpiration.String())
		fmt.Println(expiredTokenMsg)
		return false, err, expiredTokenMsg
	}

	return true, nil, string(tkj.Username)
}
