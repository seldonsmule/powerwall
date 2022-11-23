package main


import (
	"os"
	"fmt"
	"flag"
       "github.com/seldonsmule/powerwall"
        "github.com/seldonsmule/logmsg"

)



const CERTFILE string =  "powerwall.cer"
const JSON_GO_FILES string = "json_go_files"

func powerwall_login_with_userid(bDebug bool, userid string, passwd string) *powerwall.Powerwall{

  p := powerwall.New(CERTFILE)

  p.SaveResponseOn()

  if(bDebug){
fmt.Println("setting debug")
    p.DebugOn()
    p.Dump()
  }else{
fmt.Println("debug is false")
}

  if(!p.Login(userid, passwd, false)){
    fmt.Printf("Login failed\n");
    return nil
  }
  

  p.SaveTokenToFile()

  return p
}


func powerwall_login_with_tokenfile(bDebug bool, tokenfilename string) *powerwall.Powerwall{

  p := powerwall.New(CERTFILE)

  p.SetTokenFileName(tokenfilename)
  p.ReadTokenFromFile()

  msg := fmt.Sprintf("Token:[%s]\n", p.GetToken())
  logmsg.Print(logmsg.Info, msg)

  return p

}

func help(){

  fmt.Println("Simple tool to get the json response data and save as go structs")

  fmt.Println("Usage getstruct -cmd [a command, see below]")
  fmt.Println()
  flag.PrintDefaults()
  fmt.Println()
  fmt.Println("cmds:")
  fmt.Println("       login - Needs the following to be uset")
  fmt.Println("             -userid email address used with the powerwall")
  fmt.Println("             -passwd password used with the powerwall")
  fmt.Println("       get - Gets an endpoint.  Needs the following to be set")
  fmt.Println("             -endpoint telsa_endpoint (ex: customer)")
  fmt.Println("             -stdout if set - prints go code to stdout")
  fmt.Println("       rawget - Gets an endpoint from a passed in url.  Needs the following to be set")
  fmt.Println("             -endpoint telsa_endpoint (ex: customer or \"meters/aggregates\") - ")
  fmt.Println("             -stdout if set - prints go code to stdout")
  fmt.Println("       array - Runs through a complied in array of URLs")
  fmt.Println("       list - list all the endpoints")
  fmt.Println("       usetokenfile - token contained in a file instead of userid/passwd")
  fmt.Println("       certfile - Get cert file")
  fmt.Println("       percent - How full is the battery")
  fmt.Println()


}

func main() {

  cmdPtr := flag.String("cmd", "help", "Command to run")
  useridPtr := flag.String("userid", "notset", "userid/email address")
  passwdPtr := flag.String("passwd", "abc123", "userid/email address")
  endpointPtr := flag.String("endpoint", "customer", "Command to run")
  filenamePtr := flag.String("filename", "default", "Command to run")
  structnamePtr := flag.String("structname", "Default", "Command to run")
  tokenPtr := flag.String("usetokenfile", "pw.token", "Name of file containing a token instead of userid/pw")
  bstdoutPtr := flag.Bool("stdout", false, "If true, print to stdou")
  certfilePtr := flag.String("certfile", "powerwall.cer", "Name to save file to")
  bsleepPtr := flag.Bool("sleepupdate", false, "If true, sleep and loop on status")

  flag.Parse()

fmt.Printf("cmd=%s\n", *cmdPtr)

  //var err error
  //var token string
  var pw *powerwall.Powerwall
  bRawGet := false

  var bUseToken bool

  if(*useridPtr == "notset"){
    bUseToken = true
  }else{
    bUseToken = false
  }

  logmsg.SetLogFile("get_structs.log");

  logmsg.Print(logmsg.Info, "cmdPtr = ", *cmdPtr)
  logmsg.Print(logmsg.Info, "useridPtr = ", *useridPtr)
  logmsg.Print(logmsg.Info, "passwdPtr = ", *passwdPtr)
  logmsg.Print(logmsg.Info, "endpointPtr = ", *endpointPtr)
  logmsg.Print(logmsg.Info, "filenamePtr = ", *filenamePtr)
  logmsg.Print(logmsg.Info, "structnamePtr = ", *structnamePtr)
  logmsg.Print(logmsg.Info, "tokenPtr = ", *tokenPtr)
  logmsg.Print(logmsg.Info, "bstdoutPtr = ", *bstdoutPtr)
  logmsg.Print(logmsg.Info, "certfilePtr = ", *certfilePtr)
  logmsg.Print(logmsg.Info, "bsleepPtr = ", *bsleepPtr)
  logmsg.Print(logmsg.Info, "tail = ", flag.Args())

  if(*cmdPtr == "help"){
    help()
    os.Exit(1)
  }

  os.MkdirAll(JSON_GO_FILES, 0755)

  //bDebug := false

  //fmt.Println("Getting json responses to build go structs")

  switch *cmdPtr {

    case "list":
      fmt.Println("Listing all endpoints")

      powerwall.List(false)

    case "login":
      fmt.Println("do login") 

      pw = powerwall_login_with_userid(true, *useridPtr, *passwdPtr)

      if(pw == nil){
        fmt.Println("powerwall_login failed")
        os.Exit(4)
      }

    case "array":
      fmt.Println("do array get") 

      if(bUseToken){
        pw = powerwall_login_with_tokenfile(false, *tokenPtr)
      }else{
        pw = powerwall_login_with_userid(false, *useridPtr, *passwdPtr)
      }

      if(pw == nil){
        fmt.Println("powerwall_login failed")
        os.Exit(4)
      }

      pw.SaveResponseOn()

      if(!pw.GetStructList(false)){
        fmt.Println("Error processing our array of endpoints, see log")
        os.Exit(4)
      }


    case "rawget":
      bRawGet = true
      fallthrough
    case "get":
      fmt.Println("Get a single endpoint") 

      if(!bRawGet){

        if(!powerwall.CheckName(*endpointPtr)){

          fmt.Printf("Not a supported name[%s]\n\n", *endpointPtr)
        
          fmt.Printf("Supported names:\n")
          powerwall.List(true)

          os.Exit(4)

        }

      }

      if(bUseToken){
        pw = powerwall_login_with_tokenfile(false, *tokenPtr)
      }else{
        pw = powerwall_login_with_userid(false, *useridPtr, *passwdPtr)
      }

      if(pw == nil){
        fmt.Println("powerwall_login failed")
        os.Exit(4)
      }

      pw.SaveResponseOn()

      if(bRawGet){
        if(!pw.RawGetStruct(*endpointPtr, *bstdoutPtr)){
          fmt.Println("Get the endpoint failed\n")
          os.Exit(4)
        }
      }else{
        if(!pw.GetStruct(*endpointPtr, *bstdoutPtr)){
          fmt.Println("Get the endpoint failed\n")
          os.Exit(4)
        }
      }

      fmt.Println()
      fmt.Printf("SUCCESS - See [%s] directory for results\n",
                  pw.GetSaveDirectory())


    default:
      help()
      os.Exit(2)

  }

  //powerwall_login(bDebug)

  os.Exit(0)
     
}
