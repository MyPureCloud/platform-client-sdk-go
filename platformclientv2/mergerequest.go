package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mergerequest
type Mergerequest struct { 
	// SourceContactId - The ID of the source contact for the merge operation
	SourceContactId *string `json:"sourceContactId,omitempty"`


	// TargetContactId - The ID of the target contact for the merge operation
	TargetContactId *string `json:"targetContactId,omitempty"`

}

func (o *Mergerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mergerequest
	
	return json.Marshal(&struct { 
		SourceContactId *string `json:"sourceContactId,omitempty"`
		
		TargetContactId *string `json:"targetContactId,omitempty"`
		*Alias
	}{ 
		SourceContactId: o.SourceContactId,
		
		TargetContactId: o.TargetContactId,
		Alias:    (*Alias)(o),
	})
}

func (o *Mergerequest) UnmarshalJSON(b []byte) error {
	var MergerequestMap map[string]interface{}
	err := json.Unmarshal(b, &MergerequestMap)
	if err != nil {
		return err
	}
	
	if SourceContactId, ok := MergerequestMap["sourceContactId"].(string); ok {
		o.SourceContactId = &SourceContactId
	}
    
	if TargetContactId, ok := MergerequestMap["targetContactId"].(string); ok {
		o.TargetContactId = &TargetContactId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mergerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
