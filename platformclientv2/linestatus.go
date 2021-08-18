package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Linestatus
type Linestatus struct { 
	// Id - The id of this line
	Id *string `json:"id,omitempty"`


	// Reachable - Indicates whether the edge can reach the line.
	Reachable *bool `json:"reachable,omitempty"`


	// AddressOfRecord - The line's address of record.
	AddressOfRecord *string `json:"addressOfRecord,omitempty"`


	// ContactAddresses - The addresses used to contact the line.
	ContactAddresses *[]string `json:"contactAddresses,omitempty"`


	// ReachableStateTime - The time the line entered its current reachable state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReachableStateTime *time.Time `json:"reachableStateTime,omitempty"`

}

func (u *Linestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Linestatus

	
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
func (o *Linestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
