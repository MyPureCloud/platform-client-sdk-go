package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailinitialconfiguration
type Emailinitialconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// To - An email address that this email is to.
	To *string `json:"to,omitempty"`

	// From - An email address that this email is from.
	From *string `json:"from,omitempty"`

	// Cc - An email addresses that this email is carbon copied to.
	Cc *[]string `json:"cc,omitempty"`

	// Bcc - An email addresses that this email is blind carbon copied to.
	Bcc *[]string `json:"bcc,omitempty"`

	// Subject - The subject for this email.
	Subject *string `json:"subject,omitempty"`

	// PreviousEmailId - UUID identifying the most recent previous email communication ID from the same participant on this email conversation. Will be null if this is a new participant.
	PreviousEmailId *string `json:"previousEmailId,omitempty"`

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
func (o *Emailinitialconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Emailinitialconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Emailinitialconfiguration
	
	return json.Marshal(&struct { 
		To *string `json:"to,omitempty"`
		
		From *string `json:"from,omitempty"`
		
		Cc *[]string `json:"cc,omitempty"`
		
		Bcc *[]string `json:"bcc,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		PreviousEmailId *string `json:"previousEmailId,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		Alerting *bool `json:"alerting,omitempty"`
		
		Inbound *bool `json:"inbound,omitempty"`
		
		InvitedBy *string `json:"invitedBy,omitempty"`
		
		AdditionalInfo *map[string]string `json:"additionalInfo,omitempty"`
		Alias
	}{ 
		To: o.To,
		
		From: o.From,
		
		Cc: o.Cc,
		
		Bcc: o.Bcc,
		
		Subject: o.Subject,
		
		PreviousEmailId: o.PreviousEmailId,
		
		Held: o.Held,
		
		Alerting: o.Alerting,
		
		Inbound: o.Inbound,
		
		InvitedBy: o.InvitedBy,
		
		AdditionalInfo: o.AdditionalInfo,
		Alias:    (Alias)(o),
	})
}

func (o *Emailinitialconfiguration) UnmarshalJSON(b []byte) error {
	var EmailinitialconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &EmailinitialconfigurationMap)
	if err != nil {
		return err
	}
	
	if To, ok := EmailinitialconfigurationMap["to"].(string); ok {
		o.To = &To
	}
    
	if From, ok := EmailinitialconfigurationMap["from"].(string); ok {
		o.From = &From
	}
    
	if Cc, ok := EmailinitialconfigurationMap["cc"].([]interface{}); ok {
		CcString, _ := json.Marshal(Cc)
		json.Unmarshal(CcString, &o.Cc)
	}
	
	if Bcc, ok := EmailinitialconfigurationMap["bcc"].([]interface{}); ok {
		BccString, _ := json.Marshal(Bcc)
		json.Unmarshal(BccString, &o.Bcc)
	}
	
	if Subject, ok := EmailinitialconfigurationMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if PreviousEmailId, ok := EmailinitialconfigurationMap["previousEmailId"].(string); ok {
		o.PreviousEmailId = &PreviousEmailId
	}
    
	if Held, ok := EmailinitialconfigurationMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if Alerting, ok := EmailinitialconfigurationMap["alerting"].(bool); ok {
		o.Alerting = &Alerting
	}
    
	if Inbound, ok := EmailinitialconfigurationMap["inbound"].(bool); ok {
		o.Inbound = &Inbound
	}
    
	if InvitedBy, ok := EmailinitialconfigurationMap["invitedBy"].(string); ok {
		o.InvitedBy = &InvitedBy
	}
    
	if AdditionalInfo, ok := EmailinitialconfigurationMap["additionalInfo"].(map[string]interface{}); ok {
		AdditionalInfoString, _ := json.Marshal(AdditionalInfo)
		json.Unmarshal(AdditionalInfoString, &o.AdditionalInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emailinitialconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
