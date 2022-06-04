package main
	
import (
    "net/http"
	"os"
	"fmt"
)

var indexstring string = `<h1> This is dummy text, click <a href=./update>here</a> to edit it </h1>
<marquee> (@_@)`;


func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, indexstring);
}

func update(w http.ResponseWriter, r *http.Request){
	
	/* Setting up headers manually */ // TODO: Setup headers from user input 

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	
	switch r.Method {
		case "GET":		
			AddForm := `
			<form method=post action=./update>
			<textarea name=indexstring rows="20" cols="200">`+indexstring+`</textarea>
			<input type=submit></input>
			</form>
			`
			fmt.Fprint(w, AddForm)
			return
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			indexstring = r.FormValue("indexstring");

			updated := `<h2> <a href=./update>Edit Page</a>  &emsp;&emsp;<a href=../>View on Index Page</a>  </h2><br>
			` + indexstring
			fmt.Fprint(w,updated)
			return 

	}
}


func main() {
	fmt.Println("Hello World");

	http.HandleFunc("/", index);
	http.HandleFunc("/update", update);

	http.ListenAndServe(":80", nil);
	os.Exit(0) 
}