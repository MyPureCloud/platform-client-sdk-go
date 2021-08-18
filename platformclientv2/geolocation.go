package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Geolocation
type Geolocation struct { 
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

func (u *Geolocation) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		Primary: u.Primary,
		
		Latitude: u.Latitude,
		
		Longitude: u.Longitude,
		
		Country: u.Country,
		
		Region: u.Region,
		
		City: u.City,
		
		Locations: u.Locations,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Geolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
