package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesdatatype
type Capabilitiesdatatype struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FlowTypes
	FlowTypes *map[string]Capabilitiesavailability `json:"flowTypes,omitempty"`

	// AllowCollection
	AllowCollection *bool `json:"allowCollection,omitempty"`

	// Characteristics
	Characteristics *Capabilitiesdatatypecharacteristicsmodel `json:"characteristics,omitempty"`

	// Direction
	Direction *Capabilitiesdatatypedirectionbyobject `json:"direction,omitempty"`

	// Facets
	Facets *Capabilitiesdatatypefacetsmodel `json:"facets,omitempty"`

	// Id
	Id *string `json:"id,omitempty"`

	// IsAbstract
	IsAbstract *bool `json:"isAbstract,omitempty"`

	// IsAnAnyType
	IsAnAnyType *Capabilitiesavailability `json:"isAnAnyType,omitempty"`

	// IsComplex
	IsComplex *bool `json:"isComplex,omitempty"`

	// IsEnum
	IsEnum *bool `json:"isEnum,omitempty"`

	// IsEnumContainer
	IsEnumContainer *bool `json:"isEnumContainer,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Properties
	Properties *[]Capabilitiesdatatypeproperty `json:"properties,omitempty"`

	// UserCreatable
	UserCreatable *Capabilitiesavailability `json:"userCreatable,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesdatatype) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesdatatype) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesdatatype
	
	return json.Marshal(&struct { 
		FlowTypes *map[string]Capabilitiesavailability `json:"flowTypes,omitempty"`
		
		AllowCollection *bool `json:"allowCollection,omitempty"`
		
		Characteristics *Capabilitiesdatatypecharacteristicsmodel `json:"characteristics,omitempty"`
		
		Direction *Capabilitiesdatatypedirectionbyobject `json:"direction,omitempty"`
		
		Facets *Capabilitiesdatatypefacetsmodel `json:"facets,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IsAbstract *bool `json:"isAbstract,omitempty"`
		
		IsAnAnyType *Capabilitiesavailability `json:"isAnAnyType,omitempty"`
		
		IsComplex *bool `json:"isComplex,omitempty"`
		
		IsEnum *bool `json:"isEnum,omitempty"`
		
		IsEnumContainer *bool `json:"isEnumContainer,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Properties *[]Capabilitiesdatatypeproperty `json:"properties,omitempty"`
		
		UserCreatable *Capabilitiesavailability `json:"userCreatable,omitempty"`
		Alias
	}{ 
		FlowTypes: o.FlowTypes,
		
		AllowCollection: o.AllowCollection,
		
		Characteristics: o.Characteristics,
		
		Direction: o.Direction,
		
		Facets: o.Facets,
		
		Id: o.Id,
		
		IsAbstract: o.IsAbstract,
		
		IsAnAnyType: o.IsAnAnyType,
		
		IsComplex: o.IsComplex,
		
		IsEnum: o.IsEnum,
		
		IsEnumContainer: o.IsEnumContainer,
		
		Name: o.Name,
		
		Properties: o.Properties,
		
		UserCreatable: o.UserCreatable,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesdatatype) UnmarshalJSON(b []byte) error {
	var CapabilitiesdatatypeMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesdatatypeMap)
	if err != nil {
		return err
	}
	
	if FlowTypes, ok := CapabilitiesdatatypeMap["flowTypes"].(map[string]interface{}); ok {
		FlowTypesString, _ := json.Marshal(FlowTypes)
		json.Unmarshal(FlowTypesString, &o.FlowTypes)
	}
	
	if AllowCollection, ok := CapabilitiesdatatypeMap["allowCollection"].(bool); ok {
		o.AllowCollection = &AllowCollection
	}
    
	if Characteristics, ok := CapabilitiesdatatypeMap["characteristics"].(map[string]interface{}); ok {
		CharacteristicsString, _ := json.Marshal(Characteristics)
		json.Unmarshal(CharacteristicsString, &o.Characteristics)
	}
	
	if Direction, ok := CapabilitiesdatatypeMap["direction"].(map[string]interface{}); ok {
		DirectionString, _ := json.Marshal(Direction)
		json.Unmarshal(DirectionString, &o.Direction)
	}
	
	if Facets, ok := CapabilitiesdatatypeMap["facets"].(map[string]interface{}); ok {
		FacetsString, _ := json.Marshal(Facets)
		json.Unmarshal(FacetsString, &o.Facets)
	}
	
	if Id, ok := CapabilitiesdatatypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IsAbstract, ok := CapabilitiesdatatypeMap["isAbstract"].(bool); ok {
		o.IsAbstract = &IsAbstract
	}
    
	if IsAnAnyType, ok := CapabilitiesdatatypeMap["isAnAnyType"].(map[string]interface{}); ok {
		IsAnAnyTypeString, _ := json.Marshal(IsAnAnyType)
		json.Unmarshal(IsAnAnyTypeString, &o.IsAnAnyType)
	}
	
	if IsComplex, ok := CapabilitiesdatatypeMap["isComplex"].(bool); ok {
		o.IsComplex = &IsComplex
	}
    
	if IsEnum, ok := CapabilitiesdatatypeMap["isEnum"].(bool); ok {
		o.IsEnum = &IsEnum
	}
    
	if IsEnumContainer, ok := CapabilitiesdatatypeMap["isEnumContainer"].(bool); ok {
		o.IsEnumContainer = &IsEnumContainer
	}
    
	if Name, ok := CapabilitiesdatatypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Properties, ok := CapabilitiesdatatypeMap["properties"].([]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if UserCreatable, ok := CapabilitiesdatatypeMap["userCreatable"].(map[string]interface{}); ok {
		UserCreatableString, _ := json.Marshal(UserCreatable)
		json.Unmarshal(UserCreatableString, &o.UserCreatable)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
