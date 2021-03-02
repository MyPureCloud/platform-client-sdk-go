package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Messagingcampaign) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
