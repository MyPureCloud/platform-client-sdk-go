package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Usertokenstopictokennotification
type Usertokenstopictokennotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User
	User *Usertokenstopicurireference `json:"user,omitempty"`

	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`

	// DateCreated
	DateCreated *string `json:"dateCreated,omitempty"`

	// TokenExpirationDate
	TokenExpirationDate *string `json:"tokenExpirationDate,omitempty"`

	// SessionId
	SessionId *string `json:"sessionId,omitempty"`

	// ClientId
	ClientId *string `json:"clientId,omitempty"`

	// TokenHash
	TokenHash *string `json:"tokenHash,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Usertokenstopictokennotification) SetField(field string, fieldValue interface{}) {
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

func (o Usertokenstopictokennotification) MarshalJSON() ([]byte, error) {
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
	type Alias Usertokenstopictokennotification
	
	return json.Marshal(&struct { 
		User *Usertokenstopicurireference `json:"user,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		TokenExpirationDate *string `json:"tokenExpirationDate,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		ClientId *string `json:"clientId,omitempty"`
		
		TokenHash *string `json:"tokenHash,omitempty"`
		Alias
	}{ 
		User: o.User,
		
		IpAddress: o.IpAddress,
		
		DateCreated: o.DateCreated,
		
		TokenExpirationDate: o.TokenExpirationDate,
		
		SessionId: o.SessionId,
		
		ClientId: o.ClientId,
		
		TokenHash: o.TokenHash,
		Alias:    (Alias)(o),
	})
}

func (o *Usertokenstopictokennotification) UnmarshalJSON(b []byte) error {
	var UsertokenstopictokennotificationMap map[string]interface{}
	err := json.Unmarshal(b, &UsertokenstopictokennotificationMap)
	if err != nil {
		return err
	}
	
	if User, ok := UsertokenstopictokennotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IpAddress, ok := UsertokenstopictokennotificationMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if DateCreated, ok := UsertokenstopictokennotificationMap["dateCreated"].(string); ok {
		o.DateCreated = &DateCreated
	}
    
	if TokenExpirationDate, ok := UsertokenstopictokennotificationMap["tokenExpirationDate"].(string); ok {
		o.TokenExpirationDate = &TokenExpirationDate
	}
    
	if SessionId, ok := UsertokenstopictokennotificationMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if ClientId, ok := UsertokenstopictokennotificationMap["clientId"].(string); ok {
		o.ClientId = &ClientId
	}
    
	if TokenHash, ok := UsertokenstopictokennotificationMap["tokenHash"].(string); ok {
		o.TokenHash = &TokenHash
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Usertokenstopictokennotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
