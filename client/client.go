package client

import "github.com/Alperen10/Kandilli-API/models"

type Client interface {
	ParseToLines() ([]models.Earthquake, error)
	Fetch()
}
