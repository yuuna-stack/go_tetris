package main

import (
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/yuuna-stack/go_tetris/wrapper"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const resourcesDir = "images"

const M = 20
const N = 10

type Point struct {
	x int
	y int
}

func init() { runtime.LockOSThread() }

func check(field [M][N]int, a [4]Point) bool {
	for i := 0; i < 4; i++ {
		if a[i].x < 0 || a[i].x >= N || a[i].y >= M {
			return false
		} else if field[a[i].y][a[i].x] != 0 {
			return false
		}
	}
	return true
}

func fullname(filename string) string {
	return path.Join(resourcesDir, filename)
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	resources := wrapper.Resources{}

	const gameWidth = 320
	const gameHeight = 480

	vm := window.NewSfVideoMode()
	defer window.DeleteSfVideoMode(vm)
	vm.SetWidth(gameWidth)
	vm.SetHeight(gameHeight)
	vm.SetBitsPerPixel(32)

	cs := window.NewSfContextSettings()
	defer window.DeleteSfContextSettings(cs)
	w := graphics.SfRenderWindow_create(vm, "Tetris!", uint(window.SfResize|window.SfClose), cs)
	defer window.SfWindow_destroy(w)

	ev := window.NewSfEvent()
	defer window.DeleteSfEvent(ev)

	s, err := wrapper.FileToSprite(fullname("tiles.png"), &resources)
	if err != nil {
		panic("Couldn't load tiles.png")
	}
	background, err := wrapper.FileToSprite(fullname("background.png"), &resources)
	if err != nil {
		panic("Couldn't load background.png")
	}
	frame, err := wrapper.FileToSprite(fullname("frame.png"), &resources)
	if err != nil {
		panic("Couldn't load frame.png")
	}

	var a [4]Point
	var b [4]Point
	var field [M][N]int

	figures := [7][4]int{
		{1, 3, 5, 7},
		{2, 4, 5, 7},
		{3, 5, 4, 6},
		{3, 5, 4, 7},
		{2, 3, 5, 7},
		{3, 5, 7, 6},
		{2, 3, 4, 5},
	}

	dx := 0
	rotate := false
	colorNum := 1
	timer := 0.0
	delay := 0.3

	timeStamp := time.Now()
	for window.SfWindow_isOpen(w) > 0 {
		for window.SfWindow_pollEvent(w, ev) > 0 {
			if ev.GetEvType() == window.SfEventType(window.SfEvtClosed) {
				return
			}

			if ev.GetEvType() == window.SfEventType(window.SfEvtKeyPressed) {
				if ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeyUp) {
					rotate = true
				} else if ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeyLeft) {
					dx -= 1
				} else if ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeyRight) {
					dx += 1
				} else if ev.GetKey().GetCode() == window.SfKeyCode(window.SfKeyDown) {
					delay = 0.05
				}
			}
		}
		deltaTime := float64(time.Now().Sub(timeStamp).Seconds())
		timeStamp = time.Now()
		timer += deltaTime

		for i := 0; i < 4; i++ {
			b[i] = a[i]
			a[i].x += dx
		}

		if !check(field, a) {
			for i := 0; i < 4; i++ {
				a[i] = b[i]
			}
		}

		if rotate {
			p := a[1]
			for i := 0; i < 4; i++ {
				x := a[i].y - p.y
				y := a[i].x - p.x
				a[i].x = p.x - x
				a[i].y = p.y + y
			}
			if !check(field, a) {
				for i := 0; i < 4; i++ {
					a[i] = b[i]
				}
			}
		}

		if timer > delay {
			for i := 0; i < 4; i++ {
				b[i] = a[i]
				a[i].y += 1
			}
			if !check(field, a) {
				for i := 0; i < 4; i++ {
					field[b[i].y][b[i].x] = colorNum
				}
				colorNum = 1 + r1.Int()%7
				n := r1.Int() % 7
				for i := 0; i < 4; i++ {
					a[i].x = figures[n][i] % 2
					a[i].y = figures[n][i] / 2
				}
			}
			timer = 0
		}

		k := M - 1
		for i := M - 1; i > 0; i-- {
			count := 0
			for j := 0; j < N; j++ {
				if field[i][j] != 0 {
					count++
				}
				field[k][j] = field[i][j]
			}
			if count < N {
				k--
			}
		}

		dx = 0
		rotate = false
		delay = 0.3

		graphics.SfRenderWindow_clear(w, graphics.GetSfBlack())

		background.Draw(w)

		for i := 0; i < M; i++ {
			for j := 0; j < N; j++ {
				if field[i][j] == 0 {
					continue
				}
				s.SetTextureRect(field[i][j]*18, 0, 18, 18)
				s.SetPosition(float32(j*18), float32(i*18))
				s.Move(float32(28), float32(31))
				s.Draw(w)
			}
		}

		for i := 0; i < 4; i++ {
			s.SetTextureRect(colorNum*18, 0, 18, 18)
			s.SetPosition(float32(a[i].x*18), float32(a[i].y*18))
			s.Move(float32(28), float32(31))
			s.Draw(w)
		}

		frame.Draw(w)

		graphics.SfRenderWindow_display(w)
	}

	resources.Clear()
}
