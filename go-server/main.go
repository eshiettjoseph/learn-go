package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
// function to handle /hello route. 
func helloHandler(w http.ResponseWriter,r *http.Request){
	// If user request path not /hello display error
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}
	// Only GET request method allowed
	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello!")

 }

func main(){
	// Checkout /static directory and use index.html
	fileServer := http.FileServer(http.Dir("./static"))
	// Handle the root(/) route by sending to fileServer var
	http.Handle("/", fileServer)
	// Handle /form route
	http.HandleFunc("/form", formHandler)
	// handle /hello route
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// This creates the server. If error not equal to nil, handle the error
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}

}
