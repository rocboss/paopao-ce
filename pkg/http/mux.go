// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package http

import (
	"net/http"
	"strings"
	"sync"
)

// ConnectMux mux used for Connect
type ConnectMux struct {
	mu sync.RWMutex
	m  muxMap[http.Handler]
}

type muxMap[T any] interface {
	set(path string, val T) bool
	get(path string) (T, bool)
	match(pattern string) (T, bool)
}

type simpleMuxMap[T any] map[string]T

type prefixMuxMap[T any] struct {
	prefix string
	m      map[string]T
}

func (m simpleMuxMap[T]) set(path string, val T) bool {
	if _, exist := m[path]; exist {
		return false
	}
	m[path] = val
	return true
}

func (m simpleMuxMap[T]) get(path string) (val T, exist bool) {
	val, exist = m[path]
	return
}

// match assume pattern like `/core.v1.AuthenticateService/login`
func (m simpleMuxMap[T]) match(pattern string) (val T, exist bool) {
	idx := strings.IndexByte(pattern[1:], '/')
	if idx < 0 {
		return
	}
	return m.get(pattern[:idx+2])
}

func (m *prefixMuxMap[T]) set(path string, val T) bool {
	if _, exist := m.m[path]; exist {
		return false
	}
	m.m[path] = val
	return true
}

func (m *prefixMuxMap[T]) get(path string) (val T, exist bool) {
	val, exist = m.m[path]
	return
}

// match assume pattern like `/core.v1.AuthenticateService/login`
func (m *prefixMuxMap[T]) match(pattern string) (val T, exist bool) {
	path, _ := strings.CutPrefix(pattern, m.prefix)
	idx := strings.IndexByte(path[1:], '/')
	if idx < 0 {
		return
	}
	return m.get(path[:idx+2])
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *ConnectMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h := mux.handler(r.URL.Path)
	h.ServeHTTP(w, r)
}

// Handle registers the handler for the given path.
// If a handler already exists for path, Handle panics.
func (mux *ConnectMux) Handle(path string, handler http.Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if path == "" {
		panic("http: invalid pattern")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if ok := mux.m.set(path, handler); !ok {
		panic("http: multiple registrations for " + path)
	}
}

// handler is the main implementation of Handler.
// The path is known to be in canonical form, except for CONNECT methods.
func (mux *ConnectMux) handler(path string) (h http.Handler) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	if h, _ = mux.m.match(path); h == nil {
		h = http.NotFoundHandler()
	}
	return
}

// NewConnectMux allocates and returns a new ConnectMux.
func NewConnectMux(pathPrefix ...string) *ConnectMux {
	var m muxMap[http.Handler] = make(simpleMuxMap[http.Handler])
	if len(pathPrefix) > 0 {
		m = &prefixMuxMap[http.Handler]{
			m:      make(map[string]http.Handler),
			prefix: pathPrefix[0],
		}
	}
	return &ConnectMux{m: m}
}
