package pager

import "strings"

//this will properly try to execute
//this allows for better control without issues happening
func (mtr *MakeTableRender) normal(c *TableConfiguration) error {
	var defaultSection string = strings.ReplaceAll(strings.ReplaceAll(mtr.table.String(), "*", " "), "\n", "\r\n") //default
	//checks if the table has gradient
	//this will then follow the gradient route
	if mtr.WantGradient() { //checks for gradient
		//performs the gradient properly and secure
		//this will ensure its done without errors happening
		defaultTab, err := mtr.GradientTable() //performs the gradient properly
		if err != nil { //error handles properly without issues happening on purpose
			return err //returns the error correctly and properly
		}
		//joins the table properly and safely
		//this will be used when rendered without issues happening
		defaultSection = strings.Join(defaultTab, "\r\n")
	} else {
		defaultSection = strings.ReplaceAll(defaultSection, "*", " ")
	}
	//tries to write to the remote host without issues
	//this will ensure its done without errors going wrong on purpose
	if err := mtr.session.Write(defaultSection+"\r\n"); err != nil {
		return err //returns the error correctly and properly without issues
	}

	if c.NewLine {
		//tries to write to the remote host without issues
		//this will ensure its done without errors going wrong on purpose
		if err := mtr.session.Write("\r\n"); err != nil {
			return err //returns the error correctly and properly without issues
		}
	}
	return nil
}