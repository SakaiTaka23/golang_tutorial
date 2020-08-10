package main

import (
	"fmt"
)

const (
	isAdmin          = 1 << iota // 1
	isHeadquarters               // 10
	canSeeFinancials             // 100

	canSeeAfrica       // 1000
	canSeeAsia         // 10000
	canSeeEurope       // 100000
	canSeeNorthAmerica // 1000000
	canSeeSouthAmerica // 10000000
)

func main() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

	roles = isHeadquarters | canSeeAfrica
	fmt.Printf("%b\n", roles)

	roles = isHeadquarters | canSeeSouthAmerica
	fmt.Printf("%b\n",roles)
}
