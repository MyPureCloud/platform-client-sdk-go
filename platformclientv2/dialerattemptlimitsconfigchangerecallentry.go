package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerattemptlimitsconfigchangerecallentry
type Dialerattemptlimitsconfigchangerecallentry struct { 
	// NbrAttempts - The number of recall attempts to make
	NbrAttempts *int `json:"nbrAttempts,omitempty"`


	// MinutesBetweenAttempts - How long to wait between recall attempts
	MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`

}

func (o *Dialerattemptlimitsconfigchangerecallentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerattemptlimitsconfigchangerecallentry
	
	return json.Marshal(&struct { 
		NbrAttempts *int `json:"nbrAttempts,omitempty"`
		
		MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`
		*Alias
	}{ 
		NbrAttempts: o.NbrAttempts,
		
		MinutesBetweenAttempts: o.MinutesBetweenAttempts,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerattemptlimitsconfigchangerecallentry) UnmarshalJSON(b []byte) error {
	var DialerattemptlimitsconfigchangerecallentryMap map[string]interface{}
	err := json.Unmarshal(b, &DialerattemptlimitsconfigchangerecallentryMap)
	if err != nil {
		return err
	}
	
	if NbrAttempts, ok := DialerattemptlimitsconfigchangerecallentryMap["nbrAttempts"].(float64); ok {
		NbrAttemptsInt := int(NbrAttempts)
		o.NbrAttempts = &NbrAttemptsInt
	}
	
	if MinutesBetweenAttempts, ok := DialerattemptlimitsconfigchangerecallentryMap["minutesBetweenAttempts"].(float64); ok {
		MinutesBetweenAttemptsInt := int(MinutesBetweenAttempts)
		o.MinutesBetweenAttempts = &MinutesBetweenAttemptsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerattemptlimitsconfigchangerecallentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
