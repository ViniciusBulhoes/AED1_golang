package main

import (
	"fmt"
	"time"
)

type List interface {
	Init(value int)
	Add(value int)
	AddPos(val int, pos int)
	Update(val int, pos int)
	RemoveLast()
	Remove(pos int)
	Get(pos int) int
	Size() int
}

type ArrayList struct {
	values   []int
	inserted int
}

func (list *ArrayList) Init(value int) {
	newArray := make([]int, value)
	list.values = newArray
}

/**
 * Duplica o vetor.
 */
func (list *ArrayList) doubleArray() {
	curSize := len(list.values)
	doubledValues := make([]int, 2*curSize)
	for i := 0; i < curSize; i++ {
		doubledValues[i] = list.values[i]
	}
	list.values = doubledValues
}

/**
 * Adiciona sempre da esquerda para a direita,
 * após o último elemento inserido anteriormente.
 *
 * Melhor caso: Ômega(1)
 *   Just: não precisa duplicar vetor, nem varrer o vetor
 *         do início para encontrar o endereço a ser Add
 * Pior caso: O(n)
 *   Just: duplicar o vetor requer copiar os elementos
 *         para o novo vetor, causando um custo computa-
 *         cional proporcional ao tamanho do vetor (n)
 */
func (list *ArrayList) Add(val int) {
	//verificar se array está lotado
	if list.inserted >= len(list.values) {
		list.doubleArray()
	}
	list.values[list.inserted] = val
	list.inserted++
}

/**
 * Adiciona elemento na posição especificada, deslocando
 * os elementos à direita dessa posição.
 *
 * Melhor caso: Ômega(1)
 *   Just: não precisa duplicar vetor, nem varrer o vetor
 *         do início para encontrar o endereço a ser Add
 * Pior caso: O(n)
 *   Just: 1) duplicar o vetor requer copiar os elementos
 *         para o novo vetor, causando um custo computa-
 *         cional proporcional ao tamanho do vetor (n)
 *         2) adicionar no início requer deslocar os n
 *            elementos do vetor para a direita
 */
func (list *ArrayList) AddPos(val int, pos int) {
	//pos eh um valor válido?
	if pos >= 0 && pos <= list.inserted {
		//verificar se array está lotado
		if list.inserted >= len(list.values) {
			list.doubleArray()
		}

		// começar do início para o fim, para aproveitar
		// a cache pode de fato melhorar o desempenho?
		for i := list.inserted; i > pos; i-- {
			list.values[i] = list.values[i-1]
		}
		list.values[pos] = val
		list.inserted++
	}
}

/**
 * Altera o valor de um elemento na posicao especificada
 *
 * Melhor caso: Ômega(1)
 *	  Just: não precisa duplicar vetor ou deslocar
 * Pior caso: O(1)
 *	  Just: é a mesma operação do melhor caso
 */
func (list *ArrayList) Update(val int, pos int) {
	if pos >= 0 && pos <= list.inserted {
		list.values[pos] = val
	}
}

/**
 * Remove o ultimo elemento
 * Melhor e pior caso: Ômega(1) e O(1)
 * 	 Just: não precisa duplicar o vetor ou deslocar
 */
func (list *ArrayList) RemoveLast() {
	list.inserted--
}

/**
 * Remove um elemento em uma determinada posicao deslocando todos
 * os elementos a direita para a esquerda
 * Melhor caso: Ômega(1)
 *	  Just: remover o ultimo elemento, não precisando deslocar
 *	        nenhum elemento
 * Pior caso: O(n)
 *	  Just: remover o primeiro elemento, precisando deslocar todos
 *	 		elementos
 */
func (list *ArrayList) Remove(pos int) {
	if pos >= 0 && pos <= list.inserted {
		for i := pos; i < list.inserted; i++ {
			list.values[i] = list.values[i+1]
		}
		list.inserted--
	}
}

func (list *ArrayList) Get(pos int) int {
	return list.values[pos]
}

func (list *ArrayList) Size() int {
	return list.inserted
}

func main() {
	var o_1_time [15]float64
	fmt.Println("Tempo para O(n):")
	for i := 1; i < 15; i++ {
		tam := i * 10000
		a := new(ArrayList)
		a.Init(tam)
		start := time.Now()
		for j := 0; j < tam; j++ {
			a.AddPos(j, 0)
		}
		stop_count := time.Since(start)
		var totaltime float64
		totaltime = float64(stop_count)
		o_1_time[i] = totaltime
		fmt.Println(stop_count)
	}
	fmt.Println("Tempo para O(1):")
	for i := 1; i < 15; i++ {
		tam := i * 10000
		a := new(ArrayList)
		a.Init(tam)
		start := time.Now()
		for j := 0; j < tam; j++ {
			a.AddPos(j, 0)
		}
		stop_count := time.Since(start)
		var totaltime float64
		totaltime = float64(stop_count)
		o_1_time[i] = totaltime
		fmt.Println(stop_count)
	}

	//fmt.Println(o_1_time)
	/**
	a := new(ArrayList)
	a.Init(1000)
	fmt.Println(a.values[0:a.inserted])
	a.Add(5)
	a.Add(6)
	a.Add(7)
	a.Add(8)
	a.Add(9)
	fmt.Println(a.values[0:a.inserted])
	a.AddPos(10, 3)
	fmt.Println(a.values[0:a.inserted])
	a.Update(11, 5)
	fmt.Println(a.values[0:a.inserted])
	a.RemoveLast()
	fmt.Println(a.values[0:a.inserted])
	a.Remove(3)
	fmt.Println(a.values[0:a.inserted])
	fmt.Println(a.Get(2))
	fmt.Println(a.Size())
	*/
}
