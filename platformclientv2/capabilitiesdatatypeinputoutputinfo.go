package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesdatatypeinputoutputinfo
type Capabilitiesdatatypeinputoutputinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// InputAny
	InputAny *map[string]Capabilitiesavailability `json:"inputAny,omitempty"`

	// InputColl
	InputColl *map[string]Capabilitiesavailability `json:"inputColl,omitempty"`

	// InputOne
	InputOne *map[string]Capabilitiesavailability `json:"inputOne,omitempty"`

	// InputOutputAny
	InputOutputAny *map[string]Capabilitiesavailability `json:"inputOutputAny,omitempty"`

	// InputOutputColl
	InputOutputColl *map[string]Capabilitiesavailability `json:"inputOutputColl,omitempty"`

	// InputOutputOne
	InputOutputOne *map[string]Capabilitiesavailability `json:"inputOutputOne,omitempty"`

	// OutputAny
	OutputAny *map[string]Capabilitiesavailability `json:"outputAny,omitempty"`

	// OutputColl
	OutputColl *map[string]Capabilitiesavailability `json:"outputColl,omitempty"`

	// OutputOne
	OutputOne *map[string]Capabilitiesavailability `json:"outputOne,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesdatatypeinputoutputinfo) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesdatatypeinputoutputinfo) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesdatatypeinputoutputinfo
	
	return json.Marshal(&struct { 
		InputAny *map[string]Capabilitiesavailability `json:"inputAny,omitempty"`
		
		InputColl *map[string]Capabilitiesavailability `json:"inputColl,omitempty"`
		
		InputOne *map[string]Capabilitiesavailability `json:"inputOne,omitempty"`
		
		InputOutputAny *map[string]Capabilitiesavailability `json:"inputOutputAny,omitempty"`
		
		InputOutputColl *map[string]Capabilitiesavailability `json:"inputOutputColl,omitempty"`
		
		InputOutputOne *map[string]Capabilitiesavailability `json:"inputOutputOne,omitempty"`
		
		OutputAny *map[string]Capabilitiesavailability `json:"outputAny,omitempty"`
		
		OutputColl *map[string]Capabilitiesavailability `json:"outputColl,omitempty"`
		
		OutputOne *map[string]Capabilitiesavailability `json:"outputOne,omitempty"`
		Alias
	}{ 
		InputAny: o.InputAny,
		
		InputColl: o.InputColl,
		
		InputOne: o.InputOne,
		
		InputOutputAny: o.InputOutputAny,
		
		InputOutputColl: o.InputOutputColl,
		
		InputOutputOne: o.InputOutputOne,
		
		OutputAny: o.OutputAny,
		
		OutputColl: o.OutputColl,
		
		OutputOne: o.OutputOne,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesdatatypeinputoutputinfo) UnmarshalJSON(b []byte) error {
	var CapabilitiesdatatypeinputoutputinfoMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesdatatypeinputoutputinfoMap)
	if err != nil {
		return err
	}
	
	if InputAny, ok := CapabilitiesdatatypeinputoutputinfoMap["inputAny"].(map[string]interface{}); ok {
		InputAnyString, _ := json.Marshal(InputAny)
		json.Unmarshal(InputAnyString, &o.InputAny)
	}
	
	if InputColl, ok := CapabilitiesdatatypeinputoutputinfoMap["inputColl"].(map[string]interface{}); ok {
		InputCollString, _ := json.Marshal(InputColl)
		json.Unmarshal(InputCollString, &o.InputColl)
	}
	
	if InputOne, ok := CapabilitiesdatatypeinputoutputinfoMap["inputOne"].(map[string]interface{}); ok {
		InputOneString, _ := json.Marshal(InputOne)
		json.Unmarshal(InputOneString, &o.InputOne)
	}
	
	if InputOutputAny, ok := CapabilitiesdatatypeinputoutputinfoMap["inputOutputAny"].(map[string]interface{}); ok {
		InputOutputAnyString, _ := json.Marshal(InputOutputAny)
		json.Unmarshal(InputOutputAnyString, &o.InputOutputAny)
	}
	
	if InputOutputColl, ok := CapabilitiesdatatypeinputoutputinfoMap["inputOutputColl"].(map[string]interface{}); ok {
		InputOutputCollString, _ := json.Marshal(InputOutputColl)
		json.Unmarshal(InputOutputCollString, &o.InputOutputColl)
	}
	
	if InputOutputOne, ok := CapabilitiesdatatypeinputoutputinfoMap["inputOutputOne"].(map[string]interface{}); ok {
		InputOutputOneString, _ := json.Marshal(InputOutputOne)
		json.Unmarshal(InputOutputOneString, &o.InputOutputOne)
	}
	
	if OutputAny, ok := CapabilitiesdatatypeinputoutputinfoMap["outputAny"].(map[string]interface{}); ok {
		OutputAnyString, _ := json.Marshal(OutputAny)
		json.Unmarshal(OutputAnyString, &o.OutputAny)
	}
	
	if OutputColl, ok := CapabilitiesdatatypeinputoutputinfoMap["outputColl"].(map[string]interface{}); ok {
		OutputCollString, _ := json.Marshal(OutputColl)
		json.Unmarshal(OutputCollString, &o.OutputColl)
	}
	
	if OutputOne, ok := CapabilitiesdatatypeinputoutputinfoMap["outputOne"].(map[string]interface{}); ok {
		OutputOneString, _ := json.Marshal(OutputOne)
		json.Unmarshal(OutputOneString, &o.OutputOne)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypeinputoutputinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
