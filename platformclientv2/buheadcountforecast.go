package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buheadcountforecast
type Buheadcountforecast struct { 
	// Entities
	Entities *[]Buplanninggroupheadcountforecast `json:"entities,omitempty"`


	// ReferenceStartDate - Reference start date for the interval values in each forecast entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

}

func (o *Buheadcountforecast) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buheadcountforecast
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	
	return json.Marshal(&struct { 
		Entities *[]Buplanninggroupheadcountforecast `json:"entities,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		ReferenceStartDate: ReferenceStartDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Buheadcountforecast) UnmarshalJSON(b []byte) error {
	var BuheadcountforecastMap map[string]interface{}
	err := json.Unmarshal(b, &BuheadcountforecastMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BuheadcountforecastMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if referenceStartDateString, ok := BuheadcountforecastMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buheadcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
