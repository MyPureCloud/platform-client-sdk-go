package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mergeoperation
type Mergeoperation struct { 
	// SourceContact - The source contact for the merge operation
	SourceContact *Addressableentityref `json:"sourceContact,omitempty"`


	// TargetContact - The target contact for the merge operation
	TargetContact *Addressableentityref `json:"targetContact,omitempty"`


	// ResultingContact - The contact created as a result of the merge operation
	ResultingContact *Addressableentityref `json:"resultingContact,omitempty"`

}

func (o *Mergeoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mergeoperation
	
	return json.Marshal(&struct { 
		SourceContact *Addressableentityref `json:"sourceContact,omitempty"`
		
		TargetContact *Addressableentityref `json:"targetContact,omitempty"`
		
		ResultingContact *Addressableentityref `json:"resultingContact,omitempty"`
		*Alias
	}{ 
		SourceContact: o.SourceContact,
		
		TargetContact: o.TargetContact,
		
		ResultingContact: o.ResultingContact,
		Alias:    (*Alias)(o),
	})
}

func (o *Mergeoperation) UnmarshalJSON(b []byte) error {
	var MergeoperationMap map[string]interface{}
	err := json.Unmarshal(b, &MergeoperationMap)
	if err != nil {
		return err
	}
	
	if SourceContact, ok := MergeoperationMap["sourceContact"].(map[string]interface{}); ok {
		SourceContactString, _ := json.Marshal(SourceContact)
		json.Unmarshal(SourceContactString, &o.SourceContact)
	}
	
	if TargetContact, ok := MergeoperationMap["targetContact"].(map[string]interface{}); ok {
		TargetContactString, _ := json.Marshal(TargetContact)
		json.Unmarshal(TargetContactString, &o.TargetContact)
	}
	
	if ResultingContact, ok := MergeoperationMap["resultingContact"].(map[string]interface{}); ok {
		ResultingContactString, _ := json.Marshal(ResultingContact)
		json.Unmarshal(ResultingContactString, &o.ResultingContact)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mergeoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
