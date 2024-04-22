package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicconversation
type Conversationeventtopicconversation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// MaxParticipants
	MaxParticipants *int `json:"maxParticipants,omitempty"`

	// Participants
	Participants *[]Conversationeventtopicparticipant `json:"participants,omitempty"`

	// RecentTransfers
	RecentTransfers *[]Conversationeventtopicrecenttransfer `json:"recentTransfers,omitempty"`

	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`

	// Address
	Address *string `json:"address,omitempty"`

	// ExternalTag
	ExternalTag *string `json:"externalTag,omitempty"`

	// UtilizationLabelId
	UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`

	// SecurePause
	SecurePause *bool `json:"securePause,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationeventtopicconversation) SetField(field string, fieldValue interface{}) {
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

func (o Conversationeventtopicconversation) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationeventtopicconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		
		Participants *[]Conversationeventtopicparticipant `json:"participants,omitempty"`
		
		RecentTransfers *[]Conversationeventtopicrecenttransfer `json:"recentTransfers,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		UtilizationLabelId *string `json:"utilizationLabelId,omitempty"`
		
		SecurePause *bool `json:"securePause,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		MaxParticipants: o.MaxParticipants,
		
		Participants: o.Participants,
		
		RecentTransfers: o.RecentTransfers,
		
		RecordingState: o.RecordingState,
		
		Address: o.Address,
		
		ExternalTag: o.ExternalTag,
		
		UtilizationLabelId: o.UtilizationLabelId,
		
		SecurePause: o.SecurePause,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationeventtopicconversation) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicconversationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicconversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MaxParticipants, ok := ConversationeventtopicconversationMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	
	if Participants, ok := ConversationeventtopicconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if RecentTransfers, ok := ConversationeventtopicconversationMap["recentTransfers"].([]interface{}); ok {
		RecentTransfersString, _ := json.Marshal(RecentTransfers)
		json.Unmarshal(RecentTransfersString, &o.RecentTransfers)
	}
	
	if RecordingState, ok := ConversationeventtopicconversationMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
    
	if Address, ok := ConversationeventtopicconversationMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if ExternalTag, ok := ConversationeventtopicconversationMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    
	if UtilizationLabelId, ok := ConversationeventtopicconversationMap["utilizationLabelId"].(string); ok {
		o.UtilizationLabelId = &UtilizationLabelId
	}
    
	if SecurePause, ok := ConversationeventtopicconversationMap["securePause"].(bool); ok {
		o.SecurePause = &SecurePause
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
