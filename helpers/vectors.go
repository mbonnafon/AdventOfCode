package helpers

type Vector3D struct {
	X, Y, Z int
}

type Vector4D struct {
	X, Y, Z, W int
}

func (v Vector3D) Neighboor() []Vector3D {
	var neighboor []Vector3D
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				nb := v
				nb.Move(Vector3D{x, y, z})
				neighboor = append(neighboor, nb)
			}
		}
	}
	return neighboor
}

// Move vector to param vector direction
func (v *Vector3D) Move(dir Vector3D) {
	v.X += dir.X
	v.Y += dir.Y
	v.Z += dir.Z
}

func (v Vector4D) Neighboor() []Vector4D {
	var neighboor []Vector4D
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					nb := v
					nb.Move(Vector4D{x, y, z, w})
					neighboor = append(neighboor, nb)
				}
			}
		}
	}
	return neighboor
}

// Move vector to param vector direction
func (v *Vector4D) Move(dir Vector4D) {
	v.X += dir.X
	v.Y += dir.Y
	v.Z += dir.Z
	v.W += dir.W
}
