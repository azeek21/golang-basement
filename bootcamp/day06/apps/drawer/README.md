# resources
* [Drawer](#Drawer)
* [Syntax](#snytax)
* [Usage](#usage)
* [Dev notes](#notes)

## Drawer
A package inspired by [xpm](https://en.wikipedia.org/wiki/X_PixMap).

Used to generate images from files or any readable source with needed content and draw them into images.

Done as part of School 21 golang bootcamp.

### syntax
For now it only has bare minimums. I may add other features and release a new version in the future.

Filename doesn't matter as long as it has needed text content but I named it `dws` stands for `drawer source`
```image-example.dws
# this is a comment. Comments can only be at the beginning of the line (no inline commments)
# reserving version keyword even we don't have other versions yet. Just for the future
# version must be at the top of the file
version=1

# colors must come before image source and must end with the `end` keyword
colors
.=0,0,0,0
o=0,255,50
end

# image source is indicated between `image` and `end` keyswords.
# note the keywords are in different lines than source itself
image
....o....
...o.o....
..ooooo...
.o.....o..
o.......o.
end
```

### usage
```go
package main

import (
	"fmt"

	"github.com/azeek21/blog/apps/drawer"
)

func main() {
	drw := drawer.NewDrawer()
	err := drw.DrawPngFromFile("./logo.dws", "out.png")
	if err != nil {
		fmt.Println(err.Error())
	}
}
```

### notes
Looks like I'll be spending too much time with this thing. There's tons of ideas flowing into my mind. But I want to get this project out, it was needed as a part of an assignment anyways.
