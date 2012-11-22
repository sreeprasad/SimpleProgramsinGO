package main

import "fmt"

func main(){

elements := make(map[string]string)
     elements["H"] = "Hydrogen"
     elements["He"] = "Helium"
     elements["Li"] = "Lithium"
     elements["Be"] = "Beryllium"
     elements["B"] = "Boron"
     elements["C"] = "Carbon"
     elements["N"] = "Nitrogen"
     elements["O"] = "Oxygen"
     elements["F"] = "Fluorine"
     elements["Ne"] = "Neon"

allElements := map[string]string{
	"H":"Hydrogen",
	"He":"Helium",
	"Li":"Lithium",
	"Be":"Berilium",
	"B":"Boron",
	"C":"Carbon",
	"N":"Nitrogen",
	"O":"Oxygen",
	"F":"Fluorine",
	"Ne":"Neon",
}

	fmt.Println(elements)
	
	/* is elements present */
	
	if _, ok := elements["Un"]; !ok {
		fmt.Println("element un found")
	}
	fmt.Println("all elements")
	fmt.Println(allElements)

	allElementsWithExplanations := map[string]map[string]string{
		"H":map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He":map[string]string{
		     "name":"Helium",
			"state":"gas",
		},
		"Li":map[string]string{
		     "name":"Lithium",
		     "state":"solid",
		},
		"Be": map[string]string{
		                "name":"Beryllium",
		                "state":"solid",
		           },
		           "B":  map[string]string{
		                "name":"Boron",
		                "state":"solid",
		           },
		           "C":  map[string]string{
		                "name":"Carbon",
		                "state":"solid",
		           },
		           "N":  map[string]string{
		                "name":"Nitrogen",
		                "state":"gas",
		           },
		           "O":  map[string]string{
		                "name":"Oxygen",
		                "state":"gas",
		           },
		           "F":  map[string]string{
		                "name":"Fluorine",
		                "state":"gas",
		           },
		           "Ne":  map[string]string{
		                "name":"Neon",
		                "state":"gas",
		           },
		
 
	}
	
	fmt.Println(allElementsWithExplanations)
	
	if _, present:=allElementsWithExplanations["F"]; present{
		fmt.Println(allElementsWithExplanations["F"])
	}

}