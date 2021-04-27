package search

func getData(query string) string {
	body, err := request.Send(searchQuery, q.ApiKey, q.ApiSecret)
	if err != nil {
		return body, fmt.Errorf("getData() failed: %s", err)
	}

}
