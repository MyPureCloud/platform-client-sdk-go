package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Teamaddmemberfailure
type Teamaddmemberfailure struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Reason - Specific reason the member could not be added.
	Reason *string `json:"reason,omitempty"`

}

func (o *Teamaddmemberfailure) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Teamaddmemberfailure
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Reason *string `json:"reason,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Reason: o.Reason,
		Alias:    (*Alias)(o),
	})
}

func (o *Teamaddmemberfailure) UnmarshalJSON(b []byte) error {
	var TeamaddmemberfailureMap map[string]interface{}
	err := json.Unmarshal(b, &TeamaddmemberfailureMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TeamaddmemberfailureMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Reason, ok := TeamaddmemberfailureMap["reason"].(string); ok {
		o.Reason = &Reason
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Teamaddmemberfailure) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
