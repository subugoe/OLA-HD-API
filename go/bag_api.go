/*
 * OLA-HD Repository
 *
 * This is the API definition for the (OCR-D) OLA-HD Repository server. You can find out more about OLA-HD [http://ocr-d.de/modulprojekte#%20OLA-HD](http://ocr-d.de/modulprojekte#%20OLA-HD). For test purposes, you can use the api key `test-key` to test the authorization     filters.
 *
 * API version: 1.0.0
 * Contact: panzer@sub.uni-goettingen.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"log"
	"net/http"
	"os"
	"io"
	//"io/ioutil"
)

// curl -X POST http://127.0.0.1:8081/lza/bag -F "bagit_file=@bagit-v4.12.3.zip" -vvv
func AddBag(w http.ResponseWriter, r *http.Request) {

	log.Printf("length: %v", r.ContentLength)
	log.Printf("url: %v", r.URL)
	log.Printf("host: %v", r.Host)


	file, handler, _:= r.FormFile("bagit_file")
	defer file.Close()

	// copy example
	f, _ := os.OpenFile("./1_"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)


	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// curl -X DELETE http://127.0.0.1:8081/lza/bag/4711
func DeleteBagById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// curl -X GET http://127.0.0.1:8081/lza/bag/4711 > a.zip
func GetBagById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetBagList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetBagVersionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetBagVersionsListById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// curl -X PUT http://127.0.0.1:8081/lza/bag -F "bagit_file=@bagit-v4.12.3.zip" -vvv
func UpdateBag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
