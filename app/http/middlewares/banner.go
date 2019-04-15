package middlewares

import (
	"crucial/banner/app/config"
	"crucial/banner/app/http/models"
	"crucial/banner/libs/mux"
	"log"
	"net/http"
)

// CheckAPIKey - проверка
func CheckAPIKey(handler mux.Handler) mux.Handler {
	return func(ctx mux.Context) {
		var (
			request  = ctx.Request()
			response models.Response
			token    string
			//err      error
		)

		token = request.Header.Get("Authorization")
		if len(token) == 0 {
			log.Println("попытка к роуту  ", request.URL, " без авторизации")
			response.Send(ctx.Response(), http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
			return
		}

		cfg := config.Peek()
		if cfg.App.APIKey != token {
			response.Send(ctx.Response(), http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
			return
		}

		handler(ctx)
	}
}
