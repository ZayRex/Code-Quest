package main

import "fmt"

func main() {
	/* array of size 4*/
	firstNames := [4]string{"Joe", "Ahmad", "Grace", "Rami"}

	for i := 0; i < len(firstNames); i++ {
		firstName := firstNames[i]
		fmt.Println(i, firstName)
	}

	/* array of size 4, decided by compiler*/
	lastNames := [...]string{"Adams", "Rays", "Greens", "Simons"}

	for i, lastName := range lastNames {
		fmt.Println(i, lastName)
	}

	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
	}

	/* slices */
	classics := languages[:3]
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // alternatively languages[7:]

	fmt.Printf("classic languagues: %v\n", classics) // classic languagues: [C Lisp C++]
	fmt.Printf("modern languages: %v\n", modern)     // modern languages: [Java Python JavaScript Ruby]
	fmt.Printf("new languages: %v\n", new)           // new languages: [Go Rust]
}
