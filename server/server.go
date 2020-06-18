package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/j4rv/cah"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var port, secureport int
var usingTLS bool
var devMode bool
var serverCert, serverPK string
var publicDir string

var usecase cah.Usecases

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	parseFlags()
	initCertificateStuff()
	initCookieStore()
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
	flag.BoolVar(&devMode, "dev", false, "Activates development mode")
	flag.StringVar(&publicDir, "dir", "static", "the directory to serve files from. Defaults to 'static'")
	flag.Parse()
}

// Start creates and starts the server with the provided usecases
func Start(uc cah.Usecases) {
	usecase = uc

	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(simpleTmplHandler(notFoundPageTmpl))

	setRestRouterHandlers(router)
	setTemplateRouterHandlers(router)

	//Static files handler
	fs := http.FileServer(http.Dir(publicDir))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// Known files:
	router.PathPrefix("/favicon.").Handler(fs)
	router.Path("/manifest.json").Handler(fs)

	StartServer(router)
}

func setRestRouterHandlers(r *mux.Router) {
	restRouter := r.PathPrefix("/api").Subrouter()

	{
		s := restRouter.PathPrefix("/user").Subrouter()
		s.HandleFunc("/login", processLogin).Methods("POST")
		s.HandleFunc("/register", processRegister).Methods("POST")
		s.HandleFunc("/logout", processLogout).Methods("POST", "GET")
	}

	{
		s := restRouter.PathPrefix("/game").Subrouter()
		s.Handle("/create", srvHandler(createGame)).Methods("POST")
		s.Handle("/{gameID}/lobby-state", srvHandler(lobbyState)).Methods("GET")
		s.Handle("/{gameID}/join", srvHandler(joinGame)).Methods("POST")
		s.Handle("/{gameID}/leave", srvHandler(leaveGame)).Methods("POST")
		s.Handle("/{gameID}/start", srvHandler(startGame)).Methods("POST")
	}

	{
		s := restRouter.PathPrefix("/gamestate/{gameStateID}").Subrouter()
		s.HandleFunc("/state-websocket", gameStateWebsocket).Methods("GET")
		s.Handle("/choose-winner", srvHandler(chooseWinner)).Methods("POST")
		s.Handle("/play-cards", srvHandler(playCards)).Methods("POST")
	}

}

func setTemplateRouterHandlers(r *mux.Router) {
	r.HandleFunc("/", loginPageHandler)
	r.HandleFunc("/login", loginPageHandler)
	r.HandleFunc("/games", gamesPageHandler)
	r.HandleFunc("/games/create", createGamePageHandler)
	r.HandleFunc("/games/{gameID}", lobbyPageHandler)
	r.HandleFunc("/games/{gameID}/ingame", ingamePageHandler)
}

// StartServer starts the server using the provided router
func StartServer(r *mux.Router) {
	// For Heroku
	envPort := os.Getenv("PORT")
	if envPort != "" {
		log.Printf("Starting http server in port %s\n", envPort)
		log.Fatal(http.ListenAndServe(":"+envPort, r))
	}

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

type srvHandler func(http.ResponseWriter, *http.Request) error

func (fn srvHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := fn(w, req); err != nil {
		log.Printf("ServeHTTP error: %s", err)
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
	}
}
