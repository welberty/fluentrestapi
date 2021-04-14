package fluentrestapi

import "net/http"

func (sc ServerContext) Get(path string, fn Handler) ServerContext {
	sc.buildHandleRoute(path, fn, http.MethodGet)
	return sc
}

func (sc ServerContext) Post(path string, fn Handler) ServerContext {
	sc.buildHandleRoute(path, fn, http.MethodPost)
	return sc
}

func (sc ServerContext) Put(path string, fn Handler) ServerContext {
	sc.buildHandleRoute(path, fn, http.MethodPut)
	return sc
}

func (sc ServerContext) Delete(path string, fn Handler) ServerContext {
	sc.buildHandleRoute(path, fn, http.MethodDelete)
	return sc
}
