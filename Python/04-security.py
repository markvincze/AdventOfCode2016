from collections import defaultdict

def parse_line(line):
    bracket_index = line.find("[")
    text = line[:bracket_index]
    checksum = line[bracket_index + 1:-1]

    sector_id = text[(text.rfind("-") + 1):]
    text = text[0:(text.rfind("-"))]

    return text, int(sector_id), checksum

def get_checksum(str):
    letter_counts = defaultdict(int)
    for c in filter(lambda c: c.isalpha(), str):
        letter_counts[c] = letter_counts[c] + 1
    
    for x in letter_counts.items():
        k = x[0]
        c = x[1]

    s = sorted(letter_counts.items(), key=lambda x: (-x[1], x[0]))[:5]

    return "".join(k for k, _ in s)

input = []

with open("04-security-input.txt") as file:
    for l in file:
        t, sid, c = parse_line(l.rstrip())
        # print(f"text: {t}, sid: {sid}, c: {c}, calc_c: {get_checksum(t)}")
        input.append((t, sid, c))

result1 = sum([sid for t, sid, c in input if get_checksum(t) == c])
print(f"Result 1: {result1}")

def decrypt(text, sector_id):
    def shift(c, n):
        if c == '-':
            return ' '
        else:
            return chr(((ord(c) - ord('a') + n) % (ord('z') - ord('a') + 1)) + ord('a'))

    return "".join([shift(c, sector_id) for c in text])

for t, sid, c in input:
    if get_checksum(t) == c:
        decrypted = decrypt(t, sid)
        if "north" in decrypted:
            print(f"{sid}: {decrypted}")
