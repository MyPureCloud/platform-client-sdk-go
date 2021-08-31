package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingcampaign
type Messagingcampaign struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// Division - The division this entity belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// CampaignStatus - The current status of the messaging campaign. A messaging campaign may be turned 'on' or 'off'.
	CampaignStatus *string `json:"campaignStatus,omitempty"`


	// CallableTimeSet - The callable time set for this messaging campaign.
	CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`


	// ContactList - The contact list that this messaging campaign will send messages for.
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// DncLists - The dnc lists to check before sending a message for this messaging campaign.
	DncLists *[]Domainentityref `json:"dncLists,omitempty"`


	// AlwaysRunning - Whether this messaging campaign is always running
	AlwaysRunning *bool `json:"alwaysRunning,omitempty"`


	// ContactSorts - The order in which to sort contacts for dialing, based on up to four columns.
	ContactSorts *[]Contactsort `json:"contactSorts,omitempty"`


	// MessagesPerMinute - How many messages this messaging campaign will send per minute.
	MessagesPerMinute *int `json:"messagesPerMinute,omitempty"`


	// Errors - A list of current error conditions associated with this messaging campaign.
	Errors *[]Resterrordetail `json:"errors,omitempty"`


	// SmsConfig - Configuration for this messaging campaign to send SMS messages.
	SmsConfig *Smsconfig `json:"smsConfig,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Messagingcampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingcampaign
	
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
		
		Division *Domainentityref `json:"division,omitempty"`
		
		CampaignStatus *string `json:"campaignStatus,omitempty"`
		
		CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`
		
		ContactList *Domainentityref `json:"contactList,omitempty"`
		
		DncLists *[]Domainentityref `json:"dncLists,omitempty"`
		
		AlwaysRunning *bool `json:"alwaysRunning,omitempty"`
		
		ContactSorts *[]Contactsort `json:"contactSorts,omitempty"`
		
		MessagesPerMinute *int `json:"messagesPerMinute,omitempty"`
		
		Errors *[]Resterrordetail `json:"errors,omitempty"`
		
		SmsConfig *Smsconfig `json:"smsConfig,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Division: o.Division,
		
		CampaignStatus: o.CampaignStatus,
		
		CallableTimeSet: o.CallableTimeSet,
		
		ContactList: o.ContactList,
		
		DncLists: o.DncLists,
		
		AlwaysRunning: o.AlwaysRunning,
		
		ContactSorts: o.ContactSorts,
		
		MessagesPerMinute: o.MessagesPerMinute,
		
		Errors: o.Errors,
		
		SmsConfig: o.SmsConfig,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagingcampaign) UnmarshalJSON(b []byte) error {
	var MessagingcampaignMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingcampaignMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagingcampaignMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MessagingcampaignMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := MessagingcampaignMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := MessagingcampaignMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := MessagingcampaignMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Division, ok := MessagingcampaignMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if CampaignStatus, ok := MessagingcampaignMap["campaignStatus"].(string); ok {
		o.CampaignStatus = &CampaignStatus
	}
	
	if CallableTimeSet, ok := MessagingcampaignMap["callableTimeSet"].(map[string]interface{}); ok {
		CallableTimeSetString, _ := json.Marshal(CallableTimeSet)
		json.Unmarshal(CallableTimeSetString, &o.CallableTimeSet)
	}
	
	if ContactList, ok := MessagingcampaignMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if DncLists, ok := MessagingcampaignMap["dncLists"].([]interface{}); ok {
		DncListsString, _ := json.Marshal(DncLists)
		json.Unmarshal(DncListsString, &o.DncLists)
	}
	
	if AlwaysRunning, ok := MessagingcampaignMap["alwaysRunning"].(bool); ok {
		o.AlwaysRunning = &AlwaysRunning
	}
	
	if ContactSorts, ok := MessagingcampaignMap["contactSorts"].([]interface{}); ok {
		ContactSortsString, _ := json.Marshal(ContactSorts)
		json.Unmarshal(ContactSortsString, &o.ContactSorts)
	}
	
	if MessagesPerMinute, ok := MessagingcampaignMap["messagesPerMinute"].(float64); ok {
		MessagesPerMinuteInt := int(MessagesPerMinute)
		o.MessagesPerMinute = &MessagesPerMinuteInt
	}
	
	if Errors, ok := MessagingcampaignMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if SmsConfig, ok := MessagingcampaignMap["smsConfig"].(map[string]interface{}); ok {
		SmsConfigString, _ := json.Marshal(SmsConfig)
		json.Unmarshal(SmsConfigString, &o.SmsConfig)
	}
	
	if SelfUri, ok := MessagingcampaignMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
