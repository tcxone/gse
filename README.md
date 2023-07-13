# Gse

Easy-to-use go network framework

## Usage

This is an example. I believe you will understand how to use it after reading it.

```
package main

import (
	"github.com/tcxone/gse"
	"net/http"
)

func main() {
	s := gse.new()

	s.use(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	})

	s.get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"message": "az",
		}
		s.sendjson(w, http.StatusOK, data)
	})

	s.post("/msg", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"message": "az",
		}
		s.sendhtml(w, http.StatusOK, "template.html", data)
	})

	s.listen("8080")
}
```