package notion

func ConvertPageResult(p PageResult) Page {
	// Get the campaign ID from the result
	campaignID := ""
	for _, richText := range p.Properties.CampaignID.RichText {
		campaignID += richText.PlainText
	}

	// Get the short ID from the result
	shortID := ""
	for _, richText := range p.Properties.Short.RichText {
		shortID += richText.PlainText
	}

	// Get the title from the result
	title := ""
	for _, richText := range p.Properties.Name.Title {
		title += richText.PlainText
	}

	// Return the formatted page object
	return Page{
		NotionID:    p.NotionID,
		Name:        title,
		ShortID:     shortID,
		RedirectURL: p.Properties.RedirectURL.URL,
		CampaignID:  campaignID,
	}
}
