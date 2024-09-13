package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Functionzipconfig - Action function zip file upload settings and state.
type Functionzipconfig struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - Status of zip upload.
	Status *string `json:"status,omitempty"`

	// Id - Zip file Identifier
	Id *string `json:"id,omitempty"`

	// Name - Zip file name
	Name *string `json:"name,omitempty"`

	// DateCreated - Date and time zip record was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ErrorMessage - Error message if upload failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// RequestId - Upload request id used for zip upload
	RequestId *string `json:"requestId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Functionzipconfig) SetField(field string, fieldValue interface{}) {
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

func (o Functionzipconfig) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Functionzipconfig
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		RequestId *string `json:"requestId,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		ErrorMessage: o.ErrorMessage,
		
		RequestId: o.RequestId,
		Alias:    (Alias)(o),
	})
}

func (o *Functionzipconfig) UnmarshalJSON(b []byte) error {
	var FunctionzipconfigMap map[string]interface{}
	err := json.Unmarshal(b, &FunctionzipconfigMap)
	if err != nil {
		return err
	}
	
	if Status, ok := FunctionzipconfigMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Id, ok := FunctionzipconfigMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FunctionzipconfigMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := FunctionzipconfigMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ErrorMessage, ok := FunctionzipconfigMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if RequestId, ok := FunctionzipconfigMap["requestId"].(string); ok {
		o.RequestId = &RequestId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Functionzipconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
