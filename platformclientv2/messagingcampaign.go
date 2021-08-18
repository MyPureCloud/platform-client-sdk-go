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

func (u *Messagingcampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingcampaign

	
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
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		Division: u.Division,
		
		CampaignStatus: u.CampaignStatus,
		
		CallableTimeSet: u.CallableTimeSet,
		
		ContactList: u.ContactList,
		
		DncLists: u.DncLists,
		
		AlwaysRunning: u.AlwaysRunning,
		
		ContactSorts: u.ContactSorts,
		
		MessagesPerMinute: u.MessagesPerMinute,
		
		Errors: u.Errors,
		
		SmsConfig: u.SmsConfig,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagingcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
