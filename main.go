package main

import (
	"github.com/gin-gonic/gin"
	"go-cms/api"
	"go-cms/internal/pkg"
	"go-cms/schema"
	"log"
)

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
	pkg.InitLog()
	router:=gin.Default()
	v1:=router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			//v := config.GetString("db.host")
			//fmt.Println(v)
			//pkg.Logger.Warn("test")
			pkg.Dao.AutoMigrate(&schema.Menu{},&schema.User{})
		})
		v1.POST("/test",api.Login)
	}

	log.Fatal(router.Run(":8091"))
}

//func say(w http.ResponseWriter,r *http.Request)  {
//	w.Write([]byte("bye bye"))
//}