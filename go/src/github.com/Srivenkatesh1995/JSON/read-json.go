package main

//Srivenkatesh Vasiraju 
//Assignment 1.5 


import (
		"encoding/json"
		"flag"
		"fmt"
		"io/ioutil"
		"os"
)

type TransactionType struct 
{
		TxHash string
		TxIn int
		TxOut int
}

func main(){
	var Cfg = flag.String("in", "in.json", "config file for this call")
	flag.Parse() // Parse CLI arguments to this program, --cfg <name>.json
	fns := flag.Args()
		if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Usage: ./read-json [-in in.json] [arg1 ...]\n")
		os.Exit(1)
		}

 		if Cfg == nil {
		fmt.Printf("--in is a required parameter\n")
		os.Exit(1)
		}
		gCfg, err := ReadConfig(*Cfg)
		if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n",*Cfg,err)
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

func ReadConfig(filename string) (rv TransactionType, err error) {
 			var buf []byte
			buf, err = ioutil.ReadFile(filename)
			if err != nil {
			return TransactionType{}, err
			}
			err = json.Unmarshal(buf, &rv)
			if err != nil {
			return TransactionType{}, err
			}
			return
}

func IndentJSON(v interface{}) string {
		s, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
		return fmt.Sprintf("Error:%s", err)
		} else {
		nap := []byte(string(s))
		ioutil.WriteFile("out2.json",nap,0644)
		if err!=nil{
			//log.Fatal(err)
		}
		return string(s)
		}
}


		
