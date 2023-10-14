package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
)

const VERSION string = "v1.10"


//all our different file types and there headers
var FileTypes map[string]string = map[string]string{
	"Go"	: "Golang", //go
	"Java"	: "Java", //java
	"Scala"	: "Scala", //scala
	"Rb"	: "Ruby", //rb
	"C"		: "C", //c
	"C++"	: "C++", //c++
	"Php"	: "PHP", //php
	"Pl"	: "Perl", //pl
	"Js"	: "JavaScript", //js
	"Sh"	: "Shell", //shell
	"Cs"	: "CSharp", //csharp
	"Fs"	: "FSharp", //fsharp
	"Vb"	: "Visual Basic", //vb
	"Py"	: "Python", //python
	"Bat"	: "Batch", //batch
	"Html"	: "Html", //html
	"Css"	: "Stylesheets", //stylesheets
	"Xml"	: "Xml", //xml
	"Json"	: "Json", //json
	"Toml"  : "Toml", //toml
	"Txt"	: "Text", //text
	"Ini"   : "INI", //ini
	"Ppk"	: "PuTTY Private Key", //putty
	"Gif"	: "GIF", //gif
	"Key"	: "Key", //key
	"Cert"	: "Certificate", //cert
	"Sql"	: "SQL Database", //db
	"Itl"	: "ITL", //itl
	"Md"	: "Markdown", //md
}


func main() {

	//invalid length passed
	//we will now resort to err
	if len(os.Args) <= 1 { //line counter issue
		log.Fatalf("LineCounter %s\r\nsyntax: %s [accessories]\r\n", VERSION, os.Args[0])
	}

	var object map[string]RecursiveType = make(map[string]RecursiveType)

	//ranges through all possible files
	//this will ensure we read all of them
	for _, path := range os.Args[1:] {
		
		//runs the reader on the object properly
		//this will ensure its done properly and safely
		if err := Recursive(path, object); err != nil {
			log.Fatalf("LineCounter %s\r\nError: %s\r\n", VERSION, err)
		}
	}


	//renders the system properly
	//this will ensure its done properly
	branding := simpletable.New() //new simpletable
	branding.SetStyle(simpletable.StyleCompactClassic)

	branding.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Type"}, {Align: simpletable.AlignCenter, Text: "Files"}, {Align: simpletable.AlignCenter, Text: "Amount"}, {Align: simpletable.AlignCenter, Text: "Bytes"}, {Align: simpletable.AlignCenter, Text: "KB"}, {Align: simpletable.AlignCenter, Text: "MB"}, {Align: simpletable.AlignCenter, Text: "GB"},
		},
	}

	//branding.SetStyle(simpletable.StyleCompactClassic)

	//ranges through all objects
	//this will ensure its done properly
	for Type, Amount := range object { //ranges

		//default type banner here
		var displayType = Type + " " + "(UNKNOWN)"

		//looks inside our valid file type
		//this will ensure its done properly
		if _, ok := FileTypes[Type]; ok { //found match
			displayType = FileTypes[Type] //renders display
		}

		Range := []*simpletable.Cell{ //renders the type and the amount properly
			{Align: simpletable.AlignLeft, Text: displayType}, {Align: simpletable.AlignLeft, Text: strconv.Itoa(Amount.Files)}, {Align: simpletable.AlignLeft, Text: strconv.Itoa(Amount.Lines)}, {Align: simpletable.AlignLeft, Text: strconv.Itoa(Amount.Bytes)}, {Align: simpletable.AlignLeft, Text: strconv.FormatFloat(float64(Amount.Bytes/125), 'f', 9, 64)}, {Align: simpletable.AlignLeft, Text: strconv.FormatFloat(float64(Amount.Bytes/125000), 'f', 9, 64)}, {Align: simpletable.AlignLeft, Text: strconv.FormatFloat(float64(Amount.Bytes/125000/1000), 'f', 9, 64)},
		}

		//saves into the array properly
		branding.Body.Cells = append(branding.Body.Cells, Range)
	}


	//outputs the linecounter information properly
	fmt.Printf("Reading %s\r\n%s\r\n%s\r\n%s", strings.Join(os.Args[1:], ", "),strings.Repeat("-", len(strings.Join(os.Args[1:], ", "))+9),branding.String(), strings.Repeat("-", len(strings.Join(os.Args[1:], ", "))+9))
}

type RecursiveType struct {
	Files, Lines, Bytes int
}

//implements recursive reading into the frame work
//allows us to access sub folders of folders
//	- users #header
//		- admin #sub
//			- confirmed.go #file
//			- prompted.go #file
//		- kick #sub
//			- success.go #file
//		- syntax.go #file
//this framework will count the amount of lines within a certain file type
//allows us to grab the amount of (example ".go", ".txt") files inside the folders
func Recursive(path string, src map[string]RecursiveType) error {

	//possible file has been located
	//this will now act as a file reader
	if strings.Count(path, ".") > 0 {
		//reads the file dir properly
		//allows us to proper handle safely
		contains, err := ioutil.ReadFile(filepath.Join(path))
		if err != nil { //err handles
			return err
		}

		//upper the first charater only properly
		upper := strings.Join(append([]string{strings.ToUpper(strings.Split(filepath.Ext(path), "")[1])}, strings.Split(filepath.Ext(path), "")[2:]...), "")

		if _, ok := src[upper]; ok { //found old key
			src[upper] = RecursiveType{Files: src[upper].Files + 1, Lines: src[upper].Lines + getLines(contains), Bytes: len(contains)} //adds onto
		} else { //else for new key needed
			src[upper] = RecursiveType{Files: 1, Lines: getLines(contains), Bytes: len(contains)} //new key
		}
		
		//this will ensure its done properly
		return nil //returns the system properly
	}

	//reads the dir properly and safely
	//this will ensure its done properly
	objective, err := ioutil.ReadDir(path)
	if err != nil { //err handles
		return err
	}

	//ranges through the subfolder properly
	//allows us to properly range through the sys
	for _, system := range objective { //ranges through

		//checks for dir properly
		//ensures its done without issues
		if strings.Count(system.Name(), ".") <= 0 {
			
			//err handles the recursive reading of dir
			if err := Recursive(filepath.Join(path, system.Name()), src); err != nil {
				return err
			}; continue
		}

		if strings.Split(system.Name(), "")[0] == "." {
			continue
		}

		//reads the file dir properly
		//allows us to proper handle safely
		contains, err := ioutil.ReadFile(filepath.Join(path, system.Name()))
		if err != nil { //err handles
			return err
		}

		//upper the first charater only properly
		upper := strings.Join(append([]string{strings.ToUpper(strings.Split(filepath.Ext(system.Name()), "")[1])}, strings.Split(filepath.Ext(system.Name()), "")[2:]...), "")


		if _, ok := src[upper]; ok { //found old key
			src[upper] = RecursiveType{Files: src[upper].Files + 1, Lines: src[upper].Lines + getLines(contains), Bytes: src[upper].Bytes + int(system.Size())} //adds onto
		} else { //else for new key needed
			src[upper] = RecursiveType{Files: 1, Lines: getLines(contains), Bytes: int(system.Size())} //new key
		}
	}

	return nil
}

//gets the amount of lines from a file
//this will ensure we can measure the amount
// 1: abc \n 
// 2: 123 \n 
// 3: def \n
// 4: 456
//above would output 3 but we add 4
//this will include the 4th line at this point
func getLines(body []byte) int { //returns an int
	return strings.Count(string(body), "\n") + 1
}
