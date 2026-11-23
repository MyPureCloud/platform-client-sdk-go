package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesdatatypeproperty
type Capabilitiesdatatypeproperty struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Availability
	Availability *Capabilitiesavailability `json:"availability,omitempty"`

	// DataTypeId
	DataTypeId *string `json:"dataTypeId,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// EnumGroup
	EnumGroup *string `json:"enumGroup,omitempty"`

	// Facets
	Facets *Capabilitiesdatatypefacetsmodel `json:"facets,omitempty"`

	// Id
	Id *string `json:"id,omitempty"`

	// IsCollection
	IsCollection *bool `json:"isCollection,omitempty"`

	// IsComputed
	IsComputed *bool `json:"isComputed,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesdatatypeproperty) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesdatatypeproperty) MarshalJSON() ([]byte, error) {
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
	type Alias Capabilitiesdatatypeproperty
	
	return json.Marshal(&struct { 
		Availability *Capabilitiesavailability `json:"availability,omitempty"`
		
		DataTypeId *string `json:"dataTypeId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		EnumGroup *string `json:"enumGroup,omitempty"`
		
		Facets *Capabilitiesdatatypefacetsmodel `json:"facets,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		IsCollection *bool `json:"isCollection,omitempty"`
		
		IsComputed *bool `json:"isComputed,omitempty"`
		
		Name *string `json:"name,omitempty"`
		Alias
	}{ 
		Availability: o.Availability,
		
		DataTypeId: o.DataTypeId,
		
		Description: o.Description,
		
		EnumGroup: o.EnumGroup,
		
		Facets: o.Facets,
		
		Id: o.Id,
		
		IsCollection: o.IsCollection,
		
		IsComputed: o.IsComputed,
		
		Name: o.Name,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesdatatypeproperty) UnmarshalJSON(b []byte) error {
	var CapabilitiesdatatypepropertyMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesdatatypepropertyMap)
	if err != nil {
		return err
	}
	
	if Availability, ok := CapabilitiesdatatypepropertyMap["availability"].(map[string]interface{}); ok {
		AvailabilityString, _ := json.Marshal(Availability)
		json.Unmarshal(AvailabilityString, &o.Availability)
	}
	
	if DataTypeId, ok := CapabilitiesdatatypepropertyMap["dataTypeId"].(string); ok {
		o.DataTypeId = &DataTypeId
	}
    
	if Description, ok := CapabilitiesdatatypepropertyMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if EnumGroup, ok := CapabilitiesdatatypepropertyMap["enumGroup"].(string); ok {
		o.EnumGroup = &EnumGroup
	}
    
	if Facets, ok := CapabilitiesdatatypepropertyMap["facets"].(map[string]interface{}); ok {
		FacetsString, _ := json.Marshal(Facets)
		json.Unmarshal(FacetsString, &o.Facets)
	}
	
	if Id, ok := CapabilitiesdatatypepropertyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IsCollection, ok := CapabilitiesdatatypepropertyMap["isCollection"].(bool); ok {
		o.IsCollection = &IsCollection
	}
    
	if IsComputed, ok := CapabilitiesdatatypepropertyMap["isComputed"].(bool); ok {
		o.IsComputed = &IsComputed
	}
    
	if Name, ok := CapabilitiesdatatypepropertyMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypeproperty) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
