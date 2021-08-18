package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopiclinestatus
type Phonechangetopiclinestatus struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Reachable
	Reachable *bool `json:"reachable,omitempty"`


	// AddressOfRecord
	AddressOfRecord *string `json:"addressOfRecord,omitempty"`


	// ContactAddresses
	ContactAddresses *[]string `json:"contactAddresses,omitempty"`


	// ReachableStateTime
	ReachableStateTime *time.Time `json:"reachableStateTime,omitempty"`

}

func (u *Phonechangetopiclinestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopiclinestatus

	
	ReachableStateTime := new(string)
	if u.ReachableStateTime != nil {
		
		*ReachableStateTime = timeutil.Strftime(u.ReachableStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReachableStateTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Reachable *bool `json:"reachable,omitempty"`
		
		AddressOfRecord *string `json:"addressOfRecord,omitempty"`
		
		ContactAddresses *[]string `json:"contactAddresses,omitempty"`
		
		ReachableStateTime *string `json:"reachableStateTime,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Reachable: u.Reachable,
		
		AddressOfRecord: u.AddressOfRecord,
		
		ContactAddresses: u.ContactAddresses,
		
		ReachableStateTime: ReachableStateTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonechangetopiclinestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
