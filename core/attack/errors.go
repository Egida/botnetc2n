package attacks

import (
	"errors"
	"github.com/bogdanovich/dns_resolver"
)

var (
	//stores if attacks are enabled
	//this will count as a default option properly
	AttacksEnabled bool = true //they are enabled as default

	//stores if api attacks are enabled
	//this will store if api sided users can attack
	APIAttacksEnabled bool = true //they are enabled as default

	//ErrMinSuccessNotMet when the minimum send was not reached
	ErrMinSuccessNotMet error = errors.New("minimum success was not met")



	Resolver *dns_resolver.DnsResolver
)
