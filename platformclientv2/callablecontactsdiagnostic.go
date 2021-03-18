package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Callablecontactsdiagnostic
type Callablecontactsdiagnostic struct { 
	// AttemptLimits - Attempt limits for the campaign's contact list
	AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`


	// DncLists - Do not call lists for the campaign
	DncLists *[]Domainentityref `json:"dncLists,omitempty"`


	// CallableTimeSet - Callable time sets for the campaign
	CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`


	// RuleSets - Rule sets for the campaign
	RuleSets *[]Domainentityref `json:"ruleSets,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callablecontactsdiagnostic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
