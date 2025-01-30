package main

import (
	"github.com/farhodm/banner-service-clean-arch/internal/delivery/handlers"
	"github.com/farhodm/banner-service-clean-arch/internal/repository/memory"
	"github.com/farhodm/banner-service-clean-arch/internal/usecase"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func routes(router *httprouter.Router) {
	repo := memory.NewInMemoryBannerRepository()
	useCase := usecase.NewBannerUseCase(repo)
	handler := handlers.NewBannerHandler(useCase)

	router.POST("/banners", handler.CreateBanner)
	router.GET("/banners", handler.GetAllBanners)
	router.GET("/banners/:id", handler.GetBanner)
	router.PUT("/banners/:id", handler.UpdateBanner)
	router.DELETE("/banners/:id", handler.DeleteBanner)

	if err := http.ListenAndServe(":4000", router); err != nil {
		panic(err)
	}
}
