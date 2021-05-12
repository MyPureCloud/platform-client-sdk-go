package platformclientv2
import (
	"time"
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


	// Errors
	Errors *[]Outboundmessagingmessagingcampaignconfigchangeerrordetail `json:"errors,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangemessagingcampaign) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
