package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentmanagementworkspacedocumentstopiclockdata
type Contentmanagementworkspacedocumentstopiclockdata struct { 
	// LockedBy
	LockedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"lockedBy,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpires
	DateExpires *time.Time `json:"dateExpires,omitempty"`

}

func (o *Contentmanagementworkspacedocumentstopiclockdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentmanagementworkspacedocumentstopiclockdata
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateExpires := new(string)
	if o.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(o.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpires = nil
	}
	
	return json.Marshal(&struct { 
		LockedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"lockedBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		*Alias
	}{ 
		LockedBy: o.LockedBy,
		
		DateCreated: DateCreated,
		
		DateExpires: DateExpires,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentmanagementworkspacedocumentstopiclockdata) UnmarshalJSON(b []byte) error {
	var ContentmanagementworkspacedocumentstopiclockdataMap map[string]interface{}
	err := json.Unmarshal(b, &ContentmanagementworkspacedocumentstopiclockdataMap)
	if err != nil {
		return err
	}
	
	if LockedBy, ok := ContentmanagementworkspacedocumentstopiclockdataMap["lockedBy"].(map[string]interface{}); ok {
		LockedByString, _ := json.Marshal(LockedBy)
		json.Unmarshal(LockedByString, &o.LockedBy)
	}
	
	if dateCreatedString, ok := ContentmanagementworkspacedocumentstopiclockdataMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateExpiresString, ok := ContentmanagementworkspacedocumentstopiclockdataMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentmanagementworkspacedocumentstopiclockdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
