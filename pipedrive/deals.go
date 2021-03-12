package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// DealService handles deals related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Deals
type DealService service

// Deal represents a Pipedrive deal.
type Deal struct {
	ID            int `json:"id"`
	CreatorUserID struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		HasPic     bool   `json:"has_pic"`
		PicHash    string `json:"pic_hash"`
		ActiveFlag bool   `json:"active_flag"`
		Value      int    `json:"value"`
	} `json:"creator_user_id"`
	UserID struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		HasPic     bool   `json:"has_pic"`
		PicHash    string `json:"pic_hash"`
		ActiveFlag bool   `json:"active_flag"`
		Value      int    `json:"value"`
	} `json:"user_id"`
	PersonID struct {
		Name  string `json:"name"`
		Email []struct {
			Value   string `json:"value"`
			Primary bool   `json:"primary"`
		} `json:"email"`
		Phone []struct {
			Value   string `json:"value"`
			Primary bool   `json:"primary"`
		} `json:"phone"`
		Value      int  `json:"value"`
		ActiveFlag bool `json:"active_flag"`
	} `json:"person_id"`
	OrgID struct {
		Name        string      `json:"name"`
		PeopleCount int         `json:"people_count"`
		OwnerID     int         `json:"owner_id"`
		Address     interface{} `json:"address"`
		ActiveFlag  bool        `json:"active_flag"`
		CcEmail     string      `json:"cc_email"`
		Value       int         `json:"value"`
	} `json:"org_id"`
	StageID                  int         `json:"stage_id"`
	Title                    string      `json:"title"`
	Value                    float64     `json:"value"`
	Currency                 string      `json:"currency"`
	AddTime                  string      `json:"add_time"`
	UpdateTime               string      `json:"update_time"`
	StageChangeTime          string      `json:"stage_change_time"`
	Active                   bool        `json:"active"`
	Deleted                  bool        `json:"deleted"`
	Status                   string      `json:"status"`
	Probability              interface{} `json:"probability"`
	NextActivityDate         interface{} `json:"next_activity_date"`
	NextActivityTime         interface{} `json:"next_activity_time"`
	NextActivityID           interface{} `json:"next_activity_id"`
	LastActivityID           int         `json:"last_activity_id"`
	LastActivityDate         string      `json:"last_activity_date"`
	LostReason               string      `json:"lost_reason"`
	VisibleTo                string      `json:"visible_to"`
	CloseTime                string      `json:"close_time"`
	PipelineID               int         `json:"pipeline_id"`
	WonTime                  interface{} `json:"won_time"`
	FirstWonTime             interface{} `json:"first_won_time"`
	LostTime                 string      `json:"lost_time"`
	ProductsCount            int         `json:"products_count"`
	FilesCount               int         `json:"files_count"`
	NotesCount               int         `json:"notes_count"`
	FollowersCount           int         `json:"followers_count"`
	EmailMessagesCount       int         `json:"email_messages_count"`
	ActivitiesCount          int         `json:"activities_count"`
	DoneActivitiesCount      int         `json:"done_activities_count"`
	UndoneActivitiesCount    int         `json:"undone_activities_count"`
	ReferenceActivitiesCount int         `json:"reference_activities_count"`
	ParticipantsCount        int         `json:"participants_count"`
	ExpectedCloseDate        interface{} `json:"expected_close_date"`
	LastIncomingMailTime     interface{} `json:"last_incoming_mail_time"`
	LastOutgoingMailTime     interface{} `json:"last_outgoing_mail_time"`
	StageOrderNr             int         `json:"stage_order_nr"`
	PersonName               string      `json:"person_name"`
	OrgName                  string      `json:"org_name"`
	NextActivitySubject      interface{} `json:"next_activity_subject"`
	NextActivityType         interface{} `json:"next_activity_type"`
	NextActivityDuration     interface{} `json:"next_activity_duration"`
	NextActivityNote         interface{} `json:"next_activity_note"`
	FormattedValue           string      `json:"formatted_value"`
	RottenTime               interface{} `json:"rotten_time"`
	WeightedValue            int         `json:"weighted_value"`
	FormattedWeightedValue   string      `json:"formatted_weighted_value"`
	OwnerName                string      `json:"owner_name"`
	CcEmail                  string      `json:"cc_email"`
	OrgHidden                bool        `json:"org_hidden"`
	PersonHidden             bool        `json:"person_hidden"`
	OfflineCommunication     string      `json:"b556c5618b88cd3d33f99b996b5b2fdbc8ba3c7e"`
	ServicePrice             float64     `json:"6906ddfb72aaef6810b35703de142db0f435c314"`
	AgencyInCharge           struct {
		Name        string      `json:"name"`
		PeopleCount int         `json:"people_count"`
		OwnerID     int         `json:"owner_id"`
		Address     interface{} `json:"address"`
		CcEmail     string      `json:"cc_email"`
		Value       int         `json:"value"`
	} `json:"eb2a2df8945c29118a01d324c58fbf6cef7bfd43"`
	DepartureCurrentCT string  `json:"ffd2a712d781417ef17a33b3540d3ecd8d945f76"`
	ArrivalNextCT      string  `json:"f5bb67b24aec5f9733be8b3402bf1fe5b1ac6ed6"`
	LeadSource         uint    `json:"5d4fbabc9b032aeb3df515d9c66994d6892ee062"`
	TemporaryLink      string  `json:"4fe88fad67d8dcbc17d18d9ee1faac55122249fd"`
	RideCosts          string  `json:"31443a48d1405182dfccac9bf378bbe8216ffc9a"`
	InvoiceCareDays    *int    `json:"c1feb6a1b169b6caf54776fa85cc8452b27d0fd6"`
	ChristmasCare      uint    `json:"9bfc1ae5cf091620979b3ac015f2f3f1a74d165c"`
	FirstCareBegin     string  `json:"b927d845c72cbb2f90844a65206776edac927684"`
	ServiceFee         float64 `json:"0825a408763de9f63cc80be74fe55ed8b625efa4"`
	Feedback           string  `json:"ab85caab3c817a41d2583c74fb9184b72647d90a"`
}

