package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyappeventsnotificationgeolocation
type Journeyappeventsnotificationgeolocation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Country
	Country *string `json:"country,omitempty"`

	// CountryName
	CountryName *string `json:"countryName,omitempty"`

	// Latitude
	Latitude *float32 `json:"latitude,omitempty"`

	// Longitude
	Longitude *float32 `json:"longitude,omitempty"`

	// Locality
	Locality *string `json:"locality,omitempty"`

	// PostalCode
	PostalCode *string `json:"postalCode,omitempty"`

	// Region
	Region *string `json:"region,omitempty"`

	// RegionName
	RegionName *string `json:"regionName,omitempty"`

	// Timezone
	Timezone *string `json:"timezone,omitempty"`

	// Source
	Source *string `json:"source,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyappeventsnotificationgeolocation) SetField(field string, fieldValue interface{}) {
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

func (o Journeyappeventsnotificationgeolocation) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyappeventsnotificationgeolocation
	
	return json.Marshal(&struct { 
		Country *string `json:"country,omitempty"`
		
		CountryName *string `json:"countryName,omitempty"`
		
		Latitude *float32 `json:"latitude,omitempty"`
		
		Longitude *float32 `json:"longitude,omitempty"`
		
		Locality *string `json:"locality,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		RegionName *string `json:"regionName,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		
		Source *string `json:"source,omitempty"`
		Alias
	}{ 
		Country: o.Country,
		
		CountryName: o.CountryName,
		
		Latitude: o.Latitude,
		
		Longitude: o.Longitude,
		
		Locality: o.Locality,
		
		PostalCode: o.PostalCode,
		
		Region: o.Region,
		
		RegionName: o.RegionName,
		
		Timezone: o.Timezone,
		
		Source: o.Source,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyappeventsnotificationgeolocation) UnmarshalJSON(b []byte) error {
	var JourneyappeventsnotificationgeolocationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyappeventsnotificationgeolocationMap)
	if err != nil {
		return err
	}
	
	if Country, ok := JourneyappeventsnotificationgeolocationMap["country"].(string); ok {
		o.Country = &Country
	}
    
	if CountryName, ok := JourneyappeventsnotificationgeolocationMap["countryName"].(string); ok {
		o.CountryName = &CountryName
	}
    
	if Latitude, ok := JourneyappeventsnotificationgeolocationMap["latitude"].(float64); ok {
		LatitudeFloat32 := float32(Latitude)
		o.Latitude = &LatitudeFloat32
	}
    
	if Longitude, ok := JourneyappeventsnotificationgeolocationMap["longitude"].(float64); ok {
		LongitudeFloat32 := float32(Longitude)
		o.Longitude = &LongitudeFloat32
	}
    
	if Locality, ok := JourneyappeventsnotificationgeolocationMap["locality"].(string); ok {
		o.Locality = &Locality
	}
    
	if PostalCode, ok := JourneyappeventsnotificationgeolocationMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if Region, ok := JourneyappeventsnotificationgeolocationMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if RegionName, ok := JourneyappeventsnotificationgeolocationMap["regionName"].(string); ok {
		o.RegionName = &RegionName
	}
    
	if Timezone, ok := JourneyappeventsnotificationgeolocationMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
    
	if Source, ok := JourneyappeventsnotificationgeolocationMap["source"].(string); ok {
		o.Source = &Source
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyappeventsnotificationgeolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
