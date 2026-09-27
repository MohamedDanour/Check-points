import csv, re

rows = list(csv.DictReader(open('progress.csv', encoding='utf-8')))
people = [p for p in rows[0].keys() if p != 'question'] if rows else []
total = len(rows)

def solved(row, person):
    return (row[person] or '').strip() != ''

def bar(n):
    pct = round(n / total * 100) if total else 0
    return f"`{'█' * (pct // 5)}{'░' * (20 - pct // 5)}` {n} / {total} ({pct}%)"

any_solved = sum(1 for r in rows if any(solved(r, p) for p in people))
all_solved = sum(1 for r in rows if all(solved(r, p) for p in people))

lines = [f'**{any_solved} / {total} solved**', '', bar(any_solved), '']
for p in people:
    n = sum(1 for r in rows if solved(r, p))
    lines.append(f'- **{p}**: {n} solved  ')
    lines.append(f'  {bar(n)}')
lines += ['', f'🤝 Solved by both: **{all_solved}**', '']

# Checklist table
lines.append('| # | Question | ' + ' | '.join(people) + ' |')
lines.append('|---:|---|' + '|'.join([':---:'] * len(people)) + '|')
for i, r in enumerate(rows, 1):
    marks = ['✅' if solved(r, p) else '⬜' for p in people]
    q = r['question'] + (' 🤝' if all(solved(r, p) for p in people) else '')
    lines.append(f'| {i} | {q} | ' + ' | '.join(marks) + ' |')

block = '\n'.join(lines)
readme = open('README.md', encoding='utf-8').read()
readme = re.sub(r'<!-- PROGRESS:START -->.*?<!-- PROGRESS:END -->',
                lambda m: '<!-- PROGRESS:START -->\n' + block + '\n<!-- PROGRESS:END -->',
                readme, flags=re.S)
open('README.md', 'w', encoding='utf-8').write(readme)