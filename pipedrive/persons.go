package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// PersonsService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons
type PersonsService service

type Email struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// Person represents a Pipedrive person.
type Person struct {
	ID        int `json:"id"`
	CompanyID int `json:"company_id"`
	OwnerID   struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		HasPic     bool   `json:"has_pic"`
		PicHash    string `json:"pic_hash"`
		ActiveFlag bool   `json:"active_flag"`
		Value      int    `json:"value"`
	} `json:"owner_id"`
	OrgID struct {
		Name        string      `json:"name"`
		PeopleCount int         `json:"people_count"`
		OwnerID     int         `json:"owner_id"`
		Address     interface{} `json:"address"`
		ActiveFlag  bool        `json:"active_flag"`
		CcEmail     string      `json:"cc_email"`
		Value       int         `json:"value"`
	} `json:"org_id"`
	Name                        string `json:"name"`
	FirstName                   string `json:"first_name"`
	LastName                    string `json:"last_name"`
	OpenDealsCount              int    `json:"open_deals_count"`
	RelatedOpenDealsCount       int    `json:"related_open_deals_count"`
	ClosedDealsCount            int    `json:"closed_deals_count"`
	RelatedClosedDealsCount     int    `json:"related_closed_deals_count"`
	ParticipantOpenDealsCount   int    `json:"participant_open_deals_count"`
	ParticipantClosedDealsCount int    `json:"participant_closed_deals_count"`
	EmailMessagesCount          int    `json:"email_messages_count"`
	ActivitiesCount             int    `json:"activities_count"`
	DoneActivitiesCount         int    `json:"done_activities_count"`
	UndoneActivitiesCount       int    `json:"undone_activities_count"`
	ReferenceActivitiesCount    int    `json:"reference_activities_count"`
	FilesCount                  int    `json:"files_count"`
	NotesCount                  int    `json:"notes_count"`
	FollowersCount              int    `json:"followers_count"`
	WonDealsCount               int    `json:"won_deals_count"`
	RelatedWonDealsCount        int    `json:"related_won_deals_count"`
	LostDealsCount              int    `json:"lost_deals_count"`
	RelatedLostDealsCount       int    `json:"related_lost_deals_count"`
	ActiveFlag                  bool   `json:"active_flag"`
	Phone                       []struct {
		Value   string `json:"value"`
		Label   string `json:"label"`
		Primary bool   `json:"primary"`
	} `json:"phone"`
	Email                           []Email     `json:"email"`
	FirstChar                       string      `json:"first_char"`
	UpdateTime                      string      `json:"update_time"`
	AddTime                         string      `json:"add_time"`
	Birthday                        string      `json:"birthday"`
	VisibleTo                       string      `json:"visible_to"`
	PictureID                       interface{} `json:"picture_id"`
	NextActivityDate                interface{} `json:"next_activity_date"`
	NextActivityTime                interface{} `json:"next_activity_time"`
	NextActivityID                  interface{} `json:"next_activity_id"`
	LastActivityID                  int         `json:"last_activity_id"`
	LastActivityDate                string      `json:"last_activity_date"`
	TimelineLastActivityTime        interface{} `json:"timeline_last_activity_time"`
	TimelineLastActivityTimeByOwner interface{} `json:"timeline_last_activity_time_by_owner"`
	LastIncomingMailTime            interface{} `json:"last_incoming_mail_time"`
	LastOutgoingMailTime            interface{} `json:"last_outgoing_mail_time"`
	OrgName                         interface{} `json:"org_name"`
	OwnerName                       string      `json:"owner_name"`
	CcEmail                         string      `json:"cc_email"`
	Label                           uint        `json:"label"`
	BillingAddress                  string      `json:"d5d6ecba25dd34146d3b9d0f1bb34dedf384143a"`
	DeliveryAddress                 string      `json:"fb3875ae1de17d63a1a0a9a7643bb677b95ae7fb"`
	LeadExportWDS                   string      `json:"71df7f223a0dd3b9314412b2cf37d3ee55a657d8"`
	A1ExpireDate                    string      `json:"df8b1b69dfad26a8a0a4b57bcf0016ce519a7031"`
	A1Upload                        string      `json:"815b394dd12a36bdbfd6215a0aad5bc4d35be6a9"`
	PersonTitle                     string      `json:"100ace766d7a2bd9c51bb1b2c3ff71410beb5821"`
}

func (p Person) String() string {
	return Stringify(p)
}

