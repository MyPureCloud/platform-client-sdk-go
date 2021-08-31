package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Daterangewithoptionalend
type Daterangewithoptionalend struct { 
	// StartBusinessUnitDate - The start date for work plan rotation or an agent, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartBusinessUnitDate *time.Time `json:"startBusinessUnitDate,omitempty"`


	// EndBusinessUnitDate - The end date for work plan rotation or an agent, interpreted in the business unit's time zone. Null denotes open ended date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EndBusinessUnitDate *time.Time `json:"endBusinessUnitDate,omitempty"`

}

func (o *Daterangewithoptionalend) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Daterangewithoptionalend
	
	StartBusinessUnitDate := new(string)
	if o.StartBusinessUnitDate != nil {
		*StartBusinessUnitDate = timeutil.Strftime(o.StartBusinessUnitDate, "%Y-%m-%d")
	} else {
		StartBusinessUnitDate = nil
	}
	
	EndBusinessUnitDate := new(string)
	if o.EndBusinessUnitDate != nil {
		*EndBusinessUnitDate = timeutil.Strftime(o.EndBusinessUnitDate, "%Y-%m-%d")
	} else {
		EndBusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		StartBusinessUnitDate *string `json:"startBusinessUnitDate,omitempty"`
		
		EndBusinessUnitDate *string `json:"endBusinessUnitDate,omitempty"`
		*Alias
	}{ 
		StartBusinessUnitDate: StartBusinessUnitDate,
		
		EndBusinessUnitDate: EndBusinessUnitDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Daterangewithoptionalend) UnmarshalJSON(b []byte) error {
	var DaterangewithoptionalendMap map[string]interface{}
	err := json.Unmarshal(b, &DaterangewithoptionalendMap)
	if err != nil {
		return err
	}
	
	if startBusinessUnitDateString, ok := DaterangewithoptionalendMap["startBusinessUnitDate"].(string); ok {
		StartBusinessUnitDate, _ := time.Parse("2006-01-02", startBusinessUnitDateString)
		o.StartBusinessUnitDate = &StartBusinessUnitDate
	}
	
	if endBusinessUnitDateString, ok := DaterangewithoptionalendMap["endBusinessUnitDate"].(string); ok {
		EndBusinessUnitDate, _ := time.Parse("2006-01-02", endBusinessUnitDateString)
		o.EndBusinessUnitDate = &EndBusinessUnitDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Daterangewithoptionalend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
