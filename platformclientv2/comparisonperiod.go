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


	// KpiResults - KPI results for each metric
	KpiResults *[]Kpiresult `json:"kpiResults,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Comparisonperiod) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Comparisonperiod
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateEnded := new(string)
	if o.DateEnded != nil {
		
		*DateEnded = timeutil.Strftime(o.DateEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnded = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		DateEnded *string `json:"dateEnded,omitempty"`
		
		KpiResults *[]Kpiresult `json:"kpiResults,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Kpi: o.Kpi,
		
		DateStarted: DateStarted,
		
		DateEnded: DateEnded,
		
		KpiResults: o.KpiResults,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Comparisonperiod) UnmarshalJSON(b []byte) error {
	var ComparisonperiodMap map[string]interface{}
	err := json.Unmarshal(b, &ComparisonperiodMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ComparisonperiodMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Kpi, ok := ComparisonperiodMap["kpi"].(string); ok {
		o.Kpi = &Kpi
	}
    
	if dateStartedString, ok := ComparisonperiodMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if dateEndedString, ok := ComparisonperiodMap["dateEnded"].(string); ok {
		DateEnded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndedString)
		o.DateEnded = &DateEnded
	}
	
	if KpiResults, ok := ComparisonperiodMap["kpiResults"].([]interface{}); ok {
		KpiResultsString, _ := json.Marshal(KpiResults)
		json.Unmarshal(KpiResultsString, &o.KpiResults)
	}
	
	if SelfUri, ok := ComparisonperiodMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Comparisonperiod) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
