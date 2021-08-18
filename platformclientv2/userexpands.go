package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userexpands
type Userexpands struct { 
	// RoutingStatus - ACD routing status
	RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`


	// Presence - Active presence
	Presence *Userpresence `json:"presence,omitempty"`


	// ConversationSummary - Summary of conversion statistics for conversation types.
	ConversationSummary *Userconversationsummary `json:"conversationSummary,omitempty"`


	// OutOfOffice - Determine if out of office is enabled
	OutOfOffice *Outofoffice `json:"outOfOffice,omitempty"`


	// Geolocation - Current geolocation position
	Geolocation *Geolocation `json:"geolocation,omitempty"`


	// Station - Effective, default, and last station information
	Station *Userstations `json:"station,omitempty"`


	// Authorization - Roles and permissions assigned to the user
	Authorization *Userauthorization `json:"authorization,omitempty"`

}

func (u *Userexpands) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userexpands

	

	return json.Marshal(&struct { 
		RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`
		
		Presence *Userpresence `json:"presence,omitempty"`
		
		ConversationSummary *Userconversationsummary `json:"conversationSummary,omitempty"`
		
		OutOfOffice *Outofoffice `json:"outOfOffice,omitempty"`
		
		Geolocation *Geolocation `json:"geolocation,omitempty"`
		
		Station *Userstations `json:"station,omitempty"`
		
		Authorization *Userauthorization `json:"authorization,omitempty"`
		*Alias
	}{ 
		RoutingStatus: u.RoutingStatus,
		
		Presence: u.Presence,
		
		ConversationSummary: u.ConversationSummary,
		
		OutOfOffice: u.OutOfOffice,
		
		Geolocation: u.Geolocation,
		
		Station: u.Station,
		
		Authorization: u.Authorization,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userexpands) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
