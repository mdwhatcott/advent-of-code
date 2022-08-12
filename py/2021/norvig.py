# Reference: https://github.com/norvig/pytudes/blob/main/ipynb/Advent-2021.ipynb
from __future__  import annotations
from collections import Counter, defaultdict, namedtuple, deque
from itertools   import permutations, combinations, chain, count as count_from, product as cross_product
from typing      import *
from statistics  import mean, median
from math        import ceil, inf
from functools   import lru_cache
import matplotlib.pyplot as plt
import re
from heapq import heappop, heappush

def answer(puzzle_number, got, expected) -> bool:
    """Verify the answer we got was the expected answer."""
    assert got == expected, f'For {puzzle_number}, expected {expected} but got {got}.'
    return True

def parse(day, parser=str, sep='\n', print_lines=7) -> tuple:
    """Split the day's input file into entries separated by `sep`, and apply `parser` to each."""
    fname = f'py/2021/day{day}.txt'
    text  = open(fname).read()
    entries = mapt(parser, text.rstrip().split(sep))
    if print_lines:
        all_lines = text.splitlines()
        lines = all_lines[:print_lines]
        head = f'{fname} ➜ {len(text)} chars, {len(all_lines)} lines; first {len(lines)} lines:'
        dash = "-" * 100
        print(f'{dash}\n{head}\n{dash}')
        for line in lines:
            print(trunc(line))
        print(f'{dash}\nparse({day}) ➜ {len(entries)} entries:\n'
              f'{dash}\n{trunc(str(entries))}\n{dash}')
    return entries

def trunc(s: str, left=70, right=25, dots=' ... ') -> str: 
    """All of string s if it fits; else left and right ends of s with dots in the middle."""
    dots = ' ... '
    return s if len(s) <= left + right + len(dots) else s[:left] + dots + s[-right:]

Char = str # Intended as the type of a one-character string
Atom = Union[float, int, str]

def ints(text: str) -> Tuple[int]:
    """A tuple of all the integers in text, ignoring non-number characters."""
    return mapt(int, re.findall(r'-?[0-9]+', text))

def digits(text: str) -> Tuple[int]:
    """A tuple of all the digits in text (as ints 0–9), ignoring non-digit characters."""
    return mapt(int, re.findall(r'[0-9]', text))

def words(text: str) -> List[str]:
    """A list of all the alphabetic words in text, ignoring non-letters."""
    return re.findall(r'[a-zA-Z]+', text)

def atoms(text: str) -> Tuple[Atom]:
    """A tuple of all the atoms (numbers or symbol names) in text."""
    return mapt(atom, re.findall(r'[a-zA-Z_0-9.+-]+', text))

def atom(text: str) -> Atom:
    """Parse text into a single float or int or str."""
    try:
        x = float(text)
        return round(x) if round(x) == x else x
    except ValueError:
        return text
    
def mapt(fn, *args) -> tuple:
    """map(fn, *args) and return the result as a tuple."""
    return tuple(map(fn, *args))

def quantify(iterable, pred=bool) -> int:
    """Count the number of items in iterable for which pred is true."""
    return sum(1 for item in iterable if pred(item))

class multimap(defaultdict):
    """A mapping of {key: [val1, val2, ...]}."""
    def __init__(self, pairs: Iterable[tuple], symmetric=False):
        """Given (key, val) pairs, return {key: [val, ...], ...}.
        If `symmetric` is True, treat (key, val) as (key, val) plus (val, key)."""
        self.default_factory = list
        for (key, val) in pairs:
            self[key].append(val)
            if symmetric:
                self[val].append(key)

def prod(numbers) -> float: # Will be math.prod in Python 3.8
    """The product formed by multiplying `numbers` together."""
    result = 1
    for x in numbers:
        result *= x
    return result

def total(counter: Counter) -> int: 
    """The sum of all the counts in a Counter."""
    return sum(counter.values())

def sign(x) -> int: return (0 if x == 0 else +1 if x > 0 else -1)

def transpose(matrix) -> list: return list(zip(*matrix))

def nothing(*args) -> None: return None

cat     = ''.join
flatten = chain.from_iterable
cache   = lru_cache(None)

Point = Tuple[int, int] # (x, y) points on a grid

neighbors4 = ((0, 1), (1, 0), (0, -1), (-1, 0))               
neighbors8 = ((1, 1), (1, -1), (-1, 1), (-1, -1)) + neighbors4

class Grid(dict):
    """A 2D grid, implemented as a mapping of {(x, y): cell_contents}."""
    def __init__(self, mapping=(), rows=(), neighbors=neighbors4):
        """Initialize with, e.g., either `mapping={(0, 0): 1, (1, 0): 2, ...}`,
        or `rows=[(1, 2, 3), (4, 5, 6)].
        `neighbors` is a collection of (dx, dy) deltas to neighboring points.`"""
        self.update(mapping if mapping else
                    {(x, y): val 
                     for y, row in enumerate(rows) 
                     for x, val in enumerate(row)})
        self.width  = max(x for x, y in self) + 1
        self.height = max(y for x, y in self) + 1
        self.deltas = neighbors
        
    def copy(self) -> Grid: return Grid(self, neighbors=self.deltas)
    
    def neighbors(self, point) -> List[Point]:
        """Points on the grid that neighbor `point`."""
        x, y = point
        return [(x+dx, y+dy) for (dx, dy) in self.deltas 
                if (x+dx, y+dy) in self]
    
    def to_rows(self) -> List[List[object]]:
        """The contents of the grid in a rectangular list of lists."""
        return [[self[x, y] for x in range(self.width)]
                for y in range(self.height)]


def Astar(start, neighbors, h_func, step_cost) -> Tuple[int, list]:
    """Find a (cost, path) tuple for the lowest-cost path from start to a goal.
    A goal is any state `s` such that `h_func(s) == 0`."""
    frontier  = [(h_func(start), start)] # A priority queue, ordered by path_cost(s) + h(s)
    previous  = {start: None}  # start state has no previous state; other states will
    path_cost = {start: 0}     # The cost of the best path to a state.
    Path      = lambda s: ([] if (s is None) else Path(previous[s]) + [s])
    while frontier:
        (f, s) = heappop(frontier)
        if h_func(s) == 0:
            return path_cost[s], Path(s)
        for s2 in neighbors(s):
            g = path_cost[s] + step_cost(s, s2)
            if s2 not in path_cost or g < path_cost[s2]:
                heappush(frontier, (g + h_func(s2), s2))
                path_cost[s2] = g
                previous[s2] = s

