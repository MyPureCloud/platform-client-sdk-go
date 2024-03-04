package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messaginginitialconfiguration
type Messaginginitialconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ToAddress - Address for the participant on receiving side of the message conversation. If the address is a phone number, E.164 format is recommended.
	ToAddress *string `json:"toAddress,omitempty"`

	// FromAddress - Address for the participant on the sending side of the message conversation. If the address is a phone number, E.164 format is recommended.
	FromAddress *string `json:"fromAddress,omitempty"`

	// MessageType - The type of message platform from which the message originated.
	MessageType *string `json:"messageType,omitempty"`

	// Held - Indicates that this communication's initial state is held.
	Held *bool `json:"held,omitempty"`

	// Alerting - Indicates that this communication's initial state is alerting. If false, the communication started in a connected state.
	Alerting *bool `json:"alerting,omitempty"`

	// Inbound - Indicates the direction of this communication with respect to the contact center. `true` means the communication is INBOUND. `false` means the communication is OUTBOUND.
	Inbound *bool `json:"inbound,omitempty"`

	// InvitedBy - The id of the communication (the \"peer\") that \"invited\" this communication, if this occurred.
	InvitedBy *string `json:"invitedBy,omitempty"`

	// AdditionalInfo - Additional metadata about this session which should be recorded by the platform but which will not be indexed or searchable. Primarily for diagnostic value. Any information that needs to be accessible through other components like Analytics should be moved to dedicated fields.
	AdditionalInfo *map[string]string `json:"additionalInfo,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messaginginitialconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Messaginginitialconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Messaginginitialconfiguration
	
	return json.Marshal(&struct { 
		ToAddress *string `json:"toAddress,omitempty"`
		
		FromAddress *string `json:"fromAddress,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		Alerting *bool `json:"alerting,omitempty"`
		
		Inbound *bool `json:"inbound,omitempty"`
		
		InvitedBy *string `json:"invitedBy,omitempty"`
		
		AdditionalInfo *map[string]string `json:"additionalInfo,omitempty"`
		Alias
	}{ 
		ToAddress: o.ToAddress,
		
		FromAddress: o.FromAddress,
		
		MessageType: o.MessageType,
		
		Held: o.Held,
		
		Alerting: o.Alerting,
		
		Inbound: o.Inbound,
		
		InvitedBy: o.InvitedBy,
		
		AdditionalInfo: o.AdditionalInfo,
		Alias:    (Alias)(o),
	})
}

func (o *Messaginginitialconfiguration) UnmarshalJSON(b []byte) error {
	var MessaginginitialconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &MessaginginitialconfigurationMap)
	if err != nil {
		return err
	}
	
	if ToAddress, ok := MessaginginitialconfigurationMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if FromAddress, ok := MessaginginitialconfigurationMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
    
	if MessageType, ok := MessaginginitialconfigurationMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if Held, ok := MessaginginitialconfigurationMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if Alerting, ok := MessaginginitialconfigurationMap["alerting"].(bool); ok {
		o.Alerting = &Alerting
	}
    
	if Inbound, ok := MessaginginitialconfigurationMap["inbound"].(bool); ok {
		o.Inbound = &Inbound
	}
    
	if InvitedBy, ok := MessaginginitialconfigurationMap["invitedBy"].(string); ok {
		o.InvitedBy = &InvitedBy
	}
    
	if AdditionalInfo, ok := MessaginginitialconfigurationMap["additionalInfo"].(map[string]interface{}); ok {
		AdditionalInfoString, _ := json.Marshal(AdditionalInfo)
		json.Unmarshal(AdditionalInfoString, &o.AdditionalInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messaginginitialconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
