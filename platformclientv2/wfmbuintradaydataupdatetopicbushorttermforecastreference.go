package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbushorttermforecastreference
type Wfmbuintradaydataupdatetopicbushorttermforecastreference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

func (o *Wfmbuintradaydataupdatetopicbushorttermforecastreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbushorttermforecastreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		WeekDate: o.WeekDate,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbushorttermforecastreference) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbushorttermforecastreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbushorttermforecastreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuintradaydataupdatetopicbushorttermforecastreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if WeekDate, ok := WfmbuintradaydataupdatetopicbushorttermforecastreferenceMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
	
	if Description, ok := WfmbuintradaydataupdatetopicbushorttermforecastreferenceMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbushorttermforecastreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