func (d Deal) String() string {
	return Stringify(d)
}

// DealsResponse represents multiple deals response.
type DealsResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Deal         `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// DealResponse represents single deal response.
type DealResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           Deal           `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// Participant represents a Pipedrive participant.
type Participant struct {
	ID       int `json:"id"`
	PersonID struct {
		Name  string `json:"name"`
		Email []struct {
			Value   string `json:"value"`
			Primary bool   `json:"primary"`
		} `json:"email"`
		Phone []struct {
			Value   string `json:"value"`
			Primary bool   `json:"primary"`
		} `json:"phone"`
		Value int `json:"value"`
	} `json:"person_id"`
	AddTime         string `json:"add_time"`
	Active          bool   `json:"active_flag"`
	RelatedItemData struct {
		DealId int    `json:"deal_id"`
		Title  string `json:"title"`
	} `json:"related_item_data"`
	RelatedItemType string `json:"related_item_type"`
	RelatedItemId   int    `json:"related_item_id"`
	Person          Person `json:"person"`
}

// DealParticipantsResponse represents multiple deals response.
type DealParticipantsResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Participant  `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// ListUpdates about a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_id_flow
func (s *DealService) ListUpdates(ctx context.Context, id int) (*DealsResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/flow", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Find deals by name.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_find
func (s *DealService) Find(ctx context.Context, term string) (*DealsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals/find", &SearchOptions{
		Term: term,
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type DealsFilterOptions struct {
	FilterID int    `url:"filter_id,omitempty"`
	Status   string `url:"status,omitempty"`
	Start    int    `url:"start,omitempty"`
	Limit    int    `url:"limit,omitempty"`
}

// List deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals
func (s *DealService) List(ctx context.Context, dfo *DealsFilterOptions) (*DealsResponse, *Response, error) {
	var err error
	var req *http.Request

	req, err = s.client.NewRequest(http.MethodGet, "/deals", dfo, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type DealParticipantsFilterOptions struct {
	Start int `url:"start,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// List deals participants.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_id_participants
