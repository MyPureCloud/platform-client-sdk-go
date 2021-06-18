package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Comparisonperiod) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
