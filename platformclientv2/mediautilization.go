package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Mediautilization
type Mediautilization struct { 
	// MaximumCapacity - Defines the maximum number of conversations of this type that an agent can handle at one time.
	MaximumCapacity *int `json:"maximumCapacity,omitempty"`


	// InterruptableMediaTypes - Defines the list of other media types that can interrupt a conversation of this media type.  Values include call, chat, email, callback, and message.
	InterruptableMediaTypes *[]string `json:"interruptableMediaTypes,omitempty"`


	// IncludeNonAcd - If true, then track non-ACD conversations against utilization
	IncludeNonAcd *bool `json:"includeNonAcd,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediautilization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
