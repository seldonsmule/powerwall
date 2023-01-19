// This file was autogenerated using https://github.com/twpayne/go-jsonstruct

package powerwall


import "github.com/seldonsmule/logmsg"

func (pP *Powerwall) GetMeters() (bool, Meters){

  var s Meters

  pP.SetObject(&s)

  if(!pP.GetStruct("meters", false)){
    logmsg.Print(logmsg.Error, "GetStruct(meters) failed")
    return false, s
  }

  return true, s

}


type Meters []struct {
	Cts []struct {
		Inverted             []bool  `json:"inverted"`
		RealPowerScaleFactor float64 `json:"real_power_scale_factor,omitempty"`
		Type                 string  `json:"type"`
		Valid                []bool  `json:"valid"`
	} `json:"cts"`
	Serial  string `json:"serial"`
	ShortID string `json:"short_id"`
	Type    string `json:"type"`
}