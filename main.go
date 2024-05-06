package main

// I have no idea how the import works, yet.
// but for now, it seems to me that it is grabbing all the necessary dependencies
// from either go.mod where all the dependencies live
// now, I have a question and since I havent installed any new modules,
// how does it know which packages to import?
// now, maybe, they are coming from the go package, it feels like they are built-in
import (
	"log"
	"net/http"
)

// the main func is like `Startup.cs` in .Net or `App.ts` in node
func main(){
	// here, we create a new server and I believe this is something like express in node.js
	// we grab from the http. I think the `http`` package comes from the go.mod
	// something I have not installed yet, but i found it here
	// The http has a HandleFunc that take the path as a string and controller/handler
	// and I have no idea why do is call mux?
	// sometimes knowing the source of the name may give epiphany about the reason why this package exists?
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// I think log comes from the go.mod
	// I think ListenAndServe is a method that comes from the http package
	// and it is a method that takes the port and the mux as parameters
	// and it is a method that starts the server
	log.Println("http://localhost:4040")
	// if we got an error return, then stop the server(exit)
	err:=http.ListenAndServe(":4040", mux)
	log.Fatal(err)
}

// build a handler/controller
func home(w http.ResponseWriter, r *http.Request){
	// write to the response writer
	w.Write([]byte("Hello from the home"))
}

