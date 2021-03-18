package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Useraggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
