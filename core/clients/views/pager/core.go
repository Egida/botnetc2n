package pager

import (
	"Nosviak2/core/clients/sessions"
	"strings"

	"github.com/alexeyco/simpletable"
)

//stores the table information
//this will help us to properly decide if we want to render
type MakeTableRender struct {
	//stores the table name properly
	//this will make sure we know if we are gradienting the table
	header string //stored in type string properly
	//stores the simpletable structure properly
	//this will ensure we have access without issues happening
	table *simpletable.Table //stored in type simpletable.Table
	//stores the session properly
	//this will allow for better handling without issues
	session *sessions.Session
}


//makes the table correctly and properly
//this will ensure its done without errors happening
func MakeTable(header string, table *simpletable.Table, s *sessions.Session) *MakeTableRender {
	return &MakeTableRender{ //returns the structure properly
		header: header, //sets the header properly
		table: table, //sets the table properly
		session: s, //sets the session properly
	}
}

//this will completely and properly texture the table
//some sections will use gcode to properly render out the ranks
func (mtr *MakeTableRender) TextureTable() error { //returns the error
	//this will properly try to render the different type
	//this loads the configuration file without issues happening
	pointer, err := mtr.GetQuery() //gets the query configuration properly
	if err != nil { //basic error handling properly without issues
		return err //returns the err
	}

	//sets the type depending on what was given inside the file
	//this makes sure its done without issues happening on request
	mtr.table.SetStyle(mtr.TypeControl(pointer.Style, false)) //sets style
	//checks if a pager is needed properly
	//this will start the pager seq without issues
	if len(strings.Split(mtr.table.String(), "\n")) >  mtr.session.Height {
		return mtr.Pager(strings.Split(strings.Join(mtr.session.Written, "\r\n"), "\033c")[len(strings.Split(strings.Join(mtr.session.Written, "\r\n"), "\033c"))-1]) //executes the pager properly
	} else { //renders the normal system properly
		return mtr.normal(pointer) //returns the standard properly
	}
}