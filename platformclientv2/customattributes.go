package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Customattributes
type Customattributes struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of the Custom Attributes record.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ConversationId - The id of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`

	// Divisions - The list of division ids that the record is visible in. If [], the record is visible to all divisions (Unassigned Division).
	Divisions *[]string `json:"divisions,omitempty"`

	// Schema - The schema that dictates which attributes can be included.
	Schema *Conversationdataschema `json:"schema,omitempty"`

	// CustomAttributes - The map of attribute values.
	CustomAttributes *map[string]interface{} `json:"customAttributes,omitempty"`

	// CustomAttributesTimestamps - The map of timestamps for when each attribute was last updated.
	CustomAttributesTimestamps *map[string]string `json:"customAttributesTimestamps,omitempty"`

	// Version - The latest version of the Custom Attributes record.
	Version *int `json:"version,omitempty"`

	// DateCreated - The date the record was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The date the record was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Customattributes) SetField(field string, fieldValue interface{}) {
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

func (o Customattributes) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Customattributes
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		Divisions *[]string `json:"divisions,omitempty"`
		
		Schema *Conversationdataschema `json:"schema,omitempty"`
		
		CustomAttributes *map[string]interface{} `json:"customAttributes,omitempty"`
		
		CustomAttributesTimestamps *map[string]string `json:"customAttributesTimestamps,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ConversationId: o.ConversationId,
		
		Divisions: o.Divisions,
		
		Schema: o.Schema,
		
		CustomAttributes: o.CustomAttributes,
		
		CustomAttributesTimestamps: o.CustomAttributesTimestamps,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Customattributes) UnmarshalJSON(b []byte) error {
	var CustomattributesMap map[string]interface{}
	err := json.Unmarshal(b, &CustomattributesMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CustomattributesMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CustomattributesMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ConversationId, ok := CustomattributesMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if Divisions, ok := CustomattributesMap["divisions"].([]interface{}); ok {
		DivisionsString, _ := json.Marshal(Divisions)
		json.Unmarshal(DivisionsString, &o.Divisions)
	}
	
	if Schema, ok := CustomattributesMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomAttributes, ok := CustomattributesMap["customAttributes"].(map[string]interface{}); ok {
		CustomAttributesString, _ := json.Marshal(CustomAttributes)
		json.Unmarshal(CustomAttributesString, &o.CustomAttributes)
	}
	
	if CustomAttributesTimestamps, ok := CustomattributesMap["customAttributesTimestamps"].(map[string]interface{}); ok {
		CustomAttributesTimestampsString, _ := json.Marshal(CustomAttributesTimestamps)
		json.Unmarshal(CustomAttributesTimestampsString, &o.CustomAttributesTimestamps)
	}
	
	if Version, ok := CustomattributesMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := CustomattributesMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := CustomattributesMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := CustomattributesMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Customattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
