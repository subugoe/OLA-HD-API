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
    "time"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s %s %s %s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
