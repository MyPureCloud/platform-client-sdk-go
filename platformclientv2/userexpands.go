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


	// IntegrationPresence - Active 3rd party presence
	IntegrationPresence *Userpresence `json:"integrationPresence,omitempty"`


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

func (o *Userexpands) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userexpands
	
	return json.Marshal(&struct { 
		RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`
		
		Presence *Userpresence `json:"presence,omitempty"`
		
		IntegrationPresence *Userpresence `json:"integrationPresence,omitempty"`
		
		ConversationSummary *Userconversationsummary `json:"conversationSummary,omitempty"`
		
		OutOfOffice *Outofoffice `json:"outOfOffice,omitempty"`
		
		Geolocation *Geolocation `json:"geolocation,omitempty"`
		
		Station *Userstations `json:"station,omitempty"`
		
		Authorization *Userauthorization `json:"authorization,omitempty"`
		*Alias
	}{ 
		RoutingStatus: o.RoutingStatus,
		
		Presence: o.Presence,
		
		IntegrationPresence: o.IntegrationPresence,
		
		ConversationSummary: o.ConversationSummary,
		
		OutOfOffice: o.OutOfOffice,
		
		Geolocation: o.Geolocation,
		
		Station: o.Station,
		
		Authorization: o.Authorization,
		Alias:    (*Alias)(o),
	})
}

func (o *Userexpands) UnmarshalJSON(b []byte) error {
	var UserexpandsMap map[string]interface{}
	err := json.Unmarshal(b, &UserexpandsMap)
	if err != nil {
		return err
	}
	
	if RoutingStatus, ok := UserexpandsMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if Presence, ok := UserexpandsMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	
	if IntegrationPresence, ok := UserexpandsMap["integrationPresence"].(map[string]interface{}); ok {
		IntegrationPresenceString, _ := json.Marshal(IntegrationPresence)
		json.Unmarshal(IntegrationPresenceString, &o.IntegrationPresence)
	}
	
	if ConversationSummary, ok := UserexpandsMap["conversationSummary"].(map[string]interface{}); ok {
		ConversationSummaryString, _ := json.Marshal(ConversationSummary)
		json.Unmarshal(ConversationSummaryString, &o.ConversationSummary)
	}
	
	if OutOfOffice, ok := UserexpandsMap["outOfOffice"].(map[string]interface{}); ok {
		OutOfOfficeString, _ := json.Marshal(OutOfOffice)
		json.Unmarshal(OutOfOfficeString, &o.OutOfOffice)
	}
	
	if Geolocation, ok := UserexpandsMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if Station, ok := UserexpandsMap["station"].(map[string]interface{}); ok {
		StationString, _ := json.Marshal(Station)
		json.Unmarshal(StationString, &o.Station)
	}
	
	if Authorization, ok := UserexpandsMap["authorization"].(map[string]interface{}); ok {
		AuthorizationString, _ := json.Marshal(Authorization)
		json.Unmarshal(AuthorizationString, &o.Authorization)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userexpands) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
