package notion

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

var (
	databaseCache    DatabaseResult
	databaseCacheSet time.Time
)

const CACHE_TIMEOUT = 1 * time.Minute

func FetchDatabase() (DatabaseResult, error) {
	if time.Since(databaseCacheSet) < CACHE_TIMEOUT {
		return databaseCache, nil
	}

	db_id := os.Getenv("NOTION_DB_ID")
	url := "https://api.notion.com/v1/databases/" + db_id + "/query"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_TOKEN"))
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := client.Do(req)
	if err != nil {
		return DatabaseResult{}, err
	}
	defer res.Body.Close()

	var j DatabaseResult
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return DatabaseResult{}, err
	}

	databaseCache = j
	databaseCacheSet = time.Now()

	return j, nil
}
