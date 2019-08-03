package math

// Max ...
func Max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Min ...
func Min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MaxInt ...
func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt ...
func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// MuliMatrix 矩阵乘法
func MuliMatrix(m1 [][]int, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i := range res {
		t := make([]int, len(m2[0]))
		res[i] = t
		for j := range t {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

// MatrixPower 矩阵次方
func MatrixPower(m [][]int, p int) [][]int {
	res := make([][]int, len(m))
	for i := range res {
		t := make([]int, len(m[0]))
		res[i] = t
		res[i][i] = 1
	}
	tmp := m
	for q := p; q != 0; q >>= 1 {
		if q&1 != 0 {
			res = MuliMatrix(tmp, tmp)
		}
	}
	return res
}

// Make2DArr ...
func Make2DArr(row int, col int) [][]int {
	res := make([][]int, row)
	for i := range res {
		t := make([]int, col)
		res[i] = t
	}
	return res
}
