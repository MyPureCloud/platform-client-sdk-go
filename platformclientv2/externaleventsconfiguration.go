package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externaleventsconfiguration
type Externaleventsconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The unique identifier for the external event configuration.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description - A description of the external event configuration.
	Description *string `json:"description,omitempty"`

	// DivisionId - The division ID (UUID) associated with this configuration.
	DivisionId *string `json:"divisionId,omitempty"`

	// DivisionIdActive - Indicates whether the divisionId field is valid.
	DivisionIdActive *bool `json:"divisionIdActive,omitempty"`

	// SchemaId - The dynamic schema ID (UUID) that defines the structure of external events.
	SchemaId *string `json:"schemaId,omitempty"`

	// SchemaActive - Indicates whether the schema is active or inactive.
	SchemaActive *bool `json:"schemaActive,omitempty"`

	// Source - The source of the external events e.g. Adobe, Salesforce.
	Source *string `json:"source,omitempty"`

	// DateLastModified - The timestamp when the configuration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastModified *time.Time `json:"dateLastModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externaleventsconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Externaleventsconfiguration) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastModified", }
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
	type Alias Externaleventsconfiguration
	
	DateLastModified := new(string)
	if o.DateLastModified != nil {
		
		*DateLastModified = timeutil.Strftime(o.DateLastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		DivisionIdActive *bool `json:"divisionIdActive,omitempty"`
		
		SchemaId *string `json:"schemaId,omitempty"`
		
		SchemaActive *bool `json:"schemaActive,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		DateLastModified *string `json:"dateLastModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DivisionId: o.DivisionId,
		
		DivisionIdActive: o.DivisionIdActive,
		
		SchemaId: o.SchemaId,
		
		SchemaActive: o.SchemaActive,
		
		Source: o.Source,
		
		DateLastModified: DateLastModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Externaleventsconfiguration) UnmarshalJSON(b []byte) error {
	var ExternaleventsconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &ExternaleventsconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternaleventsconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ExternaleventsconfigurationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ExternaleventsconfigurationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if DivisionId, ok := ExternaleventsconfigurationMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if DivisionIdActive, ok := ExternaleventsconfigurationMap["divisionIdActive"].(bool); ok {
		o.DivisionIdActive = &DivisionIdActive
	}
    
	if SchemaId, ok := ExternaleventsconfigurationMap["schemaId"].(string); ok {
		o.SchemaId = &SchemaId
	}
    
	if SchemaActive, ok := ExternaleventsconfigurationMap["schemaActive"].(bool); ok {
		o.SchemaActive = &SchemaActive
	}
    
	if Source, ok := ExternaleventsconfigurationMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if dateLastModifiedString, ok := ExternaleventsconfigurationMap["dateLastModified"].(string); ok {
		DateLastModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastModifiedString)
		o.DateLastModified = &DateLastModified
	}
	
	if SelfUri, ok := ExternaleventsconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externaleventsconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
