package baidu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 提交请求
func Draw(prompt, image string) (res string, err error) {

	defer func() {
		if p := recover(); p != nil {
			err = errors.New("draw failed")
		}
	}()

	var accessToken string
	if accessToken, err = GetAccessToken(); err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", AI_DRAW_URL, accessToken)

	postData := fmt.Sprintf(`{"prompt":"%s","version":"v2","width":512,"height":512,"image_num":1,"image":"%s","change_degree":1}`, prompt, image)

	var resp *http.Response
	if resp, err = http.Post(url, "application/json", strings.NewReader(postData)); err != nil {
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

	taskId, _ := data["data"].(map[string]interface{})["task_id"]

	res = taskId.(string)

	return
}

// 查询结果
func GetDrawUrl(taskId string) (res string, err error) {

	var accessToken string
	if accessToken, err = GetAccessToken(); err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", GET_AI_DRAW_URL, accessToken)

	postData := fmt.Sprintf(`{"task_id":"%s"}`, taskId)

	var resp *http.Response
	if resp, err = http.Post(url, "application/json", strings.NewReader(postData)); err != nil {
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

	imgURL := data["data"].(map[string]interface{})["sub_task_result_list"].([]interface{})[0].(map[string]interface{})["final_image_list"].([]interface{})[0].(map[string]interface{})["img_url"]

	res = imgURL.(string)

	return
}
