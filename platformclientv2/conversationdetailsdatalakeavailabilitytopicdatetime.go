package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdetailsdatalakeavailabilitytopicdatetime
type Conversationdetailsdatalakeavailabilitytopicdatetime struct { 
	// IMillis
	IMillis *int `json:"iMillis,omitempty"`


	// BeforeNow
	BeforeNow *bool `json:"beforeNow,omitempty"`


	// AfterNow
	AfterNow *bool `json:"afterNow,omitempty"`


	// EqualNow
	EqualNow *bool `json:"equalNow,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationdetailsdatalakeavailabilitytopicdatetime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
