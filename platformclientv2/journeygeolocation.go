package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeygeolocation
type Journeygeolocation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Country - Geolocation's ISO 3166-1 alpha-2 country code.
	Country *string `json:"country,omitempty"`

	// CountryName - Geolocation's country name.
	CountryName *string `json:"countryName,omitempty"`

	// Latitude - Geolocation's latitude.
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude - Geolocation's longitude.
	Longitude *float64 `json:"longitude,omitempty"`

	// Locality - Geolocation's locality or city.
	Locality *string `json:"locality,omitempty"`

	// PostalCode - Geolocation's postal code or ZIP code.
	PostalCode *string `json:"postalCode,omitempty"`

	// Region - Geolocation's ISO-3166-2 region code.
	Region *string `json:"region,omitempty"`

	// RegionName - Geolocation's region name.
	RegionName *string `json:"regionName,omitempty"`

	// Source - The source that was used to determine the geolocation information.
	Source *string `json:"source,omitempty"`

	// Timezone - Geolocation's timezone.
	Timezone *string `json:"timezone,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeygeolocation) SetField(field string, fieldValue interface{}) {
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

func (o Journeygeolocation) MarshalJSON() ([]byte, error) {
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
	type Alias Journeygeolocation
	
	return json.Marshal(&struct { 
		Country *string `json:"country,omitempty"`
		
		CountryName *string `json:"countryName,omitempty"`
		
		Latitude *float64 `json:"latitude,omitempty"`
		
		Longitude *float64 `json:"longitude,omitempty"`
		
		Locality *string `json:"locality,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		RegionName *string `json:"regionName,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
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
		
		Source: o.Source,
		
		Timezone: o.Timezone,
		Alias:    (Alias)(o),
	})
}

func (o *Journeygeolocation) UnmarshalJSON(b []byte) error {
	var JourneygeolocationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneygeolocationMap)
	if err != nil {
		return err
	}
	
	if Country, ok := JourneygeolocationMap["country"].(string); ok {
		o.Country = &Country
	}
    
	if CountryName, ok := JourneygeolocationMap["countryName"].(string); ok {
		o.CountryName = &CountryName
	}
    
	if Latitude, ok := JourneygeolocationMap["latitude"].(float64); ok {
		o.Latitude = &Latitude
	}
    
	if Longitude, ok := JourneygeolocationMap["longitude"].(float64); ok {
		o.Longitude = &Longitude
	}
    
	if Locality, ok := JourneygeolocationMap["locality"].(string); ok {
		o.Locality = &Locality
	}
    
	if PostalCode, ok := JourneygeolocationMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if Region, ok := JourneygeolocationMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if RegionName, ok := JourneygeolocationMap["regionName"].(string); ok {
		o.RegionName = &RegionName
	}
    
	if Source, ok := JourneygeolocationMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Timezone, ok := JourneygeolocationMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeygeolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
