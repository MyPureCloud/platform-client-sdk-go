package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistorychangemetadata
type Buagentschedulehistorychangemetadata struct { 
	// DateModified - The timestamp of the schedule change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The user that made the schedule change
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

}

func (o *Buagentschedulehistorychangemetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistorychangemetadata
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		*Alias
	}{ 
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulehistorychangemetadata) UnmarshalJSON(b []byte) error {
	var BuagentschedulehistorychangemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulehistorychangemetadataMap)
	if err != nil {
		return err
	}
	
	if dateModifiedString, ok := BuagentschedulehistorychangemetadataMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := BuagentschedulehistorychangemetadataMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorychangemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
