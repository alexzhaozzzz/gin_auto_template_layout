package logic

import (
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"os"
	"text/template"

	"github.com/sirupsen/logrus"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewTemplate() *Template {
	return &Template{}
}

type Template struct {
}

type TemplateReq struct {
	FunName      string `json:"fun_name"`
	DbStructName string `json:"db_struct_name"`
	ParamName    string `json:"param_name"`
	FileName     string `json:"file_name"`
}

func (s Template) Add(c *ginx.Context) {
	req := &TemplateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	tDataPath := "./scripts/template/data.go.template"
	tDataOutPath := "./internal/data/" + req.FileName + ".go"
	td, err := template.ParseFiles(tDataPath)
	if err != nil {
		logrus.Errorf("template.ParseFiles Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}
	// 准备数据
	data := map[string]string{
		"FunName":      req.FunName,
		"DbStructName": req.DbStructName,
		"ParamName":    req.ParamName,
	}

	// 打开一个文件用于写入
	file, err := os.Create(tDataOutPath)
	if err != nil {
		logrus.Errorf("template.ParseFiles Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}
	defer file.Close()

	// 使用模板渲染数据，并将结果写入文件
	err = td.Execute(file, data)
	if err != nil {
		logrus.Errorf("template.ParseFiles Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	tLogicPath := "./scripts/template/logic.go.template"
	tLogicOutPath := "./internal/logic/" + req.FileName + ".go"
	tl, err := template.ParseFiles(tLogicPath)
	if err != nil {
		logrus.Errorf("template.ParseFiles Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	// 打开一个文件用于写入
	tlFile, err := os.Create(tLogicOutPath)
	if err != nil {
		logrus.Errorf("template.ParseFiles Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}
	defer tlFile.Close()

	// 使用模板渲染数据，并将结果写入文件
	err = tl.Execute(tlFile, data)
	if err != nil {
		logrus.Errorf("template.ParseFiles Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Template) Test(c *ginx.Context) {
	logrus.WithFields(logrus.Fields{
		"SpendTime": 1,
		"IP":        2,
	}).Info("--------------------test---------------------")

	logrus.Info("--------------------test---------------------")

	c2 := conn.GetMerchantRDb()
	if c2 == nil {
		logrus.Info("redis is nil")
	}

	c1 := conn.GetClickhouseDB()
	if c1 == nil {
		logrus.Info("cliclhouse is nil")
	}

	c3, err := conn.GetNSQCli()
	if err != nil {

		logrus.Info("nsq is err", err)
	}
	if c3 == nil {
		logrus.Info("nsq is nil")
	}

	c.RenderSuccess(nil)
	return
}
