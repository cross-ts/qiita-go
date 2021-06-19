package qiita

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type HomeArticleTrendFeed struct {
	Trend *Trend  `json:"trend"`
	Scope *string `json:"scope"`
	Type  *string `json:"type"`
}

type Trend struct {
	Edges *[]Edge `json:"edges"`
}

type Edge struct {
	IsLikeByViewer   *bool     `json:"isLikedByViewer"`
	IsNewArrival     *bool     `json:"isNewArrival"`
	FollowingLinkers *[]string `json:"followingLikers"`
	Node             *Node     `json:"node"`
}

type Node struct {
	EncryptedId             *string   `json:"encryptedId"`
	IsLikeByViewer          *bool     `json:"isLikedByViewer"`
	IsStockableByViewer     *bool     `json:"isStockableByViewer"`
	IsStockedByViewer       *bool     `json:"isStockedByViewer"`
	LikesCount              *int      `json:"likesCount"`
	LinkUrl                 *string   `json:"linkUrl"`
	PublishedAt             *string   `json:"publishedAt"`
	Title                   *string   `json:"title"`
	Uuid                    *string   `json:"uuid"`
	Author                  *Author   `json:"author"`
	Organization            *string   `json:"organization"`            // TBD
	RecentlyFollowingLikers *[]string `json:"recentlyFollowingLikers"` // TBD
	Tags                    *[]Tag    `json:"tags"`
}

type Author struct {
	ProfileImageUrl *string `json:"profileImageUrl"`
	UrlName         *string `json:"urlName"`
}

type Tag struct {
	UrlName *string `json:"urlName"`
	Name    *string `json:"name"`
}

type TrendScraper struct {
	Url   string
	Query string
}

func (s *TrendScraper) scrape() (HomeArticleTrendFeed, error) {
	document, err := goquery.NewDocument(s.Url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return HomeArticleTrendFeed{}, err
	}
	raw_json := document.Find(s.Query).Text()
	var homeArticleTrendFeed HomeArticleTrendFeed
	json.Unmarshal([]byte(raw_json), &homeArticleTrendFeed)
	return homeArticleTrendFeed, nil
}

func newTrendScraper() *TrendScraper {
	return &TrendScraper{
		Url:   "https://qiita.com/",
		Query: "script[data-component-name=HomeArticleTrendFeed]",
	}
}
