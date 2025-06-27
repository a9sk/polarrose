## NOTE: some reasoning to decide how to calculate which points to fill

```
    theta := 2 * math.Pi * float64(i) / float64(steps)
    r := a * math.Cos(float64(k)*theta)
    x := r * math.Cos(theta)
    y := r * math.Sin(theta)

    points = append(points, models.Point{X: x, Y: y})
```

```
    []models.Point
```

```
    type Point struct {
        X float64
        Y float64
    }
```