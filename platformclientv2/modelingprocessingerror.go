package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Modelingprocessingerror
type Modelingprocessingerror struct { 
	// InternalErrorCode - An internal code representing the type of error. ModelInputMissing for 'Model Builder inputs not found.' ModelInputInvalid for 'Model Builder inputs are invalid. Ensure the input data format is correct.' ModelFailed for 'An error occured while building the model with the given input.'
	InternalErrorCode *string `json:"internalErrorCode,omitempty"`


	// Description - A text description of the error
	Description *string `json:"description,omitempty"`

}

func (o *Modelingprocessingerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Modelingprocessingerror
	
	return json.Marshal(&struct { 
		InternalErrorCode *string `json:"internalErrorCode,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		InternalErrorCode: o.InternalErrorCode,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Modelingprocessingerror) UnmarshalJSON(b []byte) error {
	var ModelingprocessingerrorMap map[string]interface{}
	err := json.Unmarshal(b, &ModelingprocessingerrorMap)
	if err != nil {
		return err
	}
	
	if InternalErrorCode, ok := ModelingprocessingerrorMap["internalErrorCode"].(string); ok {
		o.InternalErrorCode = &InternalErrorCode
	}
    
	if Description, ok := ModelingprocessingerrorMap["description"].(string); ok {
		o.Description = &Description
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Modelingprocessingerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
