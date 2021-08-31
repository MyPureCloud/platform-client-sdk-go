package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingtemplaterequest
type Messagingtemplaterequest struct { 
	// ResponseId - A Response Management response identifier for a messaging template defined response
	ResponseId *string `json:"responseId,omitempty"`


	// Parameters - A list of Response Management response substitutions for the response's messaging template
	Parameters *[]Templateparameter `json:"parameters,omitempty"`

}

func (o *Messagingtemplaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingtemplaterequest
	
	return json.Marshal(&struct { 
		ResponseId *string `json:"responseId,omitempty"`
		
		Parameters *[]Templateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		ResponseId: o.ResponseId,
		
		Parameters: o.Parameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagingtemplaterequest) UnmarshalJSON(b []byte) error {
	var MessagingtemplaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingtemplaterequestMap)
	if err != nil {
		return err
	}
	
	if ResponseId, ok := MessagingtemplaterequestMap["responseId"].(string); ok {
		o.ResponseId = &ResponseId
	}
	
	if Parameters, ok := MessagingtemplaterequestMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingtemplaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
