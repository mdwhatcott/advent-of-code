from collections import defaultdict
import fileinput
import re


tasks = set()
deps = defaultdict(set)

for line in fileinput.input():
    a, b = re.findall(r' ([A-Z]) ', line)
    tasks |= {a, b}
    deps[b].add(a)

# part 1
done = []
for _ in tasks:
    done.append(min(x for x in tasks if x not in done and deps[x] <= set(done)))

print(''.join(done))

# part 2
done = set()
order = []
seconds = 0       # total seconds elapsed
counts = [0] * 5  # seconds remaining for worker `i` to finish its current task
work = [''] * 5   # which task worker `i` is performing

while True:
    for i, count in enumerate(counts):
        if count == 1:
            done.add(work[i])
            order.append(work[i])
        counts[i] = max(0, count - 1)

    while 0 in counts:
        idle = counts.index(0)
        ready = [x for x in tasks if deps[x] <= done]  # deps[x] is equal to or subset of done; all prereqs satisfied
        if not ready:
            break

        task = min(ready)
        tasks.remove(task)

        # have the worker start the selected task
        counts[idle] = ord(task) - ord('A') + 61
        work[idle] = task

    if sum(counts) == 0:
        break

    seconds += 1


print(seconds)
print(''.join(order))