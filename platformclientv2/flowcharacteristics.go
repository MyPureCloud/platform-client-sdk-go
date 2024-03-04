package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowcharacteristics - This is a set of enabled characteristics for the loglevel
type Flowcharacteristics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExecutionItems - Whether to report execution data about individual actions, menus, states, tasks, etc. etc. that ran during execution of the flow.
	ExecutionItems *bool `json:"executionItems,omitempty"`

	// ExecutionInputOutputs - Whether to report input setting input setting values and output data values for individual execution items above.  For example, if you have FlowExecutionInputOutputs and a Call Data Action ran in a flow, if FlowExecutionItems was enabled you'd see the fact a Call Data Action ran and the output path it took but nothing about which Data Action it ran, the input data sent to it at flow runtime and the data returned from it.  If you enable this characteristic, execution data will contain this additional detail.
	ExecutionInputOutputs *bool `json:"executionInputOutputs,omitempty"`

	// Communications - Communications are either audio or digital communications sent to or received from a participant.  An example here would be the initial greeting in an inbound call flow where it plays a greeting message to the participant.
	Communications *bool `json:"communications,omitempty"`

	// EventError - Whether to report flow error events.
	EventError *bool `json:"eventError,omitempty"`

	// EventWarning - Whether to report flow warning events.
	EventWarning *bool `json:"eventWarning,omitempty"`

	// EventOther - Whether to report events other than errors or warnings such as a language change, loop event.
	EventOther *bool `json:"eventOther,omitempty"`

	// Variables - Whether to report assignment of values to variables in flow execution data. It's important to remember there is a difference between variable value assignments and output data from an action.  If you have a Call Digital Bot flow action in an Inbound Message flow and there is no variable bound to the Exit Reason output but FlowExecutionInputOutputs is enabled, you will still be able to see the exit reason from the digital bot flow in execution data even though it is not bound to a variable.
	Variables *bool `json:"variables,omitempty"`

	// Names - This characteristic specifies whether or not name information should be emitted in execution data such as action, task, state or even the flow name itself.  Names are very handy from a readability standpoint but they do take up additional space in flow execution data instances.
	Names *bool `json:"names,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowcharacteristics) SetField(field string, fieldValue interface{}) {
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

func (o Flowcharacteristics) MarshalJSON() ([]byte, error) {
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
	type Alias Flowcharacteristics
	
	return json.Marshal(&struct { 
		ExecutionItems *bool `json:"executionItems,omitempty"`
		
		ExecutionInputOutputs *bool `json:"executionInputOutputs,omitempty"`
		
		Communications *bool `json:"communications,omitempty"`
		
		EventError *bool `json:"eventError,omitempty"`
		
		EventWarning *bool `json:"eventWarning,omitempty"`
		
		EventOther *bool `json:"eventOther,omitempty"`
		
		Variables *bool `json:"variables,omitempty"`
		
		Names *bool `json:"names,omitempty"`
		Alias
	}{ 
		ExecutionItems: o.ExecutionItems,
		
		ExecutionInputOutputs: o.ExecutionInputOutputs,
		
		Communications: o.Communications,
		
		EventError: o.EventError,
		
		EventWarning: o.EventWarning,
		
		EventOther: o.EventOther,
		
		Variables: o.Variables,
		
		Names: o.Names,
		Alias:    (Alias)(o),
	})
}

func (o *Flowcharacteristics) UnmarshalJSON(b []byte) error {
	var FlowcharacteristicsMap map[string]interface{}
	err := json.Unmarshal(b, &FlowcharacteristicsMap)
	if err != nil {
		return err
	}
	
	if ExecutionItems, ok := FlowcharacteristicsMap["executionItems"].(bool); ok {
		o.ExecutionItems = &ExecutionItems
	}
    
	if ExecutionInputOutputs, ok := FlowcharacteristicsMap["executionInputOutputs"].(bool); ok {
		o.ExecutionInputOutputs = &ExecutionInputOutputs
	}
    
	if Communications, ok := FlowcharacteristicsMap["communications"].(bool); ok {
		o.Communications = &Communications
	}
    
	if EventError, ok := FlowcharacteristicsMap["eventError"].(bool); ok {
		o.EventError = &EventError
	}
    
	if EventWarning, ok := FlowcharacteristicsMap["eventWarning"].(bool); ok {
		o.EventWarning = &EventWarning
	}
    
	if EventOther, ok := FlowcharacteristicsMap["eventOther"].(bool); ok {
		o.EventOther = &EventOther
	}
    
	if Variables, ok := FlowcharacteristicsMap["variables"].(bool); ok {
		o.Variables = &Variables
	}
    
	if Names, ok := FlowcharacteristicsMap["names"].(bool); ok {
		o.Names = &Names
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowcharacteristics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
