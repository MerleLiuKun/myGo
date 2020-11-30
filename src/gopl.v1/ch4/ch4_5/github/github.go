package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time
	Body string  // in Markdown format
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil{
		return nil, err
	}

	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil{
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}