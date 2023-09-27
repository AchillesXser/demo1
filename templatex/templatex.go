package templatex

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

type TemplateX struct {
	htmlTpl *template.Template
}

func (t TemplateX) Parse(w http.ResponseWriter, path string) (TemplateX, error) {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("解析模板错误: %v", err)
		http.Error(w, "解析模板错误", http.StatusInternalServerError)
		return TemplateX{
			htmlTpl: nil,
		}, errors.New("解析模板错误")
	}

	return TemplateX{htmlTpl: tpl}, nil
}

func (t TemplateX) Execute(w http.ResponseWriter, data interface{}) error {
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("渲染模板错误: %v", err)
		http.Error(w, "渲染模板错误", http.StatusInternalServerError)
		return errors.New("渲染模板错误")
	}

	w.Header().Set("Content-Type", "text/html;charset=utf8")
	return nil
}
