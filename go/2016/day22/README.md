--- Day 22: Grid Computing ---

You gain access to a massive storage cluster arranged in a grid;
each storage node is only connected to the four nodes directly
adjacent to it (three if the node is on an edge, two if it's in
a corner).

You can directly access data only on node /dev/grid/node-x0-y0,
but you can perform some limited actions on the other nodes:

You can get the disk usage of all nodes (via df). The result of
doing this is in your puzzle input.

You can instruct a node to move (not copy) all of its data to an
adjacent node (if the destination node has enough space to receive
the data). The sending node is left empty after this operation.

Nodes are named by their position: the node named node-x10-y10 is
adjacent to nodes node-x9-y10, node-x11-y10, node-x10-y9, and node-x10-y11.

Before you begin, you need to understand the arrangement of data
on these nodes. Even though you can only move data between directly
connected nodes, you're going to need to rearrange a lot of the data
to get access to the data you need. Therefore, you need to work out
how you might be able to shift data around.

To do this, you'd like to count the number of viable pairs of nodes.
A viable pair is any two nodes (A,B), regardless of whether they
are directly connected, such that:

- Node A is not empty (its Used is not zero).
- Nodes A and B are not the same node.
- The data on node A (its Used) would fit on node B (its Avail).

How many viable pairs of nodes are there?

Your puzzle answer was 976.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---

Now that you have a better understanding of the grid, it's time to get to work.

Your goal is to gain access to the data which begins in the node with
y=0 and the highest x (that is, the node in the top-right corner).

For example, suppose you have the following grid:

```
Filesystem            Size  Used  Avail  Use%
/dev/grid/node-x0-y0   10T    8T     2T   80%
/dev/grid/node-x0-y1   11T    6T     5T   54%
/dev/grid/node-x0-y2   32T   28T     4T   87%
/dev/grid/node-x1-y0    9T    7T     2T   77%
/dev/grid/node-x1-y1    8T    0T     8T    0%
/dev/grid/node-x1-y2   11T    7T     4T   63%
/dev/grid/node-x2-y0   10T    6T     4T   60%
/dev/grid/node-x2-y1    9T    8T     1T   88%
/dev/grid/node-x2-y2    9T    6T     3T   66%
```

In this example, you have a storage grid 3 nodes wide and 3 nodes tall.
The node you can access directly, node-x0-y0, is almost full.
The node containing the data you want to access, node-x2-y0
(because it has y=0 and the highest x value), contains 6 terabytes of data -
enough to fit on your node, if only you could make enough space to move it there.

Fortunately, node-x1-y1 looks like it has enough free space to enable
you to move some of this data around. In fact, it seems like all of
 the nodes have enough space to hold any node's data (except node-x0-y2,
 which is much larger, very full, and not moving any time soon). So,
 initially, the grid's capacities and connections look like this:

```
( 8T/10T) --  7T/ 9T -- [ 6T/10T]
    |           |           |
  6T/11T  --  0T/ 8T --   8T/ 9T
    |           |           |
 28T/32T  --  7T/11T --   6T/ 9T
```

The node you can access directly is in parentheses;
the data you want starts in the node marked by square brackets.

In this example, most of the nodes are interchangable:
they're full enough that no other node's data would fit,
but small enough that their data could be moved around.
Let's draw these nodes as `.`. The exceptions are the empty node,
which we'll draw as `_`, and the very large, very full node,
which we'll draw as `#`. Let's also draw the goal data as `G`.
Then, it looks like this:

```
(.) .  G
 .  _  .
 #  .  .
```

The goal is to move the data in the top right, G, to the node in parentheses.
To do this, we can issue some commands to the grid and rearrange the data:

Move data from node-y0-x1 to node-y1-x1, leaving node node-y0-x1 empty:

```
(.) _  G
 .  .  .
 #  .  .
```

Move the goal data from node-y0-x2 to node-y0-x1:

```
(.) G  _
 .  .  .
 #  .  .
```

