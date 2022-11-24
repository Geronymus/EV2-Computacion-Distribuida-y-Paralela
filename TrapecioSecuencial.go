package main

import (
	"fmt"
	"time"
)

type Trapecio struct {
	base float64
	BASE float64
	aX   float64
	bX   float64
	h    float64
	AREA float64
}

func (t Trapecio) Calcular() float64 {
	t.h = t.bX - t.aX
	t.AREA = ((t.base + t.BASE) * t.h) / 2
	return t.AREA
}

func main() {
	a := 5.0
	b := 20.0
	h := b - a
	var a2 float64
	nTrapecios := 100000
	var teo []Trapecio
	aTOTAL := 0.0
	aPARCIAL := 0.0
	delta := (h / float64(nTrapecios))
	startMain := time.Now()
	for i := 1; i <= nTrapecios; i++ {
		fmt.Println("-------------------------------")
		fmt.Println("a: ", a)
		a2 = a
		a = a + (delta)
		b = a
		fmt.Println("b: ", b)
		//
		start := time.Now()
		//
		teo = append(teo, Trapecio{base: funcion(a2), BASE: funcion(b), aX: a2, bX: b})
		aPARCIAL = teo[i-1].Calcular()
		//
		tiempo := time.Since(start).Nanoseconds()
		//
		fmt.Println("Area Trapecio (", i, ") :", aPARCIAL)
		fmt.Println("Tiempo: ", tiempo)
		aTOTAL = aTOTAL + aPARCIAL
	}
	fmt.Println("-------------------------------")
	fmt.Println("AREA TOTAL (", nTrapecios, ") trapecios : ", aTOTAL)
	endMain := time.Since(startMain).Nanoseconds()
	fmt.Println("Tiempo total: ", endMain)
}

func funcion(x float64) float64 {
	return (x*x + 1) / 2
}
