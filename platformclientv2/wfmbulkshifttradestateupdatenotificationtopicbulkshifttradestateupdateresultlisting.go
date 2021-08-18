package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting struct { 
	// Entities
	Entities *[]Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult `json:"entities,omitempty"`

}

func (u *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting

	

	return json.Marshal(&struct { 
		Entities *[]Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
