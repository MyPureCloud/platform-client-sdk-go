package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fulldaytimeoffmarker
type Fulldaytimeoffmarker struct { 
	// BusinessUnitDate - The date of the time off marker, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitDate *time.Time `json:"businessUnitDate,omitempty"`


	// LengthMinutes - The length of the time off marker in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Description - The description of the time off marker
	Description *string `json:"description,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with the time off marker
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Paid - Whether the time off marker is paid
	Paid *bool `json:"paid,omitempty"`


	// TimeOffRequestId - The ID of the time off request
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

}

func (o *Fulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fulldaytimeoffmarker
	
	BusinessUnitDate := new(string)
	if o.BusinessUnitDate != nil {
		*BusinessUnitDate = timeutil.Strftime(o.BusinessUnitDate, "%Y-%m-%d")
	} else {
		BusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		BusinessUnitDate *string `json:"businessUnitDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		*Alias
	}{ 
		BusinessUnitDate: BusinessUnitDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Paid: o.Paid,
		
		TimeOffRequestId: o.TimeOffRequestId,
		Alias:    (*Alias)(o),
	})
}

func (o *Fulldaytimeoffmarker) UnmarshalJSON(b []byte) error {
	var FulldaytimeoffmarkerMap map[string]interface{}
	err := json.Unmarshal(b, &FulldaytimeoffmarkerMap)
	if err != nil {
		return err
	}
	
	if businessUnitDateString, ok := FulldaytimeoffmarkerMap["businessUnitDate"].(string); ok {
		BusinessUnitDate, _ := time.Parse("2006-01-02", businessUnitDateString)
		o.BusinessUnitDate = &BusinessUnitDate
	}
	
	if LengthMinutes, ok := FulldaytimeoffmarkerMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Description, ok := FulldaytimeoffmarkerMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if ActivityCodeId, ok := FulldaytimeoffmarkerMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
	
	if Paid, ok := FulldaytimeoffmarkerMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
	
	if TimeOffRequestId, ok := FulldaytimeoffmarkerMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Fulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
