package main

import "go-cms/internal/pkg"
func main()  {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/",say)
	//
	//server:=&http.Server{
	//	Addr:":8089",
	//	Handler:mux,
	//}
	//log.Fatal(server.ListenAndServe())
	//http.HandleFunc("/",say)
	//log.Fatal(http.ListenAndServe(":8091",nil))

	pkg.Info.Println("222")
}

//func say(w http.ResponseWriter,r *http.Request)  {
//	w.Write([]byte("bye bye"))
//}