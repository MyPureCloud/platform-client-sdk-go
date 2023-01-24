package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Initialconfiguration
type Initialconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AudioState - Indicates the initial audio state for the communication.
	AudioState *Audiostate `json:"audioState,omitempty"`

	// Alerting - Indicates that this communication's initial state is alerting. If false, the communication started in a connected state.
	Alerting *bool `json:"alerting,omitempty"`

	// Inbound - Indicates the direction of this communication with respect to the contact center. `true` means the communication is INBOUND. `false` means the communication is OUTBOUND.
	Inbound *bool `json:"inbound,omitempty"`

	// InvitedBy - The id of the communication (the \"peer\") that \"invited\" this communication, if this occurred.
	InvitedBy *string `json:"invitedBy,omitempty"`

	// RecordingActive - Indicates whether recording is active for this communication at creation.
	RecordingActive *bool `json:"recordingActive,omitempty"`

	// AdditionalInfo - Additional metadata about this session which should be recorded by the platform but which will not be indexed or searchable. Primarily for diagnostic value. Any information that needs to be accessible through other components like Analytics should be moved to dedicated fields.
	AdditionalInfo *map[string]string `json:"additionalInfo,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Initialconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Initialconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Initialconfiguration
	
	return json.Marshal(&struct { 
		AudioState *Audiostate `json:"audioState,omitempty"`
		
		Alerting *bool `json:"alerting,omitempty"`
		
		Inbound *bool `json:"inbound,omitempty"`
		
		InvitedBy *string `json:"invitedBy,omitempty"`
		
		RecordingActive *bool `json:"recordingActive,omitempty"`
		
		AdditionalInfo *map[string]string `json:"additionalInfo,omitempty"`
		Alias
	}{ 
		AudioState: o.AudioState,
		
		Alerting: o.Alerting,
		
		Inbound: o.Inbound,
		
		InvitedBy: o.InvitedBy,
		
		RecordingActive: o.RecordingActive,
		
		AdditionalInfo: o.AdditionalInfo,
		Alias:    (Alias)(o),
	})
}

func (o *Initialconfiguration) UnmarshalJSON(b []byte) error {
	var InitialconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &InitialconfigurationMap)
	if err != nil {
		return err
	}
	
	if AudioState, ok := InitialconfigurationMap["audioState"].(map[string]interface{}); ok {
		AudioStateString, _ := json.Marshal(AudioState)
		json.Unmarshal(AudioStateString, &o.AudioState)
	}
	
	if Alerting, ok := InitialconfigurationMap["alerting"].(bool); ok {
		o.Alerting = &Alerting
	}
    
	if Inbound, ok := InitialconfigurationMap["inbound"].(bool); ok {
		o.Inbound = &Inbound
	}
    
	if InvitedBy, ok := InitialconfigurationMap["invitedBy"].(string); ok {
		o.InvitedBy = &InvitedBy
	}
    
	if RecordingActive, ok := InitialconfigurationMap["recordingActive"].(bool); ok {
		o.RecordingActive = &RecordingActive
	}
    
	if AdditionalInfo, ok := InitialconfigurationMap["additionalInfo"].(map[string]interface{}); ok {
		AdditionalInfoString, _ := json.Marshal(AdditionalInfo)
		json.Unmarshal(AdditionalInfoString, &o.AdditionalInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Initialconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
