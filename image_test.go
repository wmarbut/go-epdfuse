// Copyright 2017 Whit Marbut. All rights reserved.
// License information may be found in the LICENSE file.

package epdfuse

import (
	"bytes"
	"image"
	"image/png"
	"testing"
)

func TestDetectScalePlacePlan(t *testing.T) {
	fuse1 := NewEpdFuse()
	fuse2 := NewCustomEpdFuse("./", 300, 100)
	fuse3 := NewCustomEpdFuse("./", 300, 96)
	fuse4 := NewCustomEpdFuse("./", 100, 96)
	img := getTestImage()

	var plan ScalePlan
	plan = fuse1.detectScalePlacePlan(img)
	if plan != SCALE_NO {
		t.Errorf("Error detecting scale plan, expected %d, but got %d", SCALE_NO, plan)
	}

	plan = fuse2.detectScalePlacePlan(img)
	if plan != SCALE_UP {
		t.Errorf("Error detecting scale plan, expected %d, but got %d", SCALE_UP, plan)
	}

	plan = fuse3.detectScalePlacePlan(img)
	if plan != SCALE_PLACE_ONLY {
		t.Errorf("Error detecting scale plan, expected %d, but got %d", SCALE_PLACE_ONLY, plan)
	}

	plan = fuse4.detectScalePlacePlan(img)
	if plan != SCALE_DOWN {
		t.Errorf("Error detecting scale plan, expected %d, but got %d", SCALE_DOWN, plan)
	}
}

func TestContrainingAxis(t *testing.T) {
	fuse1 := NewEpdFuse()                    //2.083 r
	fuse2 := NewCustomEpdFuse("./", 100, 96) //1.041 r
	fuse3 := NewCustomEpdFuse("./", 300, 96) //3.125 r
	img := getTestImage()                    //2.083 r

	var axis Axis

	axis = fuse1.constrainingAxis(img)
	if axis != AXIS_Y {
		t.Errorf("Error detecting constraining axis. Expected %d, but got %d", AXIS_Y, axis)
	}

	axis = fuse2.constrainingAxis(img)
	if axis != AXIS_X {
		t.Errorf("Error detecting constraining axis. Expected %d, but got %d", AXIS_X, axis)
	}

	axis = fuse3.constrainingAxis(img)
	if axis != AXIS_Y {
		t.Errorf("Error detecting constraining axis. Expected %d, but got %d", AXIS_Y, axis)
	}
}

func TestScaleFactor(t *testing.T) {
	fuse1 := NewEpdFuse() //2.083 r 200x96
	fuse2 := NewCustomEpdFuse("./", 100, 48)
	img := getTestImage() //200 x 96

	var factor float64

	factor = fuse1.scaleFactor(img, AXIS_X)
	if factor != 1.0 {
		t.Errorf("Error calculating scale factor. Expected %f, but got %f", float64(1), factor)
	}
	factor = fuse1.scaleFactor(img, AXIS_Y)
	if factor != 1.0 {
		t.Errorf("Error calculating scale factor. Expected %f, but got %f", float64(1), factor)
	}

	factor = fuse2.scaleFactor(img, AXIS_X)
	if factor != 0.5 {
		t.Errorf("Error calculating scale factor. Expected %f, but got %f", float64(0.5), factor)
	}
	factor = fuse2.scaleFactor(img, AXIS_Y)
	if factor != 0.5 {
		t.Errorf("Error calculating scale factor. Expected %f, but got %f", float64(0.5), factor)
	}
}

func TestScaleDown(t *testing.T) {
	//fuse1 := NewEpdFuse() //2.083 r 200x96
	fuse2 := NewCustomEpdFuse("./", 100, 48)
	img := getTestImage() //200 x 96

	var testImg image.Image
	testImg = fuse2.scale(img)
	if testImg == nil || testImg.Bounds().Dx() != 100 || testImg.Bounds().Dy() != 48 {
		t.Errorf("Error scaling image down. Expected 100x48, but got %dx%d", testImg.Bounds().Dx(), testImg.Bounds().Dy())
	}
}

func TestScaleUp(t *testing.T) {
	fuse2 := NewCustomEpdFuse("./", 400, 192)
	img := getTestImage() //200 x 96

	var testImg image.Image
	testImg = fuse2.scale(img)
	if testImg == nil || testImg.Bounds().Dx() != 400 || testImg.Bounds().Dy() != 192 {
		t.Errorf("Error scaling image down. Expected 100x48, but got %dx%d", testImg.Bounds().Dx(), testImg.Bounds().Dy())
	}
}

func getTestImage() image.Image {

	imgByt, err := Asset("test.png")
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(bytes.NewReader(imgByt))
	if err != nil {
		panic(err)
	}
	return img
}
