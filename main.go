package main

import (
	"log"

	"Project/datastore/author"
	author3 "Project/delivery/author"
	"Project/driver"
	author2 "Project/service/author"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {

	db, err := driver.Connection()
	if err != nil {
		log.Printf("in main db connection:%v", err)
		return
	}

	authorDatastore := author.New(db)
	authorServices := author2.New(authorDatastore)
	authorHandler := author3.New(authorServices)

	//bookDatastore := book.New(db)
	//bookServices := book2.New(bookDatastore, authorDatastore)
	//bookHandler := book3.New(bookServices)

	//route := mux.NewRouter()
	////routes for author http methods
	//route.HandleFunc("/author", authorHandler.PostAuthor).Methods(http.MethodPost)
	//route.HandleFunc("/author", authorHandler.PutAuthor).Methods("PUT")
	////route.HandleFunc("/author/{id}", authorHandler.DeleteAuthor).Methods("DELETE")
	//
	////routes for book  http methods
	//route.HandleFunc("/book", bookHandler.PostBook).Methods(http.MethodPost)
	//route.HandleFunc("/book", bookHandler.GetAllBook).Methods("GET")
	//route.HandleFunc("/book/{id}", bookHandler.GetBookById).Methods("GET")
	//route.HandleFunc("/book/{id}", bookHandler.PutBook).Methods("PUT")
	//route.HandleFunc("/book/{id}", bookHandler.DeleteBook).Methods("DELETE")
	//
	//fmt.Println("server started and listening on 8000..")
	//log.Fatal(http.ListenAndServe(":8000", route))

	k := gofr.New()
	k.POST("/author", authorHandler.PostAuthor)
	k.PUT("/author/{id}", authorHandler.PutAuthor)
	k.DELETE("/author/{id}", authorHandler.DeleteAuthor)

	k.Start()

}
