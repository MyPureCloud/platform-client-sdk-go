package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsdataretentionresponse
type Analyticsdataretentionresponse struct { 
	// RetentionDays - Analytics data retention period in days for the organization.
	RetentionDays *int `json:"retentionDays,omitempty"`


	// DateCreated - Date and time when the analytics data retention was set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date and time when the analytics data retention was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

}

func (o *Analyticsdataretentionresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsdataretentionresponse
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		RetentionDays *int `json:"retentionDays,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		*Alias
	}{ 
		RetentionDays: o.RetentionDays,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsdataretentionresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsdataretentionresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsdataretentionresponseMap)
	if err != nil {
		return err
	}
	
	if RetentionDays, ok := AnalyticsdataretentionresponseMap["retentionDays"].(float64); ok {
		RetentionDaysInt := int(RetentionDays)
		o.RetentionDays = &RetentionDaysInt
	}
	
	if dateCreatedString, ok := AnalyticsdataretentionresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := AnalyticsdataretentionresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsdataretentionresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
