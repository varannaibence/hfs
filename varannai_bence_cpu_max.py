import multiprocessing
import math
import os

def _sine_cruncher(core_id: int) -> None:
    seed = core_id + 1
    angle = 0.1234 * seed
    while True:
        angle = math.sin(angle) * math.cos(angle + seed) * math.tan(angle + seed * 0.5) + seed

def cpu_max_sine():
    for core in range(multiprocessing.cpu_count()):
        p = multiprocessing.Process(target=_sine_cruncher, args=(core,), daemon=False)
        p.start()
    print(f"CPU-maxoló elindítva minden magon (PID: {os.getpid()})")

if __name__ == "__main__":
    cpu_max_sine()
