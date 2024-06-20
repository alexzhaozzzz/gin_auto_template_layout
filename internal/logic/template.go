package logic

import (
	"github.com/sirupsen/logrus"
	"os"
	"text/template"

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
}

func (s Template) Add(c *ginx.Context) {
	req := &TemplateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	tDataPath := "./scripts/template/data.go.template"
	tDataOutPath := "./internal/data/" + req.ParamName + ".go"
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
	tLogicOutPath := "./internal/logic/" + req.ParamName + ".go"
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
