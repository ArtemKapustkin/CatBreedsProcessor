package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtemKapustkin/CatBreedsProcessor/pkg/utils"
)

type App struct {
	apiURL   string
	filename string
}

func New() *App {
	return &App{
		apiURL:   "https://catfact.ninja/breeds",
		filename: "out.json",
	}
}

func (a *App) Run() error {
	response, err := http.Get(a.apiURL)
	if err != nil {
		return fmt.Errorf("error http get request: %w", err)
	}

	defer response.Body.Close()

	cats, err := utils.GetCatsFromResponse(response)
	if err != nil {
		return fmt.Errorf("error get cats from response: %w", err)
	}

	groupedCats := utils.ProcessCats(cats)

	err = utils.CreateOutJSON(groupedCats, a.filename)
	if err != nil {
		return fmt.Errorf("error create out.json: %w", err)
	}

	log.Println("out.json file created")

	return nil
}
