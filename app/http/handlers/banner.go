package handlers

import (
	"crucial/banner/app/domain/banner/entity"
	"crucial/banner/app/http/models"
	"crucial/banner/libs/mux"
	"log"
	"net/http"
	"strconv"
)

// Create - сохранения баннера
func (this *BannerHandler) Create(ctx mux.Context) {
	var (
		response models.Response
		request  models.SaveBannerRequest
		err      error
	)

	err = ctx.ParseJsonPayload(&request)
	if err != nil {
		response.Send(ctx.Response(), http.StatusBadRequest, "данные плохого формата", nil)
		return
	}
	ok, message := request.Validate()
	if !ok {
		response.Send(ctx.Response(), http.StatusBadRequest, message, nil)
		return
	}

	err = this.usecase.Create(request)
	if err != nil {
		log.Println("packege handlers, file banner.go, method Create() error : ", err)
		response.Send(ctx.Response(), http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Send(ctx.Response(), http.StatusOK, http.StatusText(http.StatusOK), nil)

}

// Destroy - удаление
func (this *BannerHandler) Destroy(ctx mux.Context) {
	var (
		response models.Response
		id       = ctx.Params().ByName("id")
		err      error
	)

	bannerID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(1, "packege handlers, file banner.go, method Destroy() error : ", err)
		response.Send(ctx.Response(), http.StatusBadRequest, "id должно являться числом", nil)
		return
	}

	err = this.usecase.Destroy(bannerID)
	if err != nil {
		log.Println(2, "packege handlers, file banner.go, method Destroy() error : ", err)
		response.Send(ctx.Response(), http.StatusBadRequest, "непарвильный id", nil)
		return
	}

	response.Send(ctx.Response(), http.StatusOK, http.StatusText(http.StatusOK), nil)

}

// Search - поиск баннеров по критерия
func (this *BannerHandler) Search(ctx mux.Context) {
	var (
		request  = ctx.Request()
		response models.Response
		err      error
	)

	queryValues := request.URL.Query()
	searchCreteria := entity.SearchBanner{
		Title: queryValues.Get("title"),
		Brand: queryValues.Get("brand"),
		Size:  queryValues.Get("size"),
	}

	banners, err := this.usecase.Search(&searchCreteria)
	if err != nil {
		log.Println("packege handlers, file banner.go, method Search() error : ", err)
		response.Send(ctx.Response(), http.StatusBadRequest, "записи не найдены", nil)
		return
	}

	response.Send(ctx.Response(), http.StatusOK, http.StatusText(http.StatusOK), banners)
}
