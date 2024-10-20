package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaign
type Campaign struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// AbandonRate - The targeted compliance abandon rate percentage. Required for power and predictive campaigns.
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

	// DynamicContactQueueingSettings - Settings for dynamic queueing of contacts.
	DynamicContactQueueingSettings *Dynamiccontactqueueingsettings `json:"dynamicContactQueueingSettings,omitempty"`

	// SkillColumns - The skill columns on the ContactList that this Campaign should take into account when dialing
	SkillColumns *[]string `json:"skillColumns,omitempty"`

	// MaxCallsPerAgent - The maximum number of calls that can be placed per agent on this campaign
	MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`

	// CallbackAutoAnswer - The option manages the auto-answer callback calls
	CallbackAutoAnswer *bool `json:"callbackAutoAnswer,omitempty"`

	// DynamicLineBalancingSettings - Dynamic line balancing settings
	DynamicLineBalancingSettings *Dynamiclinebalancingsettings `json:"dynamicLineBalancingSettings,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaign) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Campaign) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaign
	
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
		
		DynamicContactQueueingSettings *Dynamiccontactqueueingsettings `json:"dynamicContactQueueingSettings,omitempty"`
		
		SkillColumns *[]string `json:"skillColumns,omitempty"`
		
		MaxCallsPerAgent *int `json:"maxCallsPerAgent,omitempty"`
		
		CallbackAutoAnswer *bool `json:"callbackAutoAnswer,omitempty"`
		
		DynamicLineBalancingSettings *Dynamiclinebalancingsettings `json:"dynamicLineBalancingSettings,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
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
		
		Errors: o.Errors,
		
		CallerName: o.CallerName,
		
		CallerAddress: o.CallerAddress,
		
		OutboundLineCount: o.OutboundLineCount,
		
		RuleSets: o.RuleSets,
		
		SkipPreviewDisabled: o.SkipPreviewDisabled,
		
		PreviewTimeOutSeconds: o.PreviewTimeOutSeconds,
		
		AlwaysRunning: o.AlwaysRunning,
		
		ContactSort: o.ContactSort,
		
		ContactSorts: o.ContactSorts,
		
		NoAnswerTimeout: o.NoAnswerTimeout,
		
		CallAnalysisLanguage: o.CallAnalysisLanguage,
		
		Priority: o.Priority,
		
		ContactListFilters: o.ContactListFilters,
		
		Division: o.Division,
		
		DynamicContactQueueingSettings: o.DynamicContactQueueingSettings,
		
		SkillColumns: o.SkillColumns,
		
		MaxCallsPerAgent: o.MaxCallsPerAgent,
		
		CallbackAutoAnswer: o.CallbackAutoAnswer,
		
		DynamicLineBalancingSettings: o.DynamicLineBalancingSettings,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Campaign) UnmarshalJSON(b []byte) error {
	var CampaignMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CampaignMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CampaignMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := CampaignMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := CampaignMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := CampaignMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ContactList, ok := CampaignMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if Queue, ok := CampaignMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if DialingMode, ok := CampaignMap["dialingMode"].(string); ok {
		o.DialingMode = &DialingMode
	}
    
	if Script, ok := CampaignMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if EdgeGroup, ok := CampaignMap["edgeGroup"].(map[string]interface{}); ok {
		EdgeGroupString, _ := json.Marshal(EdgeGroup)
		json.Unmarshal(EdgeGroupString, &o.EdgeGroup)
	}
	
	if Site, ok := CampaignMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if CampaignStatus, ok := CampaignMap["campaignStatus"].(string); ok {
		o.CampaignStatus = &CampaignStatus
	}
    
	if PhoneColumns, ok := CampaignMap["phoneColumns"].([]interface{}); ok {
		PhoneColumnsString, _ := json.Marshal(PhoneColumns)
		json.Unmarshal(PhoneColumnsString, &o.PhoneColumns)
	}
	
	if AbandonRate, ok := CampaignMap["abandonRate"].(float64); ok {
		o.AbandonRate = &AbandonRate
	}
    
	if DncLists, ok := CampaignMap["dncLists"].([]interface{}); ok {
		DncListsString, _ := json.Marshal(DncLists)
		json.Unmarshal(DncListsString, &o.DncLists)
	}
	
	if CallableTimeSet, ok := CampaignMap["callableTimeSet"].(map[string]interface{}); ok {
		CallableTimeSetString, _ := json.Marshal(CallableTimeSet)
		json.Unmarshal(CallableTimeSetString, &o.CallableTimeSet)
	}
	
	if CallAnalysisResponseSet, ok := CampaignMap["callAnalysisResponseSet"].(map[string]interface{}); ok {
		CallAnalysisResponseSetString, _ := json.Marshal(CallAnalysisResponseSet)
		json.Unmarshal(CallAnalysisResponseSetString, &o.CallAnalysisResponseSet)
	}
	
	if Errors, ok := CampaignMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if CallerName, ok := CampaignMap["callerName"].(string); ok {
		o.CallerName = &CallerName
	}
    
	if CallerAddress, ok := CampaignMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
    
	if OutboundLineCount, ok := CampaignMap["outboundLineCount"].(float64); ok {
		OutboundLineCountInt := int(OutboundLineCount)
		o.OutboundLineCount = &OutboundLineCountInt
	}
	
	if RuleSets, ok := CampaignMap["ruleSets"].([]interface{}); ok {
		RuleSetsString, _ := json.Marshal(RuleSets)
		json.Unmarshal(RuleSetsString, &o.RuleSets)
	}
	
	if SkipPreviewDisabled, ok := CampaignMap["skipPreviewDisabled"].(bool); ok {
		o.SkipPreviewDisabled = &SkipPreviewDisabled
	}
    
	if PreviewTimeOutSeconds, ok := CampaignMap["previewTimeOutSeconds"].(float64); ok {
		PreviewTimeOutSecondsInt := int(PreviewTimeOutSeconds)
		o.PreviewTimeOutSeconds = &PreviewTimeOutSecondsInt
	}
	
	if AlwaysRunning, ok := CampaignMap["alwaysRunning"].(bool); ok {
		o.AlwaysRunning = &AlwaysRunning
	}
    
	if ContactSort, ok := CampaignMap["contactSort"].(map[string]interface{}); ok {
		ContactSortString, _ := json.Marshal(ContactSort)
		json.Unmarshal(ContactSortString, &o.ContactSort)
	}
	
	if ContactSorts, ok := CampaignMap["contactSorts"].([]interface{}); ok {
		ContactSortsString, _ := json.Marshal(ContactSorts)
		json.Unmarshal(ContactSortsString, &o.ContactSorts)
	}
	
	if NoAnswerTimeout, ok := CampaignMap["noAnswerTimeout"].(float64); ok {
		NoAnswerTimeoutInt := int(NoAnswerTimeout)
		o.NoAnswerTimeout = &NoAnswerTimeoutInt
	}
	
	if CallAnalysisLanguage, ok := CampaignMap["callAnalysisLanguage"].(string); ok {
		o.CallAnalysisLanguage = &CallAnalysisLanguage
	}
    
	if Priority, ok := CampaignMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if ContactListFilters, ok := CampaignMap["contactListFilters"].([]interface{}); ok {
		ContactListFiltersString, _ := json.Marshal(ContactListFilters)
		json.Unmarshal(ContactListFiltersString, &o.ContactListFilters)
	}
	
	if Division, ok := CampaignMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if DynamicContactQueueingSettings, ok := CampaignMap["dynamicContactQueueingSettings"].(map[string]interface{}); ok {
		DynamicContactQueueingSettingsString, _ := json.Marshal(DynamicContactQueueingSettings)
		json.Unmarshal(DynamicContactQueueingSettingsString, &o.DynamicContactQueueingSettings)
	}
	
	if SkillColumns, ok := CampaignMap["skillColumns"].([]interface{}); ok {
		SkillColumnsString, _ := json.Marshal(SkillColumns)
		json.Unmarshal(SkillColumnsString, &o.SkillColumns)
	}
	
	if MaxCallsPerAgent, ok := CampaignMap["maxCallsPerAgent"].(float64); ok {
		MaxCallsPerAgentInt := int(MaxCallsPerAgent)
		o.MaxCallsPerAgent = &MaxCallsPerAgentInt
	}
	
	if CallbackAutoAnswer, ok := CampaignMap["callbackAutoAnswer"].(bool); ok {
		o.CallbackAutoAnswer = &CallbackAutoAnswer
	}
    
	if DynamicLineBalancingSettings, ok := CampaignMap["dynamicLineBalancingSettings"].(map[string]interface{}); ok {
		DynamicLineBalancingSettingsString, _ := json.Marshal(DynamicLineBalancingSettings)
		json.Unmarshal(DynamicLineBalancingSettingsString, &o.DynamicLineBalancingSettings)
	}
	
	if SelfUri, ok := CampaignMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Campaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
