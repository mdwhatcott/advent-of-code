import sys

arrows = open('day03.txt').read().strip()

sys.setrecursionlimit(len(arrows)+10)

def move(at, arrow):
	if arrow == 'v': return (at[0], at[1]-1)
	if arrow == '^': return (at[0], at[1]+1)
	if arrow == '<': return (at[0]-1, at[1])
	if arrow == '>': return (at[0]+1, at[1])
	raise("bad instruction")

def visit(houses, arrows):
	if len(houses) == len(arrows):
		return set(houses)

	houses.append(move(houses[-1], arrows[len(houses)-1]))
	return visit(houses, arrows)

print('Part 1:', len(set(visit([(0, 0)], arrows))))

def filter(moves, pred):
	return [m for e, m in enumerate(moves) if pred(e)]

def odd(e):  return     e % 2
def even(e): return not e % 2

santa = set(visit([(0, 0)], filter(arrows, even)))
robot = set(visit([(0, 0)], filter(arrows, odd)))
print('Part 2:', len(santa | robot))