package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ArtemKapustkin/CatBreedsProcessor/internal/core"
)

func GetCatsFromResponse(response *http.Response) ([]core.Cat, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var catDataWrapper core.CatDataWrapper

	err = json.Unmarshal(body, &catDataWrapper)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return catDataWrapper.Data, nil
}

func ProcessCats(cats []core.Cat) core.GroupedCats {
	groupedCats := core.GroupCatsByCountry(cats)

	core.SortCatsByBreedName(groupedCats)

	return groupedCats
}
