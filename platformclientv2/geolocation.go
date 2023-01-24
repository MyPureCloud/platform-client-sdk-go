package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Geolocation
type Geolocation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// VarType - A string used to describe the type of client the geolocation is being updated from e.g. ios, android, web, etc.
	VarType *string `json:"type,omitempty"`

	// Primary - A boolean used to tell whether or not to set this geolocation client as the primary on a PATCH
	Primary *bool `json:"primary,omitempty"`

	// Latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude
	Longitude *float64 `json:"longitude,omitempty"`

	// Country
	Country *string `json:"country,omitempty"`

	// Region
	Region *string `json:"region,omitempty"`

	// City
	City *string `json:"city,omitempty"`

	// Locations
	Locations *[]Locationdefinition `json:"locations,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Geolocation) SetField(field string, fieldValue interface{}) {
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

func (o Geolocation) MarshalJSON() ([]byte, error) {
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
	type Alias Geolocation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		Latitude *float64 `json:"latitude,omitempty"`
		
		Longitude *float64 `json:"longitude,omitempty"`
		
		Country *string `json:"country,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		Locations *[]Locationdefinition `json:"locations,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Primary: o.Primary,
		
		Latitude: o.Latitude,
		
		Longitude: o.Longitude,
		
		Country: o.Country,
		
		Region: o.Region,
		
		City: o.City,
		
		Locations: o.Locations,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Geolocation) UnmarshalJSON(b []byte) error {
	var GeolocationMap map[string]interface{}
	err := json.Unmarshal(b, &GeolocationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GeolocationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := GeolocationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := GeolocationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Primary, ok := GeolocationMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
    
	if Latitude, ok := GeolocationMap["latitude"].(float64); ok {
		o.Latitude = &Latitude
	}
    
	if Longitude, ok := GeolocationMap["longitude"].(float64); ok {
		o.Longitude = &Longitude
	}
    
	if Country, ok := GeolocationMap["country"].(string); ok {
		o.Country = &Country
	}
    
	if Region, ok := GeolocationMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if City, ok := GeolocationMap["city"].(string); ok {
		o.City = &City
	}
    
	if Locations, ok := GeolocationMap["locations"].([]interface{}); ok {
		LocationsString, _ := json.Marshal(Locations)
		json.Unmarshal(LocationsString, &o.Locations)
	}
	
	if SelfUri, ok := GeolocationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Geolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
