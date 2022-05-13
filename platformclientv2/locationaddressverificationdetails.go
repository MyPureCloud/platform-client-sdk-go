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

func (o *Locationaddressverificationdetails) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationaddressverificationdetails
	
	DateFinished := new(string)
	if o.DateFinished != nil {
		
		*DateFinished = timeutil.Strftime(o.DateFinished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateFinished = nil
	}
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Status: o.Status,
		
		DateFinished: DateFinished,
		
		DateStarted: DateStarted,
		
		Service: o.Service,
		Alias:    (*Alias)(o),
	})
}

func (o *Locationaddressverificationdetails) UnmarshalJSON(b []byte) error {
	var LocationaddressverificationdetailsMap map[string]interface{}
	err := json.Unmarshal(b, &LocationaddressverificationdetailsMap)
	if err != nil {
		return err
	}
	
	if Status, ok := LocationaddressverificationdetailsMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if dateFinishedString, ok := LocationaddressverificationdetailsMap["dateFinished"].(string); ok {
		DateFinished, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateFinishedString)
		o.DateFinished = &DateFinished
	}
	
	if dateStartedString, ok := LocationaddressverificationdetailsMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if Service, ok := LocationaddressverificationdetailsMap["service"].(string); ok {
		o.Service = &Service
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Locationaddressverificationdetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
