package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Destination
type Destination struct { 
	// Address - Address or phone number.
	Address *string `json:"address,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`

}

func (u *Destination) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Destination

	

	return json.Marshal(&struct { 
		Address *string `json:"address,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		*Alias
	}{ 
		Address: u.Address,
		
		Name: u.Name,
		
		UserId: u.UserId,
		
		QueueId: u.QueueId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Destination) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
