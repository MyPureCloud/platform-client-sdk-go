package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataschema
type Dataschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the schema.  Only required if a schema is used for custom fields during external entity creation or updates.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Version - The schema's version, a positive integer. Required for updates.
	Version *int `json:"version,omitempty"`

	// AppliesTo - One of \"CONTACT\" or \"EXTERNAL_ORGANIZATION\".  Indicates the built-in entity type to which this schema applies.
	AppliesTo *[]string `json:"appliesTo,omitempty"`

	// Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
	Enabled *bool `json:"enabled,omitempty"`

	// CreatedBy - The URI of the user that created this schema.
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`

	// DateCreated - The date and time this schema was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// JsonSchema - A JSON schema defining the extension to the built-in entity type.
	JsonSchema *Jsonschemadocument `json:"jsonSchema,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dataschema) SetField(field string, fieldValue interface{}) {
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

func (o Dataschema) MarshalJSON() ([]byte, error) {
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
	type Alias Dataschema
	
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
		
		AppliesTo *[]string `json:"appliesTo,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		JsonSchema *Jsonschemadocument `json:"jsonSchema,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		AppliesTo: o.AppliesTo,
		
		Enabled: o.Enabled,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		JsonSchema: o.JsonSchema,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Dataschema) UnmarshalJSON(b []byte) error {
	var DataschemaMap map[string]interface{}
	err := json.Unmarshal(b, &DataschemaMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DataschemaMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DataschemaMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := DataschemaMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if AppliesTo, ok := DataschemaMap["appliesTo"].([]interface{}); ok {
		AppliesToString, _ := json.Marshal(AppliesTo)
		json.Unmarshal(AppliesToString, &o.AppliesTo)
	}
	
	if Enabled, ok := DataschemaMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if CreatedBy, ok := DataschemaMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := DataschemaMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if JsonSchema, ok := DataschemaMap["jsonSchema"].(map[string]interface{}); ok {
		JsonSchemaString, _ := json.Marshal(JsonSchema)
		json.Unmarshal(JsonSchemaString, &o.JsonSchema)
	}
	
	if SelfUri, ok := DataschemaMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dataschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
