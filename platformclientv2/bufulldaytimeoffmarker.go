package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bufulldaytimeoffmarker
type Bufulldaytimeoffmarker struct { 
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

func (o *Bufulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bufulldaytimeoffmarker
	
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

func (o *Bufulldaytimeoffmarker) UnmarshalJSON(b []byte) error {
	var BufulldaytimeoffmarkerMap map[string]interface{}
	err := json.Unmarshal(b, &BufulldaytimeoffmarkerMap)
	if err != nil {
		return err
	}
	
	if businessUnitDateString, ok := BufulldaytimeoffmarkerMap["businessUnitDate"].(string); ok {
		BusinessUnitDate, _ := time.Parse("2006-01-02", businessUnitDateString)
		o.BusinessUnitDate = &BusinessUnitDate
	}
	
	if LengthMinutes, ok := BufulldaytimeoffmarkerMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Description, ok := BufulldaytimeoffmarkerMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCodeId, ok := BufulldaytimeoffmarkerMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Paid, ok := BufulldaytimeoffmarkerMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    
	if TimeOffRequestId, ok := BufulldaytimeoffmarkerMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bufulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
