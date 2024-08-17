package url

import (
	"fmt"

	"gorm.io/gorm"
)

type UrlRepository struct {
	Db *gorm.DB
}

type UrlDTO struct {
	Value string `json:"value"`
	Alias string `json:"alias,omitempty"`
}

type Url struct {
	gorm.Model
	Value string
	Alias string
}

func (rep *UrlRepository) SaveUrl(dto UrlDTO) (uint, error) {
	const op = "model.saved_url_repo.SaveUrl"

	newUrl := Url{Value: dto.Value, Alias: dto.Alias}

	result := rep.Db.Create(&newUrl)

	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, result.Error)
	}

	return newUrl.ID, nil
}

func (rep *UrlRepository) GetRandomUrl() (*UrlDTO, error) {
	const op = "entity.saved_url_repo.GetRandomUrl"
	var randomUrl Url

	result := rep.Db.Take(&randomUrl)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, result.Error)
	}

	return &UrlDTO{Value: randomUrl.Value, Alias: randomUrl.Alias}, nil
}

func (rep *UrlRepository) GetUrlByAlias(alias string) (*UrlDTO, error) {
	const op = "entity.saved_url_repo.GetUrlByAlias"
	var urlToReturn Url

	result := rep.Db.Find(&urlToReturn, Url{Alias: alias})

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, result.Error)
	}

	return &UrlDTO{Value: urlToReturn.Value, Alias: urlToReturn.Alias}, nil
}

func (rep *UrlRepository) DeleteAlias(alias string) (bool, error) {
	const op = "entity.saved_url_repo.GetUrlByAlias"
	urlToDelete := Url{Alias: alias}

	result := rep.Db.Delete(urlToDelete)

	if result.Error != nil {
		return false, fmt.Errorf("%s: %w", op, result.Error)
	}

	return true, nil
}
