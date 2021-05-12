package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Mergeoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
