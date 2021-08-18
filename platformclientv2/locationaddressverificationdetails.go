package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationaddressverificationdetails
type Locationaddressverificationdetails struct { 
	// Status - Status of address verification process
	Status *string `json:"status,omitempty"`


	// DateFinished - Finished time of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateFinished *time.Time `json:"dateFinished,omitempty"`


	// DateStarted - Time started of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`


	// Service - Third party service used for address verification
	Service *string `json:"service,omitempty"`

}

func (u *Locationaddressverificationdetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationaddressverificationdetails

	
	DateFinished := new(string)
	if u.DateFinished != nil {
		
		*DateFinished = timeutil.Strftime(u.DateFinished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateFinished = nil
	}
	
	DateStarted := new(string)
	if u.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(u.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		DateFinished *string `json:"dateFinished,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		Service *string `json:"service,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		
		DateFinished: DateFinished,
		
		DateStarted: DateStarted,
		
		Service: u.Service,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Locationaddressverificationdetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
