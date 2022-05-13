package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailsetup
type Emailsetup struct { 
	// RootDomain - The root PureCloud domain that all sub-domains are created from.
	RootDomain *string `json:"rootDomain,omitempty"`

}

func (o *Emailsetup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailsetup
	
	return json.Marshal(&struct { 
		RootDomain *string `json:"rootDomain,omitempty"`
		*Alias
	}{ 
		RootDomain: o.RootDomain,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailsetup) UnmarshalJSON(b []byte) error {
	var EmailsetupMap map[string]interface{}
	err := json.Unmarshal(b, &EmailsetupMap)
	if err != nil {
		return err
	}
	
	if RootDomain, ok := EmailsetupMap["rootDomain"].(string); ok {
		o.RootDomain = &RootDomain
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Emailsetup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
