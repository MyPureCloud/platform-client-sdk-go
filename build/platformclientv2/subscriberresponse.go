package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Subscriberresponse
type Subscriberresponse struct { 
	// MessageReturned - Suggested valid addresses
	MessageReturned *[]string `json:"messageReturned,omitempty"`


	// Status - http status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Subscriberresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
