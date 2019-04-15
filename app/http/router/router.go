package router

import (
	"crucial/banner/app/config"
	"crucial/banner/app/db"
	"crucial/banner/app/domain/banner/usecase"
	"crucial/banner/app/http/handlers"
	"crucial/banner/app/http/middlewares"
	"crucial/banner/libs/mux"
	"log"
)

var router mux.Router

// Peek provides secure access to router instance.
func Peek() mux.Router {
	return router
}

func init() {
	var (
		cfg      = config.Peek().Server
		version  = "/v1"
		prefix   = cfg.Prefix + version
		settings mux.RouterHTTPSettings
		err      error

		// инициализация репозитория для стуктуры хэндлера
		bannerRepository = db.NewPsqlBannerRepository()
		bannerUsecase    = usecase.NewBannerUsecase(bannerRepository)
		bannerHandlers   = handlers.NewBannerHandler(bannerRepository, bannerUsecase)
	)

	settings = mux.NewRouterHTTPSettings()
	router, err = mux.NewRouter(settings)
	if err != nil {
		log.Panicln("router: " + err.Error())
	}

	router.POST(prefix+"/banner", bannerHandlers.Create, middlewares.CheckAPIKey)
	router.DELETE(prefix+"/banner/:id", bannerHandlers.Destroy, middlewares.CheckAPIKey)
	router.GET(prefix+"/banner", bannerHandlers.Search)

	log.Printf("router: assigned all endpoints to 127.0.0.1%s\n", cfg.Port+prefix)
}
