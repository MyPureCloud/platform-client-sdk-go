package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shorttermforecastreference - A pointer to a short term forecast
type Shorttermforecastreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - The description of the short term forecast
	Description *string `json:"description,omitempty"`

}

func (u *Shorttermforecastreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shorttermforecastreference

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		SelfUri: u.SelfUri,
		
		WeekDate: u.WeekDate,
		
		Description: u.Description,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shorttermforecastreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
