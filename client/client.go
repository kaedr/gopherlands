package client

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

// StartClient gets the game client running
func StartClient() {

	// Create application and scene
	a := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	clrs := make([]*math32.Color, 6)
	clrs[0] = math32.NewColor("Red")
	clrs[1] = math32.NewColor("DarkOrange")
	clrs[2] = math32.NewColor("Yellow")
	clrs[3] = math32.NewColor("Green")
	clrs[4] = math32.NewColor("Blue")
	clrs[5] = math32.NewColor("DarkViolet")
	mats := make([]*material.Standard, 6)

	txtrs := make([]*texture.Texture2D, 6)
	txtrs[0], _ = texture.NewTexture2DFromImage("assets/test/0_one.png")
	txtrs[1], _ = texture.NewTexture2DFromImage("assets/test/1_two.png")
	txtrs[2], _ = texture.NewTexture2DFromImage("assets/test/2_three.png")
	txtrs[3], _ = texture.NewTexture2DFromImage("assets/test/3_four.png")
	txtrs[4], _ = texture.NewTexture2DFromImage("assets/test/4_five.png")
	txtrs[5], _ = texture.NewTexture2DFromImage("assets/test/5_six.png")
	mats2 := make([]*material.Standard, 6)

	// Create a colored cube and add it to the scene
	geom := geometry.NewCube(1)
	mesh := graphic.NewMesh(geom, nil)
	for i := 0; i < 6; i++ {
		mats[i] = material.NewStandard(clrs[i])
		mesh.AddGroupMaterial(mats[i], i)
	}
	mesh.SetPosition(1, 0, 0)
	scene.Add(mesh)

	// Create a colored cube and add it to the scene
	geom2 := geometry.NewCube(1)
	mesh2 := graphic.NewMesh(geom2, nil)
	for i := 0; i < 6; i++ {
		mats2[i] = material.NewStandard(math32.NewColor("white"))
		mats2[i].AddTexture(txtrs[i])
		mesh2.AddGroupMaterial(mats2[i], i)
	}
	mesh2.SetPosition(0, 1, 0)
	scene.Add(mesh2)

	// keep track of current color offset
	colOffset := 0
	// Create and add a button to the scene
	btn := gui.NewButton("Advance Color")
	btn.SetPosition(200, 40)
	btn.SetSize(40, 40)
	btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		colOffset++
		for i := 0; i < 6; i++ {
			mats[i].SetColor(clrs[(i+colOffset)%6])
		}
	})
	scene.Add(btn)

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{R: 1.0, G: 1.0, B: 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{R: 1, G: 1, B: 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	scene.Add(pointLight)

	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))

	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}
