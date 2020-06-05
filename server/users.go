package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/j4rv/cah"
)

/*
	TEMPLATE HANDLERS
*/

func loginPageHandler(w http.ResponseWriter, req *http.Request) {
	execTemplate(loginPageTmpl, w, nil)
}

func processLogin(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form["username"]
	password := req.Form["password"]
	if len(username) != 1 || len(password) != 1 {
		http.Error(w, "Unexpected amount of form vals.", http.StatusUnauthorized)
		return
	}
	u, ok := usecase.User.Login(username[0], password[0])
	if !ok {
		log.Printf("%s tried to login using user '%s'", req.RemoteAddr, username)
		http.Error(w, "The username and password you entered did not match our records.", http.StatusUnauthorized)
		return
	}
	session, err := cookies.Get(req, sessionid)
	if err != nil {
		log.Printf("Failed at getting the cookie.")
		http.Error(w, "Failed at getting the cookie", http.StatusBadRequest)
		return
	}
	session.Values[userid] = u.ID
	session.Save(req, w)
	log.Printf("User %s with id %d just logged in!", u.Username, u.ID)
	// everything ok, back to index with your brand new session!
	http.Redirect(w, req, "/", http.StatusFound)
}

func processRegister(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form["username"]
	password := req.Form["password"]
	if len(username) != 1 || len(password) != 1 {
		http.Error(w, "Unexpected amount of form vals.", http.StatusUnauthorized)
		return
	}
	u, err := usecase.User.Register(username[0], password[0])
	if err != nil {
		log.Printf("%s tried to register using user '%s'", req.RemoteAddr, username[0])
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	session, err := cookies.Get(req, sessionid)
	session.Values[userid] = u.ID
	session.Save(req, w)
	log.Printf("User %s with id %d just registered!", u.Username, u.ID)
	// everything ok, back to index with your brand new session!
	http.Redirect(w, req, "/", http.StatusFound)
}

func processLogout(w http.ResponseWriter, req *http.Request) {
	session, err := cookies.Get(req, sessionid)
	if err != nil {
		http.Error(w, "There was a problem while getting the session cookie", http.StatusInternalServerError)
	}
	session.Values = make(map[interface{}]interface{})
	session.Save(req, w)
	http.Redirect(w, req, "/", http.StatusFound)
}

func validCookie(w http.ResponseWriter, req *http.Request) {
	u, err := userFromSession(req)
	if err != nil {
		http.Error(w, "you dont own a valid cookie", http.StatusUnauthorized)
		return
	}
	writeResponse(w, u)
}

/*
	SESSIONS STUFF
*/

var cookies *sessions.CookieStore

const sessionid = "session_token"
const userid = "user_id"

func init() {
	skey := os.Getenv("SESSION_KEY")
	encKey := os.Getenv("COOKIE_ENCRIPTION_KEY")

	if skey == "" {
		panic("Please set SESSION_KEY environment variable; it is needed to have secure cookies")
	}
	if encKey == "" && !devMode {
		panic("Please set COOKIE_ENCRIPTION_KEY for production instances. Must be 16 or 32 bytes long.")
	}

	if encKey == "" {
		cookies = sessions.NewCookieStore([]byte(skey))
	} else {
		cookies = sessions.NewCookieStore([]byte(skey), []byte(encKey))
	}
}

func userFromSession(r *http.Request) (cah.User, error) {
	session, err := cookies.Get(r, sessionid)
	if err != nil {
		return cah.User{}, err
	}
	val, ok := session.Values[userid]
	if !ok {
		return cah.User{}, fmt.Errorf("Tried to get user from session without an id")
	}
	id, ok := val.(int)
	if !ok {
		log.Printf("Session with non int id value: '%v'", session.Values)
		return cah.User{}, fmt.Errorf("Session with non int id value")
	}
	u, ok := usecase.User.ByID(id)
	if !ok {
		return u, fmt.Errorf("No user found with ID %d", id)
	}
	return u, nil
}
