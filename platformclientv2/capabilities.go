package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilities
type Capabilities struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Actions - Architect flow actions, mapped by action ID
	Actions *map[string]Capabilitiesaction `json:"actions,omitempty"`

	// DataTypes - Architect flow data types, mapped by data type ID
	DataTypes *map[string]Capabilitiesdatatype `json:"dataTypes,omitempty"`

	// DependencyTypes - Architect dependency types, mapped by dependency type ID
	DependencyTypes *map[string]Capabilitiesdependencytype `json:"dependencyTypes,omitempty"`

	// Expansions - Expansion metadata. e.g. flowType all to a list of all of the flow types
	Expansions *map[string]interface{} `json:"expansions,omitempty"`

	// Features - Architect features status, mapped by feature ID
	Features *map[string]Capabilitiesavailability `json:"features,omitempty"`

	// FeatureSupport - Architect features metadata
	FeatureSupport *Capabilitiesfeaturesupport `json:"featureSupport,omitempty"`

	// FlowTypes - Architect flow types, mapped by flow type ID
	FlowTypes *map[string]Capabilitiesflowtype `json:"flowTypes,omitempty"`

	// Functions - Architect functions, mapped by function ID
	Functions *map[string]Capabilitiesfunction `json:"functions,omitempty"`

	// Languages - Architect languages, mapped by language ID
	Languages *map[string]Capabilitieslanguage `json:"languages,omitempty"`

	// Organization - Information about the organization. e.g. list of available products
	Organization *Capabilitiesorganization `json:"organization,omitempty"`

	// ETag
	ETag *string `json:"eTag,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilities) SetField(field string, fieldValue interface{}) {
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

func (o Capabilities) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilities
	
	return json.Marshal(&struct { 
		Actions *map[string]Capabilitiesaction `json:"actions,omitempty"`
		
		DataTypes *map[string]Capabilitiesdatatype `json:"dataTypes,omitempty"`
		
		DependencyTypes *map[string]Capabilitiesdependencytype `json:"dependencyTypes,omitempty"`
		
		Expansions *map[string]interface{} `json:"expansions,omitempty"`
		
		Features *map[string]Capabilitiesavailability `json:"features,omitempty"`
		
		FeatureSupport *Capabilitiesfeaturesupport `json:"featureSupport,omitempty"`
		
		FlowTypes *map[string]Capabilitiesflowtype `json:"flowTypes,omitempty"`
		
		Functions *map[string]Capabilitiesfunction `json:"functions,omitempty"`
		
		Languages *map[string]Capabilitieslanguage `json:"languages,omitempty"`
		
		Organization *Capabilitiesorganization `json:"organization,omitempty"`
		
		ETag *string `json:"eTag,omitempty"`
		Alias
	}{ 
		Actions: o.Actions,
		
		DataTypes: o.DataTypes,
		
		DependencyTypes: o.DependencyTypes,
		
		Expansions: o.Expansions,
		
		Features: o.Features,
		
		FeatureSupport: o.FeatureSupport,
		
		FlowTypes: o.FlowTypes,
		
		Functions: o.Functions,
		
		Languages: o.Languages,
		
		Organization: o.Organization,
		
		ETag: o.ETag,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilities) UnmarshalJSON(b []byte) error {
	var CapabilitiesMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesMap)
	if err != nil {
		return err
	}
	
	if Actions, ok := CapabilitiesMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if DataTypes, ok := CapabilitiesMap["dataTypes"].(map[string]interface{}); ok {
		DataTypesString, _ := json.Marshal(DataTypes)
		json.Unmarshal(DataTypesString, &o.DataTypes)
	}
	
	if DependencyTypes, ok := CapabilitiesMap["dependencyTypes"].(map[string]interface{}); ok {
		DependencyTypesString, _ := json.Marshal(DependencyTypes)
		json.Unmarshal(DependencyTypesString, &o.DependencyTypes)
	}
	
	if Expansions, ok := CapabilitiesMap["expansions"].(map[string]interface{}); ok {
		ExpansionsString, _ := json.Marshal(Expansions)
		json.Unmarshal(ExpansionsString, &o.Expansions)
	}
	
	if Features, ok := CapabilitiesMap["features"].(map[string]interface{}); ok {
		FeaturesString, _ := json.Marshal(Features)
		json.Unmarshal(FeaturesString, &o.Features)
	}
	
	if FeatureSupport, ok := CapabilitiesMap["featureSupport"].(map[string]interface{}); ok {
		FeatureSupportString, _ := json.Marshal(FeatureSupport)
		json.Unmarshal(FeatureSupportString, &o.FeatureSupport)
	}
	
	if FlowTypes, ok := CapabilitiesMap["flowTypes"].(map[string]interface{}); ok {
		FlowTypesString, _ := json.Marshal(FlowTypes)
		json.Unmarshal(FlowTypesString, &o.FlowTypes)
	}
	
	if Functions, ok := CapabilitiesMap["functions"].(map[string]interface{}); ok {
		FunctionsString, _ := json.Marshal(Functions)
		json.Unmarshal(FunctionsString, &o.Functions)
	}
	
	if Languages, ok := CapabilitiesMap["languages"].(map[string]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if Organization, ok := CapabilitiesMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if ETag, ok := CapabilitiesMap["eTag"].(string); ok {
		o.ETag = &ETag
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilities) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
