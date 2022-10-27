# NewRelic Middleware for Gin

Gin middleware for tracking NewRelic web requests.

```go

    r := gin.New()

    appName := viper.GetString("NEW_RELIC_APP_NAME")
	licenseKey := viper.GetString("NEW_RELIC_LICENSE_KEY")

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = true
		},
	) 

    r.Use(ginnewrelic.NewRelicMiddleware(app))

    r.GET("/", func(c *gin.Context) {
        txn := newrelic.FromContext(c.Request.Context())

        // code...
    })

```
