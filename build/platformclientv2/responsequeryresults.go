package platformclientv2
import (
	"encoding/json"
)

// Responsequeryresults - Used to return response query results
type Responsequeryresults struct { 
	// Results - Contains the query results
	Results *Responseentitylist `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Responsequeryresults) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
