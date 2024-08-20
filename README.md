# gravitation-model

This project provide a simulation of movement of N particles in a flat space
determined by their mass, initial position, inititial speed, and gravity force.

The scale of representation is `1 pixel : 1e12 meters`

Mass can be represented using constants

- `TON` = 1e3 kg
- `KTON` = 1e6 kg
- `MT` (Mass of earth) = 5.972e24 kg

Gravitational constant is expressed in N m^2 / kg^2

- `G` = 6.67430e-11

# How to use

On the main.go file you should

1. Create a new system

```go
sys := system.New()
```

2. Define your particles, providing position and velocity for both axis, mass, a string key and a color to trace. They should be added on config.json file

3. Your particles will be automatically added to the system

```go
sys.AddParticle(p1)
```

4. Then, the system will be initialized

```go
sys.Build()
```

# Run

In order to run the project, you need to have go intalled on your machine. Then you need to run

`make tidy` to install the dependencies

`make run` to run the go application

# Result

The resulting images will be saved on /images directory, and an output video will be created from this source
