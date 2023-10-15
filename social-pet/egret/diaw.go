package egret

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"social-pet/api/baidu"
	"social-pet/global"
	"social-pet/models"

	"github.com/astaxie/beego/logs"
)

func Draw() {
	var err error
	defer func() {
		if p := recover(); p != nil {

		}
	}()

	var count int64
	list := make([]models.DrawInfo, 0)
	if count, err = global.DB.Table(new(models.DrawInfo).TableName()).
		Where("status = 0").
		FindAndCount(&list); err != nil {
		logs.Error(err.Error())
	}

	if count > 0 {
		logs.Info("draw working...")

		for k, _ := range list {
			var path string
			if path, err = baidu.GetDrawUrl(list[k].TaskId); err != nil {
				logs.Error(err.Error())
			}

			var resp *http.Response
			if resp, err = http.Get(path); err != nil {
				logs.Error(err.Error())
			}
			defer resp.Body.Close()

			// 创建一个文件用于保存
			save := fmt.Sprintf("./static/draw/%s.png", list[k].TaskId)
			var out *os.File
			if out, err = os.Create(save); err != nil {
				logs.Error(err.Error())
			}
			defer out.Close()

			// 然后将响应流和文件流对接起来
			if _, err = io.Copy(out, resp.Body); err != nil {
				logs.Error(err.Error())
			}

			// var identify string
			// if identify, err = baidu.AnimalIdentify(path); err != nil {
			// 	logs.Error(err.Error())
			// }

			if err == nil {
				list[k].ImageUrl = save[1:]
				// list[k].AnimalIdentify = identify
				list[k].Status = 1

				if _, err = global.DB.Table(new(models.DrawInfo).TableName()).
					Where("task_id = ?", list[k].TaskId).
					AllCols().
					Update(&list[k]); err != nil {
					logs.Error(err.Error())
				}
			}

		}

		logs.Info("draw finish...")
	}

}
