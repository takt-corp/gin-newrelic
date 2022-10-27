package ginnewrelic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// NewRelicMiddleware the Gin Handler the same way NewRelic wraps a raw HTTP request
// https://github.com/newrelic/go-agent/blob/v3.15.2/v3/newrelic/instrumentation.go#L31-L46
func NewRelicMiddleware(client *newrelic.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		if client != nil {
			// format the the request
			txnName := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)

			// create the transaction
			txn := client.StartTransaction(txnName)
			defer txn.End()

			// add the request id to the request attributes
			requestID := c.GetHeader("X-Request-ID")

			if requestID != "" {
				txn.AddAttribute("request_id", requestID)
			}

			// set the transaction to the request context
			ctx := c.Request.Context()
			ctx = newrelic.NewContext(ctx, txn)
			c.Request = c.Request.WithContext(ctx)

			// configure the request parameters
			// see https://github.com/newrelic/go-agent/blob/4b46fc5b40e5398695413097b5c461679a74eafd/context.go
			txn.SetWebRequestHTTP(c.Request)
			txn.SetWebResponse(c.Writer)
		}
	}
}
