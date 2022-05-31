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

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln("Error at ListenAndServe: ", err)
	} else {
		log.Printf("Server stared at localhost:%s\n", port)
	}
}
