package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ArtemKapustkin/CatBreedsProcessor/internal/core"
)

func CreateOutJSON(groupedCats core.GroupedCats, filename string) error {
	var data []map[string][]core.Cat

	for country, countryCats := range groupedCats {
		countryMap := map[string][]core.Cat{
			country: countryCats,
		}
		data = append(data, countryMap)
	}

	outJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating a file: %w", err)
	}

	defer file.Close()

	_, err = file.Write(outJSON)
	if err != nil {
		return fmt.Errorf("error writing json to file: %w", err)
	}

	return nil
}
