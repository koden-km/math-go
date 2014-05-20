package vec3

import (
	"encoding/json"
	"fmt"
	"math"
)

// This type is based on github.com/cespare/frosty/vec.go
// Vector ops are modeled after math/big.

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) String() string {
	return fmt.Sprintf("%#v", v)
}

func New(x, y, z float64) *Vec3 {
	return &Vec3{x, y, z}
}

func NewCopy(o *Vec3) *Vec3 {
	return &Vec3{o.X, o.Y, o.Z}
}

func Zero() *Vec3 {
	return &Vec3{0, 0, 0}
}

func One() *Vec3 {
	return &Vec3{1, 1, 1}
}

func Up() *Vec3 {
	return &Vec3{0, 1, 0}
}

func Forward() *Vec3 {
	return &Vec3{0, 0, 1}
}

func Left() *Vec3 {
	return &Vec3{1, 0, 0}
}

func (v *Vec3) Copy(o *Vec3) *Vec3 {
	v.X = o.X
	v.Y = o.Y
	v.Z = o.Z
	return v
}

func (v *Vec3) Add(a, b *Vec3) *Vec3 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	return v
}

func (v *Vec3) Sub(a, b *Vec3) *Vec3 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
	return v
}

// Div sets v to be u / x for some scalar x and returns v.
func (v *Vec3) Div(u *Vec3, x float64) *Vec3 {
	v.X = u.X / x
	v.Y = u.Y / x
	v.Z = u.Z / x
	return v
}

// Mul sets v to be u * x for some scalar x and returns v.
func (v *Vec3) Mul(u *Vec3, x float64) *Vec3 {
	v.X = u.X * x
	v.Y = u.Y * x
	v.Z = u.Z * x
	return v
}

// Inv sets v to be the inverse of u and returns v.
func (v *Vec3) Inv(u *Vec3) *Vec3 {
	return v.Mul(v, -1)
}

// Mag returns the magnitude of v.
func (v *Vec3) Mag() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Normalize sets v to be the normalized (unit) vector of u and returns v.
func (v *Vec3) Normalize(u *Vec3) *Vec3 {
	return v.Div(u, u.Mag())
}

// Cross returns the cross product of u and v as a newly allocated vector.
// (This function does not follow the math/big pattern because it wouldn't work if the result vector were also
// one of the operands.)
func (v *Vec3) Cross(u *Vec3) *Vec3 {
	return &Vec3{
		X: (v.Y * u.Z) - (v.Z * u.Y),
		Y: (v.Z * u.X) - (v.X * u.Z),
		Z: (v.X * u.Y) - (v.Y * u.X),
	}
}

// Dot returns the dot product of u and v.
func (v *Vec3) Dot(u *Vec3) float64 {
	return (v.X * u.X) + (v.Y * u.Y) + (v.Z * u.Z)
}

func (v *Vec3) UnmarshalJSON(b []byte) error {
	a := [3]float64{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	v.X, v.Y, v.Z = a[0], a[1], a[2]
	return nil
}
