package pipedrive

import (
	"context"
	"net/http"
)

// ItemSearchService handles search results related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ItemSearch
type ItemSearchService service

// SearchResult represents a Pipedrive search result.
type SearchResult struct {
	ResultScore float64 `json:"result_score"`
	Item        struct {
		ID        int      `json:"id"`
		Type      string   `json:"type"`
		Name      string   `json:"name"`
		Phones    []string `json:"phones"`
		Emails    []string `json:"emails"`
		VisibleTo int      `json:"visible_to"`
		Owner     struct {
			ID int `json:"id"`
		} `json:"owner"`
	} `json:"item"`
}

func (sr SearchResult) String() string {
	return Stringify(sr)
}

type SearchResultItems struct {
	Items []SearchResult `json:"items"`
}

// ItemSearch represents multiple search results response.
type ItemSearch struct {
	Success        bool              `json:"success"`
	Data           SearchResultItems `json:"data"`
	AdditionalData AdditionalData    `json:"additional_data"`
}

// ItemSearchOptions specifices the optional parameters to the
// ItemSearchService.Search method.
type ItemSearchOptions struct {
	Term                  string `url:"term,omitempty"`
	ItemType              string `url:"item_type,omitempty"`
	Fields                string `url:"fields,omitempty"`
	SearchForRelatedItems bool   `url:"search_for_related_items,omitempty"`
	ExactMatch            bool   `url:"exact_match,omitempty"`
	IncludeFiealds        string `url:"include_fields,omitempty"`
	Start                 uint   `url:"start,omitempty"`
	Limit                 uint   `url:"limit,omitempty"`
}

// Search performs a search across the account and returns ItemSearch.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/ItemSearch/get_itemSearch
func (s *ItemSearchService) Search(ctx context.Context, opt *ItemSearchOptions) (*ItemSearch, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/itemSearch", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ItemSearch

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
