def ZipString(s):
    d = {}
    for i in s:
        d[i] = 0
    for i in s:
        if i in d.keys():
            d[i]+=1
    out = ""
    for i in d.keys():
        out+=str(d[i])
        out+=i
    return out
print(ZipString("YouuungFellllas"))