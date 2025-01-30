package memory

import (
	"github.com/farhodm/banner-service-clean-arch/internal/domain"
	"sync"
)

type InMemoryBannerRepository struct {
	banners map[int]*domain.Banner
	mu      sync.RWMutex
}

func NewInMemoryBannerRepository() *InMemoryBannerRepository {
	return &InMemoryBannerRepository{
		banners: make(map[int]*domain.Banner),
	}
}

func (r *InMemoryBannerRepository) GetByID(id int) (*domain.Banner, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	banner, exists := r.banners[id]
	if !exists {
		return nil, domain.ErrBannerNotFound
	}

	return banner, nil
}

func (r *InMemoryBannerRepository) Create(banner *domain.Banner) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.banners[banner.ID] = banner

	return nil
}

func (r *InMemoryBannerRepository) Update(banner *domain.Banner) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.banners[banner.ID]
	if !exists {
		return domain.ErrBannerNotFound
	}

	r.banners[banner.ID] = banner

	return nil
}

func (r *InMemoryBannerRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.banners[id]
	if !exists {
		return domain.ErrBannerNotFound
	}

	delete(r.banners, id)

	return nil
}

func (r *InMemoryBannerRepository) GetAll() []domain.Banner {
	var banners []domain.Banner
	for _, banner := range r.banners {
		banners = append(banners, *banner)
	}

	return banners
}
