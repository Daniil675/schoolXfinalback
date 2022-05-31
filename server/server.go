package server

import (
	"log"
	"net/http"
)

func (s Server) Start(port string) {
	http.HandleFunc("/", indexHandler)

	//router.POST("/event", Wrap(cors.Then(s.eventAddHandler)))
	//router.OPTIONS("/event", s.optionsHandler)

	//user
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/users", usersHandler)

	//image
	http.HandleFunc("/image", imageHandler)
	http.HandleFunc("/images", imagesHandler)

	//user
	http.HandleFunc("/favorite", favoriteHandler)
	//http.HandleFunc("/favorites", imagesHandler)

	http.HandleFunc("/upload", uploadHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln("Error at ListenAndServe: ", err)
	} else {
		log.Printf("Server stared at localhost:%s\n", port)
	}
}
