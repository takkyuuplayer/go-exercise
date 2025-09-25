package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	{
		eg := e.Group("/cookie")
		eg.Use(session.MiddlewareWithConfig(session.Config{
			Store: sessions.NewCookieStore([]byte("secret")),
		}))
		sessionRoutes(eg, "cookie-session")
	}
	{
		var db *redis.Client
		if host, ok := os.LookupEnv("REDIS_HOST"); ok && host != "" {
			db = redis.NewClient(&redis.Options{
				Addr: host,
			})
		} else {
			db = redis.NewClient(&redis.Options{})
		}
		store, err := redisstore.NewRedisStore(context.Background(), db)
		if err != nil {
			log.Fatal("failed to create redis store: ", err)
		}

		eg := e.Group("/redis")
		eg.Use(session.MiddlewareWithConfig(session.Config{
			Store: store,
		}))
		sessionRoutes(eg, "redis-session")
	}

	e.Logger.Fatal(e.Start("localhost:0"))
}

func sessionRoutes(eg *echo.Group, key string) {
	eg.GET("/create-session", func(c echo.Context) error {
		sess, err := session.Get(key, c)
		if err != nil {
			return err
		}
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["nonce"] = uuid.New().String()
		sess.Values["state"] = uuid.New().String()
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}
		return c.NoContent(http.StatusOK)
	})
	eg.GET("/read-session", func(c echo.Context) error {
		sess, err := session.Get(key, c)
		if err != nil {
			return err
		}
		defer func() {
			sess.Options.MaxAge = -1
			sess.Save(c.Request(), c.Response())
		}()

		return c.String(http.StatusOK, fmt.Sprintf("nonce=%v\nstate=%v\n", sess.Values["nonce"], sess.Values["state"]))
	})
}
