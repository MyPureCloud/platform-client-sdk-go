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

func (u *Outboundmessagingmessagingcampaignconfigchangemessagingcampaign) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangemessagingcampaign

	
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
		Id: u.Id,
		
		Division: u.Division,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		CampaignStatus: u.CampaignStatus,
		
		CallableTimeSet: u.CallableTimeSet,
		
		ContactList: u.ContactList,
		
		DncLists: u.DncLists,
		
		ContactListFilters: u.ContactListFilters,
		
		AlwaysRunning: u.AlwaysRunning,
		
		ContactSorts: u.ContactSorts,
		
		MessagesPerMinute: u.MessagesPerMinute,
		
		SmsConfig: u.SmsConfig,
		
		EmailConfig: u.EmailConfig,
		
		Errors: u.Errors,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangemessagingcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
