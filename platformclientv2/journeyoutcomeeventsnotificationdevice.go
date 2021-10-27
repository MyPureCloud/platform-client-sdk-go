package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationdevice
type Journeyoutcomeeventsnotificationdevice struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// IsMobile
	IsMobile *bool `json:"isMobile,omitempty"`


	// ScreenHeight
	ScreenHeight *int `json:"screenHeight,omitempty"`


	// ScreenWidth
	ScreenWidth *int `json:"screenWidth,omitempty"`


	// Fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`


	// OsFamily
	OsFamily *string `json:"osFamily,omitempty"`


	// OsVersion
	OsVersion *string `json:"osVersion,omitempty"`


	// Category
	Category *string `json:"category,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationdevice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationdevice
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		IsMobile *bool `json:"isMobile,omitempty"`
		
		ScreenHeight *int `json:"screenHeight,omitempty"`
		
		ScreenWidth *int `json:"screenWidth,omitempty"`
		
		Fingerprint *string `json:"fingerprint,omitempty"`
		
		OsFamily *string `json:"osFamily,omitempty"`
		
		OsVersion *string `json:"osVersion,omitempty"`
		
		Category *string `json:"category,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		IsMobile: o.IsMobile,
		
		ScreenHeight: o.ScreenHeight,
		
		ScreenWidth: o.ScreenWidth,
		
		Fingerprint: o.Fingerprint,
		
		OsFamily: o.OsFamily,
		
		OsVersion: o.OsVersion,
		
		Category: o.Category,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationdevice) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationdeviceMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationdeviceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneyoutcomeeventsnotificationdeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if IsMobile, ok := JourneyoutcomeeventsnotificationdeviceMap["isMobile"].(bool); ok {
		o.IsMobile = &IsMobile
	}
	
	if ScreenHeight, ok := JourneyoutcomeeventsnotificationdeviceMap["screenHeight"].(float64); ok {
		ScreenHeightInt := int(ScreenHeight)
		o.ScreenHeight = &ScreenHeightInt
	}
	
	if ScreenWidth, ok := JourneyoutcomeeventsnotificationdeviceMap["screenWidth"].(float64); ok {
		ScreenWidthInt := int(ScreenWidth)
		o.ScreenWidth = &ScreenWidthInt
	}
	
	if Fingerprint, ok := JourneyoutcomeeventsnotificationdeviceMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
	
	if OsFamily, ok := JourneyoutcomeeventsnotificationdeviceMap["osFamily"].(string); ok {
		o.OsFamily = &OsFamily
	}
	
	if OsVersion, ok := JourneyoutcomeeventsnotificationdeviceMap["osVersion"].(string); ok {
		o.OsVersion = &OsVersion
	}
	
	if Category, ok := JourneyoutcomeeventsnotificationdeviceMap["category"].(string); ok {
		o.Category = &Category
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
