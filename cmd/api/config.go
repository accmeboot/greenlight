package main

import (
	"flag"
	"os"
	"strings"
)

type config struct {
	port int
	env  string

	db struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}

	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}

	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}

	cors struct {
		trustedOrigins []string
	}
}

func (c *config) Parse() {
	defaultDsn := os.Getenv("GREENLIGHT_DB_DSN")

	flag.IntVar(&c.port, "port", 4000, "API server port")
	flag.StringVar(&c.env, "env", "development", "Environment (development|staging|production")
	flag.StringVar(&c.db.dsn, "db-dsn", defaultDsn, "PostgreSQL DSN")

	flag.IntVar(&c.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&c.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&c.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max idle connection time")

	flag.Float64Var(&c.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&c.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&c.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&c.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&c.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&c.smtp.username, "smtp-username", "8d564615350280", "SMTP username")
	flag.StringVar(&c.smtp.password, "smtp-password", "0c0a2366e54055", "SMTP password")
	flag.StringVar(&c.smtp.sender, "smtp-sender", "Greenlight <no-reply@greenlight.accme.net>", "SMTP sender")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		c.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()
}
