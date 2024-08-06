package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/claudiolillo/gravitation-model/internal/system"
	"github.com/icza/mjpeg"
)

var wg sync.WaitGroup

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func addPixel(image *image.NRGBA ,particle *system.Particle, stroke int){
	for i:=0;i<stroke;i++{
		for j:=0;j<stroke;j++{
			image.Set(int(particle.X) + i, int(particle.Y) + j, particle.Color)
		}
	}
}

func getPercentaje(fn string, total int) (int){
	start := strings.Index(fn, "-") + 1
	end := strings.Index(fn,".")
	value, err := strconv.Atoi(fn[start:end])
	if err != nil {
		fmt.Println("Not able to get percentaje")
	}
	return (value + 1) * 100/total
}

func main() {

	jsonFile, err := os.Open("config.json")

	if err!=nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, readErr := io.ReadAll(jsonFile)

	if readErr!=nil{
		fmt.Println(readErr)
	}

	var config system.Config

	json.Unmarshal(byteValue, &config)

	sys := system.New()

	for _,value := range(config.Particles){
		p := system.GetParticleFromConfig(&value)
		sys.AddParticle(&p)
	}

	// The system is build
	sys.Build()

	runSystem := func(c chan *system.System, sys *system.System ){
		defer wg.Done()
		for i:=0; i< config.Iterations; i++ {
			sys.Next()
			c <- sys
		}
		close(c)
	}

	states := make(chan *system.System)
	wg.Add(1)
	go runSystem(states, sys)

	saveImages := func(c chan string, states chan *system.System){
		defer wg.Done()
		count := 0
		for state := range states {
			rect := image.Rect(0,0,config.Video.Size.X,config.Video.Size.Y)
			img := image.NewNRGBA(rect)
			for _, particle := range state.Particles {
				addPixel(img, particle, 5)
			}
			fn := fmt.Sprintf("images/image-%s.jpg",strconv.Itoa(count))
			imgFile,_ := os.Create(fn)
			defer imgFile.Close()
			c <- fn
	
		jpeg.Encode(imgFile, img, nil)
		count++
	}
		close(c)
	}

	fileNames := make(chan string)
	wg.Add(1)
	go saveImages(fileNames, states)
	
	aw, err := mjpeg.New(config.Video.Filename, int32(config.Video.Size.X),int32(config.Video.Size.Y),int32(config.Video.FPS))

	if err!=nil {
		panic(err)
	}


	for f := range fileNames {
		percentaje := getPercentaje(f, config.Iterations)
		bars := percentaje/2
		for c:=0;c<bars;c++{
			fmt.Print("▮")
		}
	
		for c:=0;c<50 - bars;c++{
			fmt.Print(" ")
		}

		fmt.Printf(" %s%% ", strconv.Itoa(percentaje))
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		data, err := os.ReadFile(f)
		
		if err!=nil {
			panic(err)
		}
		checkErr(aw.AddFrame(data))
	}

	wg.Wait()
	checkErr(aw.Close())

		
			
		
	
		

		
		
		
	

	

	

	
}
