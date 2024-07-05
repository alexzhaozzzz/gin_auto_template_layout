package system

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/util"
	"path/filepath"
	"time"
)

func NewCommon() *Common {
	return &Common{}
}

type Common struct {
}

// Upload ...
// @Summary 上传文件
// @Tags System
// @Accept multipart/form-data
// @Produce  json
// @Param file formData file true "file"
// @Success 200 {object} ginx.Result "成功"
// @Failure 400 {string} string "bad request"
// @Router /upload [post]
func (s Common) Upload(c *ginx.Context) {
	maxFileSize := 10 << 20 // 限制文件大小为10MB
	if err := c.Request.ParseMultipartForm(int64(maxFileSize)); err != nil {
		logrus.Errorf("Common Upload ParseMultipartForm Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Errorf("Common Upload Read FormFile Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	// 随机生成文件名
	ext := filepath.Ext(file.Filename)
	date := time.Now().Format(time.DateOnly)
	// 指定生成的随机字符串长度
	length := 10
	// 生成随机文件名
	randomFileName := "file_" + util.RandomString(length) + ext

	// 指定上传的文件保存路径
	savePath := "./uploads/" + date + "/" + randomFileName
	// 保存文件到指定的路径
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		logrus.Errorf("Common Upload SaveUploadedFile Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	resp := ginx.InfoResponses{
		Info: savePath,
	}
	c.RenderSuccess(resp)
	return
}
