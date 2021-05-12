package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingappointmenttopiccoachingappointmentexternallink
type Wemcoachingappointmenttopiccoachingappointmentexternallink struct { 
	// ExternalLink
	ExternalLink *string `json:"externalLink,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentexternallink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
