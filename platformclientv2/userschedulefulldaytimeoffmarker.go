package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userschedulefulldaytimeoffmarker
type Userschedulefulldaytimeoffmarker struct { 
	// ManagementUnitDate - The date associated with the time off request that this marker corresponds to.  Date only, in ISO-8601 format.
	ManagementUnitDate *string `json:"managementUnitDate,omitempty"`


	// ActivityCodeId - The id for the activity code.  Look up a map of activity codes with the activities route
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// IsPaid - Whether this is paid time off
	IsPaid *bool `json:"isPaid,omitempty"`


	// LengthInMinutes - The length in minutes of this time off marker
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Description - The description associated with the time off request that this marker corresponds to
	Description *string `json:"description,omitempty"`


	// Delete - If marked true for updating an existing full day time off marker, it will be deleted
	Delete *bool `json:"delete,omitempty"`

}

func (o *Userschedulefulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userschedulefulldaytimeoffmarker
	
	return json.Marshal(&struct { 
		ManagementUnitDate *string `json:"managementUnitDate,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		IsPaid *bool `json:"isPaid,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		*Alias
	}{ 
		ManagementUnitDate: o.ManagementUnitDate,
		
		ActivityCodeId: o.ActivityCodeId,
		
		IsPaid: o.IsPaid,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Description: o.Description,
		
		Delete: o.Delete,
		Alias:    (*Alias)(o),
	})
}

func (o *Userschedulefulldaytimeoffmarker) UnmarshalJSON(b []byte) error {
	var UserschedulefulldaytimeoffmarkerMap map[string]interface{}
	err := json.Unmarshal(b, &UserschedulefulldaytimeoffmarkerMap)
	if err != nil {
		return err
	}
	
	if ManagementUnitDate, ok := UserschedulefulldaytimeoffmarkerMap["managementUnitDate"].(string); ok {
		o.ManagementUnitDate = &ManagementUnitDate
	}
    
	if ActivityCodeId, ok := UserschedulefulldaytimeoffmarkerMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if IsPaid, ok := UserschedulefulldaytimeoffmarkerMap["isPaid"].(bool); ok {
		o.IsPaid = &IsPaid
	}
    
	if LengthInMinutes, ok := UserschedulefulldaytimeoffmarkerMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Description, ok := UserschedulefulldaytimeoffmarkerMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Delete, ok := UserschedulefulldaytimeoffmarkerMap["delete"].(bool); ok {
		o.Delete = &Delete
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userschedulefulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
