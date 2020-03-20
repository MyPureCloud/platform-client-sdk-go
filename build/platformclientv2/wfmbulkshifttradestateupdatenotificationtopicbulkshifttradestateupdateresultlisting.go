package platformclientv2
import (
	"encoding/json"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting struct { 
	// Entities
	Entities *[]Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
