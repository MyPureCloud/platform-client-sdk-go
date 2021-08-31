package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Geolocationsettings
type Geolocationsettings struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// MapboxKey
	MapboxKey *string `json:"mapboxKey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Geolocationsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Geolocationsettings
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MapboxKey *string `json:"mapboxKey,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Enabled: o.Enabled,
		
		MapboxKey: o.MapboxKey,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Geolocationsettings) UnmarshalJSON(b []byte) error {
	var GeolocationsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &GeolocationsettingsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GeolocationsettingsMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := GeolocationsettingsMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Enabled, ok := GeolocationsettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if MapboxKey, ok := GeolocationsettingsMap["mapboxKey"].(string); ok {
		o.MapboxKey = &MapboxKey
	}
	
	if SelfUri, ok := GeolocationsettingsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Geolocationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
