package main

import route2 "github.com/DiegoVSouza/imersaofsfc2-simulator/application/route"
import "fmt"

func main(){
	route := route2.Route{
		ID: "1",
		ClientID: "1"
	}
	route.LoadPostions()
	stringjson, := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}