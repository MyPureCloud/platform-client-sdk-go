package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Logcaptureuserconfiguration
type Logcaptureuserconfiguration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// DateExpired - Indicates when log capture will be turned off for the user. (Must be within 24 hours). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Logcaptureuserconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Logcaptureuserconfiguration
	
	DateExpired := new(string)
	if o.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(o.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DateExpired: DateExpired,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Logcaptureuserconfiguration) UnmarshalJSON(b []byte) error {
	var LogcaptureuserconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &LogcaptureuserconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LogcaptureuserconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if dateExpiredString, ok := LogcaptureuserconfigurationMap["dateExpired"].(string); ok {
		DateExpired, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiredString)
		o.DateExpired = &DateExpired
	}
	
	if SelfUri, ok := LogcaptureuserconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Logcaptureuserconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
