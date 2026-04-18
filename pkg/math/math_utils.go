package math

import "math"

// Distance2 计算二维平面两点距离
func Distance2(x1, y1, x2, y2 int32) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	return math.Sqrt(dx*dx + dy*dy)
}

// DistanceSquared2 计算二维平面两点距离的平方（避免开方运算）
func DistanceSquared2(x1, y1, x2, y2 int32) int64 {
	dx := int64(x2 - x1)
	dy := int64(y2 - y1)
	return dx*dx + dy*dy
}

// Distance3 计算三维空间两点距离
func Distance3(x1, y1, z1, x2, y2, z2 int32) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	dz := float64(z2 - z1)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// DistanceSquared3 计算三维空间两点距离的平方（避免开方运算）
func DistanceSquared3(x1, y1, z1, x2, y2, z2 int32) int64 {
	dx := int64(x2 - x1)
	dy := int64(y2 - y1)
	dz := int64(z2 - z1)
	return dx*dx + dy*dy + dz*dz
}

// PointInCircle 判断点是否在圆形区域内
func PointInCircle(px, py int32, cx, cy, radius int32) bool {
	return DistanceSquared2(px, py, cx, cy) <= int64(radius)*int64(radius)
}

// PointInRect 判断点是否在矩形区域内
// (x, y) 测试点, (rx, ry) 矩形左上角, (rw, rh) 矩形宽高
func PointInRect(x, y, rx, ry, rw, rh int32) bool {
	return x >= rx && x <= rx+rw && y >= ry && y <= ry+rh
}

// PointInPolygon 判断点是否在多边形内（射线法）
// 使用射线投射算法：从点向右发射射线，计算与多边形边界的交点数
// 如果交点数为奇数，点在多边形内；否则在外部
func PointInPolygon(px, py int32, points []*Vector2) bool {
	if len(points) < 3 {
		return false
	}

	n := len(points)
	inside := false

	for i := 0; i < n; i++ {
		j := (i + 1) % n

		xi, yi := points[i].X, points[i].Y
		xj, yj := points[j].X, points[j].Y

		// 检查点是否在边的Y范围内
		if ((yi > py) != (yj > py)) {
			// 计算射线与边的交点X坐标
			intersectX := int64(xj-xi)*int64(py-yi)/int64(yj-yi) + int64(xi)
			if int64(px) <= intersectX {
				inside = !inside
			}
		}
	}

	return inside
}

// AbsInt32 返回int32的绝对值
func AbsInt32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// AbsInt64 返回int64的绝对值
func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// MinInt32 返回两个int32中的较小值
func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// MaxInt32 返回两个int32中的较大值
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// MinInt64 返回两个int64中的较小值
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MaxInt64 返回两个int64中的较大值
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// ClampInt32 将值限制在[min, max]范围内
func ClampInt32(value, min, max int32) int32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampFloat64 将值限制在[min, max]范围内
func ClampFloat64(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}