package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

var NameEnqury = "http://52.168.85.231/thirdpartyapiservice/apiservice/Transfer/NameEnquiry"

func GetNameEnquiry(authToken, accountNumber, bankCode string) (map[string]interface{}, error) {
	resp, err := http.PostForm(NameEnqury, url.Values{
		"AccountNumber": {accountNumber},
		"BankCode":      {bankCode},
		"Token":         {authToken},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data["Status"] != "success" {
		return nil, errors.New("unable to get account details")
	}

	return data, nil
}
