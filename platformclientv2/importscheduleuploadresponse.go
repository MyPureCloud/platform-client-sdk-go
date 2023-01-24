package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Importscheduleuploadresponse
type Importscheduleuploadresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UploadKey - The key to pass to the secondary request to start processing of the upload
	UploadKey *string `json:"uploadKey,omitempty"`

	// Url - The url to which to PUT the upload body
	Url *string `json:"url,omitempty"`

	// Headers - Required headers for the PUT request to the url
	Headers *map[string]string `json:"headers,omitempty"`

	// UploadBodySchema - Always null. Defines the schema of the json body to be PUT to the url. The json body should be gzip encoded before uploading
	UploadBodySchema *Importscheduleuploadschema `json:"uploadBodySchema,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Importscheduleuploadresponse) SetField(field string, fieldValue interface{}) {
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

func (o Importscheduleuploadresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Importscheduleuploadresponse
	
	return json.Marshal(&struct { 
		UploadKey *string `json:"uploadKey,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		
		UploadBodySchema *Importscheduleuploadschema `json:"uploadBodySchema,omitempty"`
		Alias
	}{ 
		UploadKey: o.UploadKey,
		
		Url: o.Url,
		
		Headers: o.Headers,
		
		UploadBodySchema: o.UploadBodySchema,
		Alias:    (Alias)(o),
	})
}

func (o *Importscheduleuploadresponse) UnmarshalJSON(b []byte) error {
	var ImportscheduleuploadresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ImportscheduleuploadresponseMap)
	if err != nil {
		return err
	}
	
	if UploadKey, ok := ImportscheduleuploadresponseMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if Url, ok := ImportscheduleuploadresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Headers, ok := ImportscheduleuploadresponseMap["headers"].(map[string]interface{}); ok {
		HeadersString, _ := json.Marshal(Headers)
		json.Unmarshal(HeadersString, &o.Headers)
	}
	
	if UploadBodySchema, ok := ImportscheduleuploadresponseMap["uploadBodySchema"].(map[string]interface{}); ok {
		UploadBodySchemaString, _ := json.Marshal(UploadBodySchema)
		json.Unmarshal(UploadBodySchemaString, &o.UploadBodySchema)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Importscheduleuploadresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
