package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkmetricsnetworktypeip
type Trunkinstancetopictrunkmetricsnetworktypeip struct { 
	// Address
	Address *string `json:"address,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsnetworktypeip) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
