package utils

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type key int

const (
	keyPrincipalID key = iota
)

// GinContextToContextMiddleware retrieves gin context and injects it into main context
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), keyPrincipalID, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

//GinContextFromContext retrieves gin context from main context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(keyPrincipalID)
	if ginContext == nil {
		return nil, errors.New("Oops, something happend, couldn't retrieve gin context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("gin context has wrong type")
	}

	return gc, nil
}
