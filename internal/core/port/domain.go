package port

import (
	"context"
	"time"

	"github.com/Geovanny0401/clubhub/internal/core/domain"
)

type DomainRepo interface {
	CreateDomain(ctx context.Context, domain domain.Domain) (int64, error)
	CreateDetailDomain(ctx context.Context, detailDomain domain.DetailDomain) error
	UpdateLastGetDomain(ctx context.Context, idDomain int64, date time.Time) error
	GetAllDomain(ctx context.Context) ([]domain.Domain, error)
	GetDomainByAddress(ctx context.Context, address string) (domain.Domain, error)
	GetDetailsByDomain(ctx context.Context, idDomain int64, countServer int) ([]domain.DetailDomain, error)
}