package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Opensocialmediarecipient - Information about the recipient the message is sent to or received from.
type Opensocialmediarecipient struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The recipient ID specific to the provider.
	Id *string `json:"id,omitempty"`

	// IdType - The recipient ID type. This is used to indicate the format used for the ID.
	IdType *string `json:"idType,omitempty"`

	// FirstName - First name of the recipient.
	FirstName *string `json:"firstName,omitempty"`

	// LastName - Last name of the recipient.
	LastName *string `json:"lastName,omitempty"`

	// Nickname - Nickname or display name of the recipient.
	Nickname *string `json:"nickname,omitempty"`

	// Image - URL of an image that represents the recipient.
	Image *string `json:"image,omitempty"`

	// AdditionalIds - List of recipient additional identifiers
	AdditionalIds *[]Opensocialmediarecipientadditionalidentifier `json:"additionalIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Opensocialmediarecipient) SetField(field string, fieldValue interface{}) {
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

func (o Opensocialmediarecipient) MarshalJSON() ([]byte, error) {
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
	type Alias Opensocialmediarecipient
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		IdType *string `json:"idType,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Nickname *string `json:"nickname,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		AdditionalIds *[]Opensocialmediarecipientadditionalidentifier `json:"additionalIds,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		IdType: o.IdType,
		
		FirstName: o.FirstName,
		
		LastName: o.LastName,
		
		Nickname: o.Nickname,
		
		Image: o.Image,
		
		AdditionalIds: o.AdditionalIds,
		Alias:    (Alias)(o),
	})
}

func (o *Opensocialmediarecipient) UnmarshalJSON(b []byte) error {
	var OpensocialmediarecipientMap map[string]interface{}
	err := json.Unmarshal(b, &OpensocialmediarecipientMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OpensocialmediarecipientMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IdType, ok := OpensocialmediarecipientMap["idType"].(string); ok {
		o.IdType = &IdType
	}
    
	if FirstName, ok := OpensocialmediarecipientMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if LastName, ok := OpensocialmediarecipientMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Nickname, ok := OpensocialmediarecipientMap["nickname"].(string); ok {
		o.Nickname = &Nickname
	}
    
	if Image, ok := OpensocialmediarecipientMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if AdditionalIds, ok := OpensocialmediarecipientMap["additionalIds"].([]interface{}); ok {
		AdditionalIdsString, _ := json.Marshal(AdditionalIds)
		json.Unmarshal(AdditionalIdsString, &o.AdditionalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Opensocialmediarecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
