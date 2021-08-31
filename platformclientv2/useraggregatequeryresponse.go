package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Useraggregatequeryresponse
type Useraggregatequeryresponse struct { 
	// SystemToOrganizationMappings - A mapping from system presence to a list of organization presence ids
	SystemToOrganizationMappings *map[string][]string `json:"systemToOrganizationMappings,omitempty"`


	// Results
	Results *[]Useraggregatedatacontainer `json:"results,omitempty"`

}

func (o *Useraggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useraggregatequeryresponse
	
	return json.Marshal(&struct { 
		SystemToOrganizationMappings *map[string][]string `json:"systemToOrganizationMappings,omitempty"`
		
		Results *[]Useraggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		SystemToOrganizationMappings: o.SystemToOrganizationMappings,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Useraggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var UseraggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UseraggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if SystemToOrganizationMappings, ok := UseraggregatequeryresponseMap["systemToOrganizationMappings"].(map[string]interface{}); ok {
		SystemToOrganizationMappingsString, _ := json.Marshal(SystemToOrganizationMappings)
		json.Unmarshal(SystemToOrganizationMappingsString, &o.SystemToOrganizationMappings)
	}
	
	if Results, ok := UseraggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Useraggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
