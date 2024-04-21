package main

import "testing"

func TestRectangleArea(t *testing.T) {
	if RectangleArea(10, 5) != "even rectangle" {
		t.Errorf("RectangleArea(10, 5) = %s; want even rectangle", RectangleArea(10, 5))
	}

	if RectangleArea(3, 7) != "odd rectangle" {
		t.Errorf("RectangleArea(3, 7) = %s; want odd rectangle", RectangleArea(3, 7))
	}
}

func TestRectanglePerimeter(t *testing.T) {
	if !RectanglePerimeter(10, 5) {
		t.Errorf("RectanglePerimeter(10, 5) = %t; want true", RectanglePerimeter(10, 5))
	}

	if RectanglePerimeter(7, 4) {
		t.Errorf("RectanglePerimeter(7, 4) = %t; want false", RectanglePerimeter(7, 4))
	}
}

func TestSquareArea(t *testing.T) {
	if SquareArea(10) != "even square" {
		t.Errorf("SquareArea(10) = %s; want even square", SquareArea(10))
	}

	if SquareArea(5) != "odd square" {
		t.Errorf("SquareArea(5) = %s; want odd square", SquareArea(5))
	}
}

func TestSquarePerimeter(t *testing.T) {
	if !SquarePerimeter(10) {
		t.Errorf("SquarePerimeter(10) = %t; want true", SquarePerimeter(10))
	}

	if SquarePerimeter(5) {
		t.Errorf("SquarePerimeter(5) = %t; want false", SquarePerimeter(5))
	}
}