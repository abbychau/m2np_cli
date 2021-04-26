package api

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func Login(user, pwd string) (string, error) {
	resp, err := http.Post("https://m2np.com/api/login", "application/json", mss{"email": user, "password": pwd}.ToJson())
	if err != nil {
		return "", err
	}
	var res = map[string]string{}
	_ = json.NewDecoder(resp.Body).Decode(&res)
	return res["token"], nil
}

func ToString(x interface{}) string {
	return fmt.Sprintf("%v", x)
}

func GetContent(method string, url string, msg mss, token string) mss {
	req, _ := http.NewRequest(method, url, msg.ToJson())
	req.Header.Set("m2np-token", token)
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return res
}

//resp, err := http.Post("https://m2np.com/api/login", "application/json", mss{"email": "ainililia@163.com", "password": "ainilili2021"}.ToJson())
//
//if err != nil {
//	log.Fatal(err)
//}
//
//var res map[string]interface{}
//
//json.NewDecoder(resp.Body).Decode(&res)
//
//fmt.Println(res["token"])
//token := ToString(res["token"])
//
//r := GetContent("GET", "https://m2np.com/api/get_inbox", mss{"username": "abbychau"}, token) //時間線
//fmt.Println(r)
//fmt.Println("----")
//r = GetContent("GET", "https://m2np.com/api/get_outbox", mss{"username": "abbychau"}, token) //username發的文
//fmt.Println(r)
//fmt.Println("----")

// API.get("/search?username=" + txt) //搜尋用戶 (auto-complete 用)
// API.post("update_reply", { id: reply.id, content: txt }) //修改回覆
// API.get("/my_emojis", {}) // 查看自己的emoji (cli 不用)
// API.post("/delete_post", { id: article.id }) // 刪除自己的文
// API.post("/update_personal_info", { //更新個人資料
// 	description: description,
// 	color: color,
// 	display_name: displayName,
//   })
// API.post("/reset-password", { email: email }) // 重設密碼 (未實現)
// API.get("/notifications") // 看自己通知
// API.get("/get_replies", { post_id: post_id }) //看文章回覆
// API.get("/get_profile", {username: username}) //看username 的個人檔案
// API.get( //看文章列表 : inbox->時間線 outbox->username 的文章
// username == null ? "/get_inbox" : "/get_outbox",
// mode == "append"
// 	? { less_than_ts: minTS, username: username } //minTS 查看>minTS 的timestamp 的文章 //分頁用
// 	: mode == "prepend"
// 	? { more_than_ts: maxTS, username: username }  //maxTS 查看<maxTS 的timestamp 的文章 //分頁用
// 	: { username: username }

// API.post("/unfollow", { user_id: user.id }) //關注
// API.post("/follow", { user_id: user.id }) //取消關注
