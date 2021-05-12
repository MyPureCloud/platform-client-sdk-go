package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Responsequeryresults - Used to return response query results
type Responsequeryresults struct { 
	// Results - Contains the query results
	Results *Responseentitylist `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Responsequeryresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
