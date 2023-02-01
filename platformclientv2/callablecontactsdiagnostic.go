package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Callablecontactsdiagnostic
type Callablecontactsdiagnostic struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AttemptLimits - Attempt limits for the campaign's contact list
	AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`

	// DncLists - Do not call lists for the campaign
	DncLists *[]Domainentityref `json:"dncLists,omitempty"`

	// CallableTimeSet - Callable time sets for the campaign
	CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`

	// RuleSets - Rule sets for the campaign
	RuleSets *[]Domainentityref `json:"ruleSets,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Callablecontactsdiagnostic) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Callablecontactsdiagnostic) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callablecontactsdiagnostic
	
	return json.Marshal(&struct { 
		AttemptLimits *Domainentityref `json:"attemptLimits,omitempty"`
		
		DncLists *[]Domainentityref `json:"dncLists,omitempty"`
		
		CallableTimeSet *Domainentityref `json:"callableTimeSet,omitempty"`
		
		RuleSets *[]Domainentityref `json:"ruleSets,omitempty"`
		Alias
	}{ 
		AttemptLimits: o.AttemptLimits,
		
		DncLists: o.DncLists,
		
		CallableTimeSet: o.CallableTimeSet,
		
		RuleSets: o.RuleSets,
		Alias:    (Alias)(o),
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
