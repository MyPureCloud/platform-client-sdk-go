package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Gdprsubject
type Gdprsubject struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name
	Name *string `json:"name,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// DialerContactId
	DialerContactId *Dialercontactid `json:"dialerContactId,omitempty"`

	// JourneyCustomer
	JourneyCustomer *Gdprjourneycustomer `json:"journeyCustomer,omitempty"`

	// SocialHandle
	SocialHandle *Socialhandle `json:"socialHandle,omitempty"`

	// ExternalId
	ExternalId *string `json:"externalId,omitempty"`

	// Addresses
	Addresses *[]string `json:"addresses,omitempty"`

	// PhoneNumbers
	PhoneNumbers *[]string `json:"phoneNumbers,omitempty"`

	// EmailAddresses
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Gdprsubject) SetField(field string, fieldValue interface{}) {
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

func (o Gdprsubject) MarshalJSON() ([]byte, error) {
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
	type Alias Gdprsubject
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		DialerContactId *Dialercontactid `json:"dialerContactId,omitempty"`
		
		JourneyCustomer *Gdprjourneycustomer `json:"journeyCustomer,omitempty"`
		
		SocialHandle *Socialhandle `json:"socialHandle,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		Addresses *[]string `json:"addresses,omitempty"`
		
		PhoneNumbers *[]string `json:"phoneNumbers,omitempty"`
		
		EmailAddresses *[]string `json:"emailAddresses,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		UserId: o.UserId,
		
		ExternalContactId: o.ExternalContactId,
		
		DialerContactId: o.DialerContactId,
		
		JourneyCustomer: o.JourneyCustomer,
		
		SocialHandle: o.SocialHandle,
		
		ExternalId: o.ExternalId,
		
		Addresses: o.Addresses,
		
		PhoneNumbers: o.PhoneNumbers,
		
		EmailAddresses: o.EmailAddresses,
		Alias:    (Alias)(o),
	})
}

func (o *Gdprsubject) UnmarshalJSON(b []byte) error {
	var GdprsubjectMap map[string]interface{}
	err := json.Unmarshal(b, &GdprsubjectMap)
	if err != nil {
		return err
	}
	
	if Name, ok := GdprsubjectMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UserId, ok := GdprsubjectMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if ExternalContactId, ok := GdprsubjectMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if DialerContactId, ok := GdprsubjectMap["dialerContactId"].(map[string]interface{}); ok {
		DialerContactIdString, _ := json.Marshal(DialerContactId)
		json.Unmarshal(DialerContactIdString, &o.DialerContactId)
	}
	
	if JourneyCustomer, ok := GdprsubjectMap["journeyCustomer"].(map[string]interface{}); ok {
		JourneyCustomerString, _ := json.Marshal(JourneyCustomer)
		json.Unmarshal(JourneyCustomerString, &o.JourneyCustomer)
	}
	
	if SocialHandle, ok := GdprsubjectMap["socialHandle"].(map[string]interface{}); ok {
		SocialHandleString, _ := json.Marshal(SocialHandle)
		json.Unmarshal(SocialHandleString, &o.SocialHandle)
	}
	
	if ExternalId, ok := GdprsubjectMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if Addresses, ok := GdprsubjectMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if PhoneNumbers, ok := GdprsubjectMap["phoneNumbers"].([]interface{}); ok {
		PhoneNumbersString, _ := json.Marshal(PhoneNumbers)
		json.Unmarshal(PhoneNumbersString, &o.PhoneNumbers)
	}
	
	if EmailAddresses, ok := GdprsubjectMap["emailAddresses"].([]interface{}); ok {
		EmailAddressesString, _ := json.Marshal(EmailAddresses)
		json.Unmarshal(EmailAddressesString, &o.EmailAddresses)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gdprsubject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
