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

func (u *Useraggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useraggregatequeryresponse

	

	return json.Marshal(&struct { 
		SystemToOrganizationMappings *map[string][]string `json:"systemToOrganizationMappings,omitempty"`
		
		Results *[]Useraggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		SystemToOrganizationMappings: u.SystemToOrganizationMappings,
		
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Useraggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
