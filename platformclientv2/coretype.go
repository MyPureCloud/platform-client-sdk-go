package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Coretype
type Coretype struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Version - A positive integer denoting the core type's version
	Version *int `json:"version,omitempty"`

	// DateCreated - The date the core type was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// Schema - The core type's built-in schema
	Schema *Schema `json:"schema,omitempty"`

	// Current - A boolean indicating if the core type's version is the current one in use by the system
	Current *bool `json:"current,omitempty"`

	// ValidationFields - An array of strings naming the fields of the core type subject to validation.  Validation constraints are specified by a schema author using the core type.
	ValidationFields *[]string `json:"validationFields,omitempty"`

	// ValidationLimits - A structure denoting the system-imposed minimum and maximum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
	ValidationLimits *Validationlimits `json:"validationLimits,omitempty"`

	// ItemValidationFields - Specific to the \"tag\" core type, this is an array of strings naming the tag item fields of the core type subject to validation
	ItemValidationFields *[]string `json:"itemValidationFields,omitempty"`

	// ItemValidationLimits - A structure denoting the system-imposed minimum and maximum string length for string-array based core types such as \"tag\" and \"enum\".  Forexample, the validationLimits for a schema field using a tag core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schema author on individual tags.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field's tags.
	ItemValidationLimits *Itemvalidationlimits `json:"itemValidationLimits,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Coretype) SetField(field string, fieldValue interface{}) {
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

func (o Coretype) MarshalJSON() ([]byte, error) {
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
	type Alias Coretype
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Schema *Schema `json:"schema,omitempty"`
		
		Current *bool `json:"current,omitempty"`
		
		ValidationFields *[]string `json:"validationFields,omitempty"`
		
		ValidationLimits *Validationlimits `json:"validationLimits,omitempty"`
		
		ItemValidationFields *[]string `json:"itemValidationFields,omitempty"`
		
		ItemValidationLimits *Itemvalidationlimits `json:"itemValidationLimits,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		Schema: o.Schema,
		
		Current: o.Current,
		
		ValidationFields: o.ValidationFields,
		
		ValidationLimits: o.ValidationLimits,
		
		ItemValidationFields: o.ItemValidationFields,
		
		ItemValidationLimits: o.ItemValidationLimits,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Coretype) UnmarshalJSON(b []byte) error {
	var CoretypeMap map[string]interface{}
	err := json.Unmarshal(b, &CoretypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CoretypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CoretypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := CoretypeMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := CoretypeMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Schema, ok := CoretypeMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if Current, ok := CoretypeMap["current"].(bool); ok {
		o.Current = &Current
	}
    
	if ValidationFields, ok := CoretypeMap["validationFields"].([]interface{}); ok {
		ValidationFieldsString, _ := json.Marshal(ValidationFields)
		json.Unmarshal(ValidationFieldsString, &o.ValidationFields)
	}
	
	if ValidationLimits, ok := CoretypeMap["validationLimits"].(map[string]interface{}); ok {
		ValidationLimitsString, _ := json.Marshal(ValidationLimits)
		json.Unmarshal(ValidationLimitsString, &o.ValidationLimits)
	}
	
	if ItemValidationFields, ok := CoretypeMap["itemValidationFields"].([]interface{}); ok {
		ItemValidationFieldsString, _ := json.Marshal(ItemValidationFields)
		json.Unmarshal(ItemValidationFieldsString, &o.ItemValidationFields)
	}
	
	if ItemValidationLimits, ok := CoretypeMap["itemValidationLimits"].(map[string]interface{}); ok {
		ItemValidationLimitsString, _ := json.Marshal(ItemValidationLimits)
		json.Unmarshal(ItemValidationLimitsString, &o.ItemValidationLimits)
	}
	
	if SelfUri, ok := CoretypeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Coretype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
