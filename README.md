# Bitwise Comparable Datetime Encoding

Solana nodes provide memcmp for filters based on byte comparisons. We introduce
bitwise comparable datetime (BCDT), that encodes date time to minute precision
in 4-bytes, an int32. YY-MM-DD are encoded in the first two bytes, HH-MM in the
second two bytes. This allows lookups by day and by day and time.

```
| year 4 bit | month 4 bit | day 5 bit | 5 bit hour | 6 bit minutes |
|       0000 |        1100 |  00011111 |   00010111 |      00111011 |
|               2 bytes                |           2 bytes          |
|            year+month+day            |           datetime         |
```

Example with seconds cut out.

```
date epoch     1672531140
date lookup    2022-12-31 23:59:00 +0000 UTC
date dec       203364155
output bin     00001100000111110001011100111011
```

the last 3 bytes are padded to align day, hour and minute on bytes 2, 3 and 4
respectively:

```
  byte 1  byte 2   byte3    byte4
|       | pad 3  | pad 3  | pad 2   |
| 4 | 4 |      5 |      5 |      6  |
```

We redefine the year 2022 to year zero and use a 4 bit encoding. The year is
therefore encodable with:

    <year> % 2022 % 2^4

The next year zero will be 2038.

Supported queries based on byte comparisons with `memcmp`:
- year + month: use byte 1 to encode year and month
- year + month + day: use byte 1,2
- year + month + day + hou: use byte 1,2,3
- year + month + day + hour + minute: use byte 1,2,3,4
