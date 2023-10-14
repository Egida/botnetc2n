package apis

import "errors"

var (
	//ErrUnwantedAuthMethod shown when invalid authmethod given
	ErrUnwantedAuthMethod error = errors.New("unknown authentication method wanted")

	//ErrUnknownUserGiven when the user input is invalid
	ErrUnknownUserGiven error = errors.New("unknown user given")

	//ErrUnknownPassword whent the password is invalid
	ErrUnknownPassword error = errors.New("unknown password")

	//ErrUnknownToken when token is invalid
	ErrUnknownToken error  = errors.New("unknown token")

	//ErrUnknownError when unknown error happened
	ErrUnknownError error = errors.New("unknown error")

	//ErrCantAccess when they cant access api system
	ErrCantAccess error = errors.New("access denied")

	//ErrMissingQueryError when querys are missing
	ErrMissingQueryError error = errors.New("missing query")

	//ErrInvalidTarget error when the target is invalid
	ErrInvalidTarget error = errors.New("invalid target")

	//ErrorTargetBlacklisted when the target is blacklisted
	ErrTargetBlacklisted error = errors.New("target blacklisted")

	//ErrInvalidPort when port is invalid
	ErrInvalidPort error = errors.New("invalid port")

	//ErrUnknownMethod when method is invalid
	ErrUnknownMethod error = errors.New("unknown method")

	//ErrInvalidDuration when the duration is invalid
	ErrInvalidDuration error = errors.New("invalid duration")

	//ErrTargetAlreadyUnderAttack when target is already being attacked
	ErrTargetAlreadyUnderAttack error = errors.New("target already under attack")

	//ErrConnLimitReached when concurrent limit reached
	ErrConnLimitReached error = errors.New("conn limit reached")

	//ErrInCooldown when user is inside cooldown
	ErrInCooldown error = errors.New("you in cooldown")

	//ErrAttackLaunchFault when the attack failed to launch
	ErrAttackLaunchFault error = errors.New("attack launch fault")

	//ErrKeyValueFault when we have failed to parse kv 
	ErrKeyValueFault error = errors.New("key value fault")

	//ErrUserLimitReached when max attacks per method reached
	ErrUserLimitReached error = errors.New("user limit reached")

	//ErrPlanExpired when your plan has expired
	ErrPlanExpired error = errors.New("plan expired")

	//ErrAPIAttacksDisabled when api attacks are disabled
	ErrAPIAttacksDisabled error = errors.New("APIAttacks disabled")
)