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

func (o *Callablecontactsdiagnostic) MarshalJSON() ([]byte, error) {
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
		AttemptLimits: o.AttemptLimits,
		
		DncLists: o.DncLists,
		
		CallableTimeSet: o.CallableTimeSet,
		
		RuleSets: o.RuleSets,
		Alias:    (*Alias)(o),
	})
}

func (o *Callablecontactsdiagnostic) UnmarshalJSON(b []byte) error {
	var CallablecontactsdiagnosticMap map[string]interface{}
	err := json.Unmarshal(b, &CallablecontactsdiagnosticMap)
	if err != nil {
		return err
	}
	
	if AttemptLimits, ok := CallablecontactsdiagnosticMap["attemptLimits"].(map[string]interface{}); ok {
		AttemptLimitsString, _ := json.Marshal(AttemptLimits)
		json.Unmarshal(AttemptLimitsString, &o.AttemptLimits)
	}
	
	if DncLists, ok := CallablecontactsdiagnosticMap["dncLists"].([]interface{}); ok {
		DncListsString, _ := json.Marshal(DncLists)
		json.Unmarshal(DncListsString, &o.DncLists)
	}
	
	if CallableTimeSet, ok := CallablecontactsdiagnosticMap["callableTimeSet"].(map[string]interface{}); ok {
		CallableTimeSetString, _ := json.Marshal(CallableTimeSet)
		json.Unmarshal(CallableTimeSetString, &o.CallableTimeSet)
	}
	
	if RuleSets, ok := CallablecontactsdiagnosticMap["ruleSets"].([]interface{}); ok {
		RuleSetsString, _ := json.Marshal(RuleSets)
		json.Unmarshal(RuleSetsString, &o.RuleSets)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callablecontactsdiagnostic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
