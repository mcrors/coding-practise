# Day of the Programmer

## Problem

Marie has a time machine. She wants to visit Russia on the Day of the Programmer. This is the 256th day of the year. The year is in the range 1700 to 2700, inclusive.

Russia used two calendar systems in this range:

- From 1700 to 1917, Russia used the Julian calendar.
- In 1918, Russia changed from the Julian calendar to the Gregorian calendar.
- From 1919, Russia used the Gregorian calendar.

In 1918, the day after January 31 was February 14. February 14 was the 32nd day of the year.

February is the only month with a variable number of days. It has 29 days in a leap year. It has 28 days in all other years.

A leap year follows one of these rules:

- Julian calendar: The year is divisible by 4.
- Gregorian calendar: The year is divisible by 400.
- Gregorian calendar: The year is divisible by 4 and not divisible by 100.

Given a year `y`, find the date of the 256th day of that year in the Russian calendar. Return the date in the format `dd.mm.yyyy`.

## Example

The year is 1984. The year is divisible by 4, so it is a leap year. The 256th day of a leap year after 1918 is September 12. The answer is `12.09.1984`.

## Function description

Complete the `dayOfProgrammer` function.

`dayOfProgrammer` has this parameter:

- `year`: an integer.

It returns:

- `string`: the date of the 256th day of the year, in the format `dd.mm.yyyy`.

## Input format

A single integer. This integer is the year `y`.

## Constraints

- 1700 <= `y` <= 2700

## Output format

Print the full date in the format `dd.mm.yyyy`.

- `dd` is the two-digit day.
- `mm` is the two-digit month.
- `yyyy` is the year.

## Sample input 0

```
2017
```

## Sample output 0

```
13.09.2017
```

### Explanation 0

The year 2017 is not a leap year. The first eight months have these days:

31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 = 243

The 256th day is 256 - 243 = 13. It is day 13 of month 9 (September). The answer is `13.09.2017`.

## Sample input 1

```
2016
```

## Sample output 1

```
12.09.2016
```

### Explanation 1

The year 2016 is a leap year. February has 29 days. The first eight months have these days:

31 + 29 + 31 + 30 + 31 + 30 + 31 + 31 = 244

The 256th day is 256 - 244 = 12. The answer is `12.09.2016`.

## Sample input 2

```
1800
```

## Sample output 2

```
12.09.1800
```

### Explanation 2

The year 1800 is a leap year in the Julian calendar. The 256th day is September 12.

## Approach

Only February changes the date. The other seven months in the first eight months have 215 days in total.

The 256th day is in September. The day number in September is:

```
day = 256 - 215 - daysInFebruary = 41 - daysInFebruary
```

Use these steps:

1. If the year is 1918, February has 15 days. The answer is `26.09.1918`.
2. If the year is before 1918, use the Julian rule for a leap year.
3. If the year is after 1918, use the Gregorian rule for a leap year.
4. If the year is a leap year, the day is 12.
5. If the year is not a leap year, the day is 13.
6. Print the day with two digits, then `.09.`, then the year.

### Why 1918 has 15 days in February

In 1918, February starts on day 14. The month ends on day 28. This gives 28 - 14 + 1 = 15 days. The year has 31 + 15 = 46 days after the two months. The day number is 41 - 15 = 26.

## Test cases

| Year | Calendar            | Leap year | Expected     |
| ---- | ------------------- | --------- | ------------ |
| 1700 | Julian              | Yes       | `12.09.1700` |
| 1800 | Julian              | Yes       | `12.09.1800` |
| 1900 | Julian              | Yes       | `12.09.1900` |
| 1917 | Julian              | No        | `13.09.1917` |
| 1918 | Transition          | Not used  | `26.09.1918` |
| 1919 | Gregorian           | No        | `13.09.1919` |
| 2000 | Gregorian           | Yes       | `12.09.2000` |
| 2016 | Gregorian           | Yes       | `12.09.2016` |
| 2017 | Gregorian           | No        | `13.09.2017` |
| 2100 | Gregorian           | No        | `13.09.2100` |
| 2700 | Gregorian           | No        | `13.09.2700` |

The years 1900 and 2100 test the century rule. The year 1900 is a leap year in the Julian calendar. The year 2100 is not a leap year in the Gregorian calendar.
