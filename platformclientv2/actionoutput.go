package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionoutput - Output definition of Action.
type Actionoutput struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SuccessSchema - JSON schema that defines the transformed, successful result that will be sent back to the caller. If the 'flatten' query parameter is omitted or false, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
	SuccessSchema *Jsonschemadocument `json:"successSchema,omitempty"`

	// SuccessSchemaUri - URI to retrieve success schema
	SuccessSchemaUri *string `json:"successSchemaUri,omitempty"`

	// ErrorSchema - JSON schema that defines the body of response when request is not successful. If the 'flatten' query parameter is omitted or false, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
	ErrorSchema *Jsonschemadocument `json:"errorSchema,omitempty"`

	// ErrorSchemaUri - URI to retrieve error schema
	ErrorSchemaUri *string `json:"errorSchemaUri,omitempty"`

	// SuccessSchemaFlattened - JSON schema that defines the transformed, successful result that will be sent back to the caller. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
	SuccessSchemaFlattened *Jsonschemadocument `json:"successSchemaFlattened,omitempty"`

	// ErrorSchemaFlattened - JSON schema that defines the body of response when request is not successful. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
	ErrorSchemaFlattened *interface{} `json:"errorSchemaFlattened,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actionoutput) SetField(field string, fieldValue interface{}) {
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

func (o Actionoutput) MarshalJSON() ([]byte, error) {
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
	type Alias Actionoutput
	
	return json.Marshal(&struct { 
		SuccessSchema *Jsonschemadocument `json:"successSchema,omitempty"`
		
		SuccessSchemaUri *string `json:"successSchemaUri,omitempty"`
		
		ErrorSchema *Jsonschemadocument `json:"errorSchema,omitempty"`
		
		ErrorSchemaUri *string `json:"errorSchemaUri,omitempty"`
		
		SuccessSchemaFlattened *Jsonschemadocument `json:"successSchemaFlattened,omitempty"`
		
		ErrorSchemaFlattened *interface{} `json:"errorSchemaFlattened,omitempty"`
		Alias
	}{ 
		SuccessSchema: o.SuccessSchema,
		
		SuccessSchemaUri: o.SuccessSchemaUri,
		
		ErrorSchema: o.ErrorSchema,
		
		ErrorSchemaUri: o.ErrorSchemaUri,
		
		SuccessSchemaFlattened: o.SuccessSchemaFlattened,
		
		ErrorSchemaFlattened: o.ErrorSchemaFlattened,
		Alias:    (Alias)(o),
	})
}

func (o *Actionoutput) UnmarshalJSON(b []byte) error {
	var ActionoutputMap map[string]interface{}
	err := json.Unmarshal(b, &ActionoutputMap)
	if err != nil {
		return err
	}
	
	if SuccessSchema, ok := ActionoutputMap["successSchema"].(map[string]interface{}); ok {
		SuccessSchemaString, _ := json.Marshal(SuccessSchema)
		json.Unmarshal(SuccessSchemaString, &o.SuccessSchema)
	}
	
	if SuccessSchemaUri, ok := ActionoutputMap["successSchemaUri"].(string); ok {
		o.SuccessSchemaUri = &SuccessSchemaUri
	}
    
	if ErrorSchema, ok := ActionoutputMap["errorSchema"].(map[string]interface{}); ok {
		ErrorSchemaString, _ := json.Marshal(ErrorSchema)
		json.Unmarshal(ErrorSchemaString, &o.ErrorSchema)
	}
	
	if ErrorSchemaUri, ok := ActionoutputMap["errorSchemaUri"].(string); ok {
		o.ErrorSchemaUri = &ErrorSchemaUri
	}
    
	if SuccessSchemaFlattened, ok := ActionoutputMap["successSchemaFlattened"].(map[string]interface{}); ok {
		SuccessSchemaFlattenedString, _ := json.Marshal(SuccessSchemaFlattened)
		json.Unmarshal(SuccessSchemaFlattenedString, &o.SuccessSchemaFlattened)
	}
	
	if ErrorSchemaFlattened, ok := ActionoutputMap["errorSchemaFlattened"].(map[string]interface{}); ok {
		ErrorSchemaFlattenedString, _ := json.Marshal(ErrorSchemaFlattened)
		json.Unmarshal(ErrorSchemaFlattenedString, &o.ErrorSchemaFlattened)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionoutput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
