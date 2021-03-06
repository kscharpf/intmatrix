package intmatrix

import (
  "fmt"
)

type IntMatrix struct {
  Fields [][]int
  RowCount int
  ColCount int
}

func NewIntMatrix() IntMatrix {
  m := IntMatrix{}
  m.RowCount = 0
  m.ColCount = 0

  return m
}

func (m IntMatrix) NumRows() int {
  return m.RowCount
}

func (m IntMatrix) NumCols() int {
  return m.ColCount
}

func (m IntMatrix) Column(idx int) []int {
  if idx >= m.NumCols() {
    panic(fmt.Sprintf("Cannot extract column index %d from matrix with %d columns\n", idx, m.NumCols()))
  }

  out := make([]int, m.NumRows())
    
  for i:=0; i<m.NumRows(); i++ {
    out[i] = m.Fields[i][idx]
  }
  return out
}

func (m IntMatrix) Row(idx int) []int {
  if idx >= m.NumRows() {
    panic(fmt.Sprintf("Cannot extract row index %d from matrix with %d columns\n", idx, m.NumRows()))
  }

  return m.Fields[idx]
}

func (m *IntMatrix) AppendColumn(col []int) {
  for i:= range m.Fields {
    m.Fields[i] = append(m.Fields[i], col[i])
  }
  m.ColCount = m.ColCount + 1
}

func (m *IntMatrix) AppendRow(row []int) {
  m.Fields = append(m.Fields, row)
  m.RowCount = m.RowCount + 1
  if m.ColCount != 0  &&  m.ColCount != len(row) {
    panic(fmt.Sprintf("Incorrect number of fields in AppendRow: %d Expected %d\n", len(row), m.ColCount))
  }
    
  m.ColCount = len(row)
}



func (m IntMatrix) Transpose() IntMatrix {
  out := NewIntMatrix()
  
  for i:=0; i<m.NumCols(); i++ {
    c := m.Column(i)
    out.AppendRow(c)
  }
  return out 
}
