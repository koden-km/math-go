math-go
========

Math library for the Go language.

```go
// Example...
package main

import (
    "fmt"
    //"github.com/koden-km/math-go/vec2"
    "github.com/koden-km/math-go/vec3"
    //"github.com/koden-km/math-go/vec4"
    //"github.com/koden-km/math-go/mat3"
    "github.com/koden-km/math-go/mat4"
    //"github.com/koden-km/math-go/quat"
)

func main() {
	var v3a := vec3.Zero()				// Make a new Vec3 that is all zero
	var v3b := vec3.New(x, y, z)		// Make a new Vec3 that has x,y,z set to the params.
	var v3c := vec3.NewCopy(v3b)		// Make a new Vec3 that is a copy of v3b

	v3c.Copy(v3b)						// Alternate way to make copy (mutates v3c to be a copy of v3b)
	v3c.Zero()							// Zero all values.




	var m4a := mat4.New()				// New all zero matrix4x4
	var m4b := mat4.Identity()		    // Like New, but default to identity
}
```
