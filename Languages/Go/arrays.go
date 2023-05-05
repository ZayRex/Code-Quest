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

	////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////

	allLangs := languages[:]                      // copy of the array
	fmt.Println(reflect.TypeOf(allLangs).Kind())   // slice

	/* Create a slice containing web frameworks */
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]          // length 4 capacity 4
	frameworks = append(frameworks, "Meteor")  // not possible with arrays

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrameworks)

	//////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////
	////////// MAPS///////////
	
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}
}
