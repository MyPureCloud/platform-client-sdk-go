package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Guestmemberinfo
type Guestmemberinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DisplayName - The display name to use for the guest member in the conversation.
	DisplayName *string `json:"displayName,omitempty"`

	// FirstName - The first name to use for the guest member in the conversation.
	FirstName *string `json:"firstName,omitempty"`

	// LastName - The last name to use for the guest member in the conversation.
	LastName *string `json:"lastName,omitempty"`

	// Email - The email address to use for the guest member in the conversation.
	Email *string `json:"email,omitempty"`

	// PhoneNumber - The phone number to use for the guest member in the conversation.
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// AvatarImageUrl - The URL to the avatar image to use for the guest member in the conversation, if any.
	AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`

	// CustomFields - Any custom fields of information, in key-value format, to attach to the guest member in the conversation.
	CustomFields *map[string]string `json:"customFields,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Guestmemberinfo) SetField(field string, fieldValue interface{}) {
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

func (o Guestmemberinfo) MarshalJSON() ([]byte, error) {
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
	type Alias Guestmemberinfo
	
	return json.Marshal(&struct { 
		DisplayName *string `json:"displayName,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`
		
		CustomFields *map[string]string `json:"customFields,omitempty"`
		Alias
	}{ 
		DisplayName: o.DisplayName,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Email: o.Email,
		
		PhoneNumber: o.PhoneNumber,
		
		AvatarImageUrl: o.AvatarImageUrl,
		
		CustomFields: o.CustomFields,
		Alias:    (Alias)(o),
	})
}

func (o *Guestmemberinfo) UnmarshalJSON(b []byte) error {
	var GuestmemberinfoMap map[string]interface{}
	err := json.Unmarshal(b, &GuestmemberinfoMap)
	if err != nil {
		return err
	}
	
	if DisplayName, ok := GuestmemberinfoMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if FirstName, ok := GuestmemberinfoMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if LastName, ok := GuestmemberinfoMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Email, ok := GuestmemberinfoMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if PhoneNumber, ok := GuestmemberinfoMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if AvatarImageUrl, ok := GuestmemberinfoMap["avatarImageUrl"].(string); ok {
		o.AvatarImageUrl = &AvatarImageUrl
	}
    
	if CustomFields, ok := GuestmemberinfoMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Guestmemberinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