func (s *DealService) ListParticipants(ctx context.Context, id int, dfo *DealParticipantsFilterOptions) (*DealParticipantsResponse, *Response, error) {
	var err error
	var req *http.Request

	uri := fmt.Sprintf("/deals/%v/participants", id)
	req, err = s.client.NewRequest(http.MethodGet, uri, dfo, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealParticipantsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Duplicate a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals_id_duplicate
func (s *DealService) Duplicate(ctx context.Context, id int) (*DealResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/duplicate", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DealsMergeOptions specifices the optional parameters to the
// DealService.Merge method.
type DealsMergeOptions struct {
	MergeWithID uint `url:"merge_with_id,omitempty"`
}

// Merge two deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id_merge
func (s *DealService) Merge(ctx context.Context, id int, opt *DealsMergeOptions) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/merge", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DealsUpdateOptions specifices the optional parameters to the
// DealService.Update method.
type DealsUpdateOptions struct {
	Title           string `json:"title,omitempty"`
	Value           string `json:"value,omitempty"`
	Currency        string `json:"currency,omitempty"`
	UserID          uint   `json:"user_id,omitempty"`
	PersonID        uint   `json:"person_id,omitempty"`
	OrganizationID  uint   `json:"org_id,omitempty"`
	StageID         uint   `json:"stage_id,omitempty"`
	Status          string `json:"status,omitempty"`
	LostReason      string `json:"lost_reason,omitempty"`
	VisibleTo       uint   `json:"visible_to,omitempty"`
	TemporaryLink   string `json:"4fe88fad67d8dcbc17d18d9ee1faac55122249fd,omitempty"`
	InvoiceCareDays *int   `json:"c1feb6a1b169b6caf54776fa85cc8452b27d0fd6,omitempty"`
	ChristmasCare   uint   `json:"9bfc1ae5cf091620979b3ac015f2f3f1a74d165c,omitempty"`
	AgencyInCharge  uint   `json:"eb2a2df8945c29118a01d324c58fbf6cef7bfd43,omitempty"`
	ServiceFee      string `json:"0825a408763de9f63cc80be74fe55ed8b625efa4,omitempty"`
	Feedback        string `json:"ab85caab3c817a41d2583c74fb9184b72647d90a,omitempty"`
}

// Update a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (s *DealService) Update(ctx context.Context, id int, opt *DealsUpdateOptions) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

type DealsUpdateForceOptions struct {
	InvoiceCareDays *int `json:"c1feb6a1b169b6caf54776fa85cc8452b27d0fd6"`
}

// Update a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (s *DealService) UpdateForce(ctx context.Context, id int, opt *DealsUpdateForceOptions) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteFollower of a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_followers_follower_id
func (s *DealService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/followers/%v", id, followerID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteMultiple deletes deals in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals
func (s *DealService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/deals", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteParticipant deletes participant in a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_participants_deal_participant_id
func (s *DealService) DeleteParticipant(ctx context.Context, dealID int, participantID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/participants/%v", dealID, participantID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id
func (s *DealService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteAttachedProduct deletes attached product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_products_product_attachment_id
func (s *DealService) DeleteAttachedProduct(ctx context.Context, dealID int, productAttachmentID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/products/%v", dealID, productAttachmentID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DealCreateOptions specifices the optional parameters to the
// DealsService.Create method.
type DealCreateOptions struct {
	Title         string    `json:"title,omitempty"`
	Value         string    `json:"value,omitempty"`
	Currency      string    `json:"currency,omitempty"`
	UserID        uint      `json:"user_id,omitempty"`
	PersonID      uint      `json:"person_id,omitempty"`
	OrgID         uint      `json:"org_id,omitempty"`
	StageID       uint      `json:"stage_id,omitempty"`
	Status        string    `json:"status,omitempty"`
	Probability   uint      `json:"probability,omitempty"`
	LostReason    string    `json:"lost_reason,omitempty"`
	AddTime       Timestamp `json:"add_time,omitempty"`
	VisibleTo     VisibleTo `json:"visible_to,omitempty"`
	TemporaryLink string    `json:"4fe88fad67d8dcbc17d18d9ee1faac55122249fd,omitempty"`
	LeadSource    uint      `json:"5d4fbabc9b032aeb3df515d9c66994d6892ee062,omitempty"`
}

// Create a new deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals
func (s *DealService) Create(ctx context.Context, opt *DealCreateOptions) (*DealResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/deals", nil, struct {
		Title         string    `json:"title,omitempty"`
		Value         string    `json:"value,omitempty"`
		Currency      string    `json:"currency,omitempty"`
		UserID        uint      `json:"user_id,omitempty"`
		PersonID      uint      `json:"person_id,omitempty"`
		OrgID         uint      `json:"org_id,omitempty"`
		StageID       uint      `json:"stage_id,omitempty"`
		Status        string    `json:"status,omitempty"`
		Probability   uint      `json:"probability,omitempty"`
		LostReason    string    `json:"lost_reason,omitempty"`
		AddTime       string    `json:"add_time,omitempty"`
		VisibleTo     VisibleTo `json:"visible_to,omitempty"`
		TemporaryLink string    `json:"4fe88fad67d8dcbc17d18d9ee1faac55122249fd,omitempty"`
		LeadSource    uint      `json:"5d4fbabc9b032aeb3df515d9c66994d6892ee062,omitempty"`
	}{
		opt.Title,
		opt.Value,
		opt.Currency,
		opt.UserID,
		opt.PersonID,
		opt.OrgID,
		opt.StageID,
		opt.Status,
		opt.Probability,
		opt.LostReason,
		opt.AddTime.FormatFull(),
		opt.VisibleTo,
		opt.TemporaryLink,
		opt.LeadSource,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *DealResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Get a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_id
func (s *DealService) Get(ctx context.Context, dealID int) (*DealResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("/deals/%d", dealID), nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type DealActivitiesOptions struct {
	ID      int    `url:"id"`
	Start   int    `url:"start,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	Done    int    `url:"done,omitempty"`
	exclude string `url:"exclude,omitempty"`
}

// Lists activities associated with a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_id_activities
func (s *DealService) GetActivities(ctx context.Context, dap *DealActivitiesOptions) (*ActivitiesReponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("/deals/%d/activities", dap.ID), dap, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivitiesReponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
