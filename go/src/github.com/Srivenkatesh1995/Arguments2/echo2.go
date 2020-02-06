package main

// Srivenkatesh Vasiraju 
// svasiraj@uwyo.edu
// Assignment 1.4 echo command line arguments and parse arguments


func main() {
import (
"encoding/json"
"flag"
"fmt"
"io/ioutil"
"os"
)
type ConfigData struct {
Name string
Value string `json:"Year"`
}
func main() {
var Cfg = flag.String("cfg", "cfg.json", "config file for this call")
flag.Parse() // Parse CLI arguments to this program, --cfg <name>.json
fns := flag.Args()
if len(fns) == 0 {
fmt.Fprintf(os.Stderr, "Usage: ./echo2 [-cfg cfg.json] arg1 ...
os.Exit(1)
}
if Cfg == nil {
fmt.Printf("--cfg is a required parameter\n")
os.Exit(1)
}
gCfg, err := ReadConfig(*Cfg)
if err != nil {
fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %
os.Exit(1)
}	

fmt.Printf("Congiguration: %+v\n", gCfg)
fmt.Printf("JSON: %+v\n", IndentJSON(gCfg))
for ii, ag := range fns {
if ii < len(fns) {
fmt.Printf("%s ", ag)
} else {
fmt.Printf("%s", ag)
}
}
fmt.Printf("\n")
}
func ReadConfig(filename string) (rv ConfigData, err error) {
var buf []byte
buf, err = ioutil.ReadFile(filename)
if err != nil {
return ConfigData{}, err
}
err = json.Unmarshal(buf, &rv)
if err != nil {
return ConfigData{}, err
}
return
}
func IndentJSON(v interface{}) string {
s, err := json.MarshalIndent(v, "", "\t")
if err != nil {
return fmt.Sprintf("Error:%s", err)
} else {
return string(s)
}
}
