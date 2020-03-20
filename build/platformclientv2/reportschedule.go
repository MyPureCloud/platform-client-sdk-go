package platformclientv2
import (
	"time"
	"encoding/json"
)

// Reportschedule
type Reportschedule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// QuartzCronExpression - Quartz Cron Expression
	QuartzCronExpression *string `json:"quartzCronExpression,omitempty"`


	// NextFireTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	NextFireTime *time.Time `json:"nextFireTime,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// TimePeriod
	TimePeriod *string `json:"timePeriod,omitempty"`


	// Interval - Interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// ReportFormat
	ReportFormat *string `json:"reportFormat,omitempty"`


	// Locale
	Locale *string `json:"locale,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// ReportId - Report ID
	ReportId *string `json:"reportId,omitempty"`


	// Parameters
	Parameters *map[string]interface{} `json:"parameters,omitempty"`


	// LastRun
	LastRun *Reportrunentry `json:"lastRun,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportschedule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
