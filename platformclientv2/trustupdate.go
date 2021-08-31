package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustupdate
type Trustupdate struct { 
	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`

}

func (o *Trustupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustupdate
	
	DateExpired := new(string)
	if o.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(o.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		DateExpired: DateExpired,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustupdate) UnmarshalJSON(b []byte) error {
	var TrustupdateMap map[string]interface{}
	err := json.Unmarshal(b, &TrustupdateMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := TrustupdateMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if dateExpiredString, ok := TrustupdateMap["dateExpired"].(string); ok {
		DateExpired, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiredString)
		o.DateExpired = &DateExpired
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
