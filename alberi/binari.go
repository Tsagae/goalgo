package alberi

import "fmt"

type Nodo struct {
	val int
	sx  *Nodo
	dx  *Nodo
}

type Albero struct {
	radice *Nodo
}

type Coda struct {
	innerCoda []*Nodo
}

type Stack struct {
	innerStack []*Nodo
}

/*
func main() {
	fmt.Println("Prova")
	albero := new(Albero)
	albero.radice = creaNodo(1, nil, nil)
	albero.radice.sx = creaNodo(2, nil, nil)
	albero.radice.dx = creaNodo(3, nil, nil)
	albero.radice.sx.sx = creaNodo(4, nil, nil)
	albero.radice.sx.dx = creaNodo(5, nil, nil)
	albero.radice.dx.sx = creaNodo(6, nil, nil)
	albero.radice.dx.dx = creaNodo(7, nil, nil)

	fmt.Println("Visita in ampiezza:")
	visitaInAmpiezza(*albero)
	fmt.Println("\n--------------------")
	fmt.Println("Visita in profondità")
	visitaInProfondita(*albero)
	fmt.Println("\n--------------------")
	fmt.Println("Visita ricorsiva:")
	visitaRicorsiva(albero.radice)
	fmt.Println("\n--------------------")

}
*/

func creaNodo(val int, sx *Nodo, dx *Nodo) *Nodo {
	newNodo := new(Nodo)
	newNodo.val = val
	newNodo.sx = sx
	newNodo.dx = dx
	return newNodo
}

func visitaRicorsiva(nodo *Nodo) {
	if nodo != nil {
		fmt.Print(nodo.val, " ")
		visitaRicorsiva(nodo.sx)
		visitaRicorsiva(nodo.dx)
	}
}

func visitaInAmpiezza(albero Albero) {
	coda := creaCoda()
	coda.enqueue(albero.radice)
	for coda.size() != 0 {
		x := coda.dequeue()
		if x != nil {
			fmt.Print(x.val, " ")
			coda.enqueue(x.sx)
			coda.enqueue(x.dx)
		}
	}
}

func visitaInProfondita(albero Albero) {
	p := creaStack()
	p.push(albero.radice)
	for p.size() != 0 {
		x := p.pop()
		if x != nil {
			fmt.Print(x.val, " ")
			p.push(x.dx)
			p.push(x.sx)
		}
	}
}

func creaCoda() *Coda {
	coda := new(Coda)
	coda.innerCoda = make([]*Nodo, 0)
	return coda
}

func (coda *Coda) size() int {
	return len(coda.innerCoda)
}

func (coda *Coda) enqueue(nodo *Nodo) {
	coda.innerCoda = append(coda.innerCoda, nodo)
}

func (coda *Coda) dequeue() *Nodo {
	nodoToRet := coda.first()
	coda.innerCoda = coda.innerCoda[1:]
	return nodoToRet
}

func (coda *Coda) first() *Nodo {
	return coda.innerCoda[0]
}

func creaStack() *Stack {
	stack := new(Stack)
	stack.innerStack = make([]*Nodo, 0)
	return stack
}

func (stack *Stack) size() int {
	return len(stack.innerStack)
}

func (stack *Stack) push(nodo *Nodo) {
	stack.innerStack = append(stack.innerStack, nodo)
}

func (stack *Stack) pop() *Nodo {
	nodoToRet := stack.peek()
	stack.innerStack = stack.innerStack[:stack.size()-1]
	return nodoToRet
}

func (stack *Stack) peek() *Nodo {
	return stack.innerStack[stack.size()-1]
}
