package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchuser
type Patchuser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`

}

func (o *Patchuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		AcdAutoAnswer: o.AcdAutoAnswer,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchuser) UnmarshalJSON(b []byte) error {
	var PatchuserMap map[string]interface{}
	err := json.Unmarshal(b, &PatchuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PatchuserMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if AcdAutoAnswer, ok := PatchuserMap["acdAutoAnswer"].(bool); ok {
		o.AcdAutoAnswer = &AcdAutoAnswer
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
