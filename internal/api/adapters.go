package api

import (
	"context"
	"github.com/dembygenesis/local.tools/internal/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . capturePagesService
type capturePagesService interface {
	ListCapturePages(ctx context.Context, filters *model.CapturePagesFilters) (*model.PaginatedCapturePages, error)
	GetCapturePageByID(ctx context.Context, id int) (*model.CapturePages, error)
}
