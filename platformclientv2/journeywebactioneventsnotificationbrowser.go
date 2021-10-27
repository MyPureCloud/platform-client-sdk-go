package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationbrowser
type Journeywebactioneventsnotificationbrowser struct { 
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

func (o *Journeywebactioneventsnotificationbrowser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationbrowser
	
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

func (o *Journeywebactioneventsnotificationbrowser) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationbrowserMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationbrowserMap)
	if err != nil {
		return err
	}
	
	if Family, ok := JourneywebactioneventsnotificationbrowserMap["family"].(string); ok {
		o.Family = &Family
	}
	
	if Version, ok := JourneywebactioneventsnotificationbrowserMap["version"].(string); ok {
		o.Version = &Version
	}
	
	if Lang, ok := JourneywebactioneventsnotificationbrowserMap["lang"].(string); ok {
		o.Lang = &Lang
	}
	
	if Fingerprint, ok := JourneywebactioneventsnotificationbrowserMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
	
	if ViewHeight, ok := JourneywebactioneventsnotificationbrowserMap["viewHeight"].(float64); ok {
		ViewHeightInt := int(ViewHeight)
		o.ViewHeight = &ViewHeightInt
	}
	
	if ViewWidth, ok := JourneywebactioneventsnotificationbrowserMap["viewWidth"].(float64); ok {
		ViewWidthInt := int(ViewWidth)
		o.ViewWidth = &ViewWidthInt
	}
	
	if FeaturesFlash, ok := JourneywebactioneventsnotificationbrowserMap["featuresFlash"].(bool); ok {
		o.FeaturesFlash = &FeaturesFlash
	}
	
	if FeaturesJava, ok := JourneywebactioneventsnotificationbrowserMap["featuresJava"].(bool); ok {
		o.FeaturesJava = &FeaturesJava
	}
	
	if FeaturesPdf, ok := JourneywebactioneventsnotificationbrowserMap["featuresPdf"].(bool); ok {
		o.FeaturesPdf = &FeaturesPdf
	}
	
	if FeaturesWebrtc, ok := JourneywebactioneventsnotificationbrowserMap["featuresWebrtc"].(bool); ok {
		o.FeaturesWebrtc = &FeaturesWebrtc
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationbrowser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
