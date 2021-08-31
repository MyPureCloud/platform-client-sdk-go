package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsuserdetailsasyncqueryresponse
type Analyticsuserdetailsasyncqueryresponse struct { 
	// UserDetails
	UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`


	// Cursor - Optional cursor to indicate where to resume the results
	Cursor *string `json:"cursor,omitempty"`


	// DataAvailabilityDate - Data available up to at least this datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DataAvailabilityDate *time.Time `json:"dataAvailabilityDate,omitempty"`

}

func (o *Analyticsuserdetailsasyncqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsuserdetailsasyncqueryresponse
	
	DataAvailabilityDate := new(string)
	if o.DataAvailabilityDate != nil {
		
		*DataAvailabilityDate = timeutil.Strftime(o.DataAvailabilityDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DataAvailabilityDate = nil
	}
	
	return json.Marshal(&struct { 
		UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`
		
		Cursor *string `json:"cursor,omitempty"`
		
		DataAvailabilityDate *string `json:"dataAvailabilityDate,omitempty"`
		*Alias
	}{ 
		UserDetails: o.UserDetails,
		
		Cursor: o.Cursor,
		
		DataAvailabilityDate: DataAvailabilityDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsuserdetailsasyncqueryresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsuserdetailsasyncqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsuserdetailsasyncqueryresponseMap)
	if err != nil {
		return err
	}
	
	if UserDetails, ok := AnalyticsuserdetailsasyncqueryresponseMap["userDetails"].([]interface{}); ok {
		UserDetailsString, _ := json.Marshal(UserDetails)
		json.Unmarshal(UserDetailsString, &o.UserDetails)
	}
	
	if Cursor, ok := AnalyticsuserdetailsasyncqueryresponseMap["cursor"].(string); ok {
		o.Cursor = &Cursor
	}
	
	if dataAvailabilityDateString, ok := AnalyticsuserdetailsasyncqueryresponseMap["dataAvailabilityDate"].(string); ok {
		DataAvailabilityDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", dataAvailabilityDateString)
		o.DataAvailabilityDate = &DataAvailabilityDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsasyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
