package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerattemptlimitsconfigchangerecallentry
type Dialerattemptlimitsconfigchangerecallentry struct { 
	// NbrAttempts
	NbrAttempts *int `json:"nbrAttempts,omitempty"`


	// MinutesBetweenAttempts
	MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialerattemptlimitsconfigchangerecallentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerattemptlimitsconfigchangerecallentry

	

	return json.Marshal(&struct { 
		NbrAttempts *int `json:"nbrAttempts,omitempty"`
		
		MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		NbrAttempts: u.NbrAttempts,
		
		MinutesBetweenAttempts: u.MinutesBetweenAttempts,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerattemptlimitsconfigchangerecallentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
