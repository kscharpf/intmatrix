package intmatrix

import (
       "testing"
)

func TestIntMatrix(t *testing.T) {
  const numRows, numCols = 5, 4
  m := NewIntMatrix(numRows, numCols)
  if m.NumRows() != numRows {
    t.Errorf("m.NumRows() = %v, want %v", m.NumRows(), numRows)
  }  
  if m.NumCols() != numCols {
    t.Errorf("m.NumCols() = %v, want %v", m.NumCols(), numCols)
  }
 
  r1 := make([]int, 4)
  r1[0] = 7
  r1[1] = 9
  r1[2] = 10
  r1[3] = 13

  m.AppendRow(r1)  

  r2 := make([]int, 4)
  r2[0] = 8
  r2[1] = 10
  r2[2] = 11
  r2[3] = 14

  m.AppendRow(r2)  

  r3 := make([]int, 4)
  r3[0] = 9
  r3[1] = 11
  r3[2] = 12
  r3[3] = 15

  m.AppendRow(r3)  

  r4 := make([]int, 4)
  r4[0] = 10
  r4[1] = 11
  r4[2] = 13
  r4[3] = 16
  m.AppendRow(r4)  
 
  checkRow := m.Row(0)
  if len(checkRow) != m.NumCols() {
    t.Errorf("len(checkRow) = %v want %v\n", len(checkRow), m.NumCols())
  } 
  if checkRow[0] != r1[0] {
    t.Errorf("checkRow[0] = %v want %v\n", checkRow[0], r1[0])
  }
  if checkRow[1] != r1[1] {
    t.Errorf("checkRow[1] = %v want %v\n", checkRow[1], r1[1])
  }
  if checkRow[2] != r1[2] {
    t.Errorf("checkRow[2] = %v want %v\n", checkRow[2], r1[2])
  }
  if checkRow[3] != r1[3] {
    t.Errorf("checkRow[3] = %v want %v\n", checkRow[3], r1[3])
  }
}
