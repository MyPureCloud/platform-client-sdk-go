package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Geolocationeventgeolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
