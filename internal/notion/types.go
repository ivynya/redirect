package notion

type DatabaseResult struct {
	Results    []PageResult `json:"results"`
	NextCursor string       `json:"next_cursor"`
	HasMore    bool         `json:"has_more"`
}

type PageResult struct {
	NotionID   string `json:"id"`
	Properties struct {
		Name struct {
			Title []struct {
				PlainText string `json:"plain_text"`
			} `json:"title"`
		} `json:"Name"`
		Short struct {
			RichText []struct {
				PlainText string `json:"plain_text"`
			} `json:"rich_text"`
		} `json:"Short"`
		RedirectURL struct {
			URL string `json:"url"`
		} `json:"RedirectURL"`
		CampaignID struct {
			RichText []struct {
				PlainText string `json:"plain_text"`
			} `json:"rich_text"`
		} `json:"CampaignID"`
	} `json:"properties"`
}

type Page struct {
	NotionID    string `json:"id"`
	Name        string `json:"name"`
	ShortID     string `json:"short"`
	RedirectURL string `json:"redirect_url"`
	CampaignID  string `json:"campaign_id"`
}
