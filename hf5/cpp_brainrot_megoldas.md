# C/C++ Brainrot Worksheet megoldások

## 1. Melyik karakter vagyok?

| Kódrészlet | Karakter neve |
| --- | --- |
| `int x = 10;` | Változó Vlad |
| `int* p = &x;` | Pointer Pete |
| `void f(int x) { x=99; }` | Xerox Chad |
| `void f(int& x) { x=99; }` | Key Dealer Kevin |
| `int* p = nullptr;` | Null Nick |
| `{ int n=5; } // n meghal` | Scope Sensei |
| `template<T> T add(T a, T b)` | Generic Greg |
| `return n * f(n-1);` | Recursion Rick |

## 2. Töltsd ki a hiányzó részeket!

### a) Xerox Chad

```cpp
void valtoztat(int x) { x = 999; }

int chad = 5;
valtoztat(chad);
// chad értéke még mindig: 5
```

Magyarázat: `x` csak másolat, ezért az eredeti `chad` nem változik.

### b) Pointer Pete

```cpp
int ertek = 55;
int* pete = &ertek;
*pete = 77;
printf("%d", ertek);  // kimenet: 77
```

Kimenet: `77`

Magyarázat: `pete` az `ertek` változó címét tárolja, ezért a `*pete = 77;` az eredeti változót írja át.

### c) Recursion Rick

```cpp
int osszeg(int n) {
    if (n == 0) return 0; // base case
    return n + osszeg(n - 1);
}
// osszeg(4) = 4+3+2+1+0 = 10
```

Hiányzó értékek:

- első üres hely: `0`
- második üres hely: `0`
- `osszeg(4)` eredménye: `10`

## 3. Melyik Boss Karakter jelenik meg?

### a)

```cpp
int* p = nullptr;
printf("%d", *p);
```

Boss: `Segfault Samu`

Miért: a pointer `nullptr`, és a program megpróbálja dereferálni, vagyis egy érvénytelen memóriacímet olvas.

### b)

```cpp
int* p = (int*)malloc(sizeof(int));
*p = 42;
// ... free() soha nem hívódik
```

Boss: `Memory Leak Miki`

Miért: a lefoglalt memória nincs felszabadítva `free()`-vel, ezért memóriaszivárgás történik.

### c)

```cpp
int vegtelen(int n) {
    return vegtelen(n); // base case?
}
```

Boss: `Stack Overflow Stan`

Miért: nincs base case, ezért a függvény végtelenül hívja önmagát, amíg a verem be nem telik.

## 4. Saját Brainrot Karakter

### Név

SFINAE Stefani

### Fogalom

`SFINAE` (Substitution Failure Is Not An Error)

### Személyiség

Stefani nem mondja azt, hogy "ez rossz", csak csendben eltünteti azokat a függvényverziókat, amelyek nem illenek a típushoz. Nem hibázik hangosan, csak úgy tesz, mintha az a lehetőség soha nem is létezett volna. Olyan, mint egy fordítási szintű klub portása.

### Idézet

"Ha nem passzolsz a sablonba, meg sem történtél."

### Kódpélda

```cpp
#include <type_traits>

template <typename T>
typename std::enable_if<std::is_integral<T>::value, void>::type
kiir(T ertek) {
    std::cout << "Egesz tipus: " << ertek << '\n';
}

template <typename T>
typename std::enable_if<!std::is_integral<T>::value, void>::type
kiir(T) {
    std::cout << "Ez nem egesz tipus.\n";
}

kiir(42);    // Egesz tipus: 42
kiir(3.14);  // Ez nem egesz tipus.
```

Magyarázat: a `SFINAE` miatt csak az a sablonpéldány marad játékban, amelyik illeszkedik az adott típusra. A többi nem fordítási hibát okoz, hanem egyszerűen kiesik a választásból.
