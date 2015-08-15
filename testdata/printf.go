// Test for formating errors

// Package foo ...
package foo

import "log"

func lintErrorsWithLogPackage() {
	log.Printf("1: %s: 2: MISSING", 1, 2)      // MATCH /.*verbs.*less.*/
	log.Printf("1: %#v: 2: MISSING", 1, 2)     // MATCH /.*verbs.*less.*/
	log.Printf("1: %#+v: 2: MISSING", 1, 2)    // MATCH /.*verbs.*less.*/
	log.Printf("1: %.2f: 2: MISSING", 1, 2)    // MATCH /.*verbs.*less.*/
	log.Printf("1: %1.2f: 2: MISSING", 1, 2)   // MATCH /.*verbs.*less.*/
	log.Printf("1: %1f: 2: MISSING", 1, 2)     // MATCH /.*verbs.*less.*/
	log.Printf("1: %-30s: 2: MISSING", 1, 2)   // MATCH /.*verbs.*less.*/
	log.Printf("1: %30s: 2: MISSING", 1, 2)    // MATCH /.*verbs.*less.*/
	log.Printf("1: MISSING: 2: MISSING", 1, 2) // MATCH /.*verbs.*less.*/
	log.Printf("1: %s: 2: %s 3: %s", 1, 2)     // MATCH /.*verbs.*more.*/

	// everything is okay
	log.Printf("1: %s: 2: %s 3: %s", 1, 2, 3)
	log.Printf("blah: %s %%", "string")
	// z is unknown verb, so should skip
	log.Printf("z? %z:")
}
