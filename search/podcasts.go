package search

import (
	"fmt"
	"net/url"
)

// ByTerm returns a slice of podcasts containing the term passed in.
func (q *API) ByTerm(term string) (string, error) {
	return getData(url.QueryEscape(fmt.Sprintf("search/byterm?q=%s", term)))
}

// ByPerson returs a slice of podcasts by a person's name.
func (q *API) ByPerson(name string) (string, error) {
	return getData(url.QueryEscape(fmt.Sprintf("search/byperson?q=%s", name)))
}

// ByFeedID returns details about a podcast by Podcast Index ID.
func (q *API) ByFeedID(id string) (string, error) {
	return getData(url.QueryEscape(fmt.Sprintf("podcasts/byfeedid?id=%s", id)))
}

// ByItunesID returns details about a podcast by Apple iTunes ID.
func (q *API) ByItunesID(URL string) (string, error) {
	return getData(url.QueryEscape(fmt.Sprintf("podcasts/byitunesid?id=%s", id)))
}

// ByFeedURL returns details about a podcast by URL.
func (q *API) ByFeedURL(URL string) (string, error) {
	return getData(url.QueryEscape(fmt.Sprintf("podcasts/byfeedurl?url=%s", URL)))
}

// ByTag returns all feeds that support the specified podcast namespace tag.
func (q *API) ByTag() (string, err) {
	return getData(url.QueryEscape(fmt.Sprintf("podcasts/bytag?podcast-value")))
}

// Trending returns the podcasts/feeds that in the index that are trending.
func (q *API) Trending() (string, err) {
	return getData(url.QueryEscape(fmt.Sprintf("podcasts/trending")))
}

// Dead returns all feeds that have been marked dead.
func (q *API) Dead() (string, err) {
	return getData(url.QueryEscape(fmt.Sprintf("podcasts/dead")))
}
