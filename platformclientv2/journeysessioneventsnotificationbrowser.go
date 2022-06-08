package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationbrowser
type Journeysessioneventsnotificationbrowser struct { 
	// Family
	Family *string `json:"family,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Lang
	Lang *string `json:"lang,omitempty"`


	// Fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`


	// ViewHeight
	ViewHeight *int `json:"viewHeight,omitempty"`


	// ViewWidth
	ViewWidth *int `json:"viewWidth,omitempty"`


	// FeaturesFlash
	FeaturesFlash *bool `json:"featuresFlash,omitempty"`


	// FeaturesJava
	FeaturesJava *bool `json:"featuresJava,omitempty"`


	// FeaturesPdf
	FeaturesPdf *bool `json:"featuresPdf,omitempty"`


	// FeaturesWebrtc
	FeaturesWebrtc *bool `json:"featuresWebrtc,omitempty"`

}

func (o *Journeysessioneventsnotificationbrowser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationbrowser
	
	return json.Marshal(&struct { 
		Family *string `json:"family,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Lang *string `json:"lang,omitempty"`
		
		Fingerprint *string `json:"fingerprint,omitempty"`
		
		ViewHeight *int `json:"viewHeight,omitempty"`
		
		ViewWidth *int `json:"viewWidth,omitempty"`
		
		FeaturesFlash *bool `json:"featuresFlash,omitempty"`
		
		FeaturesJava *bool `json:"featuresJava,omitempty"`
		
		FeaturesPdf *bool `json:"featuresPdf,omitempty"`
		
		FeaturesWebrtc *bool `json:"featuresWebrtc,omitempty"`
		*Alias
	}{ 
		Family: o.Family,
		
		Version: o.Version,
		
		Lang: o.Lang,
		
		Fingerprint: o.Fingerprint,
		
		ViewHeight: o.ViewHeight,
		
		ViewWidth: o.ViewWidth,
		
		FeaturesFlash: o.FeaturesFlash,
		
		FeaturesJava: o.FeaturesJava,
		
		FeaturesPdf: o.FeaturesPdf,
		
		FeaturesWebrtc: o.FeaturesWebrtc,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationbrowser) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationbrowserMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationbrowserMap)
	if err != nil {
		return err
	}
	
	if Family, ok := JourneysessioneventsnotificationbrowserMap["family"].(string); ok {
		o.Family = &Family
	}
    
	if Version, ok := JourneysessioneventsnotificationbrowserMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if Lang, ok := JourneysessioneventsnotificationbrowserMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    
	if Fingerprint, ok := JourneysessioneventsnotificationbrowserMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if ViewHeight, ok := JourneysessioneventsnotificationbrowserMap["viewHeight"].(float64); ok {
		ViewHeightInt := int(ViewHeight)
		o.ViewHeight = &ViewHeightInt
	}
	
	if ViewWidth, ok := JourneysessioneventsnotificationbrowserMap["viewWidth"].(float64); ok {
		ViewWidthInt := int(ViewWidth)
		o.ViewWidth = &ViewWidthInt
	}
	
	if FeaturesFlash, ok := JourneysessioneventsnotificationbrowserMap["featuresFlash"].(bool); ok {
		o.FeaturesFlash = &FeaturesFlash
	}
    
	if FeaturesJava, ok := JourneysessioneventsnotificationbrowserMap["featuresJava"].(bool); ok {
		o.FeaturesJava = &FeaturesJava
	}
    
	if FeaturesPdf, ok := JourneysessioneventsnotificationbrowserMap["featuresPdf"].(bool); ok {
		o.FeaturesPdf = &FeaturesPdf
	}
    
	if FeaturesWebrtc, ok := JourneysessioneventsnotificationbrowserMap["featuresWebrtc"].(bool); ok {
		o.FeaturesWebrtc = &FeaturesWebrtc
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationbrowser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
