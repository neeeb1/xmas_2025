package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/neeeb1/xmas_2025/internal/app"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	grinch          app.Grinch
	lastHungerDecay time.Time
	lastEnergyDecay time.Time
}

// Called every tick (1/60s by default)
func (g *Game) Update() error {
	minutesElapsed := int(time.Since(g.grinch.Created).Minutes())
	if minutesElapsed > g.grinch.Age {
		g.grinch.Age = minutesElapsed
	}

	if int(time.Since(g.lastHungerDecay).Seconds()) > app.HungerDecayInterval {
		g.grinch.Hunger -= 1
		g.lastHungerDecay = time.Now()
	}

	if int(time.Since(g.lastEnergyDecay).Seconds()) > app.EnergyDecayInterval {
		g.grinch.Energy -= 1
		g.lastEnergyDecay = time.Now()
	}

	return nil
}

// Called every frame (1/60s for 60hz display)
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill((color.RGBA{192, 192, 192, 255}))
	app.DrawProgressBar(screen, 5, 15, 100, 10, float64(g.grinch.Energy), 100, "Energy")
	app.DrawProgressBar(screen, 5, 40, 100, 10, float64(g.grinch.Hunger), 100, "Hunger")
	app.DrawProgressBar(screen, 5, 65, 100, 10, float64(g.grinch.Happiness), 100, "Happiness")
	text.Draw(screen, fmt.Sprintf("Age: %d", g.grinch.Age), basicfont.Face7x13, 5, 90, color.Black)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	game.grinch = app.Grinch{
		Age:       0,
		Hunger:    100,
		Happiness: 50,
		Energy:    100,
		Stage:     app.StageBaby,
		Created:   time.Now(),
		IsAlive:   true,
	}
	game.lastEnergyDecay = time.Now()
	game.lastHungerDecay = time.Now()

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(860, 860)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
