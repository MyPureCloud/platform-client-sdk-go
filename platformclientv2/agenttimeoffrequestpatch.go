package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Agenttimeoffrequestpatch
type Agenttimeoffrequestpatch struct { 
	// MarkedAsRead - Whether this request has been read by the agent
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`


	// Status - The status of this time off request. Can only be canceled if the requested date has not already passed
	Status *string `json:"status,omitempty"`


	// Notes - Notes about the time off request. Can only be edited while the request is still pending
	Notes *string `json:"notes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agenttimeoffrequestpatch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
