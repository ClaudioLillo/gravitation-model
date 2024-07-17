package system

import (
	"image/color"
	"log"
	"math"

	"github.com/claudiolillo/gravitation-model/internal/constants"
	"github.com/claudiolillo/gravitation-model/internal/utils"
)

// Position will be expressed in 1:1e9 meters
// Time expressed in seconds
// Mass expressed in 1:1e12 kg
// Speed expressed in 1:1e3 m/s

// Time interval = 1 year
var T = 1.0

var SCALE = 1e-12

type Particle struct {
	X, Y, Vx, Vy float64
	Color color.Color
	Key string
	Context []string
	Mass float64
}

type System struct {
	Particles map[string]*Particle
}


type SystemInterface interface {
	AddParticle(p *Particle)
}

type Modifier struct {
	AddX float64
	AddY float64
	AddVX float64
	AddVY float64
}


func (s *System) AddParticle(p *Particle ) {
	s.Particles[p.Key] = p
}

func (s *System) Describe(){
	log.SetPrefix("* ")
	for key, value := range s.Particles {
		log.Println(key)
		log.Printf("p: %s, x: %f, y: %f, vx: %f, vy: %f", value.Key, value.X, value.Y, value.Vx, value.Vy)
		log.Println("Context: ", value.Context)
	}
}

func New() (*System) {
	s := &System{}
	s.Particles = make(map[string]*Particle)
	return s
}

func (s *System) Build() {
	for key, value := range s.Particles {
		for key2, _ := range s.Particles {
			if key != key2 {
				value.Context = append(value.Context, key2)
			}
		}
	}
}

func (s *System) Next() {
	modifMap := make(map[string]Modifier)
	axSum := 0.0
	aySum := 0.0
	for key, value := range s.Particles {
		for _, key2 := range value.Context {
			p2 := s.Particles[key2]
			fx, fy := Force(value, p2)
			axSum += fx / value.Mass
			aySum += fy / value.Mass
		}
		// This allows a simetrical calculation based on the current position and speed
		modifier := Modifier{
			AddX: utils.Truncate((value.Vx * T) + (axSum * T * T / 2) * SCALE, 0.001),
			AddY: utils.Truncate((value.Vy * T) + (aySum * T * T / 2) * SCALE, 0.001),
			AddVX: utils.Truncate(axSum * T * SCALE, 0.001),
			AddVY: utils.Truncate(aySum * T * SCALE, 0.001),
		}

		modifMap[key] = modifier
		axSum = 0
		aySum = 0	
	}

	for key, val := range modifMap {
		p := s.Particles[key]
		p.X += val.AddX
		p.Y += val.AddY
		p.Vx += val.AddVX
		p.Vy += val.AddVY
	}

}

func Force(p1, p2 *Particle) (float64, float64) {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	g := constants.G

	// Adds a limitation to ensure that no divission by zero occurs
	if dx < 10 && dx > -10 {
		dx = 10 * math.Abs(dx)
	}

	if dy < 10 && dy > -10 {
		dy = 10 * math.Abs(dx)
	}

	r2 := (dx * dx) + (dy * dy)
	r := math.Sqrt(r2)

	f := g * p1.Mass  * p2.Mass / r2

	fx := f * dx / r
	fy := f * dy / r
	
	return fx, fy
}
