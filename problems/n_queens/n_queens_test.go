package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomSolution(t *testing.T) {
	p := problem(t)
	s := p.RandomSolution()
	assert.Len(t, s.(board).Columns(), 10)
}


func TestColumnClashes(t *testing.T) {
	b := newBoard([]int{0, 1, 1, 2, 2, 3, 4, 7})
	assert.Equal(t, 2, b.rowClashes())
}

func TestDiagonalClashes(t *testing.T) {
	b := newBoard([]int{0, 1, 4, 3, 7, 0, 5, 0})
	assert.Equal(t, 4, b.diagonalClashes())
	b = newBoard([]int{0, 1, 0, 1, 1, 0, 0, 0})
	assert.Equal(t, 4, b.diagonalClashes())
	b = newBoard([]int{0, 1, 0, 1, 0, 1, 0, 1})
	assert.Equal(t, 7, b.diagonalClashes())
	b = newBoard([]int{0, 1, 2, 3, 4, 5, 6, 7})
	assert.Equal(t, 7, b.diagonalClashes())
}

func problem(t *testing.T) NQueensProblem {
	t.Helper()
	return NQueensProblem{checkboard_size: 10, rnd: rnd}
}


/*

zmienić układ współrzędnych?

7 ....Q...
6 ........
5 ......Q.
4 ..Q.....
3 ...Q....
2 ........
1 .Q......
0 Q....Q.Q
 */

