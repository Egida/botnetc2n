package clients

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/layouts/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

//tries to listen to the addr given
//this will ensure its done properly without errors
func ProduceClient() error { //returns error properly


	//creates the configuration for the server
	//this will ensure its done without issues happening
	Config := ssh.ServerConfig{
		//stores information about the password callback
		//this will be the main authentication control without errors
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			//tries to find the user inside the database
			//this will make sure its done properly without issues
			user, err := database.Conn.FindUser(conn.User())

			//tries to correctly insert into the database
			//this will make sure its logged properly without issues
			if err := database.Conn.LoginAttempt(strings.Join(strings.Split(conn.RemoteAddr().String(), ":")[:len(strings.Split(conn.RemoteAddr().String(), ":"))-1], ":"), string(conn.ClientVersion()), conn.User(), err == nil && user != nil && user.Password == database.HashProduct(string(password))); err != nil {
				return nil, fmt.Errorf(fmt.Sprintf("invalid authentication attempt from %s for %s", conn.RemoteAddr().String(), conn.User())) //invalid login attempt properly
			}

			//tries to see if the user can access
			//this will ensure its done without issues happening
			if err != nil || user == nil || user.Password != database.HashProduct(string(password)) { //error handles the request properly
				return nil, fmt.Errorf("invalid authentication attempt from %s for %s", conn.RemoteAddr().String(), conn.User())
			}
			//this will properly get the password and the hashed product for the password given
			//this will be mainly used for authenticating users into the systems without issues happening
			if user.Password == database.HashProduct(string(password)) {
				return nil, nil  //returns nil and nil as its valid!
			}

			//returns the invalid auth error properly
			//this will ensure its done without any errors
			return nil, fmt.Errorf("invalid authentication attempt from %s for %s", conn.RemoteAddr().String(), conn.User())
		},

		//sets the default serverVersion
		//this will make sure its done without issues
		ServerVersion: fmt.Sprintf("SSH-2.0-OpenSSH_8.6p1 %s@%s", "Nosviak2", deployment.Version),
	}


	//sets the max auth attempts properly
	//this will make sure its not ignored without issues
	Config.MaxAuthTries = json.ConfigSettings.Masters.MaxAuthAttempts

	//checks for the custom authentication menu
	//this will enable no client auth properly without issues
	if json.ConfigSettings.Masters.Server.DynamicAuth { //checks for dynamic auth
		Config.NoClientAuth = true //sets the dynamic auth option properly
	}

	//checks to see if the client is using before auth
	//this will allow the client to display a message of there choice inside the value
	if json.ConfigSettings.Masters.BeforePasswdPrompt.Status { //detects if its been enabled
		//function which returns the value for the system
		//this will return the custom callback message for the server
		Config.BannerCallback = func(conn ssh.ConnMetadata) string {
			//returns the custom message properly without issues
			return json.ConfigSettings.Masters.BeforePasswdPrompt.Message
		}
	}


	//reads the server key file properly without issues
	//this will control the servers master key properly without issues
	serverMaster, err := ioutil.ReadFile(filepath.Join(deployment.Assets, json.ConfigSettings.Masters.ServerKey))
	if err != nil { //error handles the request properly without issues
		return err //returns the error correctly
	}

	//tries to parse the ssh signature
	//this will be used for the server properly without issues
	signature, err := ssh.ParsePrivateKey(serverMaster)
	if err != nil { //error handles the request properly without issues
		return err //returns the error correctly
	}

	//saves the signature into the structure
	//this will be used inside the ssh listener properly
	Config.AddHostKey(signature)

	//starts the ssh protocol listener properly
	//this will make sure its done correctly without issues happening
	return CreateListener(Config)
}