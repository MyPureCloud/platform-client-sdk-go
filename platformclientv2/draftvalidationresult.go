package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Draftvalidationresult - Validation results
type Draftvalidationresult struct { 
	// Valid - Indicates if configuration is valid
	Valid *bool `json:"valid,omitempty"`


	// Errors - List of errors causing validation failure
	Errors *[]Errorbody `json:"errors,omitempty"`

}

func (o *Draftvalidationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Draftvalidationresult
	
	return json.Marshal(&struct { 
		Valid *bool `json:"valid,omitempty"`
		
		Errors *[]Errorbody `json:"errors,omitempty"`
		*Alias
	}{ 
		Valid: o.Valid,
		
		Errors: o.Errors,
		Alias:    (*Alias)(o),
	})
}

func (o *Draftvalidationresult) UnmarshalJSON(b []byte) error {
	var DraftvalidationresultMap map[string]interface{}
	err := json.Unmarshal(b, &DraftvalidationresultMap)
	if err != nil {
		return err
	}
	
	if Valid, ok := DraftvalidationresultMap["valid"].(bool); ok {
		o.Valid = &Valid
	}
    
	if Errors, ok := DraftvalidationresultMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Draftvalidationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
