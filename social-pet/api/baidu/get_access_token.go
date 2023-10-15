package baidu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() (res string, err error) {
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	
	var resp *http.Response
	if resp, err = http.Post(ACCESS_TOKEN_URL, "application/json", strings.NewReader(postData)); err != nil {
		return
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)

	res = accessTokenObj["access_token"]

	return
}
