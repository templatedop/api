package route

import "net/http"

func GET[Req, Res any](path string, h HandlerFunc[Req, Res], ds ...int) Route {
	return New[Req, Res](http.MethodGet, path, h, ds...)
}

func POST[Req, Res any](path string, h HandlerFunc[Req, Res], ds ...int) Route {
	return New[Req, Res](http.MethodPost, path, h, ds...)
}

func PATCH[Req, Res any](path string, h HandlerFunc[Req, Res], ds ...int) Route {
	return New[Req, Res](http.MethodPatch, path, h, ds...)
}

func PUT[Req, Res any](path string, h HandlerFunc[Req, Res], ds ...int) Route {
	return New[Req, Res](http.MethodPut, path, h, ds...)
}

func DELETE[Req, Res any](path string, h HandlerFunc[Req, Res], ds ...int) Route {
	return New[Req, Res](http.MethodDelete, path, h, ds...)
}
