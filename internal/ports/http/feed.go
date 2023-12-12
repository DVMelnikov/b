package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/storage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type FeedHandler struct {
	r    handler.Renderer
	feed storage.Storage
}

func NewFeedHandler(r handler.Renderer, feed storage.Storage) *FeedHandler {
	return &FeedHandler{
		r:    r,
		feed: feed,
	}
}

func (h *FeedHandler) BasePrefix() string {
	return "/feed"
}

func (h *FeedHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllFeed))
	r.Get("/{id}", h.r.Wrap(h.GetFeed))

	return r
}

func (h *FeedHandler) GetAllFeed(w http.ResponseWriter, r *http.Request) {
	res, err := h.feed.GetAllFeed(context.Background())

	if err != nil {
		log.WithField("", "GetAllFeed").Error(err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *FeedHandler) GetFeed(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		log.WithField("", "GetFeed").Error(err)
		return handler.BadRequestResponse()
	}

	res, err := h.feed.GetFeed(context.Background(), id)

	if err != nil {
		log.WithField("", "GetFeed").Error(err)
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}
