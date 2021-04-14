package fluentrestapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ServerContext struct {
	server *http.Server
	router *mux.Router
}

type HandlerResponse struct {
	StatusCode   int
	JsonResponse interface{}
}

type Handler func(r Request) HandlerResponse

type Request struct {
	Params map[string]string
	Body   io.ReadCloser
}

func responseWriter(rw http.ResponseWriter, responseHandler HandlerResponse) {
	rw.Header().Add("Content-type", "application/json")
	rw.WriteHeader(responseHandler.StatusCode)
	json.NewEncoder(rw).Encode(responseHandler.JsonResponse)
}

func (sc ServerContext) buildHandleRoute(path string, fn Handler, method string) {
	sc.router.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		req := Request{Params: mux.Vars(r), Body: r.Body}
		responseHandler := fn(req)
		responseWriter(rw, responseHandler)
	}).Methods(method)
}

func (sc ServerContext) Start(port int, server http.Server) error {

	server.Addr = ":" + strconv.Itoa(port)
	server.Handler = sc.router
	sc.server = &server

	return sc.server.ListenAndServe()
}

func Init() ServerContext {
	return ServerContext{router: mux.NewRouter()}
}
