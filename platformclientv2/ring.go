package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Ring
type Ring struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExpansionCriteria - The conditions that will trigger conversations to move to the next bullseye ring.
	ExpansionCriteria *[]Expansioncriterium `json:"expansionCriteria,omitempty"`

	// Actions - The actions that will be performed just before moving conversations to the next bullseye ring.
	Actions *Actions `json:"actions,omitempty"`

	// MemberGroups - The groups of agents associated with the ring, if any.  Ring membership will update to match group membership changes.
	MemberGroups *[]Membergroup `json:"memberGroups,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Ring) SetField(field string, fieldValue interface{}) {
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

func (o Ring) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Ring
	
	return json.Marshal(&struct { 
		ExpansionCriteria *[]Expansioncriterium `json:"expansionCriteria,omitempty"`
		
		Actions *Actions `json:"actions,omitempty"`
		
		MemberGroups *[]Membergroup `json:"memberGroups,omitempty"`
		Alias
	}{ 
		ExpansionCriteria: o.ExpansionCriteria,
		
		Actions: o.Actions,
		
		MemberGroups: o.MemberGroups,
		Alias:    (Alias)(o),
	})
}

func (o *Ring) UnmarshalJSON(b []byte) error {
	var RingMap map[string]interface{}
	err := json.Unmarshal(b, &RingMap)
	if err != nil {
		return err
	}
	
	if ExpansionCriteria, ok := RingMap["expansionCriteria"].([]interface{}); ok {
		ExpansionCriteriaString, _ := json.Marshal(ExpansionCriteria)
		json.Unmarshal(ExpansionCriteriaString, &o.ExpansionCriteria)
	}
	
	if Actions, ok := RingMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if MemberGroups, ok := RingMap["memberGroups"].([]interface{}); ok {
		MemberGroupsString, _ := json.Marshal(MemberGroups)
		json.Unmarshal(MemberGroupsString, &o.MemberGroups)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