// PersonsResponse represents multiple persons response.
type PersonsResponse struct {
	Success        bool           `json:"success"`
	Data           []Person       `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// PersonResponse represents single person response.
type PersonResponse struct {
	Success        bool           `json:"success"`
	Data           Person         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// PersonAddFollowerResponse represents add follower response.
type PersonAddFollowerResponse struct {
	Success bool `json:"success"`
	Data    struct {
		UserID   int    `json:"user_id"`
		ID       int    `json:"id"`
		PersonID int    `json:"person_id"`
		AddTime  string `json:"add_time"`
	} `json:"data"`
}

// PersonSearch represents a search item.
type PersonSearch struct {
	ResultScore float32 `json:"result_score"`
	Item        struct {
		ID    int    `json:"id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Phone []struct {
			Value string `json:"value"`
		} `json:"phone"`
		Email []Email `json:"email"`
		Owner struct {
			ID int `json:"id"`
		} `json:"owner"`
		VisibleTo    string `json:"visible_to"`
		Organization struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"organization"`
	} `json:"item"`
}

// PersonSearchResponse represents search response.
type PersonSearchResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Items []PersonSearch `json:"items"`
	} `json:"data`
	AdditionalData AdditionalData `json:"additional_data"`
}

type PersonFilterOptions struct {
	FilterID int `url:"filter_id,omitempty"`
	Start    int `url:"start,omitempty"`
	Limit    int `url:"limit,omitempty"`
}

