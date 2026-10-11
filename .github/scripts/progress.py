import csv, re

META = {'question', 'level'}

rows = list(csv.DictReader(open('progress.csv', encoding='utf-8')))
people = [p for p in rows[0].keys() if p not in META] if rows else []
has_level = bool(rows) and 'level' in rows[0]
total = len(rows)

def solved(row, person):
    return (row[person] or '').strip() != ''

def bar(n):
    pct = round(n / total * 100) if total else 0
    return f"`{'█' * (pct // 5)}{'░' * (20 - pct // 5)}` {n} / {total} ({pct}%)"

all_solved = sum(1 for r in rows if all(solved(r, p) for p in people))

lines = []
for p in people:
    n = sum(1 for r in rows if solved(r, p))
    lines.append(f'- **{p}** {bar(n)}')
lines += ['', f'🤝 Solved by both: **{all_solved} / {total}**', '']

# Checklist table
level_head = ' Level |' if has_level else ''
level_align = ':---:|' if has_level else ''
lines.append('| # | Question |' + level_head + ' ' + ' | '.join(people) + ' |')
lines.append('|---:|---|' + level_align + '|'.join([':---:'] * len(people)) + '|')
for i, r in enumerate(rows, 1):
    marks = ['✅' if solved(r, p) else '⬜' for p in people]
    q = r['question'] + (' 🤝' if all(solved(r, p) for p in people) else '')
    level = f" {r['level']} |" if has_level else ''
    lines.append(f'| {i} | {q} |{level} ' + ' | '.join(marks) + ' |')

block = '\n'.join(lines)
readme = open('README.md', encoding='utf-8').read()
readme = re.sub(r'<!-- PROGRESS:START -->.*?<!-- PROGRESS:END -->',
                lambda m: '<!-- PROGRESS:START -->\n' + block + '\n<!-- PROGRESS:END -->',
                readme, flags=re.S)
open('README.md', 'w', encoding='utf-8').write(readme)