//Robot
package main

import (
	"fmt"
	"time"
)

struct Robot{
	name string
}

var name []string

alphabet:=make([]string, 0, 26) 
var digit=make([]int,0,10)

func init(){
	var ch byte
	//n:="&5"
	nameList=append(nameList,n)
	//loop over A to Z and keep on appending 
	for ch= 'A'; ch<='Z'; ch++{
		alphabet=append(alphabet,string(ch))
	}
	for ch=0;ch<10;ch++{
		digit=append(digit,int(ch))
	}
	//random number generator, adding a timestamp
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	nums:=make([]int,0)
	for len(nums)<count{
		//generate random numbers
		num:=r.Intn((end-start))+start
	}
		//to be continued...
	}
}