package base

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	status int
	body   interface{}
	http.ResponseWriter
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *loggingResponseWriter) Write(body []byte) (int, error) {
	var data interface{}
	_ = json.Unmarshal(body, &data)
	w.body = data
	return w.ResponseWriter.Write(body)
}

func LoggerRequestResponse(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		// Work / inspect body. You may even modify it!
		var data interface{}
		_ = json.Unmarshal(body, &data)
		reqMap := map[string]interface{}{
			"host":       r.Host,
			"requestUri": r.RequestURI,
			"header":     r.Header,
			"method":     r.Method,
			"body":       data,
			//"summary": a,
		}

		log.Println("[Request]", reqMap)

		// And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		// Create a response wrapper:
		loggingRW := &loggingResponseWriter{
			ResponseWriter: w,
		}
		// Call next handler, passing the response wrapper:
		h.ServeHTTP(loggingRW, r)
		resMap := map[string]interface{}{
			"status": loggingRW.status,
			"body":   loggingRW.body,
		}

		log.Println("[Response]", resMap)
	})
}

func LoggerHttpClient(logType string, input interface{}) {
	log.Println(logType, input)
}
