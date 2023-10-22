package usecase

import (
	"encoding/json"
	"errors"
	"net/http"

	"effective/internal/entity"
	"effective/internal/repository/postgres"

	resty "github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

const (
	ageURI         = "https://api.agify.io/?name="
	sexURI         = "https://api.genderize.io/?name="
	nationalityURI = "https://api.nationalize.io/?name="
)

// HumansUseCase contains humans repository and humans services.
type HumansUseCase struct {
	repo *postgres.HumansRepository
}

// NewHumansUseCase creates new *HumansUseCase.
func NewHumansUseCase(repo *postgres.HumansRepository) *HumansUseCase {
	return &HumansUseCase{
		repo: repo,
	}
}

func parseAge(name string) (int, error) {
	rs := resty.New()
	res, err := rs.R().Get(ageURI + name)
	if err != nil {
		return 0, err
	} else if res.StatusCode() != http.StatusOK {
		return 0, errors.New(string(res.Body()))
	}

	type AgeResponse struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}

	var data AgeResponse
	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return 0, err
	}

	return data.Age, nil
}

func parseGender(name string) (string, error) {
	rs := resty.New()
	res, err := rs.R().Get(sexURI + name)
	if err != nil {
		return "", err
	} else if res.StatusCode() != http.StatusOK {
		return "", errors.New(string(res.Body()))
	}

	type GenderResponse struct {
		Count       int     `json:"count"`
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Probability float32 `json:"probability"`
	}

	var data GenderResponse
	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return "", err
	}

	return data.Gender, nil
}

func parseNationality(name string) (string, error) {
	rs := resty.New()
	res, err := rs.R().Get(nationalityURI + name)
	if err != nil {
		return "", err
	} else if res.StatusCode() != http.StatusOK {
		return "", errors.New(string(res.Body()))
	}

	type NationalityResponse struct {
		Count   int    `json:"count"`
		Name    string `json:"name"`
		Country []struct {
			CountryID   string  `json:"country_id"`
			Probability float32 `json:"probability"`
		} `json:"country"`
	}

	var data NationalityResponse
	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return "", err
	}

	return data.Country[0].CountryID, nil
}

// Create new human.
func (uc *HumansUseCase) Create(humanData *entity.CreateHuman) (*entity.Human, error) {
	human := &entity.Human{
		Name:       humanData.Name,
		Surname:    humanData.Surname,
		Patronymic: humanData.Patronymic,
	}

	// Age
	age, err := parseAge(human.Name)
	if err != nil {
		log.Error().Err(err).Msg("error while parsing age")
		return nil, err
	}
	human.Age = age

	// Gender
	gender, err := parseGender(human.Name)
	if err != nil {
		log.Error().Err(err).Msg("error while parsing sex")
		return nil, err
	}
	human.Gender = gender

	// Nationality
	nationality, err := parseNationality(human.Name)
	if err != nil {
		log.Error().Err(err).Msg("error while parsing nationality")
		return nil, err
	}
	human.Nation = nationality

	return uc.repo.CreateHuman(human)
}

// Update existing human.
func (uc *HumansUseCase) Update(id int, humanData *entity.UpdateHuman) (*entity.Human, error) {
	human, err := uc.repo.GetHuman(id)
	if err != nil {
		log.Error().Err(err).Msgf("error while getting human; id: %v", id)
		return nil, err
	}

	if humanData.Name != nil {
		human.Name = *humanData.Name

		human.Age, err = parseAge(human.Name)
		if err != nil {
			log.Error().Err(err).Msg("error while parsing age")
			return nil, err
		}

		human.Gender, err = parseGender(human.Name)
		if err != nil {
			log.Error().Err(err).Msg("error while parsing sex")
			return nil, err
		}

		human.Nation, err = parseNationality(human.Name)
		if err != nil {
			log.Error().Err(err).Msg("error while parsing nationality")
		}
	}

	if humanData.Surname != nil {
		human.Surname = *humanData.Surname
	}
	if humanData.Patronymic != nil {
		human.Patronymic = *humanData.Patronymic
	}

	return human, uc.repo.UpdateHuman(human)
}

// Delete existing human.
func (uc *HumansUseCase) Delete(id int) error {
	return uc.repo.DeleteHuman(id)
}

// Get all humans by filters.
func (uc *HumansUseCase) Get(filter *entity.HumanFilter) (*entity.HumansList, error) {
	// Default values
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.PageSize == 0 {
		filter.PageSize = 10
	}

	return uc.repo.GetHumans(filter)
}
