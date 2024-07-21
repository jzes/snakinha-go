package form

import "io"

type Terraformer interface {
	Form(io.Writer, string)
}
