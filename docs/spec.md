# VSOP87-GO Library Specification (draft)

- [VSOP87-GO Library Specification (draft)](#vsop87-go-library-specification-draft)
  - [Overview](#overview)
  - [Frames \& Reductions](#frames--reductions)
  - [Time Scales \& ΔT](#time-scales--δt)
  - [Sources \& Models](#sources--models)
  - [Units \& Precision](#units--precision)
  - [Public API Surface (stable)](#public-api-surface-stable)
  - [Assumptions \& Limits](#assumptions--limits)



## Overview
VSOP87-GO provides apparent ecliptic positions **of date** for the Sun, Moon, and planets, following **Jean Meeus, _Astronomical Algorithms_ (2nd ed.)**. Planetary series are taken from full VSOP87 datasets.

This document defines frames, reductions, data sources, time scales, units/precision, the stable API surface, and known limits.

---

## Frames & Reductions
- **Output frame:** Ecliptic & equinox **of date** (apparent-of-date).
- **Reduction order:** light-time → aberration → nutation (per Meeus AA 2e).
- **Precession/nutation:** applied as required to transform intermediate values to the output frame.

## Time Scales & ΔT
- Inputs may be provided in **JD_UTC** or **JD_TT** (depending on API).
- The library derives the other scale internally when needed.
- **ΔT model:** Espenak-Meeus as in Meeus AA 2e, extended with data from NASA
  _"Five Millennium Canon of Solar Eclipses: −1999 to +3000"_.

## Sources & Models
- **Planets:** **VSOP87D** series.
  - Source repository: `ctdk/vsop87` ([@faa1189 (2016-08-31)](https://github.com/ctdk/vsop87/commits/master/)).
- **Pluto:** Meeus Earth2000 (J2000) model, subsequently transformed to **of-date**
  via precession and nutation to match the output frame.
- **Moon:** Formulas per Meeus AA 2e.
  - Expected accuracy (per Meeus): ~10″ in longitude and ~4″ in latitude.

## Units & Precision
- Longitudes/latitudes (λ, β): **degrees** (decimal).
- True obliquity of date (ε): **degrees**.
- Radius vector (r): **astronomical units (AU)**.
- Time: **Julian Day** in **TT** and/or **UT**, as appropriate.
- Suggested printing precision (for downstream tools):  
  λ, β, ε — 9 decimal places; r — 9 decimal places; JD — 6 decimal places.

## Public API Surface (stable)
The following packages/types are considered **public and stable** (names illustrative; align to current code):
- `ephem` — high-level entry points to compute apparent ecliptic positions of date.
- `ephem.Body` — enumeration of supported bodies
  (`Sun, Moon, Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto`).
- Result types/structs carrying λ, β, r (and ε if exposed).

Everything under `internal/*` is **not** part of the stable API and may change without notice.

## Assumptions & Limits
- No topocentric
