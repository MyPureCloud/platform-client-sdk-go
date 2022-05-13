package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationbrowser
type Journeyoutcomeeventsnotificationbrowser struct { 
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

func (o *Journeyoutcomeeventsnotificationbrowser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationbrowser
	
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

func (o *Journeyoutcomeeventsnotificationbrowser) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationbrowserMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationbrowserMap)
	if err != nil {
		return err
	}
	
	if Family, ok := JourneyoutcomeeventsnotificationbrowserMap["family"].(string); ok {
		o.Family = &Family
	}
    
	if Version, ok := JourneyoutcomeeventsnotificationbrowserMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if Lang, ok := JourneyoutcomeeventsnotificationbrowserMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    
	if Fingerprint, ok := JourneyoutcomeeventsnotificationbrowserMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if ViewHeight, ok := JourneyoutcomeeventsnotificationbrowserMap["viewHeight"].(float64); ok {
		ViewHeightInt := int(ViewHeight)
		o.ViewHeight = &ViewHeightInt
	}
	
	if ViewWidth, ok := JourneyoutcomeeventsnotificationbrowserMap["viewWidth"].(float64); ok {
		ViewWidthInt := int(ViewWidth)
		o.ViewWidth = &ViewWidthInt
	}
	
	if FeaturesFlash, ok := JourneyoutcomeeventsnotificationbrowserMap["featuresFlash"].(bool); ok {
		o.FeaturesFlash = &FeaturesFlash
	}
    
	if FeaturesJava, ok := JourneyoutcomeeventsnotificationbrowserMap["featuresJava"].(bool); ok {
		o.FeaturesJava = &FeaturesJava
	}
    
	if FeaturesPdf, ok := JourneyoutcomeeventsnotificationbrowserMap["featuresPdf"].(bool); ok {
		o.FeaturesPdf = &FeaturesPdf
	}
    
	if FeaturesWebrtc, ok := JourneyoutcomeeventsnotificationbrowserMap["featuresWebrtc"].(bool); ok {
		o.FeaturesWebrtc = &FeaturesWebrtc
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationbrowser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
