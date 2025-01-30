package usecase

import (
	"context"
	"github.com/farhodm/banner-service-clean-arch/internal/domain"
	"math/rand"
	"time"
)

type BannerUseCase struct {
	repo domain.BannerRepository
}

func NewBannerUseCase(repo domain.BannerRepository) *BannerUseCase {
	return &BannerUseCase{
		repo: repo,
	}
}

func (uc *BannerUseCase) GetBanner(ctx context.Context, id int) (*domain.Banner, error) {
	return uc.repo.GetByID(id)
}

func (uc *BannerUseCase) CreateBanner(ctx context.Context, title, content string) (*domain.Banner, error) {
	n := rand.Intn(10)
	banner := &domain.Banner{
		ID:        n,
		Title:     title,
		Content:   content,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.repo.Create(banner); err != nil {
		return nil, err
	}
	return banner, nil
}

func (uc *BannerUseCase) UpdateBanner(ctx context.Context, id int, title, content string) error {
	banner, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}

	banner.Title = title
	banner.Content = content
	banner.UpdatedAt = time.Now()

	return uc.repo.Update(banner)
}

func (uc *BannerUseCase) GetAllBanners(ctx context.Context) []domain.Banner {
	return uc.repo.GetAll()
}

func (uc *BannerUseCase) DeleteBanner(ctx context.Context, id int) error {
	return uc.repo.Delete(id)
}
