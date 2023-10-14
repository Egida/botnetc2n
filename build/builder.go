package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

/*
	creates a brandnew Nosviak2 compiled build set for client
	this will automatically compile and obf the needed files for success
*/

func main() {

	target_fingerprint := flag.String("fingerprint", "", "used for markdown mainly") //license fingerprint
	target_username    := flag.String("username", "", "used for markdown mainly") 
	target_version     := flag.String("version", "1.0", "used for markdown mainly") 
	srcPath		       := flag.String("srcpath", "", "location of the source code folders")
	markdown		   := flag.String("markdown", "", "location of the markdown folder needed for compression")
	GoRoot			   := flag.String("goroot", "", "location of the goroot")
	GoPath			   := flag.String("gopath", "", "location of the GoPath")
	FormatFiles		   := flag.Bool("compile", true, "Controls weather if you want to compile or not")
	flag.Parse() //parses the flags properly

	//creates our build id properly and outputs to terminal
	build := hex.EncodeToString(md5.New().Sum([]byte(*target_username)))
	fmt.Printf("\x1b[34m[1/11]\x1b[0m: Created BuildID: %s\r\n", build)

	//changes the buildid properly within the build structure
	if err := ChangeBUILD(filepath.Join(*srcPath, "core", "configs", "version.go"), build); err != nil {
		fmt.Printf("\x1b[34m[2/11]\x1b[0m: Fault updating version.go: %s\r\n", err.Error())
		//return
	} else {
		fmt.Printf("\x1b[34m[2/11]\x1b[0m: correctly updated the BuildID found inside version.go\r\n")
	}

	//opens the markdown dir properly and safely
	markdownDir, err := ioutil.ReadDir(*markdown)
	if err != nil {
		fmt.Printf("\x1b[34m[3/11]\x1b[0m: error opening %s: %s\r\n", *markdown,err.Error())
		//return
	}

	target := "Nosviak2-"+*target_username+"@"+*target_version
	os.Mkdir(target, os.ModePerm) //tries to create copy compile

	//ranges through all files properly
	for rank, system := range markdownDir { //reads the file body properly
		contains, err := ioutil.ReadFile(filepath.Join(*markdown, system.Name()))
		if err != nil {
			continue
		}

		//replaces all information we want properly
		contains = []byte(strings.ReplaceAll(string(contains), "<<$fingerprint>>", *target_fingerprint))
		contains = []byte(strings.ReplaceAll(string(contains), "<<$username>>",    *target_username))
		contains = []byte(strings.ReplaceAll(string(contains), "<<$version>>",     *target_version))
		
		//opens the future possible file used for docs
		peck, err := os.Create(target+"/"+system.Name())
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		//writes the bytes
		peck.Write(contains) //renders a little message for success
		fmt.Printf("\t\x1b[34m[3][%d/%d]\x1b[0m moved %s\r\n", rank+1, len(markdownDir), filepath.Join(*markdown, system.Name()))
	}

	if !*FormatFiles {
		return
	}

	//moves the cursor properly
	if err := os.Chdir(*GoPath); err != nil {
		fmt.Printf("\x1b[34m[4/11]\x1b[0m: error when moving current outlook to %s: %s\r\n", *GoPath, err.Error()); //return
	} else {
		fmt.Printf("\x1b[34m[4/11]\x1b[0m: moved current outlook to %s\r\n", *GoPath)
	}

	//tries to remove current driver there
	if err := os.RemoveAll(*GoRoot+"/src/Nosviak2"); err != nil {
		fmt.Printf("\x1b[34m[6/11]\x1b[0m: error when removing %s: %s\r\n", *GoRoot+"/src/Nosviak2", err.Error())
	} else {
		fmt.Printf("\x1b[34m[6/11]\x1b[0m: removed %s\r\n", *GoRoot+"/src/Nosviak2")
	}

	//creates the dir properly and safely
	if err := os.Mkdir(*GoRoot+"/src/Nosviak2", os.ModePerm); err != nil {
		fmt.Printf("\x1b[34m[7/11]\x1b[0m: error when creating dir %s: %s\r\n", *GoRoot+"/src/Nosviak2", err.Error())
	} else {
		fmt.Printf("\x1b[34m[7/11]\x1b[0m: correctly made the dir %s\r\n", *GoRoot+"/src/Nosviak2")
	}

	//copies the information properly
	if err := copy(*GoPath+"/src/Nosviak2", *GoRoot+"/src/Nosviak2"); err != nil {
		fmt.Printf("\x1b[34m[8/11]\x1b[0m: error when coping current outlook to %s -> %s: %s\r\n", *GoPath+"/src/Nosviak2", *GoRoot+"/src/Nosviak2", err.Error()); //return
	} else {
		fmt.Printf("\x1b[34m[8/11]\x1b[0m: moved %s -> %s\r\n", *GoPath+"/src/Nosviak2", *GoRoot+"/src/Nosviak2")
	}

	//moves the cursor properly
	if err := os.Chdir("/"+*GoPath+"/src/Nosviak2/compression"); err != nil {
		fmt.Printf("\x1b[34m[9/11]\x1b[0m: error when moving current outlook to %s: %s\r\n", *GoPath+"/src/Nosviak2", err.Error()); //return
	} else {
		fmt.Printf("\x1b[34m[9/11]\x1b[0m: moved current outlook to %s\r\n", *GoPath+"/src/Nosviak2")
	}

	//copies the information properly
	if err := copy("/"+*GoPath+"/src/Nosviak2/compression/garble", "/"+*GoPath+"/src/Nosviak2"); err != nil {
		fmt.Printf("\x1b[34m[10/11]\x1b[0m: error when coping current outlook to %s -> %s: %s\r\n", "/"+*GoPath+"/src/Nosviak2/compression/garble", "/"+*GoPath+"/src/Nosviak2", err.Error()); //return
	} else {
		fmt.Printf("\x1b[34m[10/11]\x1b[0m: moved %s -> %s\r\n", "/"+*GoPath+"/src/Nosviak2/compression/garble", "/"+*GoPath+"/src/Nosviak2")
	}

	//moves the cursor properly
	if err := os.Chdir("/"+*GoPath+"/src/Nosviak2"); err != nil {
		fmt.Printf("\x1b[34m[11/11]\x1b[0m: error when moving current outlook to %s: %s\r\n", "/"+*GoPath+"/src/Nosviak2", err.Error()); //return
	} else {
		fmt.Printf("\x1b[34m[11/11]\x1b[0m: moved current outlook to %s\r\n", "/"+*GoPath+"/src/Nosviak2")
	}

	fmt.Printf("\x1b[34m[xx/11]\x1b[0m: \x1b[38;5;10mCompleted\x1b[0m the environment setup for install!\r\n")
	fmt.Printf("         Run:\r\n")
	fmt.Printf("           > cd Nosviak2")
	fmt.Printf("           > ./garble -literals -debug -tiny -seed=random build -ldflags=\"-w -s\" Nosviak2\r\n")
	fmt.Printf("           > strip Nosviak2\r\n")
	fmt.Printf("           > upx -9 -1 Nosviak2\r\n")
}


