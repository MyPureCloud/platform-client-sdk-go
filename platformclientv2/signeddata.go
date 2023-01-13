package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Signeddata
type Signeddata struct { 
	// Jwt
	Jwt *string `json:"jwt,omitempty"`

}

func (o *Signeddata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Signeddata
	
	return json.Marshal(&struct { 
		Jwt *string `json:"jwt,omitempty"`
		*Alias
	}{ 
		Jwt: o.Jwt,
		Alias:    (*Alias)(o),
	})
}

func (o *Signeddata) UnmarshalJSON(b []byte) error {
	var SigneddataMap map[string]interface{}
	err := json.Unmarshal(b, &SigneddataMap)
	if err != nil {
		return err
	}
	
	if Jwt, ok := SigneddataMap["jwt"].(string); ok {
		o.Jwt = &Jwt
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Signeddata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
