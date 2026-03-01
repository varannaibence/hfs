kulcs = b"42069333"
with open("/Users/varannaibence/Documents/GitHub/varannai_bence_cpu_ram_max_hf/hf3/2/bence.txt", "rb") as f:
    adat = f.read()

titkos = bytes(b ^ kulcs[i % len(kulcs)] for i, b in enumerate(adat))

with open("titkos.bin", "wb") as f:
    f.write(titkos)