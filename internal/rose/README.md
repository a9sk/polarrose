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

This will be a graphical view of what the code should do and how to decide how to do it. The roses are always arrays of points.

Given the previous rose as an input, the output to this algorithm should be an array of points as follows:

```
                                   
                          ▓▓▓▓▓▓▓▓   
                        ▓▓▓▓▓▓▓▓▓▓▓▓  
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓  
                      ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓  
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
                        ▓▓▓▓▓▓▓▓▓▓▓▓  
                         ▓▓▓▓▓▓▓▓▓▓                    
      ▓▓▓▓▓▓▓▓▓▓▓▓         ▓▓▓▓▓▓         ▓▓▓▓▓▓▓▓▓▓▓▓     
  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓    ▓▓▓▓    ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  
 ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓      ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
 ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓      ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
  ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓    ▓▓▓▓    ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓  
      ▓▓▓▓▓▓▓▓▓▓▓▓         ▓▓▓▓▓▓         ▓▓▓▓▓▓▓▓▓▓▓▓     
                         ▓▓▓▓▓▓▓▓▓▓                    
                        ▓▓▓▓▓▓▓▓▓▓▓▓  
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓  
                      ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓  
                       ▓▓▓▓▓▓▓▓▓▓▓▓▓▓ 
                        ▓▓▓▓▓▓▓▓▓▓▓▓  
                          ▓▓▓▓▓▓▓▓   
                                   
```

By drawing both the external and internal point the final picture should be:

```
                         ::::::::::
                       :::▓▓▓▓▓▓▓▓:::
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
                     :▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::
     ::::::::::::::    ::▓▓▓▓▓▓▓▓▓▓::    ::::::::::::::
 :::::▓▓▓▓▓▓▓▓▓▓▓▓:::::::::▓▓▓▓▓▓:::::::::▓▓▓▓▓▓▓▓▓▓▓▓:::::
::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::▓▓▓▓::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
:▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
:▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::▓▓▓▓::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
 :::::▓▓▓▓▓▓▓▓▓▓▓▓:::::::::▓▓▓▓▓▓:::::::::▓▓▓▓▓▓▓▓▓▓▓▓:::::
     ::::::::::::::    ::▓▓▓▓▓▓▓▓▓▓::    ::::::::::::::
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
                     :▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::
                       :::▓▓▓▓▓▓▓▓:::
                         ::::::::::
```

The main thing we notice straight away is that the parts that are between the external points and one of the two lateral ends of the rose shall never be filled.


```
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::::::::
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:::        :::
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::            ::
etcetera              :              :
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

So the first idea was to just look at a point and set as "toFill" the coordinates after it untill we find the next external point. If two external points are close we just iterate one more until we find a space, the next external point will be where we stop the filling process.


```
                         :::::::::: (this line has nothing to fill)
we find this point ->  :::▓▓▓▓▓▓▓▓::: <- we finish here
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

The output to this approach is not right tho.

```
                         ::::::::::
                       :::▓▓▓▓▓▓▓▓:::
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
                     :▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓: now this is not the most obvious of errors
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓:: but with some more petals it is worse
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:  |
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::  v
     ::::::::::::::▓▓▓▓::▓▓▓▓▓▓▓▓▓▓::▓▓▓▓::::::::::::::
 :::::▓▓▓▓▓▓▓▓▓▓▓▓:::::::::▓▓▓▓▓▓:::::::::▓▓▓▓▓▓▓▓▓▓▓▓:::::
::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::▓▓▓▓::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
:▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
:▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::::▓▓▓▓::::▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
 :::::▓▓▓▓▓▓▓▓▓▓▓▓:::::::::▓▓▓▓▓▓:::::::::▓▓▓▓▓▓▓▓▓▓▓▓:::::
     ::::::::::::::▓▓▓▓::▓▓▓▓▓▓▓▓▓▓::▓▓▓▓::::::::::::::
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
                     :▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                     ::▓▓▓▓▓▓▓▓▓▓▓▓▓▓::
                      :▓▓▓▓▓▓▓▓▓▓▓▓▓▓:
                      ::▓▓▓▓▓▓▓▓▓▓▓▓::
                       :::▓▓▓▓▓▓▓▓:::
                         ::::::::::
```

We need to therefore find a way to avoid filling those points. At first i thought of a way to map the external points using a diffusion selector algoritm, and then when drawing you could just check if the point which you want to be drawing to is not an external one. This is not even close to the best solution.

The idea proposed is pretty simple to implement. Given an "empty rose" (only the external points) we could map it to a matrix and then use the matrix to find all of the "internal points of the rose".

Given that the rose is formed by polar coordinates and results in a smooth, enclosed shape, it’s safe to assume the following:

All external points are on the "outside".
The corners of the screen are guaranteed to be outside the shape.
Any space that is not reachable from the outside is therefore "inside".

So instead of scanning and flipping flags (like traditional scanline fill algorithms), we can make a much simpler and more robust assumption:

## The flood fill idea

If we treat the entire grid as a 2D matrix of points, we can:

Mark all the external points of the rose on a matrix.
Flood fill starting from the corners of the screen.
Any point that was *not* reached by the flood and is *not* on the boundary must be inside.

We use a simple BFS to propagate from the corners. This is guaranteed to fill only the outside and never go inside, as long as the rose boundary is properly closed and connected.