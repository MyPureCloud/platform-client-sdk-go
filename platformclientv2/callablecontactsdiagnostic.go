package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Callablecontactsdiagnostic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callablecontactsdiagnostic

	

	return json.Marshal(&struct { 
		AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`
		
		DncLists *[]Domainentityref `json:"dncLists,omitempty"`
		
		CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`
		
		RuleSets *[]Domainentityref `json:"ruleSets,omitempty"`
		*Alias
	}{ 
		AttemptLimits: u.AttemptLimits,
		
		DncLists: u.DncLists,
		
		CallableTimeSet: u.CallableTimeSet,
		
		RuleSets: u.RuleSets,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callablecontactsdiagnostic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
