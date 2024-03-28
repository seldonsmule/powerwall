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

func printMetersData(pw *powerwall.Powerwall, tokenPtr *string) bool{


  //fmt.Println(ma)

  bFristTime := true

  var  SiteImportedStart float64
  var  SiteExportedStart float64
  var  BattImportedStart float64
  var  BattExportedStart float64
  var  SolarImportedStart float64
  var  SolarExportedStart float64
  var  LoadImportedStart float64
  var  LoadExportedStart float64


/*
  fmt.Println("figuring out seconds/minutes stuff")

  fmt.Println(time.Now().Second())

  if(time.Now().Second() != 0){

    st := 60 - time.Now().Second()

    fmt.Println("not 00 wait ", st)

    val := time.Duration(st) * time.Second

    fmt.Println(time.Now())
    fmt.Println("sleep val: ", val)

    time.Sleep(val)

  }

  fmt.Println(time.Now())
  fmt.Println("Finally: ", time.Now().Second())
  */

//  os.Exit(1)





  for {


    pw = powerwall_login_with_tokenfile(false, *tokenPtr)

    if(pw == nil){
      fmt.Println("powerwall_login - bad token??")
      return false
    }

    worked, ma := pw.GetMetersAggregates()

    if(!worked){
      fmt.Println("Get the endpoint failed - metersaggregates\n")
      return false
    }

    worked, soe := pw.GetSystemStatusSoe()

    if(!worked){
      fmt.Println("Get the endpoint failed - systemstatussoe\n")
      return false
    }


    battPercent := soe.Percentage

   
    battkW := ma.Battery.InstantPower / 1000
    loadkW := ma.Load.InstantPower / 1000
    solkW := ma.Solar.InstantPower / 1000
    sitekW := ma.Site.InstantPower / 1000

    

    //csi := int64(ma.Site.EnergyImported)
    //cse := int64(ma.Site.EnergyExported)
    csi := ma.Site.EnergyImported
    cse := ma.Site.EnergyExported

    cbi := ma.Battery.EnergyImported
    cbe := ma.Battery.EnergyExported

    csoli := ma.Solar.EnergyImported
    csole := ma.Solar.EnergyExported

    cli := ma.Load.EnergyImported
    cle := ma.Load.EnergyExported

    t := time.Now()

    // test if midnight
    //if(t.Hour() == 9 && t.Minute() == 21){
    if(t.Hour() == 0 && t.Minute() == 0){
      fmt.Println("Midnight - reset counters")
      bFristTime = true
    }

    if(bFristTime){
      bFristTime = false

      SiteImportedStart = csi
      SiteExportedStart = cse

      BattImportedStart = cbi
      BattExportedStart = cbe

      SolarImportedStart = csoli
      SolarExportedStart = csole

      LoadImportedStart = cli
      LoadExportedStart = cle


      /*
      BattImportedStart = BattImported
      BattExportedStart = BattExported
      */

    }

    fmt.Println(t.Format("01-02-2006 15:04:05 "))
    fmt.Printf("          Solar (Solar Roof) kW(%.1f)\n", solkW)
    fmt.Printf("Site (Grid) kW(%.1f)         Load (Home) kW(%.1f)\n", sitekW, loadkW)
    fmt.Printf("          Batt (Powerwall) kW(%.1f)\n", battkW)
    fmt.Printf("          Battery Percentage: %.1f\n", battPercent)

    //fmt.Println("figuring out seconds/minutes stuff")

    fmt.Println(time.Now().Second())

//    if(time.Now().Second() == 0){
    if(true){

      SiteImported := csi - SiteImportedStart
      SiteExported := cse - SiteExportedStart
      SiteDiff := SiteImported - SiteExported

      BattImported := cbi - BattImportedStart
      BattExported := cbe - BattExportedStart
      BattDiff := BattImported - BattExported

      SolarImported := csoli - SolarImportedStart
      SolarExported := csole - SolarExportedStart
      SolarDiff := SolarImported - SolarExported
  
      LoadImported := cli - LoadImportedStart
      LoadExported := cle - LoadExportedStart
      LoadDiff := LoadImported - LoadExported

      //fmt.Println(time.Now().Format("01-02-2006 15:04:05 "))
      fmt.Println(t.Format("01-02-2006 15:04:05 "))

      //  fmt.Printf("Usage: Solar kW [%.1f]\n", solarkW)

      fmt.Printf("Site (Grid) kW(%.1f) StartImport %f StartExport %f CurrentImport %f CurrentExport %f Imported %f Exported %f Diff %f\n", sitekW, SiteImportedStart, SiteExportedStart, csi, cse, SiteImported, SiteExported, SiteDiff)
      fmt.Printf("Batt (Powerwall) kW(%.1f) StartImport %f StartExport %f CurrentImport %f CurrentExport %f Imported %f Exported %f Diff %f\n", battkW, BattImportedStart, BattExportedStart, cbi, cbe, BattImported, BattExported, BattDiff)
      fmt.Printf("Solar (Solar Roof) kW(%.1f) StartImport %f StartExport %f CurrentImport %f CurrentExport %f Imported %f Exported %f Diff %f\n", solkW, SolarImportedStart, SolarExportedStart, csoli, csole, SolarImported, SolarExported, SolarDiff)
      fmt.Printf("Load (Home) kW(%.1f) StartImport %f StartExport %f CurrentImport %f CurrentExport %f Imported %f Exported %f Diff %f\n", loadkW, LoadImportedStart, LoadExportedStart, cli, cle, LoadImported, LoadExported, LoadDiff)
      fmt.Printf("Battery Percentage: %.1f\n", battPercent)



    } // end if seconds == 0

    fmt.Println("---------------------------------------------")

    //time.Sleep(1 * time.Minute)
    val := time.Duration(1) * time.Second

    //fmt.Println(time.Now())
    //fmt.Println("sleep val: ", val)
    time.Sleep(val)

  } // for


  return true
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

    case "metersdata":
      fmt.Println("do metersdata")

      readconf(*confPtr)

      pw = powerwall_login_with_userid(false, gMyconf.Userid, gMyconf.Passwd)

      if(pw == nil){
        fmt.Println("powerwall_login failed")
        os.Exit(4)
      }

      printMetersData(pw, tokenPtr)

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
