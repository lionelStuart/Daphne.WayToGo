package Session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	sessionKey string
	mux        *http.ServeMux
	store      *sessions.CookieStore
	writer     *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	fmt.Println("end test")
	os.Exit(code)
}

func setUp() {
	sessionKey = "test"
	store = sessions.NewCookieStore([]byte(sessionKey))
	mux = http.NewServeMux()
	mux.HandleFunc("/test", Handler)
	mux.HandleFunc("/gest", Handler2)
	writer = httptest.NewRecorder()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session1")
	session2, _ := store.Get(r, "session2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session2.Values["foo"] = "bar"
	session2.Values[42] = 1
	session.Values["foo"] = "bar"
	session.Values[42] = 1
	sessions.Save(r, w)
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session1")
	var value = session.Values["foo"]
	fmt.Println("results", value)
}

func TestGorillaSession(t *testing.T) {
	request, _ := http.NewRequest("GET", "/test", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	for j, i := range writer.Result().Cookies() {
		t.Log(j, i)
	}

	t.Log("get cookie")

	request2, _ := http.NewRequest("GET", "/gest", nil)
	for _, j := range writer.Result().Cookies() {
		request2.AddCookie(j)
	}
	mux.ServeHTTP(writer, request2)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

}
