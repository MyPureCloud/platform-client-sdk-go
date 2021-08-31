package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Faxsummary
type Faxsummary struct { 
	// ReadCount
	ReadCount *int `json:"readCount,omitempty"`


	// UnreadCount
	UnreadCount *int `json:"unreadCount,omitempty"`


	// TotalCount
	TotalCount *int `json:"totalCount,omitempty"`

}

func (o *Faxsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxsummary
	
	return json.Marshal(&struct { 
		ReadCount *int `json:"readCount,omitempty"`
		
		UnreadCount *int `json:"unreadCount,omitempty"`
		
		TotalCount *int `json:"totalCount,omitempty"`
		*Alias
	}{ 
		ReadCount: o.ReadCount,
		
		UnreadCount: o.UnreadCount,
		
		TotalCount: o.TotalCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Faxsummary) UnmarshalJSON(b []byte) error {
	var FaxsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &FaxsummaryMap)
	if err != nil {
		return err
	}
	
	if ReadCount, ok := FaxsummaryMap["readCount"].(float64); ok {
		ReadCountInt := int(ReadCount)
		o.ReadCount = &ReadCountInt
	}
	
	if UnreadCount, ok := FaxsummaryMap["unreadCount"].(float64); ok {
		UnreadCountInt := int(UnreadCount)
		o.UnreadCount = &UnreadCountInt
	}
	
	if TotalCount, ok := FaxsummaryMap["totalCount"].(float64); ok {
		TotalCountInt := int(TotalCount)
		o.TotalCount = &TotalCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Faxsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
