# NOTE: some reasoning to decide how to calculate which points to fill

## How external points work

Short breakdown on how i am calculating the rose points:

This is what a point is like, just an X and a Y:
```
    type Point struct {
        X float64
        Y float64
    }
```
We are using an array of those points:
```
    []models.Point
```
Which is formed as follows:
```
    theta := 2 * math.Pi * float64(i) / float64(steps)
    r := a * math.Cos(float64(k)*theta)
    x := r * math.Cos(theta)
    y := r * math.Sin(theta)

    points = append(points, models.Point{X: x, Y: y})
```

So the main problem here is that the points are appended to the points array without any real ordering based on the x or y values. I belive i could fix that and make the whole thing easier but i want it to be more of a challenge with myself.

## Considerations

I want to implement an algorithm that, given as an input a *[]models.Point* returns another *[]models.Point* containing all the internal points for the drawing.

For the whole considerations i will be using this 2 petals (and therefore really 4 petals) rose as an example:

```
                         ::::::::::
                       :::        :::
                      ::            ::
                      :              :
                     ::              ::
                     :                :
                     ::              ::
                      :              :
                      ::            ::
     ::::::::::::::    ::          ::    ::::::::::::::
 :::::            :::::::::      :::::::::            :::::
::                      ::::    ::::                      ::
:                          ::::::                          :
:                          ::::::                          :
::                      ::::    ::::                      ::
 :::::            :::::::::      :::::::::            :::::
     ::::::::::::::    ::          ::    ::::::::::::::
                      ::            ::
                      :              :
                     ::              ::
                     :                :
                     ::              ::
                      :              :
                      ::            ::
                       :::        :::
                         ::::::::::
```