package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueobservationqueryresponse
type Queueobservationqueryresponse struct { 
	// SystemToOrganizationMappings - A mapping from system presence to a list of organization presence ids
	SystemToOrganizationMappings *map[string][]string `json:"systemToOrganizationMappings,omitempty"`


	// Results
	Results *[]Queueobservationdatacontainer `json:"results,omitempty"`

}

func (o *Queueobservationqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueobservationqueryresponse
	
	return json.Marshal(&struct { 
		SystemToOrganizationMappings *map[string][]string `json:"systemToOrganizationMappings,omitempty"`
		
		Results *[]Queueobservationdatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		SystemToOrganizationMappings: o.SystemToOrganizationMappings,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueobservationqueryresponse) UnmarshalJSON(b []byte) error {
	var QueueobservationqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &QueueobservationqueryresponseMap)
	if err != nil {
		return err
	}
	
	if SystemToOrganizationMappings, ok := QueueobservationqueryresponseMap["systemToOrganizationMappings"].(map[string]interface{}); ok {
		SystemToOrganizationMappingsString, _ := json.Marshal(SystemToOrganizationMappings)
		json.Unmarshal(SystemToOrganizationMappingsString, &o.SystemToOrganizationMappings)
	}
	
	if Results, ok := QueueobservationqueryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueobservationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
