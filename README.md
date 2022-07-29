# Bitwise Comparable Datetime Encoding

Solana nodes provide memcmp for filters based on byte level comparisons. We
introduce the bitwise comparable datetime (BCDT). BCDT encodes datetime to
minute precision in 4-bytes, an int32 while supporting byte queries by year and
month, by day and by time. Year and month are encoded in the first byte, the
next three bytes each eoncode day, hour and minute. The last 3 bytes are zero
padded to align day, hour and minute on bytes 2, 3 and 4 respectively.

```
| 4 bit year | 4 bit month | 5 bit day | 5 bit hour | 6 bit minute |
|       0000 |        1100 | 000.11111 |  000.10111 |    00.111011 |
|          1 byte          |   1 byte  |   1 byte   |    1 byte    |
|    query year+month      | query day | query hour | query minute |
```

Example encoding from Unix epoch:

```
date epoch     1672531140
date lookup    2022-12-31 23:59:00 +0000 UTC
date dec       203364155
output bin     00001100000111110001011100111011
```

We redefine the year 2022 to year zero and use a 4 bit encoding. The year is
therefore encodable with:

    <year> % 2022 % 2^4

The next year zero will be 2038.

Supported queries based on byte comparisons with `memcmp`:
- year + month: use byte 1 to encode year and month
- year + month + day: byte 1,2
- year + month + day + hour: byte 1,2,3
- year + month + day + hour + minute: byte 1,2,3,4
