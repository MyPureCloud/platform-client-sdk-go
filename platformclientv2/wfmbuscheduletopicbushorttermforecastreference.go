package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbushorttermforecastreference
type Wfmbuscheduletopicbushorttermforecastreference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

func (o *Wfmbuscheduletopicbushorttermforecastreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbushorttermforecastreference
	
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

func (o *Wfmbuscheduletopicbushorttermforecastreference) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicbushorttermforecastreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicbushorttermforecastreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduletopicbushorttermforecastreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if WeekDate, ok := WfmbuscheduletopicbushorttermforecastreferenceMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
	
	if Description, ok := WfmbuscheduletopicbushorttermforecastreferenceMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbushorttermforecastreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
