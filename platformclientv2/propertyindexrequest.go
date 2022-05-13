package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Propertyindexrequest
type Propertyindexrequest struct { 
	// SessionId - Attach properties to a segment in the indicated session
	SessionId *string `json:"sessionId,omitempty"`


	// TargetDate - Attach properties to a segment covering a specific point in time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	TargetDate *time.Time `json:"targetDate,omitempty"`


	// Properties - The list of properties to index
	Properties *[]Analyticsproperty `json:"properties,omitempty"`

}

func (o *Propertyindexrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Propertyindexrequest
	
	TargetDate := new(string)
	if o.TargetDate != nil {
		
		*TargetDate = timeutil.Strftime(o.TargetDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TargetDate = nil
	}
	
	return json.Marshal(&struct { 
		SessionId *string `json:"sessionId,omitempty"`
		
		TargetDate *string `json:"targetDate,omitempty"`
		
		Properties *[]Analyticsproperty `json:"properties,omitempty"`
		*Alias
	}{ 
		SessionId: o.SessionId,
		
		TargetDate: TargetDate,
		
		Properties: o.Properties,
		Alias:    (*Alias)(o),
	})
}

func (o *Propertyindexrequest) UnmarshalJSON(b []byte) error {
	var PropertyindexrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PropertyindexrequestMap)
	if err != nil {
		return err
	}
	
	if SessionId, ok := PropertyindexrequestMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if targetDateString, ok := PropertyindexrequestMap["targetDate"].(string); ok {
		TargetDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", targetDateString)
		o.TargetDate = &TargetDate
	}
	
	if Properties, ok := PropertyindexrequestMap["properties"].([]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Propertyindexrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
