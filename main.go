package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000....")
	http.ListenAndServe(":3000", r)
}
