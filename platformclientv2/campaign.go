package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaign
type Campaign struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Campaign.
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// ContactList - The ContactList for this Campaign to dial.
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// Queue - The Queue for this Campaign to route calls to. Required for all dialing modes except agentless.
	Queue *Domainentityref `json:"queue,omitempty"`


	// DialingMode - The strategy this Campaign will use for dialing.
	DialingMode *string `json:"dialingMode,omitempty"`


	// Script - The Script to be displayed to agents that are handling outbound calls. Required for all dialing modes except agentless.
	Script *Domainentityref `json:"script,omitempty"`


	// EdgeGroup - The EdgeGroup that will place the calls. Required for all dialing modes except preview.
	EdgeGroup *Domainentityref `json:"edgeGroup,omitempty"`


	// Site - The identifier of the site to be used for dialing; can be set in place of an edge group.
	Site *Domainentityref `json:"site,omitempty"`


	// CampaignStatus - The current status of the Campaign. A Campaign may be turned 'on' or 'off'. Required for updates.
	CampaignStatus *string `json:"campaignStatus,omitempty"`


	// PhoneColumns - The ContactPhoneNumberColumns on the ContactList that this Campaign should dial.
	PhoneColumns *[]Phonecolumn `json:"phoneColumns,omitempty"`


	// AbandonRate - The targeted abandon rate percentage. Required for progressive, power, and predictive campaigns.
	AbandonRate *float64 `json:"abandonRate,omitempty"`


	// DncLists - DncLists for this Campaign to check before placing a call.
	DncLists *[]Domainentityref `json:"dncLists,omitempty"`


	// CallableTimeSet - The callable time set for this campaign to check before placing a call.
	CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`


	// CallAnalysisResponseSet - The call analysis response set to handle call analysis results from the edge. Required for all dialing modes except preview.
	CallAnalysisResponseSet *Domainentityref `json:"callAnalysisResponseSet,omitempty"`


	// Errors - A list of current error conditions associated with the campaign.
	Errors *[]Resterrordetail `json:"errors,omitempty"`


	// CallerName - The caller id name to be displayed on the outbound call.
	CallerName *string `json:"callerName,omitempty"`


	// CallerAddress - The caller id phone number to be displayed on the outbound call.
	CallerAddress *string `json:"callerAddress,omitempty"`


	// OutboundLineCount - The number of outbound lines to be concurrently dialed. Only applicable to non-preview campaigns; only required for agentless.
	OutboundLineCount *int `json:"outboundLineCount,omitempty"`


	// RuleSets - Rule sets to be applied while this campaign is dialing.
	RuleSets *[]Domainentityref `json:"ruleSets,omitempty"`


	// SkipPreviewDisabled - Whether or not agents can skip previews without placing a call. Only applicable for preview campaigns.
	SkipPreviewDisabled *bool `json:"skipPreviewDisabled,omitempty"`


	// PreviewTimeOutSeconds - The number of seconds before a call will be automatically placed on a preview. A value of 0 indicates no automatic placement of calls. Only applicable to preview campaigns.
	PreviewTimeOutSeconds *int `json:"previewTimeOutSeconds,omitempty"`


	// AlwaysRunning - Indicates (when true) that the campaign will remain on after contacts are depleted, allowing additional contacts to be appended/added to the contact list and processed by the still-running campaign. The campaign can still be turned off manually.
	AlwaysRunning *bool `json:"alwaysRunning,omitempty"`


	// ContactSort - The order in which to sort contacts for dialing, based on a column.
	ContactSort *Contactsort `json:"contactSort,omitempty"`


	// ContactSorts - The order in which to sort contacts for dialing, based on up to four columns.
	ContactSorts *[]Contactsort `json:"contactSorts,omitempty"`


	// NoAnswerTimeout - How long to wait before dispositioning a call as 'no-answer'. Default 30 seconds. Only applicable to non-preview campaigns.
	NoAnswerTimeout *int `json:"noAnswerTimeout,omitempty"`


	// CallAnalysisLanguage - The language the edge will use to analyze the call.
	CallAnalysisLanguage *string `json:"callAnalysisLanguage,omitempty"`


	// Priority - The priority of this campaign relative to other campaigns that are running on the same queue. 5 is the highest priority, 1 the lowest.
	Priority *int `json:"priority,omitempty"`


	// ContactListFilters - Filter to apply to the contact list before dialing. Currently a campaign can only have one filter applied.
	ContactListFilters *[]Domainentityref `json:"contactListFilters,omitempty"`


	// Division - The division this campaign belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Campaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaign

	
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
		
		ContactList *Domainentityref `json:"contactList,omitempty"`
		
		Queue *Domainentityref `json:"queue,omitempty"`
		
		DialingMode *string `json:"dialingMode,omitempty"`
		
		Script *Domainentityref `json:"script,omitempty"`
		
		EdgeGroup *Domainentityref `json:"edgeGroup,omitempty"`
		
		Site *Domainentityref `json:"site,omitempty"`
		
		CampaignStatus *string `json:"campaignStatus,omitempty"`
		
		PhoneColumns *[]Phonecolumn `json:"phoneColumns,omitempty"`
		
		AbandonRate *float64 `json:"abandonRate,omitempty"`
		
		DncLists *[]Domainentityref `json:"dncLists,omitempty"`
		
		CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`
		
		CallAnalysisResponseSet *Domainentityref `json:"callAnalysisResponseSet,omitempty"`
		
		Errors *[]Resterrordetail `json:"errors,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		OutboundLineCount *int `json:"outboundLineCount,omitempty"`
		
		RuleSets *[]Domainentityref `json:"ruleSets,omitempty"`
		
		SkipPreviewDisabled *bool `json:"skipPreviewDisabled,omitempty"`
		
		PreviewTimeOutSeconds *int `json:"previewTimeOutSeconds,omitempty"`
		
		AlwaysRunning *bool `json:"alwaysRunning,omitempty"`
		
		ContactSort *Contactsort `json:"contactSort,omitempty"`
		
		ContactSorts *[]Contactsort `json:"contactSorts,omitempty"`
		
		NoAnswerTimeout *int `json:"noAnswerTimeout,omitempty"`
		
		CallAnalysisLanguage *string `json:"callAnalysisLanguage,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		ContactListFilters *[]Domainentityref `json:"contactListFilters,omitempty"`
		
		Division *Domainentityref `json:"division,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
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
		
		Errors: u.Errors,
		
		CallerName: u.CallerName,
		
		CallerAddress: u.CallerAddress,
		
		OutboundLineCount: u.OutboundLineCount,
		
		RuleSets: u.RuleSets,
		
		SkipPreviewDisabled: u.SkipPreviewDisabled,
		
		PreviewTimeOutSeconds: u.PreviewTimeOutSeconds,
		
		AlwaysRunning: u.AlwaysRunning,
		
		ContactSort: u.ContactSort,
		
		ContactSorts: u.ContactSorts,
		
		NoAnswerTimeout: u.NoAnswerTimeout,
		
		CallAnalysisLanguage: u.CallAnalysisLanguage,
		
		Priority: u.Priority,
		
		ContactListFilters: u.ContactListFilters,
		
		Division: u.Division,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
