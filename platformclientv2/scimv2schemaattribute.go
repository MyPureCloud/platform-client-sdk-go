package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2schemaattribute - A complex type that defines service provider attributes or subattributes and their qualities.
type Scimv2schemaattribute struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the attribute.
	Name *string `json:"name,omitempty"`

	// VarType - The data type of the attribute.
	VarType *string `json:"type,omitempty"`

	// SubAttributes - The list of subattributes for an attribute of the type \"complex\". Uses the same schema as \"attributes\".
	SubAttributes *[]Scimv2schemaattribute `json:"subAttributes,omitempty"`

	// MultiValued - Indicates whether an attribute contains multiple values.
	MultiValued *bool `json:"multiValued,omitempty"`

	// Description - The description of the attribute.
	Description *string `json:"description,omitempty"`

	// Required - Indicates whether an attribute is required.
	Required *bool `json:"required,omitempty"`

	// CanonicalValues - The list of standard values that service providers may use. Service providers may ignore unsupported values.
	CanonicalValues *[]string `json:"canonicalValues,omitempty"`

	// CaseExact - Indicates whether a string attribute is case-sensitive. If set to \"true\", the server preserves case sensitivity. If set to \"false\", the server may change the case. The server also uses case sensitivity when evaluating filters. See section 3.4.2.2 \"Filtering\" in RFC 7644 for details.
	CaseExact *bool `json:"caseExact,omitempty"`

	// Mutability - The circumstances under which an attribute can be defined or redefined. The default is \"readWrite\".
	Mutability *string `json:"mutability,omitempty"`

	// Returned - The circumstances under which an attribute and its values are returned in response to a GET, PUT, POST, or PATCH request.
	Returned *string `json:"returned,omitempty"`

	// Uniqueness - The method by which the service provider enforces the uniqueness of an attribute value. A server can reject a value by returning the HTTP response code 400 (Bad Request). A client can enforce uniqueness to a greater degree than the server provider enforces. For example, a client could make a value unique even though the server has \"uniqueness\" set to \"none\".
	Uniqueness *string `json:"uniqueness,omitempty"`

	// ReferenceTypes - The list of SCIM resource types that may be referenced. Only applies when \"type\" is set to \"reference\".
	ReferenceTypes *[]string `json:"referenceTypes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimv2schemaattribute) SetField(field string, fieldValue interface{}) {
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

func (o Scimv2schemaattribute) MarshalJSON() ([]byte, error) {
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
	type Alias Scimv2schemaattribute
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		SubAttributes *[]Scimv2schemaattribute `json:"subAttributes,omitempty"`
		
		MultiValued *bool `json:"multiValued,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		
		CanonicalValues *[]string `json:"canonicalValues,omitempty"`
		
		CaseExact *bool `json:"caseExact,omitempty"`
		
		Mutability *string `json:"mutability,omitempty"`
		
		Returned *string `json:"returned,omitempty"`
		
		Uniqueness *string `json:"uniqueness,omitempty"`
		
		ReferenceTypes *[]string `json:"referenceTypes,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		
		SubAttributes: o.SubAttributes,
		
		MultiValued: o.MultiValued,
		
		Description: o.Description,
		
		Required: o.Required,
		
		CanonicalValues: o.CanonicalValues,
		
		CaseExact: o.CaseExact,
		
		Mutability: o.Mutability,
		
		Returned: o.Returned,
		
		Uniqueness: o.Uniqueness,
		
		ReferenceTypes: o.ReferenceTypes,
		Alias:    (Alias)(o),
	})
}

func (o *Scimv2schemaattribute) UnmarshalJSON(b []byte) error {
	var Scimv2schemaattributeMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2schemaattributeMap)
	if err != nil {
		return err
	}
	
	if Name, ok := Scimv2schemaattributeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := Scimv2schemaattributeMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if SubAttributes, ok := Scimv2schemaattributeMap["subAttributes"].([]interface{}); ok {
		SubAttributesString, _ := json.Marshal(SubAttributes)
		json.Unmarshal(SubAttributesString, &o.SubAttributes)
	}
	
	if MultiValued, ok := Scimv2schemaattributeMap["multiValued"].(bool); ok {
		o.MultiValued = &MultiValued
	}
    
	if Description, ok := Scimv2schemaattributeMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Required, ok := Scimv2schemaattributeMap["required"].(bool); ok {
		o.Required = &Required
	}
    
	if CanonicalValues, ok := Scimv2schemaattributeMap["canonicalValues"].([]interface{}); ok {
		CanonicalValuesString, _ := json.Marshal(CanonicalValues)
		json.Unmarshal(CanonicalValuesString, &o.CanonicalValues)
	}
	
	if CaseExact, ok := Scimv2schemaattributeMap["caseExact"].(bool); ok {
		o.CaseExact = &CaseExact
	}
    
	if Mutability, ok := Scimv2schemaattributeMap["mutability"].(string); ok {
		o.Mutability = &Mutability
	}
    
	if Returned, ok := Scimv2schemaattributeMap["returned"].(string); ok {
		o.Returned = &Returned
	}
    
	if Uniqueness, ok := Scimv2schemaattributeMap["uniqueness"].(string); ok {
		o.Uniqueness = &Uniqueness
	}
    
	if ReferenceTypes, ok := Scimv2schemaattributeMap["referenceTypes"].([]interface{}); ok {
		ReferenceTypesString, _ := json.Marshal(ReferenceTypes)
		json.Unmarshal(ReferenceTypesString, &o.ReferenceTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2schemaattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
