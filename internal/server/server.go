package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cah "github.com/j4rv/cah/internal/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var port, secureport int
var devMode bool

var config cah.Config
var usecase cah.Usecases

var logError = log.New(os.Stderr, "[ERROR]", log.LstdFlags)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func parseFlags() {
	flag.IntVar(&port, "port", 80, "Server port for serving HTTP")
	flag.IntVar(&secureport, "secureport", 443, "Server port for serving HTTPS")
	flag.BoolVar(&devMode, "dev", false, "Activates development mode")
	flag.Parse()
}

// Start creates and starts the server with the provided usecases
func Start(cfg cah.Config, uc cah.Usecases) {
	config = cfg
	usecase = uc

	parseFlags()
	initCookieStore()
	initTemplates()

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(simpleTmplHandler(notFoundPageTmpl, http.StatusNotFound))
	setRestRouterHandlers(router)
	setTemplateRouterHandlers(router)

	//Static files handler
	fs := http.FileServer(http.Dir(config.StaticPath))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// Known files:
	router.PathPrefix("/favicon.").Handler(fs)
	router.Path("/manifest.json").Handler(fs)

	StartServer(router)
}

const loginPath = "/login"
const gamesPath = "/games"

func setRestRouterHandlers(r *mux.Router) {
	restRouter := r.PathPrefix("/api").Subrouter()

	{
		s := restRouter.PathPrefix("/user").Subrouter()
		s.HandleFunc(loginPath, processLogin).Methods("POST")
		s.HandleFunc("/register", processRegister).Methods("POST")
		s.HandleFunc("/logout", processLogout).Methods("POST", "GET")
	}

	{
		s := restRouter.PathPrefix("/game").Subrouter()
		s.Handle("/create", loggedInHandler(createGame)).Methods("POST")
		s.Handle("/{gameID}/lobby-state", gameHandler(lobbyState)).Methods("GET")
		s.Handle("/{gameID}/join", gameHandler(joinGame)).Methods("POST")
		s.Handle("/{gameID}/leave", gameHandler(leaveGame)).Methods("POST")
		s.Handle("/{gameID}/start", loggedInHandler(startGame)).Methods("POST")
	}

	{
		s := restRouter.PathPrefix("/game/{gameID}/state").Subrouter()
		s.Handle("/websocket", gameHandler(gameStateWebsocket)).Methods("GET")
		s.Handle("/choose-winner", gameHandler(chooseWinner)).Methods("POST")
		s.Handle("/play-cards", gameHandler(playCards)).Methods("POST")
	}

}

func setTemplateRouterHandlers(r *mux.Router) {
	r.HandleFunc("/", loginPageHandler)
	r.HandleFunc(loginPath, loginPageHandler)
	r.Handle(gamesPath, loggedInHandler(gamesPageHandler))
	r.Handle(gamesPath+"/create", loggedInHandler(createGamePageHandler))
	r.Handle(gamesPath+"/{gameID}", gameHandler(lobbyPageHandler))
	r.Handle(gamesPath+"/{gameID}/ingame", gameHandler(ingamePageHandler))
}

// StartServer starts the server using the provided router
func StartServer(r *mux.Router) {
	// For Heroku
	envPort := os.Getenv("PORT")
	if envPort != "" {
		log.Printf("Starting http server in port %s\n", envPort)
		log.Fatal(http.ListenAndServe(":"+envPort, r))
	}

	log.Printf("Starting http server in port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

// Handler wrappers

// To handle errors returned by the HandlerFunc
type srvHandler func(http.ResponseWriter, *http.Request) error

func (fn srvHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := fn(w, req); err != nil {
		logError.Printf("trying to ServeHTTP: %s", err)
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
	}
}

// For handlers that required a logged user
type loggedInHandler func(cah.User, http.ResponseWriter, *http.Request) error

func (fn loggedInHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	srvHandler(func(w http.ResponseWriter, req *http.Request) error {
		user, err := userFromSession(w, req)
		if err != nil {
			addFlashMsg(notLoggedInMsg, w, req)
			execTemplate(loginPageTmpl, w, http.StatusUnauthorized, getFlashes(w, req))
			return nil
		}
		return fn(user, w, req)
	}).ServeHTTP(w, req)
}

// For handlers that require a game (and a user)
type gameHandler func(cah.Game, cah.User, http.ResponseWriter, *http.Request) error

func (fn gameHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	loggedInHandler(func(user cah.User, w http.ResponseWriter, req *http.Request) error {
		game, err := gameFromRequest(req)
		if err != nil {
			addFlashMsg(gameDoesntExistMsg, w, req)
			http.Redirect(w, req, gamesPath, http.StatusFound)
			return err
		}
		return fn(game, user, w, req)
	}).ServeHTTP(w, req)
}
