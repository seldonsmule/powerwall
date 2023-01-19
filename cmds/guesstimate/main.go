package main


import (
	"os"
	"fmt"
	"flag"
	"strings"
	"math"
       "github.com/seldonsmule/powerwall"
       "github.com/seldonsmule/simpleconffile"
        "time"
        "github.com/seldonsmule/logmsg"

)

const CERTFILE string =  "powerwall.cer"

type Configuration struct {

  Userid string
  Passwd string
  Filename string
  Encrypted bool

}


const COMPILE_IN_KEY = "example key 1234"

var gMyconf Configuration

func is_grid_up(pw *powerwall.Powerwall, tokenPtr *string) bool {

   pw = powerwall_login_with_tokenfile(false, *tokenPtr)

   if(pw == nil){
      fmt.Println("powerwall_login - bad token??")
      return false
    }

   worked, ss := pw.GetSystemStatus()

   if(!worked){
     fmt.Println("Get the endpoint failed - system_status\n")
     return true
   }

   // there might be more than one battery pack, but we only need 
   // to check the 1st one for grid status


  maxbattery := ss.NominalFullPackEnergy / 1000
  currentcharge := ss.NominalEnergyRemaining / 1000

  fmt.Printf("Max battery: %f",  maxbattery)
  fmt.Printf("Current charge: %f",  currentcharge)


  return true

}

func get_battery_size(pw *powerwall.Powerwall, tokenPtr *string) (float64, float64){

   pw = powerwall_login_with_tokenfile(false, *tokenPtr)

   if(pw == nil){
      fmt.Println("powerwall_login - bad token??")
      return -1, -1
    }

   worked, ss := pw.GetSystemStatus()

   if(!worked){
     fmt.Println("Get the endpoint failed - system_status\n")
     return -1, -1
   }


  maxbattery := ss.NominalFullPackEnergy / 1000
  currentcharge := ss.NominalEnergyRemaining / 1000

  return maxbattery, currentcharge

}

func get_kWs(pw *powerwall.Powerwall, tokenPtr *string) (float64, float64, float64, float64){

      pw = powerwall_login_with_tokenfile(false, *tokenPtr)

      if(pw == nil){
        fmt.Println("powerwall_login - bad token??")
        return -1, -1, -1, -1
      }

      worked, ma := pw.GetMetersAggregates()

      if(!worked){
        fmt.Println("Get the endpoint failed - metersaggregates\n")
        return -1, -1, -1, -1
      }

//fmt.Println(ma)

  batterykW := ma.Battery.InstantPower / 1000
  housekW := ma.Load.InstantPower / 1000
  solarkW := ma.Solar.InstantPower / 1000
  gridkW := ma.Site.InstantPower / 1000

  return batterykW, housekW, solarkW, gridkW
}

func get_percent(pw *powerwall.Powerwall, tokenPtr *string) float64{

  pw = powerwall_login_with_tokenfile(false, *tokenPtr)

  if(pw == nil){
     fmt.Println("powerwall_login - bad token??")
     return -1
  }

  worked, soe := pw.GetSystemStatusSoe()

  if(!worked){
     fmt.Println("Get the endpoint failed - metersaggregates\n")
     return -1
  }

  return(soe.Percentage)

}

func print_pw_status(pw *powerwall.Powerwall, tokenPtr *string) {

      currentTime := time.Now()

      fmt.Println(currentTime)

      batterykW, homekW, solarkW, gridkW := get_kWs(pw, tokenPtr)
      batterySize, currentCharge := get_battery_size(pw, tokenPtr)

      fmt.Printf("Usage: Solar kW [%.1f]\n", solarkW)
      fmt.Printf("Usage: Home kW [%.1f]\n", homekW)
      fmt.Printf("Usage: Battery kW [%.1f]\n", batterykW)

      fmt.Printf("Usage: Grid kW [%.1f]\n", gridkW)

      fmt.Println("-------------------")

      bat_percent := get_percent(pw, tokenPtr)
      fmt.Printf("Battery Level [%.3f]\n", bat_percent)
      //fmt.Printf("Battery Level [%f]\n", bat_percent)


      batterykW = math.Abs(batterykW) // need it not negative

      batterykWneeded := batterySize - currentCharge

//fmt.Printf("batterykWneeded[%f] / batterkW[%f]\n", batterykWneeded, batterykW)

      hrtofull := batterykWneeded / batterykW


      fmt.Printf("Max Battery kW [%f]\n", batterySize)
      fmt.Printf("Remaining Battery kW [%f]\n", currentCharge)

      if(bat_percent < 99){
        fmt.Printf("Charging needed kW [%f]\n", batterykWneeded)

        if(solarkW < homekW){

          fmt.Printf("Not enough solar to charge battery right now\n")
          return
        }

        fmt.Printf("Time to full - hours [%f]\n", hrtofull)

        roundmin := math.Round(hrtofull * 60)

        minutes := int(roundmin)

        fmt.Printf("Time to full - minutes [%f]\n", hrtofull * 60)
        fmt.Printf("Time to full - minutes [%d]\n", minutes)
        fmt.Printf("Time to full - round minutes [%f]\n", roundmin)



        mymin := time.Duration(minutes)
 

       fmt.Println(currentTime.Add(+time.Minute * mymin) )

      }

}

