# Funkcionális Paradigma Brainrot Worksheet megoldások

## 1. Fogalomfelismerés

| Kódrészlet / leírás | Fogalom / Karakter |
| --- | --- |
| `\x -> x * 2` | lambda / névtelen függvény, `Lambda Lali` |
| `f . g` | függvényösszetétel, `Függvényösszetétel Fanni` |
| `add 5` | currying / részleges alkalmazás, `Curry Csaba` |
| `map f xs`, `filter p xs` | higher-order függvények, `Higher-Order Hugó` |
| `foldl (+) 0 [1,2,3]` | fold / reduce, `Fold Frida` |
| `f [] = 0   f (x:xs) = ...` | mintaillesztés, `Mintaillesztő Miska` |
| `take 10 [1..]` | lusta kiértékelés, `Lusta László` |
| Ugyanarra a bemenetre mindig ugyanaz az eredmény | tiszta függvény, `Tiszta Tivadar` |
| A kifejezés értéke nem függ helyétől a kódban | referenciális átlátszóság, `Hivatkozási Hanna` |
| Az érték, amiből már nem lehet tovább redukálni | normál forma, `Normál Forma Norbert` |

## 2. Redukciós lépések

### a) Beta-redukció

```haskell
(\x -> x + x) 5
-- 1. lépés (β): 5 + 5
-- 2. lépés:     10
-- Normál forma: 10
```

### b) Mi a normál forma?

- `(a) 2 + 3 * 4` → normál forma: `14`
- `(b) map (*2) [1,2]` → normál forma: `[2,4]`
- `(c) let x = x in x` → normál forma: `nincs`

Indok: önmagára hivatkozik, ezért végtelen redukcióba fut, nem jut el normál formáig.

## 3. Tail-rekurzió vs sima rekurzió

```haskell
-- A:
szorzat []     = 1
szorzat (x:xs) = x * szorzat xs
```

Tail? `NEM-TAIL`

Magyarázat: itt a rekurzív hívás még nincs kész önmagában, mert utána a programnak visszafelé még végre kell hajtania a szorzásokat. Emiatt ez nem tail-rekurzió.

```haskell
-- B:
szorzatT [] acc     = acc
szorzatT (x:xs) acc = szorzatT xs (x * acc)
```

Tail? `TAIL`

Magyarázat: ebben a verzióban minden részszámolás rögtön bekerül az `acc` változóba, így a rekurzív hívás tényleg az utolsó lépés. Ezért tail-rekurzív.

```haskell
-- C: tail-recursive verzió
osszegT [] acc     = acc
osszegT (x:xs) acc = osszegT xs (x + acc)
```

Ha teljes függvényként kell:

```haskell
osszeg xs = osszegT xs 0
```

## 4. Curry Csaba és függvényösszetétel

### a) Curry

```haskell
add x y = x + y

add3     = add 3
```

- `add3` típusa: `Num a => a -> a`
- `add3 7` → `10`
- `map (add 10) [1,2,3]` → `[11,12,13]`

### b) Függvényösszetétel

```haskell
duplaPlusz1 = (+1) . (*2)
```

- `duplaPlusz1 4` lépései: `4 * 2 = 8`, `+1 = 9`, eredmény: `9`

```haskell
kombinalt = (^2) . subtract 1 . (*3)
```

- `kombinalt 5` → `196`

Lépések:

- `5 * 3 = 15`
- `15 - 1 = 14`
- `14 ^ 2 = 196`

## 5. Lusta László vs Mohó Mónika

| Szituáció | LUSTA / MOHÓ / MINDKETTŐ |
| --- | --- |
| `take 5 [1..]` — végtelen lista első 5 eleme | `LUSTA` |
| `sum [1..100]` — véges lista összege | `MINDKETTŐ` |
| Fibonacci sorozat végtelen listája | `LUSTA` |
| Ha egy függvény argumentumát soha nem használjuk fel | `LUSTA` |

## 6. Saját Brainrot Karakter

### Név

Functor Feri

### Fogalom

Functor

### Személyiség

Feri bármin végig tud menni, amiben van érték, de a csomagolást békén hagyja. Nem szedi szét a dobozt, nem cseréli le, csak azt intézi el, ami benne van. Kicsit olyan, mint valaki, aki ugyanabban a tálcában hozza vissza az ételt, csak közben újrafűszerezte.

### Idézet

"A keret marad. A belsejét intézem."

### Kódpélda

```haskell
fmap (+1) (Just 4)      -- Just 5
fmap (*2) [1,2,3]       -- [2,4,6]
fmap reverse ["ab","cd"] -- ["ba","dc"]
```

Magyarázat: a `Functor` lényege, hogy a belső értéket át tudod alakítani anélkül, hogy közben elveszne a körülötte lévő struktúra. Vagyis nem a dobozt cseréled le, csak azt, ami benne van.
