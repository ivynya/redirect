package csv

import (
	"encoding/csv"
	"os"

	"github.com/ivynya/redirect/internal/notion"
)

func GetPagesFromFile(path string) ([]notion.Page, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	pages := make([]notion.Page, len(records)-1)
	for i, record := range records[1:] {
		pages[i] = notion.Page{
			ShortID:     record[4],
			RedirectURL: record[3],
			CampaignID:  record[1],
		}
	}

	return pages, nil
}
