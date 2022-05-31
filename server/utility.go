package server

import (
	"encoding/json"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func makeError(w http.ResponseWriter, errCode ErrorCode) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	e := Err{Code: errCode}
	responseJSON(w, e)
}

func getURLValue(values url.Values, k string) (string, bool) {
	keys, ok := values[k]
	if !ok || len(keys[0]) < 1 {
		return "", false
	}
	key := keys[0]
	return key, ok
}

func getURLValueInt(values url.Values, k string) (int, bool) {
	value, ok := getURLValue(values, k)
	if !ok {
		return 0, false
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return valueInt, true
}

func newSize(fWidth, fHeight, nWidth, nHeight int) (uint, uint) {
	newWidth, newHeight := nWidth, nHeight
	ratio := fWidth / fHeight
	if newWidth/nHeight > ratio {
		newWidth = newHeight * ratio
	} else {
		newHeight = newWidth / ratio
	}
	return uint(newWidth), uint(newHeight)
}

func resizeImage(img image.Image, folder, prefix, ext string, w, h int) {
	newWidth, newHeight := newSize(img.Bounds().Size().X, img.Bounds().Size().Y, w, h)
	newImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	out, err := os.Create(uploadPath + folder + prefix + ext)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()

	// write new image to file
	switch ext {
	case ".png":
		png.Encode(out, newImg)
	case ".jpg", ".jpeg":
		jpeg.Encode(out, newImg, nil)
	}
}
