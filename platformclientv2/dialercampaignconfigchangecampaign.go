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
	// ContactList
	ContactList *Dialercampaignconfigchangeurireference `json:"contactList,omitempty"`


	// Queue - A UriReference for a resource
	Queue *Dialercampaignconfigchangeurireference `json:"queue,omitempty"`


	// DialingMode - dialing mode of the campaign
	DialingMode *string `json:"dialingMode,omitempty"`


	// Script - A UriReference for a resource
	Script *Dialercampaignconfigchangeurireference `json:"script,omitempty"`


	// EdgeGroup - A UriReference for a resource
	EdgeGroup *Dialercampaignconfigchangeurireference `json:"edgeGroup,omitempty"`


	// Site - A UriReference for a resource
	Site *Dialercampaignconfigchangeurireference `json:"site,omitempty"`


	// CampaignStatus
	CampaignStatus *string `json:"campaignStatus,omitempty"`


	// PhoneColumns - the contact list phone columns to be called for the campaign
	PhoneColumns *[]Dialercampaignconfigchangephonecolumn `json:"phoneColumns,omitempty"`


	// AbandonRate - the targeted abandon rate percentage
	AbandonRate *float32 `json:"abandonRate,omitempty"`


	// DncLists - identifiers of the do not call lists
	DncLists *[]Dialercampaignconfigchangeurireference `json:"dncLists,omitempty"`


	// CallableTimeSet - A UriReference for a resource
	CallableTimeSet *Dialercampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`


	// CallAnalysisResponseSet - A UriReference for a resource
	CallAnalysisResponseSet *Dialercampaignconfigchangeurireference `json:"callAnalysisResponseSet,omitempty"`


	// CallerName - caller id name to be displayed on the outbound call
	CallerName *string `json:"callerName,omitempty"`


	// CallerAddress - caller id phone number to be displayed on the outbound call
	CallerAddress *string `json:"callerAddress,omitempty"`


	// OutboundLineCount - for agentless campaigns, the number of outbound lines to be concurrently dialed
	OutboundLineCount *int `json:"outboundLineCount,omitempty"`


	// Errors - a list of current error conditions associated with the campaign
	Errors *[]Dialercampaignconfigchangeresterrordetail `json:"errors,omitempty"`


	// RuleSets - identifiers of the rule sets
	RuleSets *[]Dialercampaignconfigchangeurireference `json:"ruleSets,omitempty"`


	// SkipPreviewDisabled - for preview campaigns, indicator of whether the agent can skip a preview without placing a call
	SkipPreviewDisabled *bool `json:"skipPreviewDisabled,omitempty"`


	// PreviewTimeOutSeconds - for preview campaigns, number of seconds before a call will be automatically placed. A value of 0 indicates no automatic placement of calls
	PreviewTimeOutSeconds *int `json:"previewTimeOutSeconds,omitempty"`


	// SingleNumberPreview - for preview campaigns with multiple phone columns, indicator if one (true) or multiple (false) phone numbers will be available to call for each preview
	SingleNumberPreview *bool `json:"singleNumberPreview,omitempty"`


	// ContactSort
	ContactSort *Dialercampaignconfigchangecontactsort `json:"contactSort,omitempty"`


	// ContactSorts - List of contact sort objects.
	ContactSorts *[]Dialercampaignconfigchangecontactsort `json:"contactSorts,omitempty"`


	// NoAnswerTimeout - for non-preview campaigns, how long to wait before dispositioning as 'no-answer', default 30 seconds
	NoAnswerTimeout *int `json:"noAnswerTimeout,omitempty"`


	// CallAnalysisLanguage - The language the edge will use to analyze the call
	CallAnalysisLanguage *string `json:"callAnalysisLanguage,omitempty"`


	// Priority - The priority of this campaign relative to other campaigns
	Priority *int `json:"priority,omitempty"`


	// ContactListFilters - List of contact filters
	ContactListFilters *[]Dialercampaignconfigchangeurireference `json:"contactListFilters,omitempty"`


	// Division - A UriReference for a resource
	Division *Dialercampaignconfigchangeurireference `json:"division,omitempty"`


	// AgentOwnedColumn - For Preview Campaigns. Name of the contact column in the contact list containing the userIds of agents to assign specific contact records to.
	AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`


	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

}

func (o *Dialercampaignconfigchangecampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangecampaign
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
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
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		ContactList: o.ContactList,
		
		Queue: o.Queue,
		
		DialingMode: o.DialingMode,
		
		Script: o.Script,
		
		EdgeGroup: o.EdgeGroup,
		
		Site: o.Site,
		
		CampaignStatus: o.CampaignStatus,
		
		PhoneColumns: o.PhoneColumns,
		
		AbandonRate: o.AbandonRate,
		
		DncLists: o.DncLists,
		
		CallableTimeSet: o.CallableTimeSet,
		
		CallAnalysisResponseSet: o.CallAnalysisResponseSet,
		
		CallerName: o.CallerName,
		
		CallerAddress: o.CallerAddress,
		
		OutboundLineCount: o.OutboundLineCount,
		
		Errors: o.Errors,
		
		RuleSets: o.RuleSets,
		
		SkipPreviewDisabled: o.SkipPreviewDisabled,
		
		PreviewTimeOutSeconds: o.PreviewTimeOutSeconds,
		
		SingleNumberPreview: o.SingleNumberPreview,
		
		ContactSort: o.ContactSort,
		
		ContactSorts: o.ContactSorts,
		
		NoAnswerTimeout: o.NoAnswerTimeout,
		
		CallAnalysisLanguage: o.CallAnalysisLanguage,
		
		Priority: o.Priority,
		
		ContactListFilters: o.ContactListFilters,
		
		Division: o.Division,
		
		AgentOwnedColumn: o.AgentOwnedColumn,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignconfigchangecampaign) UnmarshalJSON(b []byte) error {
	var DialercampaignconfigchangecampaignMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignconfigchangecampaignMap)
	if err != nil {
		return err
	}
	
	if ContactList, ok := DialercampaignconfigchangecampaignMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if Queue, ok := DialercampaignconfigchangecampaignMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if DialingMode, ok := DialercampaignconfigchangecampaignMap["dialingMode"].(string); ok {
		o.DialingMode = &DialingMode
	}
    
	if Script, ok := DialercampaignconfigchangecampaignMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if EdgeGroup, ok := DialercampaignconfigchangecampaignMap["edgeGroup"].(map[string]interface{}); ok {
		EdgeGroupString, _ := json.Marshal(EdgeGroup)
		json.Unmarshal(EdgeGroupString, &o.EdgeGroup)
	}
	
	if Site, ok := DialercampaignconfigchangecampaignMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if CampaignStatus, ok := DialercampaignconfigchangecampaignMap["campaignStatus"].(string); ok {
		o.CampaignStatus = &CampaignStatus
	}
    
	if PhoneColumns, ok := DialercampaignconfigchangecampaignMap["phoneColumns"].([]interface{}); ok {
		PhoneColumnsString, _ := json.Marshal(PhoneColumns)
		json.Unmarshal(PhoneColumnsString, &o.PhoneColumns)
	}
	
	if AbandonRate, ok := DialercampaignconfigchangecampaignMap["abandonRate"].(float64); ok {
		AbandonRateFloat32 := float32(AbandonRate)
		o.AbandonRate = &AbandonRateFloat32
	}
    
	if DncLists, ok := DialercampaignconfigchangecampaignMap["dncLists"].([]interface{}); ok {
		DncListsString, _ := json.Marshal(DncLists)
		json.Unmarshal(DncListsString, &o.DncLists)
	}
	
	if CallableTimeSet, ok := DialercampaignconfigchangecampaignMap["callableTimeSet"].(map[string]interface{}); ok {
		CallableTimeSetString, _ := json.Marshal(CallableTimeSet)
		json.Unmarshal(CallableTimeSetString, &o.CallableTimeSet)
	}
	
	if CallAnalysisResponseSet, ok := DialercampaignconfigchangecampaignMap["callAnalysisResponseSet"].(map[string]interface{}); ok {
		CallAnalysisResponseSetString, _ := json.Marshal(CallAnalysisResponseSet)
		json.Unmarshal(CallAnalysisResponseSetString, &o.CallAnalysisResponseSet)
	}
	
	if CallerName, ok := DialercampaignconfigchangecampaignMap["callerName"].(string); ok {
		o.CallerName = &CallerName
	}
    
	if CallerAddress, ok := DialercampaignconfigchangecampaignMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if OutboundLineCount, ok := DialercampaignconfigchangecampaignMap["outboundLineCount"].(float64); ok {
		OutboundLineCountInt := int(OutboundLineCount)
		o.OutboundLineCount = &OutboundLineCountInt
	}
	
	if Errors, ok := DialercampaignconfigchangecampaignMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if RuleSets, ok := DialercampaignconfigchangecampaignMap["ruleSets"].([]interface{}); ok {
		RuleSetsString, _ := json.Marshal(RuleSets)
		json.Unmarshal(RuleSetsString, &o.RuleSets)
	}
	
	if SkipPreviewDisabled, ok := DialercampaignconfigchangecampaignMap["skipPreviewDisabled"].(bool); ok {
		o.SkipPreviewDisabled = &SkipPreviewDisabled
	}
    
	if PreviewTimeOutSeconds, ok := DialercampaignconfigchangecampaignMap["previewTimeOutSeconds"].(float64); ok {
		PreviewTimeOutSecondsInt := int(PreviewTimeOutSeconds)
		o.PreviewTimeOutSeconds = &PreviewTimeOutSecondsInt
	}
	
	if SingleNumberPreview, ok := DialercampaignconfigchangecampaignMap["singleNumberPreview"].(bool); ok {
		o.SingleNumberPreview = &SingleNumberPreview
	}
    
	if ContactSort, ok := DialercampaignconfigchangecampaignMap["contactSort"].(map[string]interface{}); ok {
		ContactSortString, _ := json.Marshal(ContactSort)
		json.Unmarshal(ContactSortString, &o.ContactSort)
	}
	
	if ContactSorts, ok := DialercampaignconfigchangecampaignMap["contactSorts"].([]interface{}); ok {
		ContactSortsString, _ := json.Marshal(ContactSorts)
		json.Unmarshal(ContactSortsString, &o.ContactSorts)
	}
	
	if NoAnswerTimeout, ok := DialercampaignconfigchangecampaignMap["noAnswerTimeout"].(float64); ok {
		NoAnswerTimeoutInt := int(NoAnswerTimeout)
		o.NoAnswerTimeout = &NoAnswerTimeoutInt
	}
	
	if CallAnalysisLanguage, ok := DialercampaignconfigchangecampaignMap["callAnalysisLanguage"].(string); ok {
		o.CallAnalysisLanguage = &CallAnalysisLanguage
	}
    
	if Priority, ok := DialercampaignconfigchangecampaignMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if ContactListFilters, ok := DialercampaignconfigchangecampaignMap["contactListFilters"].([]interface{}); ok {
		ContactListFiltersString, _ := json.Marshal(ContactListFilters)
		json.Unmarshal(ContactListFiltersString, &o.ContactListFilters)
	}
	
	if Division, ok := DialercampaignconfigchangecampaignMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if AgentOwnedColumn, ok := DialercampaignconfigchangecampaignMap["agentOwnedColumn"].(string); ok {
		o.AgentOwnedColumn = &AgentOwnedColumn
	}
    
	if Id, ok := DialercampaignconfigchangecampaignMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialercampaignconfigchangecampaignMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialercampaignconfigchangecampaignMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialercampaignconfigchangecampaignMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialercampaignconfigchangecampaignMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangecampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
