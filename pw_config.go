// This file was autogenerated using https://github.com/twpayne/go-jsonstruct

package powerwall


import "github.com/seldonsmule/logmsg"

func (pP *Powerwall) GetConfig() (bool, Config){

  var s Config

  pP.SetObject(&s)

  if(!pP.GetStruct("config", false)){
    logmsg.Print(logmsg.Error, "GetStruct(config) failed")
    return false, s
  }

  return true, s

}


type Config struct {
	Vin string `json:"vin"`
}