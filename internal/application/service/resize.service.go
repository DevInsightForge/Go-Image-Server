package service

import (
	"image"
	"image/jpeg"
	"net/http"

	"image-server/internal/application/utility"
	"image-server/internal/domain/dto"

	"github.com/disintegration/imaging"
)

func ResizeHandler(w http.ResponseWriter, r *http.Request) {
	requestQueryDto := dto.ResizeRequestQueryDto{}
	queryParams := r.URL.Query()
	utility.ParseQueryParams(queryParams, &requestQueryDto)

	if requestQueryDto.ImageUrl == "" || requestQueryDto.Width <= 0 || requestQueryDto.Height <= 0 {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	if requestQueryDto.Quality == 0 {
		requestQueryDto.Quality = 75
	}

	resp, err := http.Get(requestQueryDto.ImageUrl)
	if err != nil {
		http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	resizedImg := imaging.Resize(img, requestQueryDto.Width, requestQueryDto.Height, imaging.Lanczos)

	w.Header().Set("Content-Type", "image/jpeg")

	err = jpeg.Encode(w, resizedImg, &jpeg.Options{Quality: requestQueryDto.Quality})
	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}
}
