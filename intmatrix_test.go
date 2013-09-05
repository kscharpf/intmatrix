package intmatrix

import (
       "testing"
)

func check_array(actual []int, expected []int, t *testing.T) {
  if len(actual) != len(expected) {
    t.Errorf("Actuals - num entries: %v Expected %v", len(actual), len(expected))
  }

  for i := range actual {
    if actual[i] != expected[i] {
      t.Errorf("Actual[%v] = %v expected %v", i, actual[i], expected[i])
    }
  }
}

  

func TestIntMatrix(t *testing.T) {
  m := NewIntMatrix()
  if m.NumRows() != 0 {
    t.Errorf("m.NumRows() = %v, want 0", m.NumRows())
  }  
  if m.NumCols() != 0 {
    t.Errorf("m.NumCols() = %v, want 0", m.NumCols())
  }
 
  r1 := make([]int, 4)
  r1[0] = 7
  r1[1] = 9
  r1[2] = 10
  r1[3] = 13

  m.AppendRow(r1)  


  if m.NumRows() != 1 {
    t.Errorf("m.NumRows() = %v, want 1", m.NumRows())
  }

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
 
  check_array(r1, m.Row(0), t)
  check_array(r2, m.Row(1), t)
  check_array(r3, m.Row(2), t)
  check_array(r4, m.Row(3), t)

  m2 := m.Transpose()
  check_array(m2.Column(0), m.Row(0), t)
  check_array(m2.Column(1), m.Row(1), t)
  check_array(m2.Column(2), m.Row(2), t)
  check_array(m2.Column(3), m.Row(3), t)
}
