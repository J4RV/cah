package server

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/j4rv/cah"
)

const wrongUserOrPassMsg = "The username or password you entered is incorrect."
const afterLoginRedirect = "/game/list/open"

/*
	TEMPLATE HANDLERS
*/

type loginPageCtx struct {
	ErrorMsg string
}

func loginPageHandler(w http.ResponseWriter, req *http.Request) {
	execTemplate(loginPageTmpl, w, nil)
}

func processLogin(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form["username"]
	password := req.Form["password"]
	if len(username) != 1 || len(password) != 1 {
		http.Redirect(w, req, "/login", http.StatusUnauthorized)
		return
	}
	u, ok := usecase.User.Login(username[0], password[0])
	if !ok {
		log.Printf("%s tried to login using user '%s'", req.RemoteAddr, username)
		execTemplate(loginPageTmpl, w, loginPageCtx{ErrorMsg: wrongUserOrPassMsg})
		return
	}
	log.Printf("User %s with id %d just logged in!", u.Username, u.ID)
	session, _ := cookies.Get(req, sessionid)
	session.Values[userid] = u.ID
	session.Save(req, w)
	// everything ok, back to index with your brand new session!
	http.Redirect(w, req, afterLoginRedirect, http.StatusFound)
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
		execTemplate(loginPageTmpl, w, loginPageCtx{ErrorMsg: err.Error()})
		return
	}
	log.Printf("User %s with id %d just registered!", u.Username, u.ID)
	session, _ := cookies.Get(req, sessionid)
	session.Values[userid] = u.ID
	session.Save(req, w)
	// everything ok, back to index with your brand new session!
	http.Redirect(w, req, afterLoginRedirect, http.StatusFound)
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
	skey := make([]byte, 64)
	encKey := make([]byte, 32)
	rand.Read(skey)
	rand.Read(encKey)
	cookies = sessions.NewCookieStore(skey, encKey)
}

func userFromSession(req *http.Request) (cah.User, error) {
	session, err := cookies.Get(req, sessionid)
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
