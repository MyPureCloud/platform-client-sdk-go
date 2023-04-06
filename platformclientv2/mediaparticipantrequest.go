package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediaparticipantrequest
type Mediaparticipantrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Wrapup - Wrap-up to assign to this participant.
	Wrapup *Wrapupinput `json:"wrapup,omitempty"`

	// State - The state to update to set for this participant's communications.  Possible values are: 'connected' and 'disconnected'.
	State *string `json:"state,omitempty"`

	// Recording - True to enable recording of this participant, otherwise false to disable recording.
	Recording *bool `json:"recording,omitempty"`

	// Muted - True to mute this conversation participant.
	Muted *bool `json:"muted,omitempty"`

	// Confined - True to confine this conversation participant.  Should only be used for ad-hoc conferences
	Confined *bool `json:"confined,omitempty"`

	// Held - True to hold this conversation participant.
	Held *bool `json:"held,omitempty"`

	// WrapupSkipped - True to skip wrap-up for this participant.
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Mediaparticipantrequest) SetField(field string, fieldValue interface{}) {
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

func (o Mediaparticipantrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Mediaparticipantrequest
	
	return json.Marshal(&struct { 
		Wrapup *Wrapupinput `json:"wrapup,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		Alias
	}{ 
		Wrapup: o.Wrapup,
		
		State: o.State,
		
		Recording: o.Recording,
		
		Muted: o.Muted,
		
		Confined: o.Confined,
		
		Held: o.Held,
		
		WrapupSkipped: o.WrapupSkipped,
		Alias:    (Alias)(o),
	})
}

func (o *Mediaparticipantrequest) UnmarshalJSON(b []byte) error {
	var MediaparticipantrequestMap map[string]interface{}
	err := json.Unmarshal(b, &MediaparticipantrequestMap)
	if err != nil {
		return err
	}
	
	if Wrapup, ok := MediaparticipantrequestMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if State, ok := MediaparticipantrequestMap["state"].(string); ok {
		o.State = &State
	}
    
	if Recording, ok := MediaparticipantrequestMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
    
	if Muted, ok := MediaparticipantrequestMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if Confined, ok := MediaparticipantrequestMap["confined"].(bool); ok {
		o.Confined = &Confined
	}
    
	if Held, ok := MediaparticipantrequestMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupSkipped, ok := MediaparticipantrequestMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mediaparticipantrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
