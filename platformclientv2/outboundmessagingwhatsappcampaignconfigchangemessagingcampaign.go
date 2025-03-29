package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingwhatsappcampaignconfigchangemessagingcampaign
type Outboundmessagingwhatsappcampaignconfigchangemessagingcampaign struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CampaignStatus
	CampaignStatus *string `json:"campaignStatus,omitempty"`

	// CallableTimeSet
	CallableTimeSet *Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`

	// ContactList - A UriReference for a resource
	ContactList *Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"contactList,omitempty"`

	// DncLists - The dnc lists to check before sending a message for this messaging campaign.
	DncLists *[]Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"dncLists,omitempty"`

	// ContactListFilters - The contact list filters to check before sending a message for this messaging campaign.
	ContactListFilters *[]Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"contactListFilters,omitempty"`

	// AlwaysRunning - Whether this messaging campaign is always running.
	AlwaysRunning *bool `json:"alwaysRunning,omitempty"`

	// ContactSorts - The order in which to sort contacts for dialing, based on up to four columns.
	ContactSorts *[]Outboundmessagingwhatsappcampaignconfigchangecontactsort `json:"contactSorts,omitempty"`

	// MessagesPerMinute - How many messages this messaging campaign will send per minute.
	MessagesPerMinute *int `json:"messagesPerMinute,omitempty"`

	// RuleSets
	RuleSets *[]Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"ruleSets,omitempty"`

	// SmsConfig
	SmsConfig *Outboundmessagingwhatsappcampaignconfigchangesmsconfig `json:"smsConfig,omitempty"`

	// EmailConfig
	EmailConfig *Outboundmessagingwhatsappcampaignconfigchangeemailconfig `json:"emailConfig,omitempty"`

	// WhatsAppConfig
	WhatsAppConfig *Outboundmessagingwhatsappcampaignconfigchangewhatsappconfig `json:"whatsAppConfig,omitempty"`

	// Errors - A list of current error conditions associated with this messaging campaign
	Errors *[]Outboundmessagingwhatsappcampaignconfigchangeerrordetail `json:"errors,omitempty"`

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

	// Division - A UriReference for a resource
	Division *Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"division,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outboundmessagingwhatsappcampaignconfigchangemessagingcampaign) SetField(field string, fieldValue interface{}) {
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

func (o Outboundmessagingwhatsappcampaignconfigchangemessagingcampaign) MarshalJSON() ([]byte, error) {
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
	type Alias Outboundmessagingwhatsappcampaignconfigchangemessagingcampaign
	
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
		CampaignStatus *string `json:"campaignStatus,omitempty"`
		
		CallableTimeSet *Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"callableTimeSet,omitempty"`
		
		ContactList *Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"contactList,omitempty"`
		
		DncLists *[]Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"dncLists,omitempty"`
		
		ContactListFilters *[]Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"contactListFilters,omitempty"`
		
		AlwaysRunning *bool `json:"alwaysRunning,omitempty"`
		
		ContactSorts *[]Outboundmessagingwhatsappcampaignconfigchangecontactsort `json:"contactSorts,omitempty"`
		
		MessagesPerMinute *int `json:"messagesPerMinute,omitempty"`
		
		RuleSets *[]Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"ruleSets,omitempty"`
		
		SmsConfig *Outboundmessagingwhatsappcampaignconfigchangesmsconfig `json:"smsConfig,omitempty"`
		
		EmailConfig *Outboundmessagingwhatsappcampaignconfigchangeemailconfig `json:"emailConfig,omitempty"`
		
		WhatsAppConfig *Outboundmessagingwhatsappcampaignconfigchangewhatsappconfig `json:"whatsAppConfig,omitempty"`
		
		Errors *[]Outboundmessagingwhatsappcampaignconfigchangeerrordetail `json:"errors,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Division *Outboundmessagingwhatsappcampaignconfigchangeurireference `json:"division,omitempty"`
		Alias
	}{ 
		CampaignStatus: o.CampaignStatus,
		
		CallableTimeSet: o.CallableTimeSet,
		
		ContactList: o.ContactList,
		
		DncLists: o.DncLists,
		
		ContactListFilters: o.ContactListFilters,
		
		AlwaysRunning: o.AlwaysRunning,
		
		ContactSorts: o.ContactSorts,
		
		MessagesPerMinute: o.MessagesPerMinute,
		
		RuleSets: o.RuleSets,
		
		SmsConfig: o.SmsConfig,
		
		EmailConfig: o.EmailConfig,
		
		WhatsAppConfig: o.WhatsAppConfig,
		
		Errors: o.Errors,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Division: o.Division,
		Alias:    (Alias)(o),
	})
}

func (o *Outboundmessagingwhatsappcampaignconfigchangemessagingcampaign) UnmarshalJSON(b []byte) error {
	var OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap)
	if err != nil {
		return err
	}
	
	if CampaignStatus, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["campaignStatus"].(string); ok {
		o.CampaignStatus = &CampaignStatus
	}
    
	if CallableTimeSet, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["callableTimeSet"].(map[string]interface{}); ok {
		CallableTimeSetString, _ := json.Marshal(CallableTimeSet)
		json.Unmarshal(CallableTimeSetString, &o.CallableTimeSet)
	}
	
	if ContactList, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if DncLists, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["dncLists"].([]interface{}); ok {
		DncListsString, _ := json.Marshal(DncLists)
		json.Unmarshal(DncListsString, &o.DncLists)
	}
	
	if ContactListFilters, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["contactListFilters"].([]interface{}); ok {
		ContactListFiltersString, _ := json.Marshal(ContactListFilters)
		json.Unmarshal(ContactListFiltersString, &o.ContactListFilters)
	}
	
	if AlwaysRunning, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["alwaysRunning"].(bool); ok {
		o.AlwaysRunning = &AlwaysRunning
	}
    
	if ContactSorts, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["contactSorts"].([]interface{}); ok {
		ContactSortsString, _ := json.Marshal(ContactSorts)
		json.Unmarshal(ContactSortsString, &o.ContactSorts)
	}
	
	if MessagesPerMinute, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["messagesPerMinute"].(float64); ok {
		MessagesPerMinuteInt := int(MessagesPerMinute)
		o.MessagesPerMinute = &MessagesPerMinuteInt
	}
	
	if RuleSets, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["ruleSets"].([]interface{}); ok {
		RuleSetsString, _ := json.Marshal(RuleSets)
		json.Unmarshal(RuleSetsString, &o.RuleSets)
	}
	
	if SmsConfig, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["smsConfig"].(map[string]interface{}); ok {
		SmsConfigString, _ := json.Marshal(SmsConfig)
		json.Unmarshal(SmsConfigString, &o.SmsConfig)
	}
	
	if EmailConfig, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["emailConfig"].(map[string]interface{}); ok {
		EmailConfigString, _ := json.Marshal(EmailConfig)
		json.Unmarshal(EmailConfigString, &o.EmailConfig)
	}
	
	if WhatsAppConfig, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["whatsAppConfig"].(map[string]interface{}); ok {
		WhatsAppConfigString, _ := json.Marshal(WhatsAppConfig)
		json.Unmarshal(WhatsAppConfigString, &o.WhatsAppConfig)
	}
	
	if Errors, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if Id, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Division, ok := OutboundmessagingwhatsappcampaignconfigchangemessagingcampaignMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingwhatsappcampaignconfigchangemessagingcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
