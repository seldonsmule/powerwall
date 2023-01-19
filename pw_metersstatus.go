// This file was autogenerated using https://github.com/twpayne/go-jsonstruct

package powerwall


import "github.com/seldonsmule/logmsg"

func (pP *Powerwall) GetMetersStatus() (bool, MetersStatus){

  var s MetersStatus

  pP.SetObject(&s)

  if(!pP.GetStruct("metersstatus", false)){
    logmsg.Print(logmsg.Error, "GetStruct(metersstatus) failed")
    return false, s
  }

  return true, s

}


type MetersStatus struct {
	Errors interface{} `json:"errors"`
	Status string      `json:"status"`
}