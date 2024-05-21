package mml

type LightType string

const (
	LightTypePoint     LightType = "point"
	LightTypeSpotlight LightType = "spotlight"
)

// TODO: generate this from the schema?
type Light struct {
	X         float64   `structs:"x,omitempty"`
	Y         float64   `structs:"y,omitempty"`
	Z         float64   `structs:"z,omitempty"`
	RX        float64   `structs:"rx,omitempty"`
	RY        float64   `structs:"ry,omitempty"`
	RZ        float64   `structs:"rz,omitempty"`
	SX        float64   `structs:"sx,omitempty"`
	SY        float64   `structs:"sy,omitempty"`
	SZ        float64   `structs:"sz,omitempty"`
	Visible   string    `structs:"visible,omitempty"` // I don't think this can be a boolean because the absence of this attr is means it _is_ vsible??
	Socket    string    `structs:"socket,omitempty"`
	Debug     bool      `structs:"debug,omitempty"`
	Color     string    `structs:"color,omitempty"` // todo make this typesafe RGB, NameColor, etc
	ID        string    `structs:"id,omitempty"`
	Class     string    `structs:"class,omitempty"`
	Intensity float64   `structs:"intensity,omitempty"`
	Distance  float64   `structs:"distance,omitempty"`
	Angle     float64   `structs:"angle,omitempty"`
	Type      LightType `structs:"type,omitempty"`
}

type Alignment string

const (
	AlignmentLeft   Alignment = "left"
	AlignmentCenter Alignment = "center"
	AlignmentRight  Alignment = "right"
)
