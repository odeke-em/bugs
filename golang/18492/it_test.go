package main

import (
	"image"
	"testing"
)

var _nooptimize int

func BenchmarkTypeSwitchInside(b *testing.B) {
	var m image.Image = image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	bounds := m.Bounds()
	for i := 0; i < b.N; i++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				switch m.(type) {
				case *image.RGBA:
					_nooptimize = 1
				case *image.YCbCr:
					_nooptimize = 2
				default:
					_nooptimize = 0
				}
			}
		}
	}
}

func BenchmarkTypeSwitchOutside(b *testing.B) {
	var m image.Image = image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	bounds := m.Bounds()
	for i := 0; i < b.N; i++ {
		_type := 0
		switch m.(type) {
		case *image.RGBA:
			_type = 1
		case *image.YCbCr:
			_type = 2
		}
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				switch _type {
				case 1:
					_nooptimize = 1
				case 2:
					_nooptimize = 2
				default:
					_nooptimize = 0
				}
			}
		}
	}
}
