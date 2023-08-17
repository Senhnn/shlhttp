package app

import "context"

// TODO
type HandlerFunc func(c context.Context, ctx *RequestContext)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc
