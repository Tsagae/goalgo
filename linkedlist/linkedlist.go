package linkedlist

import "fmt"

type Lista struct {
	testa *Nodo
}

type Nodo struct {
	val  int
	next *Nodo
}

/*
	func main() {
		fmt.Println("prova")
		lista := new(Lista)
		fmt.Println(lista.toString())

		for i := 5; i < 10; i++ {
			lista.add(i)
		}
		fmt.Println(lista.toString())
		lista.insert(2, creaNodo(0, nil))
		fmt.Println(lista.toString())
		nodo := lista.trovaValore(5)
		fmt.Println("valore: ", nodo.val)
		fmt.Println(lista.toString())
		lista.remove(2)
		fmt.Println(lista.toString())

}
*/
func (lista *Lista) toString() string {
	outStr := ""
	currentNodo := lista.testa
	var i int = 0
	for currentNodo != nil {
		outStr += fmt.Sprintf("| %d - %d ", i, currentNodo.val)
		currentNodo = currentNodo.next
		i++
	}
	outStr += "|"
	return outStr
}

func (lista *Lista) add(val int) {
	lastNodo := lista.getLast()
	newNodo := creaNodo(val, nil)
	if lastNodo == nil {
		lista.testa = newNodo
	} else {
		lastNodo.next = newNodo
	}
}

func creaNodo(val int, next *Nodo) *Nodo {
	newNodo := new(Nodo)
	newNodo.val = val
	newNodo.next = next
	return newNodo
}

func (lista *Lista) insert(index int, nodoToInsert *Nodo) {
	if index < 0 {
		return
	}
	if index == 0 {
		nodoToInsert.next = lista.testa
		lista.testa = nodoToInsert
	} else {
		nodo := lista.get(index - 1)
		if nodo == nil {
			return
		}
		nodoToInsert.next = nodo.next
		nodo.next = nodoToInsert
	}
}

// Restituisce nil se la lista è vuota
func (lista *Lista) getLast() *Nodo {
	currentNodo := lista.testa
	if currentNodo == nil {
		return nil
	}
	for currentNodo.next != nil {
		currentNodo = currentNodo.next
	}
	return currentNodo
}

func (lista *Lista) trovaValore(val int) *Nodo {
	currentNodo := lista.testa
	if currentNodo == nil {
		return nil
	}
	for currentNodo != nil {
		if currentNodo.val == val {
			break
		}
		currentNodo = currentNodo.next
	}
	return currentNodo
}

// return nil se non trovato
func (lista *Lista) get(index int) *Nodo {
	if index < 0 {
		return nil
	}
	currentNodo := lista.testa
	if currentNodo == nil {
		return nil
	}
	var i int = 0
	for currentNodo != nil {
		if i == index {
			break
		}
		currentNodo = currentNodo.next
		i++
	}
	return currentNodo
}

func (lista *Lista) remove(index int) {
	if index == 0 {
		if lista.testa != nil {
			lista.testa = lista.testa.next
		}
	} else {
		nodo := lista.get(index - 1)
		nodoToRemove := nodo.next
		if nodoToRemove != nil {
			nodo.next = nodoToRemove.next
		}
	}
}