//ChangeBUILD will change the current builds id
func ChangeBUILD(path string, id string) error {

	//reads the path towards the file
	contains, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	//stores the array of commanders properly
	var caught []string = make([]string, 0)

	//ranges through the file line by line looking for it
	for _, line := range strings.Split(string(contains), "\r\n") {
		
		//checks for the build id inside here
		if strings.Contains(line, "BuildID") {
			//updates the line and ensures its done without errors
			line = strings.SplitAfter(line, "=")[0] + " " + "\"" + id + "\""
		}

		//saves into the array properly
		caught = append(caught, line)
	}

	//creates the file which we will edit
	writer, err := os.Create(path)
	if err != nil {
		return err
	}

	defer writer.Close()
	//writes into the file properly
	if _, err := writer.WriteString(strings.Join(caught, "\r\n")); err != nil {
		return err
	}

	return nil
}

func copy(source, destination string) error {
	var proc int = 0
    var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
        var relPath string = strings.Replace(path, source, "", 1)
        if relPath == "" {
            return nil
        }
		proc++
		fmt.Printf("\t\x1b[34m[8/11][%d]\x1b[0m: moved %s -> %s\r\n", proc, path, destination)
        if info.IsDir() {
            return os.Mkdir(filepath.Join(destination, relPath), 0755)
        } else {
            var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
            if err1 != nil {
                return err1
            }
            return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
        }
    })
    return err
}

func generate(n int) string {
    var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
    str := make([]rune, n)
    for i := range str {
        str[i] = chars[rand.Intn(len(chars))]
    }
    return string(str)
}