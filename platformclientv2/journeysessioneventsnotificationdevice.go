package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationdevice
type Journeysessioneventsnotificationdevice struct { 
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

func (o *Journeysessioneventsnotificationdevice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationdevice
	
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

func (o *Journeysessioneventsnotificationdevice) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationdeviceMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationdeviceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneysessioneventsnotificationdeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IsMobile, ok := JourneysessioneventsnotificationdeviceMap["isMobile"].(bool); ok {
		o.IsMobile = &IsMobile
	}
    
	if ScreenHeight, ok := JourneysessioneventsnotificationdeviceMap["screenHeight"].(float64); ok {
		ScreenHeightInt := int(ScreenHeight)
		o.ScreenHeight = &ScreenHeightInt
	}
	
	if ScreenWidth, ok := JourneysessioneventsnotificationdeviceMap["screenWidth"].(float64); ok {
		ScreenWidthInt := int(ScreenWidth)
		o.ScreenWidth = &ScreenWidthInt
	}
	
	if Fingerprint, ok := JourneysessioneventsnotificationdeviceMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if OsFamily, ok := JourneysessioneventsnotificationdeviceMap["osFamily"].(string); ok {
		o.OsFamily = &OsFamily
	}
    
	if OsVersion, ok := JourneysessioneventsnotificationdeviceMap["osVersion"].(string); ok {
		o.OsVersion = &OsVersion
	}
    
	if Category, ok := JourneysessioneventsnotificationdeviceMap["category"].(string); ok {
		o.Category = &Category
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
