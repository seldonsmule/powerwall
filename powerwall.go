package powerwall

import (

  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  "strings"
  "sort"
  "github.com/seldonsmule/restapi"
  "github.com/seldonsmule/logmsg"
)


type PW_Endpoints struct{

  Endpoint string    // the final piece of the URL
  Structname string  // Name to save in a created go file for the struct
  Filename string    // Name of the file to save it in

  bHidden bool // if true, will be visiable in list and saving actions


}

type Powerwall struct {

  sToken string
  bHaveToken bool
  sCertfile string
  sTokenFilename string

  oObject interface{}
  bSaveResponse bool

  bDebug bool

  bLoggedIn bool

  bRawGetStruct bool

  sSaveDirectory string

  mEndpoints map[string]PW_Endpoints

}


const PWURL string = "https://powerwall/api"

func New(certfile string) *Powerwall {

  p := new(Powerwall)

  p.sToken = "not gotten"
  p.bHaveToken = false
  p.sCertfile = certfile
  p.sTokenFilename = "pw.token"

  p.bSaveResponse = false
  p.oObject = nil

  p.bLoggedIn = false

  p.bRawGetStruct = false

  p.sSaveDirectory = "tesla_json"

  p.mEndpoints = make(map[string]PW_Endpoints)

  p.mEndpoints["config"] = PW_Endpoints{ "config", "Config", "config", false}

  p.mEndpoints["customer"] = PW_Endpoints{ "customer", "Customer", "customer", false}

  p.mEndpoints["devicesvitals"] = PW_Endpoints{ "devices/vitals", "DeviceVitals", "devicevitals", true}

  p.mEndpoints["generators"] = PW_Endpoints{ "generators", "Generators", "generators", false}

  p.mEndpoints["generatorsdisconnecttypes"] = PW_Endpoints{ "generators/disconnect_types", "DisconnectTypes", "generator_disconnect_types", false}

  p.mEndpoints["installer"] = PW_Endpoints{ "installer", "Installer", "installer", false}

  p.mEndpoints["installercompanies"] = PW_Endpoints{ "installer/companies", "InstallerCompanies", "installercompanies", false}

  p.mEndpoints["login"] = PW_Endpoints{ "login/Basic", "Login", "login", true}

  p.mEndpoints["meters"] = PW_Endpoints{ "meters", "Meters", "meters", false}

  p.mEndpoints["metersaggregates"] = PW_Endpoints{ "meters/aggregates", "MetersAggregates", "metersaggregates", false}

  p.mEndpoints["metersreadings"] = PW_Endpoints{ "meters/readings", "MetersReadings", "metersreadings", false}

  p.mEndpoints["metersstatus"] = PW_Endpoints{ "meters/status", "MetersStatus", "metersstatus", false}

  p.mEndpoints["networks"] = PW_Endpoints{ "networks", "Networks", "networks", false}

  p.mEndpoints["powerwalls"] = PW_Endpoints{ "powerwalls", "Powerwalls", "powerwalls", false}

  p.mEndpoints["site_info"] = PW_Endpoints{ "site_info", "SiteInfo", "siteinfo", false}

  p.mEndpoints["system_status"] = PW_Endpoints{ "system_status", "SystemStatus", "systemstatus", false}
  p.mEndpoints["system_status_gridfaults"] = PW_Endpoints{ "system_status/grid_faults", "SystemStatusGridfaults", "systemstatus_gridfaults", false}
  p.mEndpoints["system_status_soe"] = PW_Endpoints{ "system_status/soe", "SystemStatusSoe", "systemstatus_soe", false}

  p.mEndpoints["meters_solar"] = PW_Endpoints{ "meters/solar", "MetersSolar", "meters_solar", false}
  p.mEndpoints["meters_site"] = PW_Endpoints{ "meters/site", "MetersSite", "meters_site", false}


  return p

}


func (pP *Powerwall) SaveResponseOn(){
  pP.bSaveResponse = true
}

func (pP *Powerwall) SaveResponseOff(){
  pP.bSaveResponse = false
}

func (pP *Powerwall) SetObject(obj interface{}){

  pP.oObject = obj

}


func (pP *Powerwall) SetSaveDirectory(dir string){
  pP.sSaveDirectory = dir
}

func (pP *Powerwall) GetSaveDirectory() string{
  return pP.sSaveDirectory
}

func (pP *Powerwall) SetTokenFileName(filename string){
  pP.sTokenFilename = filename
}

func (pP *Powerwall) GetTokenFileName() string{
  return pP.sTokenFilename
}

