// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "ComingServer": pods TestHelpers
//
// Command:
// $ goagen
// --design=github.com/TinyKitten/ComingServer/design
// --out=$(GOPATH)/src/github.com/TinyKitten/ComingServer
// --version=v1.3.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// AddPodsBadRequest runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddPodsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, payload *app.AddPodsPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	addCtx, __err := app.NewAddPodsContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	addCtx.Payload = payload

	// Perform action
	__err = ctrl.Add(addCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// AddPodsCreated runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddPodsCreated(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, payload *app.AddPodsPayload) (http.ResponseWriter, *app.Pod) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	addCtx, __err := app.NewAddPodsContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	addCtx.Payload = payload

	// Perform action
	__err = ctrl.Add(addCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	var mt *app.Pod
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.Pod)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Pod", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// AddPodsInternalServerError runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddPodsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, payload *app.AddPodsPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	addCtx, __err := app.NewAddPodsContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	addCtx.Payload = payload

	// Perform action
	__err = ctrl.Add(addCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ListPodsInternalServerError runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListPodsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, limit int, offset int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/pods"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	listCtx, _err := app.NewListPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ListPodsNotFound runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListPodsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, limit int, offset int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/pods"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	listCtx, _err := app.NewListPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ListPodsOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListPodsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, limit int, offset int) (http.ResponseWriter, app.PodCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(offset)}
		query["offset"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/v1/pods"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(offset)}
		prms["offset"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	listCtx, _err := app.NewListPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.PodCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.PodCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.PodCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// RegenerateTokenPodsBadRequest runs the method RegenerateToken of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func RegenerateTokenPodsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v/token", id),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	regenerateTokenCtx, _err := app.NewRegenerateTokenPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.RegenerateToken(regenerateTokenCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// RegenerateTokenPodsInternalServerError runs the method RegenerateToken of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func RegenerateTokenPodsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v/token", id),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	regenerateTokenCtx, _err := app.NewRegenerateTokenPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.RegenerateToken(regenerateTokenCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// RegenerateTokenPodsNotFound runs the method RegenerateToken of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func RegenerateTokenPodsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v/token", id),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	regenerateTokenCtx, _err := app.NewRegenerateTokenPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.RegenerateToken(regenerateTokenCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// RegenerateTokenPodsOK runs the method RegenerateToken of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func RegenerateTokenPodsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int) (http.ResponseWriter, *app.Token) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v/token", id),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	regenerateTokenCtx, _err := app.NewRegenerateTokenPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.RegenerateToken(regenerateTokenCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Token
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Token)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Token", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowPodsInternalServerError runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowPodsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	showCtx, _err := app.NewShowPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowPodsNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowPodsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	showCtx, _err := app.NewShowPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowPodsOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowPodsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int) (http.ResponseWriter, *app.Pod) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	showCtx, _err := app.NewShowPodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Pod
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Pod)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Pod", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// UpdatePodsBadRequest runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdatePodsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int, payload *app.UpdatePodsPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v", id),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	updateCtx, _err := app.NewUpdatePodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}
	updateCtx.Payload = payload

	// Perform action
	_err = ctrl.Update(updateCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// UpdatePodsInternalServerError runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdatePodsInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int, payload *app.UpdatePodsPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v", id),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	updateCtx, _err := app.NewUpdatePodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}
	updateCtx.Payload = payload

	// Perform action
	_err = ctrl.Update(updateCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// UpdatePodsNoContent runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdatePodsNoContent(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int, payload *app.UpdatePodsPayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v", id),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	updateCtx, _err := app.NewUpdatePodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	updateCtx.Payload = payload

	// Perform action
	_err = ctrl.Update(updateCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

	// Return results
	return rw
}

// UpdatePodsNotFound runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdatePodsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.PodsController, id int, payload *app.UpdatePodsPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/v1/pods/%v", id),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PodsTest"), rw, req, prms)
	updateCtx, _err := app.NewUpdatePodsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}
	updateCtx.Payload = payload

	// Perform action
	_err = ctrl.Update(updateCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}