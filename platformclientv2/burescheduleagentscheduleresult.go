package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Burescheduleagentscheduleresult
type Burescheduleagentscheduleresult struct { 
	// ManagementUnit - The management unit to which this part of the result applies
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// DownloadResult - The agent schedules.  Result will always come via the downloadUrl; however the schema is included for documentation
	DownloadResult *Murescheduleresultwrapper `json:"downloadResult,omitempty"`


	// DownloadUrl - The download URL from which to fetch the result
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Burescheduleagentscheduleresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Burescheduleagentscheduleresult
	
	return json.Marshal(&struct { 
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		DownloadResult *Murescheduleresultwrapper `json:"downloadResult,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		ManagementUnit: o.ManagementUnit,
		
		DownloadResult: o.DownloadResult,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Burescheduleagentscheduleresult) UnmarshalJSON(b []byte) error {
	var BurescheduleagentscheduleresultMap map[string]interface{}
	err := json.Unmarshal(b, &BurescheduleagentscheduleresultMap)
	if err != nil {
		return err
	}
	
	if ManagementUnit, ok := BurescheduleagentscheduleresultMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if DownloadResult, ok := BurescheduleagentscheduleresultMap["downloadResult"].(map[string]interface{}); ok {
		DownloadResultString, _ := json.Marshal(DownloadResult)
		json.Unmarshal(DownloadResultString, &o.DownloadResult)
	}
	
	if DownloadUrl, ok := BurescheduleagentscheduleresultMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Burescheduleagentscheduleresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
