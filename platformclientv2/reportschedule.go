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

func (o *Reportschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportschedule
	
	NextFireTime := new(string)
	if o.NextFireTime != nil {
		
		*NextFireTime = timeutil.Strftime(o.NextFireTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		NextFireTime = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		QuartzCronExpression: o.QuartzCronExpression,
		
		NextFireTime: NextFireTime,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Description: o.Description,
		
		TimeZone: o.TimeZone,
		
		TimePeriod: o.TimePeriod,
		
		Interval: o.Interval,
		
		ReportFormat: o.ReportFormat,
		
		Locale: o.Locale,
		
		Enabled: o.Enabled,
		
		ReportId: o.ReportId,
		
		Parameters: o.Parameters,
		
		LastRun: o.LastRun,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportschedule) UnmarshalJSON(b []byte) error {
	var ReportscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &ReportscheduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ReportscheduleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ReportscheduleMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if QuartzCronExpression, ok := ReportscheduleMap["quartzCronExpression"].(string); ok {
		o.QuartzCronExpression = &QuartzCronExpression
	}
	
	if nextFireTimeString, ok := ReportscheduleMap["nextFireTime"].(string); ok {
		NextFireTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", nextFireTimeString)
		o.NextFireTime = &NextFireTime
	}
	
	if dateCreatedString, ok := ReportscheduleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := ReportscheduleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Description, ok := ReportscheduleMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if TimeZone, ok := ReportscheduleMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if TimePeriod, ok := ReportscheduleMap["timePeriod"].(string); ok {
		o.TimePeriod = &TimePeriod
	}
	
	if Interval, ok := ReportscheduleMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if ReportFormat, ok := ReportscheduleMap["reportFormat"].(string); ok {
		o.ReportFormat = &ReportFormat
	}
	
	if Locale, ok := ReportscheduleMap["locale"].(string); ok {
		o.Locale = &Locale
	}
	
	if Enabled, ok := ReportscheduleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if ReportId, ok := ReportscheduleMap["reportId"].(string); ok {
		o.ReportId = &ReportId
	}
	
	if Parameters, ok := ReportscheduleMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if LastRun, ok := ReportscheduleMap["lastRun"].(map[string]interface{}); ok {
		LastRunString, _ := json.Marshal(LastRun)
		json.Unmarshal(LastRunString, &o.LastRun)
	}
	
	if SelfUri, ok := ReportscheduleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
