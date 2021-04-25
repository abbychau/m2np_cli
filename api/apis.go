package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type mss map[string]interface{}

func (x mss) ToJson() *bytes.Buffer {
	data, err := json.Marshal(x)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewBuffer(data)
}

func Login(user, pwd string) (string, error){
	resp, err := http.Post("https://m2np.com/api/login", "application/json", mss{"email": user, "password": pwd}.ToJson())
	if err != nil{
		return "", err
	}
	var res = map[string]string{}
	_ = json.NewDecoder(resp.Body).Decode(&res)
	return res["token"], nil
}