package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Results struct {
	Data []struct {
		Type             string `json:"type"`
		ID               string `json:"id"`
		URL              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		Rating           string `json:"rating"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			Original struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
			} `json:"original"`
			Downsized struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_large"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_medium"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"downsized_small"`
			DownsizedStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_still"`
			FixedHeight struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height_downsampled"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_height_small"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_height_small_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_height_still"`
			FixedWidth struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width_downsampled"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
			} `json:"fixed_width_small"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_width_small_still"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"fixed_width_still"`
			Looping struct {
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"looping"`
			OriginalStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"original_still"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"original_mp4"`
			Preview struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"preview"`
			PreviewGif struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"preview_gif"`
			PreviewWebp struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"preview_webp"`
			Four80WStill struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"480w_still"`
		} `json:"images"`
		User struct {
			AvatarURL    string `json:"avatar_url"`
			BannerImage  string `json:"banner_image"`
			BannerURL    string `json:"banner_url"`
			ProfileURL   string `json:"profile_url"`
			Username     string `json:"username"`
			DisplayName  string `json:"display_name"`
			Description  string `json:"description"`
			InstagramURL string `json:"instagram_url"`
			WebsiteURL   string `json:"website_url"`
			IsVerified   bool   `json:"is_verified"`
		} `json:"user"`
		AnalyticsResponsePayload string `json:"analytics_response_payload"`
		Analytics                struct {
			Onload struct {
				URL string `json:"url"`
			} `json:"onload"`
			Onclick struct {
				URL string `json:"url"`
			} `json:"onclick"`
			Onsent struct {
				URL string `json:"url"`
			} `json:"onsent"`
		} `json:"analytics"`
	} `json:"data"`
	Pagination struct {
		TotalCount int `json:"total_count"`
		Count      int `json:"count"`
		Offset     int `json:"offset"`
	} `json:"pagination"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}

type Client struct {
	http        *http.Client
	giphyApiKey string
}

func (c *Client) Search(query string) (*Results, error) {

	const GIPHY_SEARCH_BASE_URL = "https://api.giphy.com/v1/gifs/search"

	giphySearchEndpoint := fmt.Sprintf("%s?api_key=%s&q=%s&limit=4&offset=0&rating=g&lang=en", GIPHY_SEARCH_BASE_URL, c.giphyApiKey, url.QueryEscape(query))

	response, err := c.http.Get(giphySearchEndpoint)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	results := &Results{}

	return results, json.Unmarshal(body, results)
}

func NewClient(httpClient *http.Client, key string) *Client {
	return &Client{httpClient, key}
}
