package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"schoolXfinalback/storage/datastore"
	"strconv"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	datastore.UserAdd(datastore.User{
		Name:     "Daniil",
		Email:    "daniil.popenko@yandex.ru",
		Phone:    "89287750060",
		Password: "1234567890",
		Created:  time.Now(),
	})
	fmt.Fprint(w, "API works")
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userGetHandler(w, r)
	case "POST":
		userAddHandler(w, r)
	case "PUT":
		userEditHandler(w, r)
	case "DELETE":
		userDeleteHandler(w, r)
	case "OPTIONS":
		optionsHandler(w, r)
	}
}
func userGetHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		makeError(w, InvalidData)
		return
	}
	key := keys[0]

	id, err := strconv.Atoi(key)
	if err != nil {
		makeError(w, InvalidData)
		return
	}

	city, ok := datastore.UserGetBy("id", id)
	if !ok {
		makeError(w, SomethingWentWrong)
		return
	}
	responseJSON(w, city)
}
func userAddHandler(w http.ResponseWriter, r *http.Request) {
	var data datastore.User
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		makeError(w, InvalidData)
		return
	}

	data.Created = time.Now()
	datastore.UserAdd(data)
}
func userEditHandler(w http.ResponseWriter, r *http.Request) {
	var data datastore.User
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		makeError(w, InvalidData)
		return
	}

	datastore.UserEdit(data)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) // u.Scheme+"://"+u.Host
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-api-token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

//func uploadHandler(w http.ResponseWriter, r *http.Request) {
//	//session, _ := r.Cookie("session")
//	//if _, ok := s.getInfoFromSession(session.Value); ok {
//	// if user := s.DB.GetUser("id", strconv.Itoa(id)); user.Role == db.Admin {
//	r.ParseMultipartForm((128 << 20) * 8)
//	file, handler, err := r.FormFile("uploadfile")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//	//fmt.Println(handler.Filename)
//	extIndex := strings.LastIndex(handler.Filename, ".")
//	var allowable bool
//	for _, ext := range AllowableFileExtensions {
//		if ext == strings.ToLower(handler.Filename[extIndex+1:]) {
//			allowable = true
//			break
//		}
//	}
//	if allowable {
//		newFilename := uuid.NewV3(uuid.Must(uuid.NewV1()), handler.Filename).String() + handler.Filename[extIndex:]
//		f, err := os.OpenFile("./server/resourses/temp/"+newFilename, os.O_WRONLY|os.O_CREATE, 0666)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		io.Copy(f, file)
//		f.Close()
//
//		responseJSON(w, Img{Name: newFilename})
//
//		ext := filepath.Ext(newFilename)
//		ext = strings.ToLower(ext)
//		//fmt.Println(ext)
//		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
//			previewFile, err := os.Open("./server/resourses/temp/" + newFilename)
//
//			//if err != nil {
//			//	log.Println(err)
//			//}
//
//			// decode jpeg into image.Image
//			//img, err := jpeg.Decode(file)
//			var (
//				img image.Image
//			)
//			switch (ext) {
//			case ".png":
//				img, err = png.Decode(previewFile)
//			case ".jpg", ".jpeg":
//				img, err = jpeg.Decode(previewFile)
//			}
//			if err != nil {
//				log.Println(err)
//			}
//			previewFile.Close()
//
//			// resize to width 1000 using Lanczos resampling
//			// and preserve aspect ratio
//			m := resize.Resize(400, 300, img, resize.Lanczos3)
//
//			out, err := os.Create("./server/resourses/temp/preview_" + newFilename)
//			if err != nil {
//				log.Println(err)
//			}
//			defer out.Close()
//
//			// write new image to file
//			switch (ext) {
//			case ".png":
//				png.Encode(out, m)
//			case ".jpg", ".jpeg":
//				jpeg.Encode(out, m, nil)
//			}
//			//
//
//			//test_png_resized_2.png
//			//preview_test_png_resized_2.png
//		}
//	} else {
//		makeError(w, UnallowableFileExtension)
//	}
//	//}
//	//}
//}