func (pP *Powerwall) ReadTokenFromFile() bool {

  data, err := os.ReadFile(pP.sTokenFilename)

  if(err != nil){
    msg := fmt.Sprintf("Error reading token file [%s]", pP.sTokenFilename)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  pP.SetToken(string(data))

  return true
}

func (pP *Powerwall) SaveTokenToFile() bool {

  f, err := os.Create(pP.sTokenFilename)

  if(err != nil){
    msg := fmt.Sprintf("Error creating token file [%s]", pP.sTokenFilename)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  defer f.Close()

  _, err = f.WriteString(pP.sToken)

  if(err != nil){

    msg := fmt.Sprintf("Error writing to token file [%s]", pP.sTokenFilename)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  return true

}

func (pP *Powerwall) GetToken() string{
 
  return pP.sToken

}

func (pP *Powerwall) SetToken(token string) {

  pP.sToken = token
  pP.bLoggedIn = true

}


func (pP *Powerwall) CheckEndpoint(endpointname string) bool {

  e := pP.mEndpoints[endpointname]

  // testing for an empty struct - i.e. , not found
  if( e == (PW_Endpoints{}) ){
    return false
  }

  return true
}


func (pP *Powerwall) RawGetStruct(endpointname string, bstdout bool) bool {

  pP.bRawGetStruct = true

  return(pP.GetStruct(endpointname, bstdout))

}

func (pP *Powerwall) GetStructList(bstdout bool) bool {

  // maps are not sorted, so to display a sorted one we need to sort
  // our keys

  keys := make([]string, 0, len(pP.mEndpoints))

  for k := range pP.mEndpoints {

    keys = append(keys, k)

  } 

  sort.Strings(keys)

  for _, k := range keys {

    if(pP.mEndpoints[k].bHidden){
      msg := fmt.Sprintf("Skipping - set to hidden - GetStruct for [%s]", k) 
      logmsg.Print(logmsg.Warning, msg)
      continue
    }

    msg := fmt.Sprintf("GetStruct for [%s]", k) 
    logmsg.Print(logmsg.Debug01, msg)
  
    if(!pP.GetStruct(k, bstdout)){
      msg := fmt.Sprintf("GetStruct for [%s] failed", k) 
      logmsg.Print(logmsg.Error, msg)
      return false
    }


  }

  return true
}

func (pP *Powerwall) GetStruct(endpointname string, bstdout bool) bool {

  var msg string

  if(pP.bRawGetStruct){
    msg = fmt.Sprintf("Getting Go Struct for shortURL[%s]\n", endpointname)
  }else{ // using our array
    msg = fmt.Sprintf("Getting Go Struct for [%s]\n", endpointname)
  }

  logmsg.Print(logmsg.Debug01, msg)

  var sStructname string
  var sUrl string
  var sFilename string

  if(pP.bRawGetStruct){
    sStructname = "RawStruct"
    sFilename = "RawFile"
    sUrl = pP.RawGetUrl(endpointname)
  }else{
    if(!pP.CheckEndpoint(endpointname)){
      msg := fmt.Sprintf("Invalid endpointname[%s]\n", endpointname)
      logmsg.Print(logmsg.Error, msg)
      return false
    }

    sStructname = pP.mEndpoints[endpointname].Structname
    //sUrl = pP.GetUrl(pP.mEndpoints[endpointname].Endpoint)
    sUrl = pP.GetUrl(endpointname)
    sFilename = pP.mEndpoints[endpointname].Filename
  }

  emsg := fmt.Sprintf("GetStrut: endpointname[%s] url: %s\n", endpointname, sUrl)
  logmsg.Print(logmsg.Debug01, emsg)

  r := restapi.NewGet(sStructname, sUrl)

  r.SetBearerAccessToken(pP.GetToken())

  if(!r.UseCert(pP.sCertfile)){
    logmsg.Print(logmsg.Error, "Failed - could not open cert file")
    return false
  }

  r.JsonOnly()

  if(!r.Send()){
    msg := fmt.Sprintf("Error getting struct for [%s]\n", endpointname)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

 if(pP.bDebug){
    fmt.Println(r.GetResponseBody())
 }

  if(!pP.SaveResponseBody(r, sFilename, sStructname, endpointname, bstdout)){
    logmsg.Print(logmsg.Error,"SaveResponseBody failed")
    return(false)
  }


  if(pP.oObject != nil){

//    fmt.Println("Attempting to unmarsal to oObject type: ", reflect.TypeOf(pP.oObject))

//fmt.Printf("--------Type: %T\n", pP.oObject)
    json.Unmarshal(r.BodyBytes, pP.oObject)

//fmt.Println("=========")
//    fmt.Println(pP.oObject)

  }



  return true

}

func (pP *Powerwall) SaveResponseBody(r *restapi.Restapi, sFilename string, sStructname string, sEndpointname string, bstdout bool) bool{


  if(!pP.bSaveResponse){
    logmsg.Print(logmsg.Debug01, "Skipping save - flag not set")
    return true
  }

 
err := os.MkdirAll(pP.sSaveDirectory, 0755)

  if(err != nil){

    msg := fmt.Sprintf("Mkdir(%s) failed [%s]\n", pP.sSaveDirectory, err)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  fullname := pP.sSaveDirectory+"/pw_"+sFilename

  //pP.SaveResponseBody(pP.sSaveDirectory+"/"+sFilename, sStructname, bstdout)
  r.SaveResponseBody(fullname, sStructname, bstdout)

  if(!pP.bRawGetStruct){ // skipping putting a helper function in the file

    fullname = fullname+".go"

    input, err := ioutil.ReadFile(fullname)
    if err != nil {
      logmsg.Print(logmsg.Error, err)
      return false
    }

    lines := strings.Split(string(input), "\n")

    //newstuff := fmt.Sprintf("package powerwall\n\n"+
    newstuff := fmt.Sprintf("\n\nimport \"github.com/seldonsmule/logmsg\"\n\n"+
                  "func (pP *Powerwall) Get%s() (bool, %s){\n\n"+
                  "  var s %s\n\n"+
                  "  pP.SetObject(&s)\n\n"+
                  "  if(!pP.GetStruct(\"%s\", false)){\n"+
                  "    logmsg.Print(logmsg.Error, \"GetStruct(%s) failed\")\n"+
                  "    return false, s\n"+
                  "  }\n\n"+
                  "  return true, s\n\n"+
                  "}\n\n", sStructname, sStructname, sStructname, sEndpointname, sEndpointname)

    structstart := fmt.Sprintf("type %s ", sStructname)

    for i, line := range lines {
      if strings.Contains(line, "package") {
        lines[i] = "package powerwall"
      }

      if strings.Contains(line, structstart){
        lines[i-1] = newstuff
      }
    }

    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(fullname, []byte(output), 0644)
    if err != nil {
      logmsg.Print(logmsg.Error, err)
      return false
    }

  } // end if !bRawGetStruct

  return true

}

func (pP *Powerwall) RawGetUrl(shortUrl string) string{

  Url := fmt.Sprintf("%s/%s", PWURL, shortUrl)

  return(Url)

}

func (pP *Powerwall) GetUrl(endpointname string) string{

  //Url := fmt.Sprintf("%s/%s", PWURL, pP.mEndpoints[endpointname].Endpoint)

//fmt.Printf("GetUrl: endpointname[%s] url[%s]\n", endpointname, pP.mEndpoints[endpointname].Endpoint)

  return( pP.RawGetUrl(pP.mEndpoints[endpointname].Endpoint))

}

func (pP *Powerwall) Login(username string, passwd string, bstdout bool) bool{

  r := restapi.NewPost(pP.mEndpoints["login"].Structname, pP.GetUrl("login"))
 
  if(!r.UseCert(pP.sCertfile)){
    msg := fmt.Sprintln("Failed - could not open cert file")
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  if(pP.bDebug){
    r.DebugOn()
  }

  jsonstr := fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\", \"email\" : \"%s\", \"force_sm_off\":false}", "customer", passwd, username)

  //fmt.Printf("[%s]\n", jsonstr)

  r.SetPostJson(jsonstr)

  if(r.Send()){

    if(pP.bDebug){
      r.Dump()
    }

  }

  if(!pP.SaveResponseBody(r, "login", "Login", "login", bstdout)){

    logmsg.Print(logmsg.Error,"SaveResponseBody failed")
    return(false)
  }

  pP.sToken = r.GetValueString("token")

  pP.bLoggedIn = true

  return true

}

func (pP *Powerwall) DebugOn(){
  pP.bDebug = true
}

func (pP *Powerwall) DebugOff(){
  pP.bDebug = false
}

func List(bnameonly bool){

  p := New("blank")

  p.ListEndpoints(bnameonly)

}

func CheckName(endpointname string) bool{

  p := New("blank")

  return p.CheckEndpoint(endpointname)

}

func (pP *Powerwall) ListEndpoints(bnameonly bool){

  // maps are not sorted, so to display a sorted one we need to sort
  // our keys

  keys := make([]string, 0, len(pP.mEndpoints))

  for k := range pP.mEndpoints {

    keys = append(keys, k)

  } 

  sort.Strings(keys)

  for _, k := range keys {

    if(pP.mEndpoints[k].bHidden){
      continue
    }

    if(bnameonly){
      fmt.Printf("Name:[%s]\n", k)
    }else{
      fmt.Printf("Name:[%s] Endpoint[%s] StructName[%s] GoFileName[%s]\n", k, 
                           pP.mEndpoints[k].Endpoint,
                           pP.mEndpoints[k].Structname,
                           pP.mEndpoints[k].Filename)
    }

  }

}

func (pP *Powerwall) Dump(){

  fmt.Println("Powerwall dump")
  //fmt.Println(pP)

  fmt.Printf("bLoggedIn: %t\n", pP.bLoggedIn)
  fmt.Printf("sToken: %s\n", pP.sToken)
  fmt.Printf("sTokenFile: %s\n", pP.sTokenFilename)
  fmt.Printf("bHaveToken: %t\n", pP.bHaveToken)
  fmt.Printf("sCertfile: %s\n", pP.sCertfile)
  fmt.Printf("sSaveDirectory: %s\n", pP.sSaveDirectory)
  fmt.Printf("bRawGetStruct: %t\n", pP.bRawGetStruct)
//  fmt.Printf("Endpoints: %s\n", pP.sCertfile)
  fmt.Println(pP.mEndpoints)

}
