package main

import (
	libraryFunction "advance-golang/library" // --> import "advance-golang/library" with name libraryFunction
	"fmt"
	"math"
	"sync"
)

//#region  //*=========== Method Region ===========
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// or

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
//#endregion  //*======== Method Region ===========

func main() {
	//#region  //*=========== Pointer Region ===========
	a := 200
	b := &a
	fmt.Println(a) 
	fmt.Println(b)
	*b++
	fmt.Println(a) 

	v := Vertex{3, 4}
	fmt.Println(v.Abs()) 

	fmt.Println(Abs(v)) 

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) 
	//#endregion  //*======== Pointer Region ===========


	//#region  //*=========== Package Access Control in Method Region ===========
	println(libraryFunction.ExtractionData(("John Wick")))
	println(libraryFunction.Substraction(10, 5))
	//#endregion  //*======== Package Access Control in Method Region ===========


	//#region  //*=========== Interface Region ===========
	l := lingkaran{14}
	s := segitiga{8, 10}
	p := persegi{10}
	fmt.Println("Luas dan Keliling Lingkaran", l.luas(), l.keliling())
	hitungLuasKeliling(l)
	fmt.Println("Luas dan Keliling Segitiga", s.luas(), s.keliling())
	hitungLuasKeliling(s)
	fmt.Println("Luas dan Keliling Persegi", p.luas(), p.keliling())
	hitungLuasKeliling(p)
	//#endregion  //*======== Interface Region ===========


	//#region  //*=========== Goroutine Region ===========
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Hello World")
		}
	}()
	var input string
	fmt.Scanln(&input)
	//#endregion  //*======== Goroutine Region ===========


	//#region  //*=========== Gotoutine with WaitGroup Region ===========
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hi")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	wg.Wait()
	//#endregion  //*======== Gotoutine with WaitGroup Region ===========
}


type lingkaran struct {
	diameter float64
}
type segitiga struct {
	alas   float64
	tinggi float64
}
type persegi struct {
	sisi float64
}
type hitung interface {
	luas() float64
	keliling() float64
}

func (l lingkaran) luas() float64 {
	return 3.14 * l.diameter * l.diameter
}
func (l lingkaran) keliling() float64 {
	return 2 * 3.14 * l.diameter
}
func (s segitiga) luas() float64 {
	return 0.5 * s.alas * s.tinggi
}
func (s segitiga) keliling() float64 {
	return 3 * s.alas
}
func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}
func (p persegi) keliling() float64 {
	return 4 * p.sisi
}
func hitungLuasKeliling(h hitung) {
	fmt.Println("Luas :", h.luas())
	fmt.Println("Keliling :", h.keliling())
}
