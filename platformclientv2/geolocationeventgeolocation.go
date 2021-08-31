package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Geolocationeventgeolocation
type Geolocationeventgeolocation struct { 
	// UserId
	UserId *string `json:"userId,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Country
	Country *string `json:"country,omitempty"`


	// Region
	Region *string `json:"region,omitempty"`


	// City
	City *string `json:"city,omitempty"`

}

func (o *Geolocationeventgeolocation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Geolocationeventgeolocation
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Country *string `json:"country,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		City *string `json:"city,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		VarType: o.VarType,
		
		Country: o.Country,
		
		Region: o.Region,
		
		City: o.City,
		Alias:    (*Alias)(o),
	})
}

func (o *Geolocationeventgeolocation) UnmarshalJSON(b []byte) error {
	var GeolocationeventgeolocationMap map[string]interface{}
	err := json.Unmarshal(b, &GeolocationeventgeolocationMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := GeolocationeventgeolocationMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if VarType, ok := GeolocationeventgeolocationMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Country, ok := GeolocationeventgeolocationMap["country"].(string); ok {
		o.Country = &Country
	}
	
	if Region, ok := GeolocationeventgeolocationMap["region"].(string); ok {
		o.Region = &Region
	}
	
	if City, ok := GeolocationeventgeolocationMap["city"].(string); ok {
		o.City = &City
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Geolocationeventgeolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
