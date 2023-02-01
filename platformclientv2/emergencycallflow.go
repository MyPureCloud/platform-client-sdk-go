package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Emergencycallflow - An emergency flow associates a call flow to use in an emergency with the ivr(s) to route to it.
type Emergencycallflow struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EmergencyFlow - The call flow to execute in an emergency.
	EmergencyFlow *Domainentityref `json:"emergencyFlow,omitempty"`

	// Ivrs - The IVR(s) to route to the call flow during an emergency.
	Ivrs *[]Domainentityref `json:"ivrs,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Emergencycallflow) SetField(field string, fieldValue interface{}) {
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

func (o Emergencycallflow) MarshalJSON() ([]byte, error) {
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
	type Alias Emergencycallflow
	
	return json.Marshal(&struct { 
		EmergencyFlow *Domainentityref `json:"emergencyFlow,omitempty"`
		
		Ivrs *[]Domainentityref `json:"ivrs,omitempty"`
		Alias
	}{ 
		EmergencyFlow: o.EmergencyFlow,
		
		Ivrs: o.Ivrs,
		Alias:    (Alias)(o),
	})
}

func (o *Emergencycallflow) UnmarshalJSON(b []byte) error {
	var EmergencycallflowMap map[string]interface{}
	err := json.Unmarshal(b, &EmergencycallflowMap)
	if err != nil {
		return err
	}
	
	if EmergencyFlow, ok := EmergencycallflowMap["emergencyFlow"].(map[string]interface{}); ok {
		EmergencyFlowString, _ := json.Marshal(EmergencyFlow)
		json.Unmarshal(EmergencyFlowString, &o.EmergencyFlow)
	}
	
	if Ivrs, ok := EmergencycallflowMap["ivrs"].([]interface{}); ok {
		IvrsString, _ := json.Marshal(Ivrs)
		json.Unmarshal(IvrsString, &o.Ivrs)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emergencycallflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
