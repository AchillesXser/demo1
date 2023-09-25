package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Header().Set("Content-Type", "text/html;charset=GBK222")
		fmt.Fprint(w, "<h1>Main Page</h1>")
	case "/contact":
		fmt.Fprint(w, "<h1>our contact</h1><a href='https://baidu.com'>contact Us</a><br>")
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {

	// 1. http.Handler - 带有 ServeHttp 方法的接口;
	// 2. http.HandlerFunc - func类型，接收与ServeHttp相同的参数，是 http.Handler的实现;

	// 3. http.HandleFunc - 是自动将 func类型的函数如 homeHandler 转换成 HandlerFunc类型的方法;
	// 	  http.HandleFunc("/", homeHandler)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/faq", faqHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Header().Set("Content-Type", "text/html;charset=GBK222")
		fmt.Fprint(w, "<h1>Main Page1</h1>")
	case "/contact":
		fmt.Fprint(w, "<h1>our contact</h1><a href='https://baidu.com'>contact Us</a><br>")
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h2>Q: Is there a free version?</h2><h2>A: Yes, We offer a free trial for 30 days on any paid plans.</h2><br><h2>Q: What are your support hours?</h2><br><h2>A: We are support staff	answering emails 24/7, though response times may be a bit slower on weekends.</h2>")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=GBK222")
	fmt.Fprint(w, "<h1>Main Page</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>our contact</h1><a href='https://baidu.com'>contact Us</a><br>")
}
