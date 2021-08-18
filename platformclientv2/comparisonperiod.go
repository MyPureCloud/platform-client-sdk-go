package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Comparisonperiod
type Comparisonperiod struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Kpi - Key Performance Indicator optimised during the comparison period.
	Kpi *string `json:"kpi,omitempty"`


	// DateStarted - Start date of the comparison period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`


	// DateEnded - End date of the comparison period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnded *time.Time `json:"dateEnded,omitempty"`


	// KpiTotalOn - Absolute metric (in which the KPI is based) total for the interactions handled by predictive routing (GPR was on)
	KpiTotalOn *int `json:"kpiTotalOn,omitempty"`


	// KpiTotalOff - Absolute metric (in which the KPI is based) total for the interactions not routed by predictive routing (GPR was off)
	KpiTotalOff *int `json:"kpiTotalOff,omitempty"`


	// InteractionCountOn - Total interactions handled by predictive routing (GPR was on)
	InteractionCountOn *int `json:"interactionCountOn,omitempty"`


	// InteractionCountOff - Total interactions not routed by predictive routing (GPR was off)
	InteractionCountOff *int `json:"interactionCountOff,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Comparisonperiod) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Comparisonperiod

	
	DateStarted := new(string)
	if u.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(u.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateEnded := new(string)
	if u.DateEnded != nil {
		
		*DateEnded = timeutil.Strftime(u.DateEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnded = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		DateEnded *string `json:"dateEnded,omitempty"`
		
		KpiTotalOn *int `json:"kpiTotalOn,omitempty"`
		
		KpiTotalOff *int `json:"kpiTotalOff,omitempty"`
		
		InteractionCountOn *int `json:"interactionCountOn,omitempty"`
		
		InteractionCountOff *int `json:"interactionCountOff,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Kpi: u.Kpi,
		
		DateStarted: DateStarted,
		
		DateEnded: DateEnded,
		
		KpiTotalOn: u.KpiTotalOn,
		
		KpiTotalOff: u.KpiTotalOff,
		
		InteractionCountOn: u.InteractionCountOn,
		
		InteractionCountOff: u.InteractionCountOff,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Comparisonperiod) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
