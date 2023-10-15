package baidu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AnimalIdentify(imgUrl string) (res string, err error) {

	defer func() {
		if p := recover(); p != nil {

		}
	}()

	var accessToken string
	if accessToken, err = GetAccessToken(); err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", ANIMAL_IDENTIFY, accessToken)

	postData := fmt.Sprintf(`url=%s&top_num=1`, imgUrl)

	var resp *http.Response
	if resp, err = http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData)); err != nil {
		return
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return
	}

	fmt.Println("data: ", data)

	name := data["result"].([]interface{})[0].(map[string]interface{})["name"]

	res = name.(string)

	return
}
