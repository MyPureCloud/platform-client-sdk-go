package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappintegrationupdaterequest
type Whatsappintegrationupdaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - WhatsApp Integration name
	Name *string `json:"name,omitempty"`

	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// MessagingSetting - Defines the message settings to be applied for this integration
	MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`

	// Action - The action used to activate and then confirm a WhatsApp Integration.
	Action *string `json:"action,omitempty"`

	// AuthenticationMethod - The authentication method used to confirm a WhatsApp Integration activation. If action is set to Activate, then authenticationMethod is a required field. 
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`

	// ConfirmationCode - The confirmation code sent by Whatsapp to you during the activation step. If action is set to Confirm, then confirmationCode is a required field.
	ConfirmationCode *string `json:"confirmationCode,omitempty"`

	// PhoneNumber - Phone number to associate with the WhatsApp integration
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Whatsappintegrationupdaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Whatsappintegrationupdaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Whatsappintegrationupdaterequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingrequestreference `json:"messagingSetting,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
		
		ConfirmationCode *string `json:"confirmationCode,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		Action: o.Action,
		
		AuthenticationMethod: o.AuthenticationMethod,
		
		ConfirmationCode: o.ConfirmationCode,
		
		PhoneNumber: o.PhoneNumber,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Whatsappintegrationupdaterequest) UnmarshalJSON(b []byte) error {
	var WhatsappintegrationupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappintegrationupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WhatsappintegrationupdaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WhatsappintegrationupdaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := WhatsappintegrationupdaterequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := WhatsappintegrationupdaterequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if Action, ok := WhatsappintegrationupdaterequestMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if AuthenticationMethod, ok := WhatsappintegrationupdaterequestMap["authenticationMethod"].(string); ok {
		o.AuthenticationMethod = &AuthenticationMethod
	}
    
	if ConfirmationCode, ok := WhatsappintegrationupdaterequestMap["confirmationCode"].(string); ok {
		o.ConfirmationCode = &ConfirmationCode
	}
    
	if PhoneNumber, ok := WhatsappintegrationupdaterequestMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if SelfUri, ok := WhatsappintegrationupdaterequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappintegrationupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