func powerwall_login_with_userid(bDebug bool, userid string, passwd string) *powerwall.Powerwall{

  p := powerwall.New(CERTFILE)

  p.SaveResponseOn()

/*
  if(bDebug){
fmt.Println("setting debug")
    p.DebugOn()
    p.Dump()
  }else{
fmt.Println("debug is false")
}
*/

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
  fmt.Println("       setconf - Setup Conf file")
  fmt.Println("             -userid email address used with the powerwall")
  fmt.Println("             -passwd password used with the powerwall")
  fmt.Println("             -conffile name of conffile (.pw.conf default)")
  fmt.Println("       readconf - Display conf info")
  fmt.Println("       login - Needs the following to be uset")
  fmt.Println("             -userid email address used with the powerwall")
  fmt.Println("             -passwd password used with the powerwall")
  fmt.Println("       status - Basic stats on the system")
  fmt.Println("             -sleepupdate if set, status will run in a loop")
  fmt.Println("       checkgrid - Test if grid power is off")
  fmt.Println()


}

func readconf(confFile string){

  simple := simpleconffile.New(COMPILE_IN_KEY, confFile)

  if(!simple.ReadConf(&gMyconf)){
    fmt.Println("Error reading conf file: ", confFile)
    os.Exit(3)
  }

  if(gMyconf.Encrypted){
    gMyconf.Userid = simple.DecryptString(gMyconf.Userid)
    gMyconf.Passwd = simple.DecryptString(gMyconf.Passwd)
  }

     
/*
  fmt.Printf("Encrypted [%v]\n", gMyconf.Encrypted)
  fmt.Printf("Userid [%v]\n", gMyconf.Userid)
  fmt.Printf("Passwd [%v]\n", gMyconf.Passwd)
*/

}

func main() {


  cmdPtr := flag.String("cmd", "help", "Command to run")
  useridPtr := flag.String("userid", "notset", "userid/email address")
  passwdPtr := flag.String("passwd", "notset", "userid/email address")
  confPtr := flag.String("conffile", ".pw.conf", "config file name")
  tokenPtr := flag.String("usetokenfile", "pw.token", "Name of file containing a token instead of userid/pw")
  bsleepPtr := flag.Bool("sleepupdate", false, "If true, sleep and loop on status")
  bdebugPtr := flag.Bool("debug", false, "If true, do debug magic")

  flag.Parse()

fmt.Printf("cmd=%s\n", *cmdPtr)

  var pw *powerwall.Powerwall

  logmsg.SetLogFile("guesstimate.log");

  logmsg.Print(logmsg.Info, "cmdPtr = ", *cmdPtr)
  logmsg.Print(logmsg.Info, "useridPtr = ", *useridPtr)
  logmsg.Print(logmsg.Info, "passwdPtr = ", *passwdPtr)
  logmsg.Print(logmsg.Info, "confPtr = ", *confPtr)
  logmsg.Print(logmsg.Info, "tokenPtr = ", *tokenPtr)
  logmsg.Print(logmsg.Info, "bsleepPtr = ", *bsleepPtr)
  logmsg.Print(logmsg.Info, "bdebugPtr = ", *bdebugPtr)
  logmsg.Print(logmsg.Info, "tail = ", flag.Args())

  if(*cmdPtr == "help"){
    help()
    os.Exit(1)
  }

  switch *cmdPtr {

    case "readconf":
      fmt.Println("Reading conf file")
      readconf(*confPtr)

    case "setconf":
      fmt.Println("Setting conf file")

      if(strings.Compare(*useridPtr, "notset") == 0){
        fmt.Println("FAIL: userid missing");
        os.Exit(2)
      }

      if(strings.Compare(*passwdPtr, "notset") == 0){
        fmt.Println("FAIL: password missing");
        os.Exit(2)
      }

      simple := simpleconffile.New(COMPILE_IN_KEY, *confPtr)

      gMyconf.Encrypted = true
      gMyconf.Userid = simple.EncryptString(*useridPtr)
      gMyconf.Passwd = simple.EncryptString(*passwdPtr)
      gMyconf.Filename = *confPtr

      simple.SaveConf(gMyconf)


    case "login":
      fmt.Println("do login") 

      pw = powerwall_login_with_userid(true, *useridPtr, *passwdPtr)

      if(pw == nil){
        fmt.Println("powerwall_login failed")
        os.Exit(4)
      }

    case "status":

      readconf(*confPtr)

      pw = powerwall_login_with_userid(false, gMyconf.Userid, gMyconf.Passwd)

      if(pw == nil){
        fmt.Println("powerwall_login failed")
        os.Exit(4)
      }

      if(*bsleepPtr){
        fmt.Println("Going into a forever loop")
        for {

          print_pw_status(pw, tokenPtr)
          fmt.Println("------------------------------------")
          time.Sleep(1 * time.Minute)
          fmt.Println()
          fmt.Println()
          fmt.Println()

        }
      }else{
        print_pw_status(pw, tokenPtr)
      }

    case "checkgrid":

      readconf(*confPtr)

      pw = powerwall_login_with_userid(false, gMyconf.Userid, gMyconf.Passwd)

      if(pw == nil){
        fmt.Println("powerwall_login failed")
        os.Exit(4)
      }

      if(is_grid_up(pw, tokenPtr)){
        fmt.Println("Grid is working")
        os.Exit(0)
      }else{
        fmt.Println("Power is out - running of batteries")
        os.Exit(1)
      }


    default:
      help()
      os.Exit(2)

  }

  os.Exit(0)
     
}
