# OOP Brainrot Worksheet megoldások

## 1. Melyik karakter vagyok?

| Kódrészlet | Fogalom / Karakter |
| --- | --- |
| `class Kutya { ... };` | Osztály, `Osztály Olga` |
| `Kutya k("Rex");` | Objektum / példányosítás, `Objektum Oszkár` vagy `Példányosító Pál` |
| `Kutya(string n) { nev=n; }` | Konstruktor, `Konstruktor Kristóf` |
| `~Kutya() { delete adat; }` | Destruktor, `Destruktor Dezső` |
| `class Kutya : public Allat` | Öröklés, `Öröklődő Olivér` |
| `virtual void hangotAd() = 0;` | Absztrakt metódus / pure virtual, `Absztrakt Anikó` |
| `void f() override { ... }` | Polimorfizmus / felüldefiniálás, `Polimorf Péter` |
| `private: int egyenleg;` | Bezárás / enkapszuláció, `Bezáró Bence` |
| `Allat* a = new Kutya(); a->hangotAd();` | Késői kötés / dynamic dispatch, `Late Binding Lajos` |

## 2. Töltsd ki a hiányzó részeket!

### a) Konstruktor Kristóf és Destruktor Dezső

```cpp
class Auto {
public:
    std::string marka;
    int* ev;

    Auto(std::string m) {
        marka = m;
        ev = new int(2024);
    }

    ~Auto() {
        delete ev;
    }
};
```

Hiányzó részek:

- első üres hely: `Auto`
- második üres hely: `~`
- harmadik üres hely: `ev`

### b) Bezáró Bence

```cpp
class BankSzamla {
private:
    int egyenleg = 0;

protected:
    int hitelkeret = 50000;

public:
    void befizet(int x);
};
```

Hiányzó kulcsszavak:

- `egyenleg` → `private`
- `hitelkeret` → `protected`
- `befizet` → `public`

### c) Polimorf Péter

```cpp
Allat* allatok[] = { new Kutya(), new Macska(), new Kutya() };
for (auto* a : allatok) a->hangotAd();
```

Kimenetek sorban:

- `Vau!`
- `Miau!`
- `Vau!`

## 3. Absztrakt Anikó vs Interfész István

`X` = interfész

Miért: az `X` csak egy szerződést ad. Pure virtual függvényekből áll, és nincs benne saját állapot vagy kész működés, amit a leszármazottak megörökölnének.

`Y` = absztrakt osztály

Miért: a `Y` már nem csak előír dolgokat, hanem ad is valamit a gyerekosztályoknak. Van benne konkrét metódus (`legzik`) és adattag (`energia`) is, ezért ez absztrakt osztály, nem tiszta interfész.

## 4. Early Eddie vs Late Binding Lajos

| Szituáció | EARLY / LATE |
| --- | --- |
| `Kutya k; k.ugat();` (`ugat` nem virtual) | `EARLY` |
| `Allat* a = new Kutya(); a->hangotAd();` (`virtual`) | `LATE` |
| `osszeAdd<int>(3,4);` (`template`) | `EARLY` |
| Függvény overloading (`f(int)` vs `f(double)`) | `EARLY` |
| `virtual void f() = 0;` leszármazotton hívva | `LATE` |

## 5. Mini osztályhierarchia

### Példa: Jármű hierarchia

`Jarmu` -> `MotorosJarmu` -> `Auto`

### Szerepek

- `Jarmu` absztrakt osztály
- `MotorosJarmu` absztrakt vagy részben konkrét köztes osztály
- `Auto` konkrét osztály

### Hozzáférések

- `private`: `alvazszam`
- `protected`: `uzemanyagSzint`
- `public`: `indit()`, `megall()`

### Virtual metódusok

```cpp
class Jarmu {
private:
    std::string alvazszam;

protected:
    int sebesseg = 0;

public:
    virtual void mozog() = 0;
    virtual ~Jarmu() = default;
};

class MotorosJarmu : public Jarmu {
protected:
    int uzemanyagSzint = 100;

public:
    void tankol(int liter) { uzemanyagSzint += liter; }
};

class Auto : public MotorosJarmu {
public:
    void mozog() override {
        // az Auto sajat megvalositasa
    }

    void dudal() {
        // konkret extra viselkedes
    }
};
```

Magyarázat: a `Jarmu` lefekteti az alapot, vagyis azt, hogy minden járműnek tudnia kell mozogni. A `MotorosJarmu` erre ráteszi a motoros közös részeket, például az üzemanyagszintet és a tankolást. Az `Auto` pedig már egy konkrét, használható osztály, ami ténylegesen megvalósítja a mozgást, és még saját plusz viselkedést is kap, például a dudálást.
