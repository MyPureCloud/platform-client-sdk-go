package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangecampaign
type Dialercampaignconfigchangecampaign struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// ContactList
	ContactList *Dialercampaignconfigchangeurireference `json:"contactList,omitempty"`


	// Queue
	Queue *Dialercampaignconfigchangeurireference `json:"queue,omitempty"`


	// DialingMode
	DialingMode *string `json:"dialingMode,omitempty"`


	// Script
	Script *Dialercampaignconfigchangeurireference `json:"script,omitempty"`


	// EdgeGroup
	EdgeGroup *Dialercampaignconfigchangeurireference `json:"edgeGroup,omitempty"`


	// Site
	Site *Dialercampaignconfigchangeurireference `json:"site,omitempty"`


	// CampaignStatus
	CampaignStatus *string `json:"campaignStatus,omitempty"`


	// PhoneColumns
	PhoneColumns *[]Dialercampaignconfigchangephonecolumn `json:"phoneColumns,omitempty"`


	// AbandonRate
	AbandonRate *float32 `json:"abandonRate,omitempty"`


	// DncLists
	DncLists *[]Dialercampaignconfigchangeurireference `json:"dncLists,omitempty"`


	// CallableTimeSet
	CallableTimeSet *Dialercampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`


	// CallAnalysisResponseSet
	CallAnalysisResponseSet *Dialercampaignconfigchangeurireference `json:"callAnalysisResponseSet,omitempty"`


	// CallerName
	CallerName *string `json:"callerName,omitempty"`


	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`


	// OutboundLineCount
	OutboundLineCount *int `json:"outboundLineCount,omitempty"`


	// Errors
	Errors *[]Dialercampaignconfigchangeresterrordetail `json:"errors,omitempty"`


	// RuleSets
	RuleSets *[]Dialercampaignconfigchangeurireference `json:"ruleSets,omitempty"`


	// SkipPreviewDisabled
	SkipPreviewDisabled *bool `json:"skipPreviewDisabled,omitempty"`


	// PreviewTimeOutSeconds
	PreviewTimeOutSeconds *int `json:"previewTimeOutSeconds,omitempty"`


	// SingleNumberPreview
	SingleNumberPreview *bool `json:"singleNumberPreview,omitempty"`


	// ContactSort
	ContactSort *Dialercampaignconfigchangecontactsort `json:"contactSort,omitempty"`


	// ContactSorts
	ContactSorts *[]Dialercampaignconfigchangecontactsort `json:"contactSorts,omitempty"`


	// NoAnswerTimeout
	NoAnswerTimeout *int `json:"noAnswerTimeout,omitempty"`


	// CallAnalysisLanguage
	CallAnalysisLanguage *string `json:"callAnalysisLanguage,omitempty"`


	// Priority
	Priority *int `json:"priority,omitempty"`


	// ContactListFilters
	ContactListFilters *[]Dialercampaignconfigchangeurireference `json:"contactListFilters,omitempty"`


	// Division
	Division *Dialercampaignconfigchangeurireference `json:"division,omitempty"`


	// AgentOwnedColumn
	AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercampaignconfigchangecampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangecampaign

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ContactList *Dialercampaignconfigchangeurireference `json:"contactList,omitempty"`
		
		Queue *Dialercampaignconfigchangeurireference `json:"queue,omitempty"`
		
		DialingMode *string `json:"dialingMode,omitempty"`
		
		Script *Dialercampaignconfigchangeurireference `json:"script,omitempty"`
		
		EdgeGroup *Dialercampaignconfigchangeurireference `json:"edgeGroup,omitempty"`
		
		Site *Dialercampaignconfigchangeurireference `json:"site,omitempty"`
		
		CampaignStatus *string `json:"campaignStatus,omitempty"`
		
		PhoneColumns *[]Dialercampaignconfigchangephonecolumn `json:"phoneColumns,omitempty"`
		
		AbandonRate *float32 `json:"abandonRate,omitempty"`
		
		DncLists *[]Dialercampaignconfigchangeurireference `json:"dncLists,omitempty"`
		
		CallableTimeSet *Dialercampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`
		
		CallAnalysisResponseSet *Dialercampaignconfigchangeurireference `json:"callAnalysisResponseSet,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		OutboundLineCount *int `json:"outboundLineCount,omitempty"`
		
		Errors *[]Dialercampaignconfigchangeresterrordetail `json:"errors,omitempty"`
		
		RuleSets *[]Dialercampaignconfigchangeurireference `json:"ruleSets,omitempty"`
		
		SkipPreviewDisabled *bool `json:"skipPreviewDisabled,omitempty"`
		
		PreviewTimeOutSeconds *int `json:"previewTimeOutSeconds,omitempty"`
		
		SingleNumberPreview *bool `json:"singleNumberPreview,omitempty"`
		
		ContactSort *Dialercampaignconfigchangecontactsort `json:"contactSort,omitempty"`
		
		ContactSorts *[]Dialercampaignconfigchangecontactsort `json:"contactSorts,omitempty"`
		
		NoAnswerTimeout *int `json:"noAnswerTimeout,omitempty"`
		
		CallAnalysisLanguage *string `json:"callAnalysisLanguage,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		ContactListFilters *[]Dialercampaignconfigchangeurireference `json:"contactListFilters,omitempty"`
		
		Division *Dialercampaignconfigchangeurireference `json:"division,omitempty"`
		
		AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		ContactList: u.ContactList,
		
		Queue: u.Queue,
		
		DialingMode: u.DialingMode,
		
		Script: u.Script,
		
		EdgeGroup: u.EdgeGroup,
		
		Site: u.Site,
		
		CampaignStatus: u.CampaignStatus,
		
		PhoneColumns: u.PhoneColumns,
		
		AbandonRate: u.AbandonRate,
		
		DncLists: u.DncLists,
		
		CallableTimeSet: u.CallableTimeSet,
		
		CallAnalysisResponseSet: u.CallAnalysisResponseSet,
		
		CallerName: u.CallerName,
		
		CallerAddress: u.CallerAddress,
		
		OutboundLineCount: u.OutboundLineCount,
		
		Errors: u.Errors,
		
		RuleSets: u.RuleSets,
		
		SkipPreviewDisabled: u.SkipPreviewDisabled,
		
		PreviewTimeOutSeconds: u.PreviewTimeOutSeconds,
		
		SingleNumberPreview: u.SingleNumberPreview,
		
		ContactSort: u.ContactSort,
		
		ContactSorts: u.ContactSorts,
		
		NoAnswerTimeout: u.NoAnswerTimeout,
		
		CallAnalysisLanguage: u.CallAnalysisLanguage,
		
		Priority: u.Priority,
		
		ContactListFilters: u.ContactListFilters,
		
		Division: u.Division,
		
		AgentOwnedColumn: u.AgentOwnedColumn,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangecampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
