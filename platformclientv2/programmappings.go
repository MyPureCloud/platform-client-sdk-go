package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Programmappings
type Programmappings struct { 
	// Program
	Program *Baseprogramentity `json:"program,omitempty"`


	// Queues
	Queues *[]Addressableentityref `json:"queues,omitempty"`


	// Flows
	Flows *[]Addressableentityref `json:"flows,omitempty"`


	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

}

func (o *Programmappings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programmappings
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Program *Baseprogramentity `json:"program,omitempty"`
		
		Queues *[]Addressableentityref `json:"queues,omitempty"`
		
		Flows *[]Addressableentityref `json:"flows,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		*Alias
	}{ 
		Program: o.Program,
		
		Queues: o.Queues,
		
		Flows: o.Flows,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		Alias:    (*Alias)(o),
	})
}

func (o *Programmappings) UnmarshalJSON(b []byte) error {
	var ProgrammappingsMap map[string]interface{}
	err := json.Unmarshal(b, &ProgrammappingsMap)
	if err != nil {
		return err
	}
	
	if Program, ok := ProgrammappingsMap["program"].(map[string]interface{}); ok {
		ProgramString, _ := json.Marshal(Program)
		json.Unmarshal(ProgramString, &o.Program)
	}
	
	if Queues, ok := ProgrammappingsMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if Flows, ok := ProgrammappingsMap["flows"].([]interface{}); ok {
		FlowsString, _ := json.Marshal(Flows)
		json.Unmarshal(FlowsString, &o.Flows)
	}
	
	if ModifiedBy, ok := ProgrammappingsMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := ProgrammappingsMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Programmappings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
