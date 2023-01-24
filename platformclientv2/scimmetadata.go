package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimmetadata - Defines the SCIM metadata.
type Scimmetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ResourceType - The type of SCIM resource.
	ResourceType *string `json:"resourceType,omitempty"`

	// LastModified - The last time that the resource was modified. Date time is represented as an \"ISO-8601 string\", for example, yyyy-MM-ddTHH:mm:ss.SSSZ. Not included with \"Schema\" and \"ResourceType\" resources.
	LastModified *time.Time `json:"lastModified,omitempty"`

	// Location - The URI of the resource.
	Location *string `json:"location,omitempty"`

	// Version - The version of the resource. Matches the ETag HTTP response header. Not included with \"Schema\" and \"ResourceType\" resources.
	Version *string `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimmetadata) SetField(field string, fieldValue interface{}) {
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

func (o Scimmetadata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "LastModified", }
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
	type Alias Scimmetadata
	
	LastModified := new(string)
	if o.LastModified != nil {
		
		*LastModified = timeutil.Strftime(o.LastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastModified = nil
	}
	
	return json.Marshal(&struct { 
		ResourceType *string `json:"resourceType,omitempty"`
		
		LastModified *string `json:"lastModified,omitempty"`
		
		Location *string `json:"location,omitempty"`
		
		Version *string `json:"version,omitempty"`
		Alias
	}{ 
		ResourceType: o.ResourceType,
		
		LastModified: LastModified,
		
		Location: o.Location,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Scimmetadata) UnmarshalJSON(b []byte) error {
	var ScimmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &ScimmetadataMap)
	if err != nil {
		return err
	}
	
	if ResourceType, ok := ScimmetadataMap["resourceType"].(string); ok {
		o.ResourceType = &ResourceType
	}
    
	if lastModifiedString, ok := ScimmetadataMap["lastModified"].(string); ok {
		LastModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastModifiedString)
		o.LastModified = &LastModified
	}
	
	if Location, ok := ScimmetadataMap["location"].(string); ok {
		o.Location = &Location
	}
    
	if Version, ok := ScimmetadataMap["version"].(string); ok {
		o.Version = &Version
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
