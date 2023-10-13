# IKT Etlap Projekt

[<img src="https://trello.com/favicon.ico" width="12"/> Trello](https://trello.com/b/OK3465So/ikt-etlap)\
[<img src="https://static.figma.com/app/icon/1/favicon.svg" width="12"/> Figma](https://www.figma.com/file/npwXQoPwjioLEQvZrfqbXX/Web-Project-0)

## Futtatás:
`php -S localhost:8080`

## PHP Fájlok:

### [index.php](./index.php)
Az URL alapján eldönti mivel válaszoljon a szerver.

### [main.php](./main.php)
Az fő oldal felépítését határozza meg és dinamikusan jeleníti meg a kívánt ételelket a következő kettő fájlnak megfelelően.

### [description.php](./pages/description.php)
Bővebb leírást ad egy adott ételről.

### [menu-cards.php](./data/menu-cards.php)
A hét napjain elérhető ételeket határozza meg.

### [cards-days.php](./data/cards-days.php)
A különböző ételek 'kártyáit', az az a hozzá tartozó adatokat határozza meg.\
Ezek az étel neve, a hozzátartózó kép és a leírás sorai.
