package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userscheduleactivity
type Userscheduleactivity struct { 
	// ActivityCodeId - The id for the activity code.  Look up a map of activity codes with the activities route
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// StartDate - Start time in UTC for this activity, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthInMinutes - Length in minutes for this activity
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Description - Description for this activity
	Description *string `json:"description,omitempty"`


	// CountsAsPaidTime - Whether this activity is paid
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// IsDstFallback - Whether this activity spans a DST fallback
	IsDstFallback *bool `json:"isDstFallback,omitempty"`


	// TimeOffRequestId - Time off request id of this activity
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

}

func (o *Userscheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userscheduleactivity
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		IsDstFallback *bool `json:"isDstFallback,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		StartDate: StartDate,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Description: o.Description,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		IsDstFallback: o.IsDstFallback,
		
		TimeOffRequestId: o.TimeOffRequestId,
		Alias:    (*Alias)(o),
	})
}

func (o *Userscheduleactivity) UnmarshalJSON(b []byte) error {
	var UserscheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &UserscheduleactivityMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := UserscheduleactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if startDateString, ok := UserscheduleactivityMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthInMinutes, ok := UserscheduleactivityMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Description, ok := UserscheduleactivityMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CountsAsPaidTime, ok := UserscheduleactivityMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if IsDstFallback, ok := UserscheduleactivityMap["isDstFallback"].(bool); ok {
		o.IsDstFallback = &IsDstFallback
	}
    
	if TimeOffRequestId, ok := UserscheduleactivityMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