At this point, we're quite close. However, we have no deletion command,
so we have to move some more data around. So, next, we move the data from node-y1-x2 to node-y0-x2:

```
(.) G  .
 .  .  _
 #  .  .
```

Move the data from node-y1-x1 to node-y1-x2:

```
(.) G  .
 .  _  .
 #  .  .
```

Move the data from node-y1-x0 to node-y1-x1:

```
(.) G  .
 _  .  .
 #  .  .
```

Next, we can free up space on our node by moving the data from node-y0-x0 to node-y1-x0:

```
(_) G  .
 .  .  .
 #  .  .
```

Finally, we can access the goal data by moving the it from node-y0-x1 to node-y0-x0:

```
(G) _  .
 .  .  .
 #  .  .
```

So, after 7 steps, we've accessed the data we want. Unfortunately, each of these moves takes time,
and we need to be efficient:

What is the fewest number of steps required to move your goal data to node-x0-y0?

Although it hasn't changed, you can still get your puzzle input.

Answer:
 [Submit]

I tried my hand at pretty-printing the disk, but got my x-axis and y-axis flipped:

73/86  65/88  68/88  71/85  73/90  65/86  73/89  72/92  68/94  71/94  71/94  65/88  71/91  65/87  68/88  65/90  68/89  72/89  65/85  73/90  68/89  70/90  73/88  71/94  67/94  64/87  72/86
69/86  65/94  66/92  73/87  70/86  73/85  70/87  71/92  73/86  70/90  70/90  67/88  71/89  69/87  65/90  71/93  65/94  72/89  65/87  65/89  66/87  68/88  67/91  68/89  68/85  65/85  68/85
64/86  73/92  72/94  72/91  65/88  70/90  64/86  68/93  73/93  69/88  66/93  69/93  70/89  73/94  71/91  66/89  69/87  65/86  66/94  64/85  67/92  71/88  71/85  69/87  65/91  66/92  64/91
64/85  66/90  68/92  67/93  72/94  73/87  66/89  65/91  66/93  66/92  70/87  70/93  68/86  64/91  72/93  69/90  66/87  68/91  72/92  71/85  71/87  71/92  69/89  66/90  70/89  70/93  72/86
65/87  68/85  73/85  71/91  64/89  67/85  65/91  69/91  69/91  68/90  68/88  73/86  69/89  72/88  65/85  69/85  65/87  65/89  66/92  68/88  68/87  68/85  71/87  66/94  69/94  69/89  70/89
64/85  68/90  64/94  66/93  67/93  66/94  66/88  72/90  68/90  65/92  73/90  67/86  72/86  71/88  65/92  67/92  68/91  66/89  73/92  65/85  66/90  68/85  65/88  66/94  67/88  69/88  72/89
69/87  73/90  64/86  73/85  73/85  69/89  67/91  68/92  65/89  65/88  72/87  71/89  68/85  72/93  70/90  69/90  64/89  67/89  69/94  70/90  68/87  66/91  64/86  73/87  67/85  70/85  73/88
67/88  68/88  64/90  65/89  66/94  65/92  64/91  70/86  67/92  68/88  69/88  68/86  65/92  64/89  69/93  73/94  70/93  71/88  67/89  64/90  69/87  72/90  70/90  64/90  65/90  67/92  64/92
70/92  70/88  72/89  67/94  70/86  65/87  69/91  64/88  67/94  68/88  65/90  71/88  66/90  67/86  73/89  68/88  69/88  73/89  68/85  68/92  67/93  65/89  71/94  65/90  64/88  73/87  73/91
67/91  72/94  65/89  71/90  66/85  68/94  73/92  73/89  68/87  69/90  68/88  72/86  68/88  68/89  67/92  71/91  68/86  68/87  73/94  68/94  72/92  66/86  66/87  65/89  64/86  67/86  72/90
73/94  68/88  64/93  71/90  64/90  65/88  68/86  68/92  72/93  70/88  72/90  66/90  73/90  72/93  68/93  71/87  65/91  64/86  71/92  64/93  69/90  73/94  70/92  73/94  64/94  64/85  69/93
71/89  73/88  65/87  70/93  65/93  68/90  72/89  70/89  73/92  66/90  71/88  67/89  64/88  66/87  73/86  65/93  66/94  69/93  70/91  68/92  68/93  64/85  73/94  72/88  67/90  66/87  70/92
66/94  70/90  68/86  73/88  73/89  71/93  68/88  73/87  71/93  69/88  65/93  71/88  69/94  65/88  70/93  64/91  70/90  68/91  67/92  65/85  68/86  73/90  67/91  72/90  73/86  65/89  64/89
68/88  68/85  72/87  64/94  71/85  71/86  72/88  65/88  67/92  65/93  69/94  68/85  68/86  65/94  65/88  67/88  64/91  67/90  64/90  64/86  69/90  73/85  67/86  64/94  65/89  66/94  72/92
73/87  65/91  64/87  72/88  65/90  69/91  66/90  67/86  71/86  71/90  71/88  68/90  73/93  69/94  67/89  69/92  65/92  73/89  64/86  72/90  73/85  69/92  68/89  71/85  64/89  68/91  69/92
69/88  67/91  499/509  70/91  71/92  67/90  65/92  73/89  70/90  65/87  67/85  70/89  69/87  69/90  64/87  70/87  70/88  66/92  65/85  72/88  69/85  64/92  73/88  73/89  69/94  66/90  65/90
69/92  70/86  491/501  70/92  67/87  68/90  67/93  70/88  67/86  65/94  65/93  71/93  69/89  67/88  73/85  70/92  72/85  70/94  69/89  66/86  70/94  65/85  72/87  70/90  72/94  69/86  72/86
73/90  69/91  491/501  68/88  66/94  72/91  64/89  70/89  67/90  67/89  67/91  70/91  72/89  70/94  68/93  68/85  64/88  71/85  71/85  69/93  69/86  68/88  70/89  71/85  64/93  70/91  68/87
69/86  68/89  498/508  67/93  67/89  72/90  70/86  71/87  69/85  71/92  64/89  72/85  67/85  70/87  70/90  67/92  73/91  73/85  72/92  69/93  72/88  71/93  69/88  69/87  67/92  64/91  67/86
67/87  66/85  496/507  64/92  70/87  71/90  71/93  64/94  67/86  67/90  72/92  72/94  67/86  71/91  66/89  65/91  69/93  66/93  68/87  69/88  64/92  68/91  66/92  68/93  71/90  69/87  65/85
66/91  70/87  492/506  64/94  65/86  68/92  0/92  70/94  68/85  73/92  67/89  69/86  73/88  64/94  64/93  65/86  70/89  68/94  64/86  67/93  69/92  67/91  71/86  71/94  66/92  70/92  69/86
65/90  73/91  493/506  64/92  70/92  68/89  68/94  73/94  65/88  65/92  71/94  70/85  67/90  69/88  67/85  72/86  64/89  73/93  70/87  64/88  72/86  69/87  68/93  71/91  66/92  66/87  69/86
73/89  65/91  494/502  67/86  69/91  65/87  67/88  70/87  72/86  67/85  66/88  68/90  68/89  68/90  68/88  72/89  65/92  66/90  64/87  64/87  69/87  73/87  65/86  67/86  69/86  67/92  64/94
70/91  73/91  490/509  64/89  64/88  66/90  65/94  68/85  66/93  66/93  64/90  64/86  70/87  73/93  72/89  66/91  64/93  67/89  64/85  65/87  69/88  64/85  64/91  72/92  69/86  69/89  66/93
64/88  73/89  496/504  72/94  68/91  64/87  73/87  70/85  69/92  68/86  65/93  69/90  70/87  69/88  68/86  73/93  71/87  71/91  68/90  66/94  71/88  71/92  73/89  73/90  71/93  68/87  71/94
68/89  73/91  491/502  69/92  72/86  72/94  71/93  72/93  70/93  73/86  69/90  67/91  64/85  68/85  71/85  71/92  64/91  67/93  66/88  64/87  66/90  65/86  69/87  65/91  64/91  71/93  73/93
71/90  69/91  494/502  73/93  66/92  67/85  65/88  69/85  73/90  64/94  66/91  67/94  70/87  72/85  69/90  73/93  73/94  66/87  71/88  73/91  69/86  65/88  70/90  72/85  71/85  64/88  69/93
66/93  72/92  495/505  69/89  70/88  64/89  65/93  66/91  65/92  65/86  71/92  64/89  67/92  66/92  67/88  64/93  66/91  66/89  73/94  72/91  71/88  67/90  65/93  69/92  68/87  67/88  72/91
68/88  66/90  497/508  72/92  72/92  73/92  69/85  69/92  68/85  69/85  72/88  70/90  66/86  72/88  68/89  67/91  68/85  65/94  68/87  65/94  67/94  66/86  68/94  67/87  70/91  68/86  73/88
68/94  69/91  498/509  70/94  64/88  69/93  71/86  67/88  73/90  66/85  66/92  69/87  66/89  65/88  67/91  68/94  70/93  70/90  73/90  69/91  68/94  69/89  70/89  66/87  65/88  73/94  69/93
68/91  68/91  495/510  73/87  73/93  64/91  67/89  69/88  69/92  67/85  71/94  66/89  67/93  73/88  64/85  71/90  69/86  68/90  73/93  71/90  71/90  66/87  69/87  64/85  70/90  70/91  70/86
66/85  66/87  490/503  65/85  70/90  70/87  68/93  73/92  69/87  64/87  64/89  65/94  70/87  65/93  69/94  66/92  72/90  71/88  68/88  71/89  69/88  72/93  73/88  71/88  69/88  69/90  70/88
66/92  70/90  497/503  67/87  70/86  68/91  71/90  69/87  64/87  69/94  72/87  72/89  69/94  70/89  71/90  64/93  65/88  73/86  73/90  66/89  71/92  73/93  71/87  73/90  68/93  67/90  71/93
66/90  67/87  498/503  69/93  64/91  70/89  68/86  65/91  64/93  68/87  73/90  64/92  65/88  71/86  71/93  69/88  69/91  70/87  67/93  73/90  68/88  68/86  70/94  65/85  67/86  67/86  73/88
67/85  71/90  491/503  68/92  64/92  67/86  67/88  66/87  70/87  72/90  66/87  69/87  66/90  66/85  64/85  72/93  73/89  64/85  69/94  71/86  69/86  70/86  72/86  68/93  68/90  67/90  70/92
70/91  64/88  492/505  69/93  68/86  73/90  67/90  64/94  70/87  68/91  64/91  73/89  68/89  69/94  70/90  69/91  72/89  67/94  64/89  64/91  64/88  72/86  65/87  73/93  64/87  66/93  70/88
70/93  70/94  498/502  65/86  64/90  71/90  67/91  66/88  72/87  68/89  68/88  66/85  65/91  69/87  70/94  67/87  72/88  69/90  65/87  67/91  73/89  67/86  68/93  64/87  69/85  73/94  70/85

This produced my first hand-solve attempt:

33 + (5 * 36) = 213 (INCORRECT)

Second attempt involved using this codepen:

https://codepen.io/anon/pen/BQEZzK

Which pretty-printed the maze more nicely and allowed you to experiment by hand:

• . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . G
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . # # # # # # # # # # # # # # # # # # # # # #
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . _ . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .

It takes 34 moves to take the empty sector around the wall and along the
top edge to the top-right corner (which displaces the G sector one space
to the left--toward the goal). Then it's a 5-move pattern to move the G
each subsequent space to the left. Repeat that 35 times and you have the
following equation:

initial_moves = 34
repeated_moves = 5 * 35
initial_moves + repeated_moves = 209

That's the right answer! You are one gold star closer to fixing the sleigh.

You have completed Day 22!