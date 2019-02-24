package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/j4rv/golang-stuff/cah"

	"github.com/gorilla/mux"
)

var port, secureport int
var usingTLS bool
var serverCert, serverPK string
var publicDir string

var usecase cah.Usecases

func init() {
	initCertificateStuff()
	parseFlags()
}

func initCertificateStuff() {
	serverCert = os.Getenv("SERVER_CERTIFICATE")
	serverPK = os.Getenv("SERVER_PRIVATE_KEY")
	usingTLS = serverCert != "" && serverPK != ""
	if serverCert == "" {
		log.Println("Server certificate not found. Environment variable: SERVER_CERTIFICATE")
	}
	if serverPK == "" {
		log.Println("Server private key not found. Environment variable: SERVER_PRIVATE_KEY")
	}
}

func parseFlags() {
	flag.IntVar(&port, "port", 80, "Server port for serving HTTP")
	flag.IntVar(&secureport, "secureport", 443, "Server port for serving HTTPS")
	flag.StringVar(&publicDir, "dir", "./frontend/build", "the directory to serve files from. Defaults to './frontend/build'")
	flag.Parse()
}

func Start(uc cah.Usecases) {
	usecase = uc

	router := mux.NewRouter()
	//Any non found paths should redirect to index. React-router will handle those.
	router.NotFoundHandler = http.HandlerFunc(serveFrontend(publicDir + "/index.html"))

	restRouter := router.PathPrefix("/rest").Subrouter()
	handleUsers(restRouter)
	handleGames(restRouter)
	handleGameStates(restRouter)

	//Static files handler
	router.PathPrefix("/static").Handler(http.FileServer(http.Dir(publicDir)))
	// Known files. We have to define them one by one since we can't use PathPrefix("/"),
	// as that would make the NotFoundHandler stop working.
	router.PathPrefix("/favicon.").Handler(http.FileServer(http.Dir(publicDir)))
	router.Path("/manifest.json").Handler(http.FileServer(http.Dir(publicDir)))

	StartServer(router)
}

func StartServer(r *mux.Router) {
	if usingTLS {
		go func() {
			log.Printf("Starting http server in port %d\n", port)
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
		}()
		log.Printf("Starting https server in port %d\n", secureport)
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", secureport), serverCert, serverPK, r))
	} else {
		log.Println("Server will listen and serve without TLS")
		log.Printf("Starting http server in port %d\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
	}
}

func serveFrontend(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, path)
	}
}

type srvHandler func(http.ResponseWriter, *http.Request) error

func (fn srvHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := fn(w, req); err != nil {
		log.Printf("ServeHTTP error: %s", err)
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
	}
}
