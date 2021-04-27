package search

// API contains the API Key and Secret for sending requests to PodcastIndex.org
type API struct {
	ApiKey    string
	ApiSecret string
}

func getData(query string) string {
	body, err := request.Send(searchQuery, q.ApiKey, q.ApiSecret)
	if err != nil {
		return body, fmt.Errorf("getData() failed: %s", err)
	}

}
