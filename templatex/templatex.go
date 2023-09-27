package templatex

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateX struct {
	Tpl string
}

func (t *TemplateX) Parse(w http.ResponseWriter, filepath string) (*template.Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("解析模板错误: %v", err)
		http.Error(w, "解析模板错误", http.StatusInternalServerError)
		return nil, err
	}

	return tpl, nil
}

func (t *TemplateX) Execute(w http.ResponseWriter, tplParsed *template.Template) {
	err := tplParsed.Execute(w, nil)
	if err != nil {
		log.Printf("渲染模板错误: %v", err)
		http.Error(w, "渲染模板错误", http.StatusInternalServerError)
	}
}
