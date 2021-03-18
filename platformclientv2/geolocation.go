package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Geolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
