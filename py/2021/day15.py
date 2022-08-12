from norvig import Grid, Tuple, Astar, parse, digits, answer

def Astar_search_grid(grid, start=(0, 0))  -> Tuple[int, list]:
    """The (risk, path) tuple of the best path from start to bottom-right on grid."""
    goal = max(grid)
    def neighbors(s): return grid.neighbors(s)  # possible moves
    def h_func(s): return sum(goal) - sum(s)    # estimated path cost from s to goal
    def step_cost(_, s2): return grid[s2]       # cost of moving to s2
    return Astar(start, neighbors, h_func, step_cost)

in15 = Grid(rows=parse(15, digits))
path = Astar_search_grid(in15)
answer(15.2, path[0], 824)
