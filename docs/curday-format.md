# Curday Format

## TOC

- [Curday Format](#curday-format)
  - [TOC](#toc)
  - [Basic overview](#basic-overview)
  - [Header](#header)
  - [Channels](#channels)
    - [Channel Bit Flags](#channel-bit-flags)
  - [Programs](#programs)
    - [Special Characters](#special-characters)

## Basic overview

- Header
- Channel Information for Channel A
- Program Information for Program A
- Channel Infomration for Channel B
- Program Information for Program B
- etc
- CHANNEL/PROGRAM INFORMATION for TV Guide copyright block (Registered as Channel PRV002)

## Header

| Number of Bytes | Description         | Format                              | Example     |
| --------------- | ------------------- | ----------------------------------- | ----------- |
| 7               | Diagnostic Settings | `[BCK][FWD][SSPD][#AD][LINES][N/L]` | `AE3366N`   |
| 2               | ?                   | ? ?                                 | `0x03 0x01` |
| 2               | Timezone            | `[TZ][DST]`                         | `4Y`        |
| 1               | Cont                | `[Y]`                               | `Y`         |
| 8               | Ad Settings         | 7x`[Y/N]` 1x`l`                     | `YNNYYNNl`  |
| 2               | Padding             |                                     | `0x00 0x00` |
| 1               | Zero                |                                     | `0x60`      |
| 6               | DREV #              | `DREV [1-5]`                        | `DREV 5`    |
| 1               | Padding             |                                     | `0x00`      |
| x               | Airport ICAO        | `ABC..`                             | `KDCA`      |
| 1               | Padding             |                                     | `0x00`      |
| x               | City Name           | `City Name...`                      | `Threeway`  |
| 1               | Padding             |                                     | `0x00`      |
| x               | Julian Date         | `123`                               | `91`        |
| 1               | Padding             |                                     | `0x00`      |
| x               | Number of Listings  | `123`                               | `92`        |
| 1               | Padding             |                                     | `0x00`      |
| x               | ?                   | `131`                               | `131`       |
| 1               | Padding             |                                     | `0x00`      |
| x               | ?                   | `1348`                              | `1348`      |
| 1               | Padding             |                                     | `0x00`      |

## Channels

### Channel Bit Flags

| Flag              | Byte | Notes                                                          |
| ----------------- | ---- | -------------------------------------------------------------- |
| None              | 0x01 | Absolutely no clue what it does                                |
| HILITESRC         | 0x02 | Red gradient highlight                                         |
| SUMBYSRC          | 0x04 |                                                                |
| VIDEO_TAG_DISABLE | 0x08 |                                                                |
| CAF_PPVSRC        | 0x10 |                                                                |
| DITTO             | 0x20 | Listings don't change often? Potentially the double arrow flag |
| ALTHILITESRC      | 0x40 | Blue gradient highlight                                        |
| STEREO            | 0x80 | Station has stereo                                             |

## Programs

### Special Characters

| Name                              | Byte |
| --------------------------------- | ---- |
| Closed Captions                   | 0x7C |
| Right Triangle                    | 0x80 |
| Upside Down Right Triangle        | 0x81 |
| Right Double Triangle             | 0x82 |
| Upside Down Double Right Triangle | 0x83 |
| Rated R                           | 0x84 |
| Rated PG                          | 0x85 |
| Adult Content                     | 0x86 |
| Rated PG-13                       | 0x87 |
| Left Triangle                     | 0x88 |
| Upside Down Left Triangle         | 0x89 |
| Left Double Triangle              | 0x8A |
| Upside Down Double Left Triangle  | 0x8B |
| No Rating                         | 0x8C |
| Rated G                           | 0x8D |
| VCR+                              | 0x8E |
| Rated NC-17                       | 0x8F |
| Rated TV-Y                        | 0x90 |
| Headphones                        | 0x91 |
| Disney                            | 0x92 |
| Rated TV-Y7                       | 0x93 |
| 16 ANS+                           | 0x94 |
| HBO                               | 0x95 |
| 18 ANS+                           | 0x96 |
| 13 ANS+                           | 0x97 |
| Cinemax                           | 0x98 |
| Rated TV-G                        | 0x99 |
| Rated TV-14                       | 0x9A |
| Rated TV-PG                       | 0x9B |
| WOW                               | 0x9C |
| G TOUS                            | 0x9D |
| Prevue                            | 0x9E |
| Special Offer                     | 0x9F |
| Rated TV-M                        | 0xA1 |
| Upside Down Exclamation Mark      | 0xA2 |
| Rated TV-MA                       | 0xA3 |
