package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationdevice
type Journeywebeventsnotificationdevice struct { 
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

func (o *Journeywebeventsnotificationdevice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationdevice
	
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

func (o *Journeywebeventsnotificationdevice) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationdeviceMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationdeviceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneywebeventsnotificationdeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if IsMobile, ok := JourneywebeventsnotificationdeviceMap["isMobile"].(bool); ok {
		o.IsMobile = &IsMobile
	}
	
	if ScreenHeight, ok := JourneywebeventsnotificationdeviceMap["screenHeight"].(float64); ok {
		ScreenHeightInt := int(ScreenHeight)
		o.ScreenHeight = &ScreenHeightInt
	}
	
	if ScreenWidth, ok := JourneywebeventsnotificationdeviceMap["screenWidth"].(float64); ok {
		ScreenWidthInt := int(ScreenWidth)
		o.ScreenWidth = &ScreenWidthInt
	}
	
	if Fingerprint, ok := JourneywebeventsnotificationdeviceMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
	
	if OsFamily, ok := JourneywebeventsnotificationdeviceMap["osFamily"].(string); ok {
		o.OsFamily = &OsFamily
	}
	
	if OsVersion, ok := JourneywebeventsnotificationdeviceMap["osVersion"].(string); ok {
		o.OsVersion = &OsVersion
	}
	
	if Category, ok := JourneywebeventsnotificationdeviceMap["category"].(string); ok {
		o.Category = &Category
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
