package cake

import (
	"context"

	"github.com/matherique/cakesie.app-backend/internal/models"
	"github.com/matherique/cakesie.app-backend/internal/repository"
)

type Creater interface {
	Create(ctx context.Context, data *models.Cake) (*models.Cake, error)
}

type Getter interface {
	GetById(ctx context.Context, id int) (*models.Cake, error)
	GetAllByStatus(ctx context.Context, status bool) ([]*models.Cake, error)
}

type Updater interface {
	ChangeStatus(ctx context.Context, id int, status bool) error
	Update(ctx context.Context, id int, data *models.Cake) (*models.Cake, error)
}

type Cake interface {
	Creater
	Getter
	Updater
}

type cake struct {
	repo repository.CakeRepository
}

func NewCakeApp(repo repository.CakeRepository) Cake {
	return &cake{
		repo: repo,
	}
}

func (c *cake) Create(ctx context.Context, data *models.Cake) (*models.Cake, error) {
	if err := validateToInsert(data); err != nil {
		return nil, err
	}

	data.Status = true

	if err := c.repo.Insert(ctx, data); err != nil {
		return nil, DefaultRepositoryError
	}

	return data, nil
}

func (c *cake) GetAllByStatus(ctx context.Context, status bool) ([]*models.Cake, error) {
	cakes, err := c.repo.GetAllByStatus(ctx, status)
	if err != nil {
		return nil, DefaultRepositoryError
	}

	return cakes, nil
}

func (c *cake) ChangeStatus(ctx context.Context, id int, status bool) error {
	if id == 0 {
		return IdRequiredError
	}

	if err := c.repo.UpdateStatus(ctx, id, status); err != nil {
		return DefaultRepositoryError
	}

	return nil
}

func (c *cake) Update(ctx context.Context, id int, data *models.Cake) (*models.Cake, error) {
	if id == 0 {
		return nil, IdRequiredError
	}

	if err := c.repo.Update(ctx, id, data); err != nil {
		return nil, DefaultRepositoryError
	}

	return data, nil
}

func (c *cake) GetById(ctx context.Context, id int) (*models.Cake, error) {
	if id == 0 {
		return nil, IdRequiredError
	}

	cake, err := c.repo.GetById(ctx, id)

	if err != nil {
		return nil, DefaultRepositoryError
	}

	return cake, nil
}
