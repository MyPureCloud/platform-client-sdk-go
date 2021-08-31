package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsuserdetail
type Analyticsuserdetail struct { 
	// UserId - The identifier for the user
	UserId *string `json:"userId,omitempty"`


	// PrimaryPresence - The presence records for the user
	PrimaryPresence *[]Analyticsuserpresencerecord `json:"primaryPresence,omitempty"`


	// RoutingStatus - The ACD routing status records for the user
	RoutingStatus *[]Analyticsroutingstatusrecord `json:"routingStatus,omitempty"`

}

func (o *Analyticsuserdetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsuserdetail
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		PrimaryPresence *[]Analyticsuserpresencerecord `json:"primaryPresence,omitempty"`
		
		RoutingStatus *[]Analyticsroutingstatusrecord `json:"routingStatus,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		PrimaryPresence: o.PrimaryPresence,
		
		RoutingStatus: o.RoutingStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsuserdetail) UnmarshalJSON(b []byte) error {
	var AnalyticsuserdetailMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsuserdetailMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := AnalyticsuserdetailMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if PrimaryPresence, ok := AnalyticsuserdetailMap["primaryPresence"].([]interface{}); ok {
		PrimaryPresenceString, _ := json.Marshal(PrimaryPresence)
		json.Unmarshal(PrimaryPresenceString, &o.PrimaryPresence)
	}
	
	if RoutingStatus, ok := AnalyticsuserdetailMap["routingStatus"].([]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsuserdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
