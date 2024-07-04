package system

import (
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
	PoFileName   string `json:"po_file_name"`
	Dir          string `json:"dir"`
	ChannelKey   string `json:"channel_key"`
	IsDay        int    `json:"is_day"`
	OnlyPo       int    `json:"only_po"`
}

func (s Template) Add(c *ginx.Context) {
	req := &TemplateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	tDataPath := "./scripts/template/data.go.template"
	tDataOutPath := "./internal/data/" + req.Dir + "/" + req.PoFileName + ".go"
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
		"Dir":          req.Dir,
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

	if req.OnlyPo <= 0 {
		tLogicPath := "./scripts/template/logic.go.template"
		tLogicOutPath := "./internal/logic/" + req.Dir + "/" + req.FileName + ".go"
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
	}

	c.RenderSuccess(nil)
	return
}

func (s Template) AbilityAdd(c *ginx.Context) {
	req := &TemplateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Template AbilityAdd ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	tDataPath := "./scripts/template/data.ability.go.template"
	tDataOutPath := "./internal/data/" + req.Dir + "/" + req.PoFileName + ".go"
	if req.IsDay > 0 {
		tDataPath = "./scripts/template/data.ability.day.go.template"
	}

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
		"Dir":          req.Dir,
		"ChannelKey":   req.ChannelKey,
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

	if req.OnlyPo <= 0 {
		tLogicPath := "./scripts/template/logic.ability.go.template"
		tLogicOutPath := "./internal/logic/" + req.Dir + "/" + req.FileName + ".go"
		if req.IsDay > 0 {
			tLogicPath = "./scripts/template/logic.ability.day.go.template"
		}

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
	}

	c.RenderSuccess(nil)
	return
}
