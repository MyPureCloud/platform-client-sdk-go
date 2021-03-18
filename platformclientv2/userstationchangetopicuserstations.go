package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userstationchangetopicuserstations
type Userstationchangetopicuserstations struct { 
	// AssociatedStation
	AssociatedStation *Userstationchangetopicuserstation `json:"associatedStation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
