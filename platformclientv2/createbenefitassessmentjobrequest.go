package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Createbenefitassessmentjobrequest
type Createbenefitassessmentjobrequest struct { 
	// DivisionIds - The list of division ids for routing queues that are to be assessed for Predictive Routing benefit.
	DivisionIds *[]string `json:"divisionIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createbenefitassessmentjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
