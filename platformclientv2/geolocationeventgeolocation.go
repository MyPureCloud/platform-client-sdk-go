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

func (u *Geolocationeventgeolocation) MarshalJSON() ([]byte, error) {
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
		UserId: u.UserId,
		
		VarType: u.VarType,
		
		Country: u.Country,
		
		Region: u.Region,
		
		City: u.City,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Geolocationeventgeolocation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
