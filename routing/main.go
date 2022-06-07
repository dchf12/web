package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	//リクエストハンドラを特定のHTTPメソッドに制限
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	//リクエストハンドラを特定のホスト名またはサブドメインに制限
	r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	//リクエストハンドラをhttp/httpsに制限
	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	//リクエストハンドラを特定のパスプレフィックスに制限
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)

	http.ListenAndServe(":80", r)
}

func CreateBook(w http.ResponseWriter, r *http.Request)      {}
func ReadBook(w http.ResponseWriter, r *http.Request)        {}
func UpdateBook(w http.ResponseWriter, r *http.Request)      {}
func DeleteBook(w http.ResponseWriter, r *http.Request)      {}
func BookHandler(w http.ResponseWriter, r *http.Request)     {}
func SecureHandler(w http.ResponseWriter, r *http.Request)   {}
func InsecureHandler(w http.ResponseWriter, r *http.Request) {}
func AllBooks(w http.ResponseWriter, r *http.Request)        {}
func GetBook(w http.ResponseWriter, r *http.Request)         {}
