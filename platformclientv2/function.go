package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Function - Action function settings.
type Function struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Function identifier.
	Id *string `json:"id,omitempty"`

	// Name - Function name. This is automatically created based on the Action Id.
	Name *string `json:"name,omitempty"`

	// Description - Description of the function. Limit 255 characters.
	Description *string `json:"description,omitempty"`

	// DateCreated - Date and time function was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ZipId - Zip file identifier.
	ZipId *string `json:"zipId,omitempty"`

	// Handler - Handler entry point into zip file to execute function. Should be path within upload function package to the invocation module without language extension, followed by a period and then exported invocation method name. e.g. 'src/index.handler'
	Handler *string `json:"handler,omitempty"`

	// Runtime - Runtime required for execution. Valid runtimes change over time as versions are deprecated. Use /api/v2/integrations/actions/functions/runtimes for current list.
	Runtime *string `json:"runtime,omitempty"`

	// TimeoutSeconds - Execution timeout to apply to function. Value is in seconds. Range allowed 1 to 60. Default value 15 seconds.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Function) SetField(field string, fieldValue interface{}) {
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

func (o Function) MarshalJSON() ([]byte, error) {
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
	type Alias Function
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ZipId *string `json:"zipId,omitempty"`
		
		Handler *string `json:"handler,omitempty"`
		
		Runtime *string `json:"runtime,omitempty"`
		
		TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		ZipId: o.ZipId,
		
		Handler: o.Handler,
		
		Runtime: o.Runtime,
		
		TimeoutSeconds: o.TimeoutSeconds,
		Alias:    (Alias)(o),
	})
}

func (o *Function) UnmarshalJSON(b []byte) error {
	var FunctionMap map[string]interface{}
	err := json.Unmarshal(b, &FunctionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FunctionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FunctionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := FunctionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateCreatedString, ok := FunctionMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ZipId, ok := FunctionMap["zipId"].(string); ok {
		o.ZipId = &ZipId
	}
    
	if Handler, ok := FunctionMap["handler"].(string); ok {
		o.Handler = &Handler
	}
    
	if Runtime, ok := FunctionMap["runtime"].(string); ok {
		o.Runtime = &Runtime
	}
    
	if TimeoutSeconds, ok := FunctionMap["timeoutSeconds"].(float64); ok {
		TimeoutSecondsInt := int(TimeoutSeconds)
		o.TimeoutSeconds = &TimeoutSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Function) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
