package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Adjustablelivespeakerdetection
type Adjustablelivespeakerdetection struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Mode - Modes to tune between speed to live speaker detection vs accuracy.
	Mode *string `json:"mode,omitempty"`

	// PreconnectDuration - ISO 8601 formatted relative duration (e.g., PT30.8427419S for 30.8 seconds), calculated on line connect.
	PreconnectDuration *string `json:"preconnectDuration,omitempty"`

	// EventName - The name of the event that triggered the ALSD evaluation (e.g., line.connect, speech.generic).
	EventName *string `json:"eventName,omitempty"`

	// IsPersonLikely - The output of the ALSD detector, evaluating whether there is likely a person on the call based on the above inputs, and if so, a person is detected early (person disposition name and speech.person analyzer result) and the associated action taken (e.g., speech.person postconnect entry in the disposition table has the action to transfer to a queue).
	IsPersonLikely *bool `json:"isPersonLikely,omitempty"`

	// TotalRingbacks - Number of tone.ring.* analyzer events detected during the call (expected mostly during pre-connect but the last ringback tone detection could potentially complete after line connect, which will increment totalRingbacks still).
	TotalRingbacks *int `json:"totalRingbacks,omitempty"`

	// LineConnected - Protocol line connect received (answered by a person, machine, busy, fax).
	LineConnected *bool `json:"lineConnected,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Adjustablelivespeakerdetection) SetField(field string, fieldValue interface{}) {
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

func (o Adjustablelivespeakerdetection) MarshalJSON() ([]byte, error) {
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
	type Alias Adjustablelivespeakerdetection
	
	return json.Marshal(&struct { 
		Mode *string `json:"mode,omitempty"`
		
		PreconnectDuration *string `json:"preconnectDuration,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		
		IsPersonLikely *bool `json:"isPersonLikely,omitempty"`
		
		TotalRingbacks *int `json:"totalRingbacks,omitempty"`
		
		LineConnected *bool `json:"lineConnected,omitempty"`
		Alias
	}{ 
		Mode: o.Mode,
		
		PreconnectDuration: o.PreconnectDuration,
		
		EventName: o.EventName,
		
		IsPersonLikely: o.IsPersonLikely,
		
		TotalRingbacks: o.TotalRingbacks,
		
		LineConnected: o.LineConnected,
		Alias:    (Alias)(o),
	})
}

func (o *Adjustablelivespeakerdetection) UnmarshalJSON(b []byte) error {
	var AdjustablelivespeakerdetectionMap map[string]interface{}
	err := json.Unmarshal(b, &AdjustablelivespeakerdetectionMap)
	if err != nil {
		return err
	}
	
	if Mode, ok := AdjustablelivespeakerdetectionMap["mode"].(string); ok {
		o.Mode = &Mode
	}
    
	if PreconnectDuration, ok := AdjustablelivespeakerdetectionMap["preconnectDuration"].(string); ok {
		o.PreconnectDuration = &PreconnectDuration
	}
    
	if EventName, ok := AdjustablelivespeakerdetectionMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if IsPersonLikely, ok := AdjustablelivespeakerdetectionMap["isPersonLikely"].(bool); ok {
		o.IsPersonLikely = &IsPersonLikely
	}
    
	if TotalRingbacks, ok := AdjustablelivespeakerdetectionMap["totalRingbacks"].(float64); ok {
		TotalRingbacksInt := int(TotalRingbacks)
		o.TotalRingbacks = &TotalRingbacksInt
	}
	
	if LineConnected, ok := AdjustablelivespeakerdetectionMap["lineConnected"].(bool); ok {
		o.LineConnected = &LineConnected
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adjustablelivespeakerdetection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
