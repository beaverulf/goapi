// Go CryptoService API.
//
// Project description.
//
//     Schemes: http
//     Version: 0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/beaverulf/goapi/controllers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var rootRouter = mux.NewRouter()

//Router for API endpoint controllers
var Router = rootRouter.PathPrefix("/v1").Subrouter()

func main() {

	setupPrometheus()
	setupCryptoEndpoints()
	log.Fatal(http.ListenAndServe(":8000", rootRouter))
}

func setupPrometheus() {
	rootRouter.Handle("/metrics", promhttp.Handler())
}
func setupCryptoEndpoints() {
	Router.HandleFunc("/crypto/aes", controllers.GetAESEncryptionServicesList).Methods("GET")
	Router.Handle("/crypto/aes/ecb/encrypt", TimedHandler(controllers.EncryptAesECBHandlerFunc, aes128ECBEncryptorResponseTimerMetric)).Methods("POST")
	Router.HandleFunc("/crypto/aes/ecb/decrypt", TimedHandler(controllers.DecryptAesECBHandlerFunc, aes128ECBDecryptorResponseTimerMetric)).Methods("POST")
}

//Prometheus metrics.
var (
	aes128ECBEncryptorResponseTimerMetric = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "api_crypto_aes_128ECB_encrypt_response_time_ms",
		Help: "Response time of AES-128-ECB Encryption",
	})

	aes128ECBDecryptorResponseTimerMetric = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "api_crypto_aes_128ECB_decrypt_response_time_",
		Help: "Response time of AES-128-ECB Decryption",
	})
)

//TimedHandler intercepts a handlefunction and performs before/after
func TimedHandler(h http.HandlerFunc, summary prometheus.Summary) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		duration := time.Now().Sub(start)

		summary.Observe(float64(duration.Seconds() * 1000))
	}
}
