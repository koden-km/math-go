// Example...
package main

import (
    "fmt"
    "github.com/koden-km/math-go/mat3"
    "github.com/koden-km/math-go/mat4"
)

func main() {
	var m3a := mat3.New()				// New all zero mat3x3
	var m3b := mat3.Identity()		    // Like New, but default to identity

	var m4a := mat4.New()				// New all zero mat4x4
	var m4b := mat4.Identity()		    // Like New, but default to identity
}
