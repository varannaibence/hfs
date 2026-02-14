import psutil
import time

# 1. VERZIÓ: Gyors lista allokáció
def ram_max_v1():
    data = []
    while True:
        data.append(' ' * 10**6)  # 1MB blokkok

# 2. VERZIÓ: Agresszív memória foglalás
def ram_max_v2():
    memory = []
    while True:
        memory.extend([0] * 10**7)  # 10M elemek

# BIZTONSÁGOS TESZT: Limitált memória használat
def ram_test_safe(target_mb=5000, duration=15):
    """Teszt mód: 'target_mb' MB-ot foglal 'duration' másodpercig"""
    print(f"Start RAM: {psutil.virtual_memory().percent}%")
    data = []
    
    for i in range(target_mb):
        data.append(' ' * 10**6)
        if i % 100 == 0:
            print(f"Foglalt: {i}MB | RAM: {psutil.virtual_memory().percent}%")
    
    print(f"Max RAM: {psutil.virtual_memory().percent}%")
    time.sleep(duration)
    print("Teszt vége, memória felszabadítva")

# TESZTELÉS:
#ram_test_safe(target_mb=1000, duration=15)
# ÉLES használat:
ram_max_v1()
