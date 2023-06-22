package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Ucthirdpartypresence - Update a Genesys Cloud user's presence from a given 3rd-party integration
type Ucthirdpartypresence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Email - Primary Email address of the associated Genesys Cloud user.
	Email *string `json:"email,omitempty"`

	// Presence - Integration presence value.
	Presence *string `json:"presence,omitempty"`

	// Message - Integration presence message.
	Message *string `json:"message,omitempty"`

	// DateModified - ISO 8601 timestamp of presence value change.
	DateModified *time.Time `json:"dateModified,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Ucthirdpartypresence) SetField(field string, fieldValue interface{}) {
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

func (o Ucthirdpartypresence) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateModified", }
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
	type Alias Ucthirdpartypresence
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Email *string `json:"email,omitempty"`
		
		Presence *string `json:"presence,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		Alias
	}{ 
		Email: o.Email,
		
		Presence: o.Presence,
		
		Message: o.Message,
		
		DateModified: DateModified,
		Alias:    (Alias)(o),
	})
}

func (o *Ucthirdpartypresence) UnmarshalJSON(b []byte) error {
	var UcthirdpartypresenceMap map[string]interface{}
	err := json.Unmarshal(b, &UcthirdpartypresenceMap)
	if err != nil {
		return err
	}
	
	if Email, ok := UcthirdpartypresenceMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if Presence, ok := UcthirdpartypresenceMap["presence"].(string); ok {
		o.Presence = &Presence
	}
    
	if Message, ok := UcthirdpartypresenceMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if dateModifiedString, ok := UcthirdpartypresenceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ucthirdpartypresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
