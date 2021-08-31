package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangemessagingcampaign
type Outboundmessagingmessagingcampaignconfigchangemessagingcampaign struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Division
	Division *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"division,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// CampaignStatus
	CampaignStatus *string `json:"campaignStatus,omitempty"`


	// CallableTimeSet
	CallableTimeSet *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`


	// ContactList
	ContactList *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"contactList,omitempty"`


	// DncLists
	DncLists *[]Outboundmessagingmessagingcampaignconfigchangeurireference `json:"dncLists,omitempty"`


	// ContactListFilters
	ContactListFilters *[]Outboundmessagingmessagingcampaignconfigchangeurireference `json:"contactListFilters,omitempty"`


	// AlwaysRunning
	AlwaysRunning *bool `json:"alwaysRunning,omitempty"`


	// ContactSorts
	ContactSorts *[]Outboundmessagingmessagingcampaignconfigchangecontactsort `json:"contactSorts,omitempty"`


	// MessagesPerMinute
	MessagesPerMinute *int `json:"messagesPerMinute,omitempty"`


	// SmsConfig
	SmsConfig *Outboundmessagingmessagingcampaignconfigchangesmsconfig `json:"smsConfig,omitempty"`


	// EmailConfig
	EmailConfig *Outboundmessagingmessagingcampaignconfigchangeemailconfig `json:"emailConfig,omitempty"`


	// Errors
	Errors *[]Outboundmessagingmessagingcampaignconfigchangeerrordetail `json:"errors,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignconfigchangemessagingcampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangemessagingcampaign
	
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
		
		Division *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"division,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		CampaignStatus *string `json:"campaignStatus,omitempty"`
		
		CallableTimeSet *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`
		
		ContactList *Outboundmessagingmessagingcampaignconfigchangeurireference `json:"contactList,omitempty"`
		
		DncLists *[]Outboundmessagingmessagingcampaignconfigchangeurireference `json:"dncLists,omitempty"`
		
		ContactListFilters *[]Outboundmessagingmessagingcampaignconfigchangeurireference `json:"contactListFilters,omitempty"`
		
		AlwaysRunning *bool `json:"alwaysRunning,omitempty"`
		
		ContactSorts *[]Outboundmessagingmessagingcampaignconfigchangecontactsort `json:"contactSorts,omitempty"`
		
		MessagesPerMinute *int `json:"messagesPerMinute,omitempty"`
		
		SmsConfig *Outboundmessagingmessagingcampaignconfigchangesmsconfig `json:"smsConfig,omitempty"`
		
		EmailConfig *Outboundmessagingmessagingcampaignconfigchangeemailconfig `json:"emailConfig,omitempty"`
		
		Errors *[]Outboundmessagingmessagingcampaignconfigchangeerrordetail `json:"errors,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		CampaignStatus: o.CampaignStatus,
		
		CallableTimeSet: o.CallableTimeSet,
		
		ContactList: o.ContactList,
		
		DncLists: o.DncLists,
		
		ContactListFilters: o.ContactListFilters,
		
		AlwaysRunning: o.AlwaysRunning,
		
		ContactSorts: o.ContactSorts,
		
		MessagesPerMinute: o.MessagesPerMinute,
		
		SmsConfig: o.SmsConfig,
		
		EmailConfig: o.EmailConfig,
		
		Errors: o.Errors,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangemessagingcampaign) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Division, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Name, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if CampaignStatus, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["campaignStatus"].(string); ok {
		o.CampaignStatus = &CampaignStatus
	}
	
	if CallableTimeSet, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["callableTimeSet"].(map[string]interface{}); ok {
		CallableTimeSetString, _ := json.Marshal(CallableTimeSet)
		json.Unmarshal(CallableTimeSetString, &o.CallableTimeSet)
	}
	
	if ContactList, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if DncLists, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["dncLists"].([]interface{}); ok {
		DncListsString, _ := json.Marshal(DncLists)
		json.Unmarshal(DncListsString, &o.DncLists)
	}
	
	if ContactListFilters, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["contactListFilters"].([]interface{}); ok {
		ContactListFiltersString, _ := json.Marshal(ContactListFilters)
		json.Unmarshal(ContactListFiltersString, &o.ContactListFilters)
	}
	
	if AlwaysRunning, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["alwaysRunning"].(bool); ok {
		o.AlwaysRunning = &AlwaysRunning
	}
	
	if ContactSorts, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["contactSorts"].([]interface{}); ok {
		ContactSortsString, _ := json.Marshal(ContactSorts)
		json.Unmarshal(ContactSortsString, &o.ContactSorts)
	}
	
	if MessagesPerMinute, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["messagesPerMinute"].(float64); ok {
		MessagesPerMinuteInt := int(MessagesPerMinute)
		o.MessagesPerMinute = &MessagesPerMinuteInt
	}
	
	if SmsConfig, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["smsConfig"].(map[string]interface{}); ok {
		SmsConfigString, _ := json.Marshal(SmsConfig)
		json.Unmarshal(SmsConfigString, &o.SmsConfig)
	}
	
	if EmailConfig, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["emailConfig"].(map[string]interface{}); ok {
		EmailConfigString, _ := json.Marshal(EmailConfig)
		json.Unmarshal(EmailConfigString, &o.EmailConfig)
	}
	
	if Errors, ok := OutboundmessagingmessagingcampaignconfigchangemessagingcampaignMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangemessagingcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
