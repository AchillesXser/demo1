package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// 1. http.Handler - 带有 ServeHttp 方法的接口;
	// 2. http.HandlerFunc - func类型，接收与ServeHttp相同的参数，是 http.Handler的实现;

	// 3. http.HandleFunc - 是自动将 func类型的函数如 homeHandler 转换成 HandlerFunc类型的方法;
	// 	  http.HandleFunc("/", homeHandler)

	r := chi.NewRouter()

	// r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/articles", articleHandler)

	r.With(middleware.Logger).Get("/articles/{id}", articleHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "页面不存在", http.StatusNotFound)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func executeTemplate(w http.ResponseWriter, filepath string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("解析模板错误: %v", err)
		http.Error(w, "解析模板错误", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("执行模板错误: %v", err)
		http.Error(w, "执行模板错误", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("cmd", "Ex1", "hello.gohtml")

	executeTemplate(w, templatePath, struct {
		Name string
		Age  int
		Mata struct {
			Visitors int
		}
		Bio string
	}{
		Name: "李毅",
		Age:  32,
		Mata: struct{ Visitors int }{Visitors: 8},
		Bio:  `<script>alert("你被黑了")</script>`,
	})

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("cmd", "Ex1", "contact.gohtml")

	executeTemplate(w, path, struct {
		ContactUrl string
	}{
		ContactUrl: "http://baidu.com",
	})
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("cmd", "Ex1", "faq.gohtml")
	executeTemplate(w, path, struct {
		Version string
		Days    int
		Hours   int
	}{
		Version: "v1.0",
		Days:    7,
		Hours:   24,
	})
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("cmd", "Ex1", "ex1.gohtml")
	executeTemplate(w, path, struct {
		Cint   int
		Cfloat float64
		Cslice []int
		Cmap   map[string]int
	}{
		Cint:   999,
		Cfloat: 12.23891829381928391283981293,
		Cslice: []int{1, 2, 3, 4, 5},
		Cmap: map[string]int{
			"a": 1,
			"b": 1,
			"c": 1,
			"d": 1,
			"e": 1,
		},
	})
}
