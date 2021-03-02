package platformclientv2
import (
	"encoding/json"
)

// Queueobservationqueryresponse
type Queueobservationqueryresponse struct { 
	// SystemToOrganizationMappings - A mapping from system presence to a list of organization presence ids
	SystemToOrganizationMappings *map[string][]string `json:"systemToOrganizationMappings,omitempty"`


	// Results
	Results *[]Queueobservationdatacontainer `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueobservationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
