package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectionrequest
type Nludetectionrequest struct { 
	// Input - The input subject to NLU detection.
	Input *Nludetectioninput `json:"input,omitempty"`


	// Context - The context for the input to NLU detection.
	Context *Nludetectioncontext `json:"context,omitempty"`

}

func (o *Nludetectionrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectionrequest
	
	return json.Marshal(&struct { 
		Input *Nludetectioninput `json:"input,omitempty"`
		
		Context *Nludetectioncontext `json:"context,omitempty"`
		*Alias
	}{ 
		Input: o.Input,
		
		Context: o.Context,
		Alias:    (*Alias)(o),
	})
}

func (o *Nludetectionrequest) UnmarshalJSON(b []byte) error {
	var NludetectionrequestMap map[string]interface{}
	err := json.Unmarshal(b, &NludetectionrequestMap)
	if err != nil {
		return err
	}
	
	if Input, ok := NludetectionrequestMap["input"].(map[string]interface{}); ok {
		InputString, _ := json.Marshal(Input)
		json.Unmarshal(InputString, &o.Input)
	}
	
	if Context, ok := NludetectionrequestMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nludetectionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
