package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"schoolXfinalback/storage/datastore"
	"schoolXfinalback/utility"
	"strings"
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
	id, ok := getURLValueInt(r.URL.Query(), "id")
	if !ok {
		makeError(w, InvalidData)
		return
	}

	user, ok := datastore.UserGetBy("id", id)
	if !ok {
		makeError(w, SomethingWentWrong)
		return
	}
	responseJSON(w, user)
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		usersGetAllHandler(w, r)
	case "OPTIONS":
		optionsHandler(w, r)
	}
}
func usersGetAllHandler(w http.ResponseWriter, r *http.Request) {
	users, ok := datastore.UsersGetAll()
	if !ok {
		makeError(w, SomethingWentWrong)
		return
	}
	responseJSON(w, users)
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
func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := getURLValueInt(r.URL.Query(), "id")
	if !ok {
		makeError(w, InvalidData)
		return
	}

	datastore.UserDeleteBy("id", id)
}

//image
func imageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		imageGetHandler(w, r)
	case "POST":
		imageAddHandler(w, r)
	case "PUT":
		imageEditHandler(w, r)
	case "DELETE":
		imageDeleteHandler(w, r)
	case "OPTIONS":
		optionsHandler(w, r)
	}
}
func imageGetHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := getURLValueInt(r.URL.Query(), "id")
	if !ok {
		makeError(w, InvalidData)
		return
	}

	image, ok := datastore.ImageGetBy("id", id)
	if !ok {
		makeError(w, SomethingWentWrong)
		return
	}
	responseJSON(w, image)
}
func imagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		imagesGetAllHandler(w, r)
	case "OPTIONS":
		optionsHandler(w, r)
	}
}
func imagesGetAllHandler(w http.ResponseWriter, r *http.Request) {
	images, ok := datastore.ImagesGetAll()
	if !ok {
		makeError(w, SomethingWentWrong)
		return
	}
	responseJSON(w, images)
}
func imageAddHandler(w http.ResponseWriter, r *http.Request) {
	var data datastore.Image

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		makeError(w, InvalidData)
		return
	}

	data.Created = time.Now()
	datastore.ImageAdd(data)
}
func imageEditHandler(w http.ResponseWriter, r *http.Request) {
	var data datastore.Image

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		makeError(w, InvalidData)
		return
	}

	datastore.ImageEdit(data)
}
func imageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := getURLValueInt(r.URL.Query(), "id")
	if !ok {
		makeError(w, InvalidData)
		return
	}

	datastore.ImageDeleteBy("id", id)
}

//favorite
func favoriteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//case "GET":
	//	imageGetHandler(w, r)
	case "POST":
		favoriteAddHandler(w, r)
	//case "PUT":
	//	imageEditHandler(w, r)
	//case "DELETE":
	//	imageDeleteHandler(w, r)
	case "OPTIONS":
		optionsHandler(w, r)
	}
}
func favoriteAddHandler(w http.ResponseWriter, r *http.Request) {
	var data datastore.Favorite

	userId, ok := getURLValueInt(r.URL.Query(), "user_id")
	if !ok {
		makeError(w, InvalidData)
		return
	}
	imageId, ok := getURLValueInt(r.URL.Query(), "image_id")
	if !ok {
		makeError(w, InvalidData)
		return
	}
	data.UserID = userId
	data.ImageID = imageId

	data.Created = time.Now()
	datastore.FavoriteAdd(data)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) // u.Scheme+"://"+u.Host
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-api-token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm((128 << 20) * 8)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//extIndex := strings.LastIndex(handler.Filename, ".")
	ext := strings.ToLower(filepath.Ext(handler.Filename))

	if !utility.Contains(AllowableFileExtensions, ext) {
		makeError(w, UnallowableFileExtension)
		return
	}

	newFoldername := uuid.NewV3(uuid.Must(uuid.NewV1()), handler.Filename).String() + "/"
	os.MkdirAll(uploadPath+newFoldername, os.ModePerm)

	newFilename := uploadPath + newFoldername + "orig" + ext
	f, err := os.OpenFile(newFilename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(f, file)
	f.Close()

	extensionsForResize := []string{".png", ".jpg", ".jpeg"}
	if utility.Contains(extensionsForResize, ext) {
		origFile, err := os.Open(newFilename)
		defer origFile.Close()

		var img image.Image

		switch ext {
		case ".png":
			img, err = png.Decode(origFile)
		case ".jpg", ".jpeg":
			img, err = jpeg.Decode(origFile)
		}
		if err != nil {
			log.Println(err)
		}

		for _, size := range sizes {
			if img.Bounds().Size().X > size.Width && img.Bounds().Size().Y > size.Height {
				resizeImage(img, newFoldername, size.Name, ext, size.Width, size.Height)
			}
		}
	}
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
//		f, err := os.OpenFile("./resources/temp/"+newFilename, os.O_WRONLY|os.O_CREATE, 0666)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		io.Copy(f, file)
//		f.Close()
//
//		//responseJSON(w, Img{Name: newFilename})
//
//		ext := filepath.Ext(newFilename)
//		ext = strings.ToLower(ext)
//		//fmt.Println(ext)
//		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
//			previewFile, err := os.Open("./resources/temp/" + newFilename)
//			var img image.Image
//
//			switch ext {
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
//			fmt.Println(img.Bounds().Size().Y)
//			// resize to width 1000 using Lanczos resampling
//			// and preserve aspect ratio
//			m := resize.Resize(400, 300, img, resize.Lanczos3)
//			//resize.Resize()
//
//			out, err := os.Create("./resources/temp/preview_" + newFilename)
//			if err != nil {
//				log.Println(err)
//			}
//			defer out.Close()
//
//			// write new image to file
//			switch ext {
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
