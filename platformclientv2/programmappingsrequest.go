package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Programmappingsrequest
type Programmappingsrequest struct { 
	// QueueIds - The program queues
	QueueIds *[]string `json:"queueIds,omitempty"`


	// FlowIds - The program flows
	FlowIds *[]string `json:"flowIds,omitempty"`

}

func (o *Programmappingsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programmappingsrequest
	
	return json.Marshal(&struct { 
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		FlowIds *[]string `json:"flowIds,omitempty"`
		*Alias
	}{ 
		QueueIds: o.QueueIds,
		
		FlowIds: o.FlowIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Programmappingsrequest) UnmarshalJSON(b []byte) error {
	var ProgrammappingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ProgrammappingsrequestMap)
	if err != nil {
		return err
	}
	
	if QueueIds, ok := ProgrammappingsrequestMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	
	if FlowIds, ok := ProgrammappingsrequestMap["flowIds"].([]interface{}); ok {
		FlowIdsString, _ := json.Marshal(FlowIds)
		json.Unmarshal(FlowIdsString, &o.FlowIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Programmappingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