// List all persons.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons
func (s *PersonsService) List(ctx context.Context, pfo *PersonFilterOptions) (*PersonsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/persons", pfo, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// AddFollower adds a follower to person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons_id_followers
func (s *PersonsService) AddFollower(ctx context.Context, id int, userID int) (*PersonAddFollowerResponse, *Response, error) {
	uri := fmt.Sprintf("/persons/%v/followers", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, struct {
		UserID int `json:"user_id"`
	}{
		userID,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *PersonAddFollowerResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// PersonCreateOptions specifices the optional parameters to the
// PersonsService.Create method.
type PersonCreateOptions struct {
	Name             string    `json:"name"`
	OwnerID          uint      `json:"owner_id"`
	OrgID            uint      `json:"org_id"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	VisibleTo        VisibleTo `json:"visible_to"`
	AddTime          Timestamp `json:"add_time"`
	Label            uint      `json:"label"`
	NewsletterStatus uint      `json:"4ba5027f64252634ad94493411cc251f1371786c,omitempty"`
	LeadExportWDS    string    `json:"71df7f223a0dd3b9314412b2cf37d3ee55a657d8,omitempty"`
	A1ExpireDate     Timestamp `json:"df8b1b69dfad26a8a0a4b57bcf0016ce519a7031,omitempty"`
}

// Create a new person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons
func (s *PersonsService) Create(ctx context.Context, opt *PersonCreateOptions) (*PersonResponse, *Response, error) {

	phone := []struct {
		Value string `json:"value"`
	}{
		{opt.Phone},
	}

	req, err := s.client.NewRequest(http.MethodPost, "/persons", nil, struct {
		Name    string `json:"name"`
		OwnerID uint   `json:"owner_id"`
		OrgID   uint   `json:"org_id"`
		Email   string `json:"email"`
		Phone   []struct {
			Value string `json:"value"`
		} `json:"phone"`
		Label            uint      `json:"label"`
		VisibleTo        VisibleTo `json:"visible_to"`
		AddTime          string    `json:"add_time"`
		NewsletterStatus uint      `json:"4ba5027f64252634ad94493411cc251f1371786c,omitempty"`
		LeadExportWDS    string    `json:"71df7f223a0dd3b9314412b2cf37d3ee55a657d8,omitempty"`
		A1ExpireDate     string    `json:"df8b1b69dfad26a8a0a4b57bcf0016ce519a7031,omitempty"`
	}{
		opt.Name,
		opt.OwnerID,
		opt.OrgID,
		opt.Email,
		phone,
		opt.Label,
		opt.VisibleTo,
		opt.AddTime.FormatFull(),
		opt.NewsletterStatus,
		opt.LeadExportWDS,
		opt.A1ExpireDate.Format(),
	})

	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// PersonUpdateOptions specifices the optional parameters to the
// PersonUpdateOptions.Update method.
type PersonUpdateOptions struct {
	Name             string    `json:"name,omitempty"`
	OwnerID          uint      `json:"owner_id,omitempty"`
	OrgID            uint      `json:"org_id,omitempty"`
	Email            []Email   `json:"email,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	VisibleTo        VisibleTo `json:"visible_to,omitempty"`
	BillingAddress   string    `json:"d5d6ecba25dd34146d3b9d0f1bb34dedf384143a,omitempty"`
	DeliveryAddress  string    `json:"fb3875ae1de17d63a1a0a9a7643bb677b95ae7fb,omitempty"`
	NewsletterStatus uint      `json:"4ba5027f64252634ad94493411cc251f1371786c,omitempty"`
	LeadExportWDS    string    `json:"71df7f223a0dd3b9314412b2cf37d3ee55a657d8,omitempty"`
	A1ExpireDate     Timestamp `json:"df8b1b69dfad26a8a0a4b57bcf0016ce519a7031,omitempty"`
	A1Upload         string    `json:"815b394dd12a36bdbfd6215a0aad5bc4d35be6a9,omitempty"`
	PersonTitle      int       `json:"100ace766d7a2bd9c51bb1b2c3ff71410beb5821,omitempty"`
}

// Update a specific person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id
func (s *PersonsService) Update(ctx context.Context, id int, opt *PersonUpdateOptions) (*PersonResponse, *Response, error) {

	uri := fmt.Sprintf("/persons/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, struct {
		Name             string    `json:"name,omitempty"`
		OwnerID          uint      `json:"owner_id,omitempty"`
		OrgID            uint      `json:"org_id,omitempty"`
		Email            []Email   `json:"email,omitempty"`
		Phone            string    `json:"phone,omitempty"`
		VisibleTo        VisibleTo `json:"visible_to,omitempty"`
		BillingAddress   string    `json:"d5d6ecba25dd34146d3b9d0f1bb34dedf384143a,omitempty"`
		DeliveryAddress  string    `json:"fb3875ae1de17d63a1a0a9a7643bb677b95ae7fb,omitempty"`
		NewsletterStatus uint      `json:"4ba5027f64252634ad94493411cc251f1371786c,omitempty"`
		LeadExportWDS    string    `json:"71df7f223a0dd3b9314412b2cf37d3ee55a657d8,omitempty"`
		A1ExpireDate     string    `json:"df8b1b69dfad26a8a0a4b57bcf0016ce519a7031,omitempty"`
		A1Upload         string    `json:"815b394dd12a36bdbfd6215a0aad5bc4d35be6a9,omitempty"`
		PersonTitle      int       `json:"100ace766d7a2bd9c51bb1b2c3ff71410beb5821,omitempty"`
	}{
		opt.Name,
		opt.OwnerID,
		opt.OrgID,
		opt.Email,
		opt.Phone,
		opt.VisibleTo,
		opt.BillingAddress,
		opt.DeliveryAddress,
		opt.NewsletterStatus,
		opt.LeadExportWDS,
		opt.A1ExpireDate.Format(),
		opt.A1Upload,
		opt.PersonTitle,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Merge selected persons.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id_merge
func (s *PersonsService) Merge(ctx context.Context, id int, mergeWithID int) (*PersonResponse, *Response, error) {
	uri := fmt.Sprintf("/persons/%v/merge", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, struct {
		MergeWithID int `json:"merge_with_id"`
	}{
		mergeWithID,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteFollower removes follower from person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id_followers_follower_id
func (s *PersonsService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
	uri := fmt.Sprintf("/persons/%v/followers/%v", id, followerID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete marks person as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id
func (s *PersonsService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/persons/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeletePicture deletes person picture.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id_picture
func (s *PersonsService) DeletePicture(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/persons/%v/picture", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteMultiple marks multiple persons as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons
func (s *PersonsService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/persons", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Get a person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_id
func (s *PersonsService) Get(ctx context.Context, personID int) (*PersonResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("/persons/%d", personID), nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type PersonDealsOptions struct {
	ID    int `url:"id"`
	Start int `url:"start,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// Lists deals associated with a person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_id_deals
func (s *PersonsService) GetDeals(ctx context.Context, pdo *PersonDealsOptions) (*DealsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("/persons/%d/deals", pdo.ID), pdo, nil)

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

type PersonFilesOptions struct {
	ID    int `url:"id"`
	Start int `url:"start,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// Lists files associated with a person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/getPersonFiles
func (s *PersonsService) GetFiles(ctx context.Context, pfo *PersonFilesOptions) (*FilesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("/persons/%d/files", pfo.ID), pfo, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *FilesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type PersonSearchOptions struct {
	Term           string `url:"term"`
	Fields         string `url:"fields,omitempty"`
	ExactMatch     bool   `url:"exact_match,omitempty"`
	OrganizationId int    `url:"organization_id,omitempty"`
	IncludeFields  string `url:"include_fields,omitempty"`
	Start          int    `url:"start,omitempty"`
	Limit          int    `url:"limit,omitempty"`
}

// Search persons.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_search
func (s *PersonsService) Search(ctx context.Context, pso *PersonSearchOptions) (*PersonSearchResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/persons/search", pso, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonSearchResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
