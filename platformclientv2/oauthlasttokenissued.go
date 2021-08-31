package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthlasttokenissued
type Oauthlasttokenissued struct { 
	// DateIssued - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateIssued *time.Time `json:"dateIssued,omitempty"`

}

func (o *Oauthlasttokenissued) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Oauthlasttokenissued
	
	DateIssued := new(string)
	if o.DateIssued != nil {
		
		*DateIssued = timeutil.Strftime(o.DateIssued, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateIssued = nil
	}
	
	return json.Marshal(&struct { 
		DateIssued *string `json:"dateIssued,omitempty"`
		*Alias
	}{ 
		DateIssued: DateIssued,
		Alias:    (*Alias)(o),
	})
}

func (o *Oauthlasttokenissued) UnmarshalJSON(b []byte) error {
	var OauthlasttokenissuedMap map[string]interface{}
	err := json.Unmarshal(b, &OauthlasttokenissuedMap)
	if err != nil {
		return err
	}
	
	if dateIssuedString, ok := OauthlasttokenissuedMap["dateIssued"].(string); ok {
		DateIssued, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateIssuedString)
		o.DateIssued = &DateIssued
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Oauthlasttokenissued) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
