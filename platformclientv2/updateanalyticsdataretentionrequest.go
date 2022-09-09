package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateanalyticsdataretentionrequest
type Updateanalyticsdataretentionrequest struct { 
	// RetentionDays - Analytics data retention period in days to set for the organization.
	RetentionDays *int `json:"retentionDays,omitempty"`

}

func (o *Updateanalyticsdataretentionrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateanalyticsdataretentionrequest
	
	return json.Marshal(&struct { 
		RetentionDays *int `json:"retentionDays,omitempty"`
		*Alias
	}{ 
		RetentionDays: o.RetentionDays,
		Alias:    (*Alias)(o),
	})
}

func (o *Updateanalyticsdataretentionrequest) UnmarshalJSON(b []byte) error {
	var UpdateanalyticsdataretentionrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateanalyticsdataretentionrequestMap)
	if err != nil {
		return err
	}
	
	if RetentionDays, ok := UpdateanalyticsdataretentionrequestMap["retentionDays"].(float64); ok {
		RetentionDaysInt := int(RetentionDays)
		o.RetentionDays = &RetentionDaysInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateanalyticsdataretentionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
