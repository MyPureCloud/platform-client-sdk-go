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

func (u *Analyticsuserdetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsuserdetail

	

	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		PrimaryPresence *[]Analyticsuserpresencerecord `json:"primaryPresence,omitempty"`
		
		RoutingStatus *[]Analyticsroutingstatusrecord `json:"routingStatus,omitempty"`
		*Alias
	}{ 
		UserId: u.UserId,
		
		PrimaryPresence: u.PrimaryPresence,
		
		RoutingStatus: u.RoutingStatus,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsuserdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
