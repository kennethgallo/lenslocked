package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:kennygthedev@gmail.com\">kennygthedev@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Frequently Asked Questions</h1><br>
	<ul>
		<li><b>Is there a free version?</b> Yes! We offer a 30 day trial on any paid plans</li>
		<li><b>What are your support hours?</b> We have support staff answering emails 24/7, though reponse times may vary</li>
		<li><b>How do I contact support?</b> Email us - <a href="mailto:kennygthedev@gmail.com"</a>kennygthedev@gmail.com</li>
	</ul>
	
	`)
	// fmt.Fprint(w, "<h2>Is there a free version?</h2>")
	// fmt.Fprint(w, "<p>Yes! We offer a 30 day trial on any paid plans</p>")
	// fmt.Fprint(w, "<h2>What are your support hours</h2>")
	// fmt.Fprint(w, "<p>We have support staff answering emails 24/7, though reponse times may vary</p>")
	// fmt.Fprint(w, "<h2>How do I contact support</h2>")
	// fmt.Fprint(w, "<p>Email us - kennygthedev@gmail.com</p>")

}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000....")
	http.ListenAndServe(":3000", router)
}
