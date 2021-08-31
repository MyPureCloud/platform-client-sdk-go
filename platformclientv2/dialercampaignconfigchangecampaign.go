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
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
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
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignconfigchangecampaign) UnmarshalJSON(b []byte) error {
	var DialercampaignconfigchangecampaignMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignconfigchangecampaignMap)
	if err != nil {
		return err
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
	
	if AdditionalProperties, ok := DialercampaignconfigchangecampaignMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangecampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
