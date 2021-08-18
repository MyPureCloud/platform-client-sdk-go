package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Mediautilization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediautilization

	

	return json.Marshal(&struct { 
		MaximumCapacity *int `json:"maximumCapacity,omitempty"`
		
		InterruptableMediaTypes *[]string `json:"interruptableMediaTypes,omitempty"`
		
		IncludeNonAcd *bool `json:"includeNonAcd,omitempty"`
		*Alias
	}{ 
		MaximumCapacity: u.MaximumCapacity,
		
		InterruptableMediaTypes: u.InterruptableMediaTypes,
		
		IncludeNonAcd: u.IncludeNonAcd,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediautilization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
