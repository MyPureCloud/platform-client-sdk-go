package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Directroutingcallmediasettings
type Directroutingcallmediasettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Toggle that enables Direct Routing for this media type.
	Enabled *bool `json:"enabled,omitempty"`

	// InboundFlow - The Direct Routing inbound flow id for this media type.
	InboundFlow *Addressableentityref `json:"inboundFlow,omitempty"`

	// VoicemailFlow - ID of the in-queue flow that handles voicemails for Direct Routing and to transfer calls to ACD voicemail.
	VoicemailFlow *Addressableentityref `json:"voicemailFlow,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Directroutingcallmediasettings) SetField(field string, fieldValue interface{}) {
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

func (o Directroutingcallmediasettings) MarshalJSON() ([]byte, error) {
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
	type Alias Directroutingcallmediasettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		InboundFlow *Addressableentityref `json:"inboundFlow,omitempty"`
		
		VoicemailFlow *Addressableentityref `json:"voicemailFlow,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		InboundFlow: o.InboundFlow,
		
		VoicemailFlow: o.VoicemailFlow,
		Alias:    (Alias)(o),
	})
}

func (o *Directroutingcallmediasettings) UnmarshalJSON(b []byte) error {
	var DirectroutingcallmediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &DirectroutingcallmediasettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := DirectroutingcallmediasettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InboundFlow, ok := DirectroutingcallmediasettingsMap["inboundFlow"].(map[string]interface{}); ok {
		InboundFlowString, _ := json.Marshal(InboundFlow)
		json.Unmarshal(InboundFlowString, &o.InboundFlow)
	}
	
	if VoicemailFlow, ok := DirectroutingcallmediasettingsMap["voicemailFlow"].(map[string]interface{}); ok {
		VoicemailFlowString, _ := json.Marshal(VoicemailFlow)
		json.Unmarshal(VoicemailFlowString, &o.VoicemailFlow)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Directroutingcallmediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
