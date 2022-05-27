// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (DELETE /envs)
	DeleteEnvs(c *gin.Context)

	// (POST /envs)
	PostEnvs(c *gin.Context)

	// (POST /envs/state)
	PostEnvsState(c *gin.Context)

	// (GET /envs/{codeSnippetID})
	GetEnvsCodeSnippetID(c *gin.Context, codeSnippetID string)

	// (GET /health)
	GetHealth(c *gin.Context)

	// (GET /sessions)
	GetSessions(c *gin.Context)

	// (POST /sessions)
	PostSessions(c *gin.Context)

	// (DELETE /sessions/{sessionID})
	DeleteSessionsSessionID(c *gin.Context, sessionID string)

	// (PUT /sessions/{sessionID}/refresh)
	PutSessionsSessionIDRefresh(c *gin.Context, sessionID string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// DeleteEnvs operation middleware
func (siw *ServerInterfaceWrapper) DeleteEnvs(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteEnvs(c)
}

// PostEnvs operation middleware
func (siw *ServerInterfaceWrapper) PostEnvs(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostEnvs(c)
}

// PostEnvsState operation middleware
func (siw *ServerInterfaceWrapper) PostEnvsState(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostEnvsState(c)
}

// GetEnvsCodeSnippetID operation middleware
func (siw *ServerInterfaceWrapper) GetEnvsCodeSnippetID(c *gin.Context) {

	var err error

	// ------------- Path parameter "codeSnippetID" -------------
	var codeSnippetID string

	err = runtime.BindStyledParameter("simple", false, "codeSnippetID", c.Param("codeSnippetID"), &codeSnippetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter codeSnippetID: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetEnvsCodeSnippetID(c, codeSnippetID)
}

// GetHealth operation middleware
func (siw *ServerInterfaceWrapper) GetHealth(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetHealth(c)
}

// GetSessions operation middleware
func (siw *ServerInterfaceWrapper) GetSessions(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetSessions(c)
}

// PostSessions operation middleware
func (siw *ServerInterfaceWrapper) PostSessions(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostSessions(c)
}

// DeleteSessionsSessionID operation middleware
func (siw *ServerInterfaceWrapper) DeleteSessionsSessionID(c *gin.Context) {

	var err error

	// ------------- Path parameter "sessionID" -------------
	var sessionID string

	err = runtime.BindStyledParameter("simple", false, "sessionID", c.Param("sessionID"), &sessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter sessionID: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteSessionsSessionID(c, sessionID)
}

// PutSessionsSessionIDRefresh operation middleware
func (siw *ServerInterfaceWrapper) PutSessionsSessionIDRefresh(c *gin.Context) {

	var err error

	// ------------- Path parameter "sessionID" -------------
	var sessionID string

	err = runtime.BindStyledParameter("simple", false, "sessionID", c.Param("sessionID"), &sessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter sessionID: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PutSessionsSessionIDRefresh(c, sessionID)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.DELETE(options.BaseURL+"/envs", wrapper.DeleteEnvs)

	router.POST(options.BaseURL+"/envs", wrapper.PostEnvs)

	router.POST(options.BaseURL+"/envs/state", wrapper.PostEnvsState)

	router.GET(options.BaseURL+"/envs/:codeSnippetID", wrapper.GetEnvsCodeSnippetID)

	router.GET(options.BaseURL+"/health", wrapper.GetHealth)

	router.GET(options.BaseURL+"/sessions", wrapper.GetSessions)

	router.POST(options.BaseURL+"/sessions", wrapper.PostSessions)

	router.DELETE(options.BaseURL+"/sessions/:sessionID", wrapper.DeleteSessionsSessionID)

	router.PUT(options.BaseURL+"/sessions/:sessionID/refresh", wrapper.PutSessionsSessionIDRefresh)

	return router
}
