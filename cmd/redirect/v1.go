package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	csv "github.com/ivynya/redirect/internal/csv"
	notion "github.com/ivynya/redirect/internal/notion"
)

func createRouter(a *fiber.App) {
	a.Get("/*", func(c *fiber.Ctx) error {
		// find the page with the matching short ID
		var page notion.Page

		// if notion db is not set, try csv mode
		db_id := os.Getenv("NOTION_DB_ID")
		csv_url := os.Getenv("CSV_URL")
		if db_id == "" && csv_url != "" {
			log.Println("NOTION_DB_ID not set, trying CSV_URL...")
			pages, err := csv.GetPagesFromFile(csv_url)
			if err != nil {
				log.Println("Error fetching CSV: " + err.Error())
				return c.SendStatus(500)
			}
			for _, p := range pages {
				if trimLeadingSlash(p.ShortID) == c.Params("*") {
					page = p
					break
				}
			}
		} else if db_id != "" {
			log.Println("Fetching page from Notion...")
			db, err := notion.FetchDatabase()
			if err != nil {
				return c.SendStatus(500)
			}
			// convert db.results into Page objects
			pages := make([]notion.Page, len(db.Results))
			for i, p := range db.Results {
				pages[i] = notion.ConvertPageResult(p)
			}
			for _, p := range pages {
				if trimLeadingSlash(p.ShortID) == c.Params("*") {
					page = p
					break
				}
			}
		} else {
			log.Println("NOTION_DB_ID and CSV_URL not set")
			return c.SendStatus(500)
		}

		// if no page was found, return 404
		if page == (notion.Page{}) {
			log.Println("Page not found: " + c.Params("*"))
			return c.Status(404).SendString("Page not found")
		}

		// if ANALYTICS_API and page.CampaignID defined, make POST request
		ANALYTICS_HOST := os.Getenv("ANALYTICS_HOST")
		ANALYTICS_VER := os.Getenv("ANALYTICS_VERSION")
		analyticsBase := "https://" + ANALYTICS_HOST + "/" + ANALYTICS_VER
		analyticsURL := analyticsBase + "/campaign/" + page.CampaignID
		if ANALYTICS_HOST != "" && ANALYTICS_VER != "" && page.CampaignID != "" {
			log.Println("[A] " + analyticsURL)
			res, err := http.Post(analyticsURL, "application/json", nil)
			if err != nil {
				log.Println("Analytics API error: " + err.Error())
				return c.Status(500).SendString("Analytics API error")
			}
			defer res.Body.Close()
		}

		// if a page was found, redirect to it
		log.Println("[>] " + page.ShortID + " -> " + page.RedirectURL + "...")
		return c.Redirect(page.RedirectURL, 302)
	})
}

func trimLeadingSlash(s string) string {
	if len(s) > 0 && s[0] == '/' {
		return s[1:]
	}
	return s
}
