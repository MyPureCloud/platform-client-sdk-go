package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportschedule
type Reportschedule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// QuartzCronExpression - Quartz Cron Expression
	QuartzCronExpression *string `json:"quartzCronExpression,omitempty"`


	// NextFireTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	NextFireTime *time.Time `json:"nextFireTime,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

func (u *Reportschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportschedule

	
	NextFireTime := new(string)
	if u.NextFireTime != nil {
		
		*NextFireTime = timeutil.Strftime(u.NextFireTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		NextFireTime = nil
	}
	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		QuartzCronExpression *string `json:"quartzCronExpression,omitempty"`
		
		NextFireTime *string `json:"nextFireTime,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		TimePeriod *string `json:"timePeriod,omitempty"`
		
		Interval *string `json:"interval,omitempty"`
		
		ReportFormat *string `json:"reportFormat,omitempty"`
		
		Locale *string `json:"locale,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		ReportId *string `json:"reportId,omitempty"`
		
		Parameters *map[string]interface{} `json:"parameters,omitempty"`
		
		LastRun *Reportrunentry `json:"lastRun,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		QuartzCronExpression: u.QuartzCronExpression,
		
		NextFireTime: NextFireTime,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Description: u.Description,
		
		TimeZone: u.TimeZone,
		
		TimePeriod: u.TimePeriod,
		
		Interval: u.Interval,
		
		ReportFormat: u.ReportFormat,
		
		Locale: u.Locale,
		
		Enabled: u.Enabled,
		
		ReportId: u.ReportId,
		
		Parameters: u.Parameters,
		
		LastRun: u.LastRun,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
