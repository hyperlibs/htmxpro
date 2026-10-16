# HTMX PRO — Deep Gap Analysis vs React/Node

**Date:** 2026-10-16 · **Scope:** repository `hyperlibs/htmxpro` @ `da021576` · **Method:** static inspection + one test run

Every claim below is tied to a file, a count, or a command result observed in the working tree.
Where a claim is a judgement rather than a measurement, it is marked *(judgement)*.

---

## 1. Verified baseline

| Fact | Evidence |
|---|---|
| 15 test files, **88 pass / 0 fail**, 749 assertions, 925 ms | `bun test tests/` → exit 0 |
| 160 HTML components in `views/components/` | `ls views/components \| wc -l` |
| **54** exported Go constructors across `go/components/*.go` | `grep -h "^func [A-Z]" go/components/*.go \| wc -l` |
| 13 client engines, 9,605 lines of TypeScript in `src/` | `wc -l src/*.ts` |
| `public/htmx-bolt.js` = **107,609 B / 3,031 lines, unminified** | `wc -c`, `wc -l` |
| `public/styles/output.css` = **196,417 B / 6,266 lines, unminified** | `wc -c`, `wc -l` |
| CI matrix: ubuntu/macos/windows, engines build, checksums, tests, `<300 KB` JS budget | `.github/workflows/ci.yml` |
| Repo pack **97.77 MiB**; `dist/` tracked (172 files, 3.1 MB) | `git count-objects -vH`, `git ls-files dist \| wc -l` |

**Baseline verdict:** the client-engine layer is genuinely solid — real tests, real checksums, real size budget, a real JSON grammar spec (`schema/htmxpro.json`). The gaps are almost entirely in the **application platform**: packaging, type safety, server-side capabilities, offline, and the promised component parity.

---

## 2. Where HTMX PRO loses to React/Node

| # | Axis | React/Node | HTMX PRO today | Verdict |
|---|---|---|---|---|
| 2.1 | **Compile-time safety** | TSX fails the build on a bad prop; type errors are impossible to ship | `hx-*` attributes are unvalidated strings. No `tsconfig.json` exists, so **the 9,605 lines of "100% Strict TypeScript" are never type-checked**. `schema/htmxpro.json` is advisory prose-equivalent — nothing runs it in CI | **React wins outright** |
| 2.2 | **Component encapsulation** | Props boundary, typed contracts, colocation | 160 components are copy-paste HTML with inline Tailwind strings; no props contract between a view and its consumers | **React wins** |
| 2.3 | **Ecosystem & talent** | npm registry, Testing Library, DevTools, millions of examples | No markup-testing story (only **1** test file references `views/components`), no component registry, no framework-agnostic testing harness | **React wins** |
| 2.4 | **Offline-first + client routing** | SPA + service worker is a documented pattern | No service worker, no manifest, no client routes. `htmx-offline.ts` uses `indexedDB` once with zero `serviceWorker`/`caches.*` references — data cache only, no app shell | **React wins** |
| 2.5 | **Interaction latency** | Local state mutation in <1 ms | Every interaction is a round-trip. `htmx-vibe`/`htmx-sim` do FLIP + physics, but sustained 60 fps local editing still routes through the server | **React wins on interaction-heavy UI** |
| 2.6 | **Cost model** | One bundle, then client CPU is free | Server CPU scales with interactions. Read-mostly apps get cheaper; chatty UI gets more expensive per user | **Tradeoff, not a loss** |
| 2.7 | **Ecosystem lock-out** | Composability with the entire JS ecosystem | 13 hand-written engines replace libraries you cannot otherwise use | **Tradeoff** — this is the whole point of the project |
| 2.8 | **Bundle + hydration** | 40–150 KB framework + hydration cost | No hydration, no VDOM, 22 KB gzipped bolt engine | **HTMX PRO wins** |
| 2.9 | **Version hell** | React 18/19, Next pages/app-router ambiguity, breaking ecosystem churn | One htmx pin + engine files | **HTMX PRO wins** |
| 2.10 | **Agent token cost** | Infinite API surface; every feature re-litigated | Constrained vocabulary is a real advantage *if* the parity and validation gaps above are closed | **HTMX PRO can win — currently unproven** (see §5.3) |

**Blunt summary:** today `htmxpro` is a superb **client-engine library** and an immature **application platform**. React loses on weight and token surface; HTMX PRO loses on safety, offline, and server-side capability.

---

## 3. Gap matrix by target application

### 3.1 PWA mobile app (Go **or** Python backend)

| Requirement | Status | Evidence |
|---|---|---|
| Web manifest / `theme-color` / `apple-mobile-web-app-*` | ❌ **missing** | `app-mobile.html` has `viewport` only |
| Service worker, precache, offline shell | ❌ **missing** | no `serviceWorker`/`sw.js` anywhere outside the browser extension (`releases/*.crx`, `hmlr.mx`) |
| Install prompt / app lifecycle | ❌ **missing** | no `beforeinstallprompt` handler |
| Local htmx copy (offline hard requirement) | ❌ **missing** | 13 references to `https://unpkg.com/htmx.org@2.0.4`, zero vendored copy |
| Touch/safe-area shell | ⚠️ partial | `safe-area-inset` present in `go/components/appz.go`; `app-*` shell exists |
| Push notifications | ❌ **missing** | no `pushsubscription` / Notification permission flow |
| Offline mutation queue + conflict resolution | ✅ present | `htmx:offlineConflict` event, IndexedDB path in `htmx-offline.ts` |
| Go backend story | ⚠️ partial | `go/htmxpro/htmxpro.go` exists (23 funcs) but **no `go:embed` anywhere** → no single binary, contradicting AGENTS.md §1 |
| Python backend story | ❌ **missing** | Python appears **only as a code sample** in `docs/BACKEND_SPEC.md`. No `.py`, no Jinja macros, no FastAPI/ASGI adapter, no PyPI artifact |

### 3.2 Operating-system UI (desktop-like shell in the browser)

| Requirement | Status | Evidence |
|---|---|---|
| Layer stacking 0–9 | ✅ declared | `<app-layer level>` per AGENTS.md §4 |
| Real custom elements | ❌ **cosmetic** | a single `defineCE` helper registers ~16 `app-*` tags as **empty classes** (`class extends HTMLElement {}`) — no lifecycle, no shadow DOM, no attributes API |
| Window manager (drag, resize, snap, z-order arbitration) | ❌ **missing** | 5 `drag/pointerdown/touch` occurrences in 1,196 lines of `appz.ts`; no window registry |
| Global keyboard / focus model (hotkeys, escape stack, roving focus) | ⚠️ partial | `htmx-a11y.ts` (300 lines) covers focus trap + matrix nav; no OS-level keymap |
| Multi-app launcher / multitask / lock screen | ⚠️ partial | `<app-launcher>`, `<app-multitask>`, `<app-lockscreen>` exist as tags + Go DSL (`go/components/appz.go`) |
| Persistence of window/app state | ❌ **missing** | no serialization of shell layout |
| Menubar, tray, notifications center | ⚠️ partial | `StatusBar`, `ControlCenter` DSL constructors exist |

### 3.3 Data-streaming dashboard

| Requirement | Status | Evidence |
|---|---|---|
| Wire protocol definition | ✅ excellent | `docs/BACKEND_SPEC.md`: OOB, `hxMatrixUpdate`, `hxStateUpdate`, `HX-Trigger` |
| Client consumes it | ⚠️ **indirect only** | `src/htmx-bolt.ts` listens for `htmx:sseMessage`; **`EventSource` count is 0 across all 16 `src/` files** |
| SSE transport actually loadable | ❌ **missing in-repo** | no view uses `hx-ext="sse"` / `hx-sse`; the htmx SSE extension is not vendored or referenced |
| WebSocket transport | ❌ **same gap** | `htmx:wsAfterMessage` is hooked, the `ws` extension is never loaded |
| Server-side examples | ✅ documented | Go `net/http`, Rust axum, Python FastAPI, .NET in `BACKEND_SPEC.md` |
| Charting / time-series (sparkline, OHLC, axes) | ❌ **missing** | no `sparkline`/`candlestick`/`time-series` engine; `htmx-canvas.ts` is a spatial node diagram |
| Backpressure / delta coalescing / reconnect + resume | ❌ **missing** | no `Last-Event-ID`, no sequence numbers, no throttle |
| 100k-row virtualized grid | ✅ claimed | `htmx-virtual`/`htmx-grid`; numbers in `benchmarks/BENCHMARKS.md` (32 DOM nodes constant) |

---

## 4. Backlog — 34 corrections, improvements, removals

Legend: **P0** blocker for a real app · **P1** high value · **P2** polish.

### A. Corrections (things that are wrong, drifted, or self-contradicting)

1. **[P0] Add `tsconfig.json` and a typecheck step.** No `tsconfig.json` exists, so 9,605 lines of TypeScript are transpiled by `bun build` without ever being type-checked — AGENTS.md §3's "100% Strict TypeScript" is unenforced. Add `strict: true`, `bun run typecheck`, and a CI job.
2. **[P0] Fix the ESM entry point.** `package.json` maps `exports["."].import` → `./src/index.ts` and `types` → `./src/index.ts`. Shipping raw `.ts` as the import target breaks every consumer that does not transpile `node_modules`. Point it at a built ESM file (`dist/index.mjs` + `.d.ts`).
3. **[P0] Single-source the engine list.** `schema/htmxpro.json` declares **11** engines (missing `sim` and `appz`); `package.json` exports **12** (missing `./canvas`); `build:engines` compiles **13**. Three lists, three different answers.
4. **[P0] Fix or retract AGENTS.md §5 (1:1 Go parity).** 54 Go constructors exist for 160 HTML components ≈ **34% parity**. Either generate the Go DSL from `views/components/*.html`, or scope the rule to the implemented subset. A rule that is 66% violated trains agents to ignore rules.
5. **[P1] Stop violating AGENTS.md §3.** Loose `.js` files in root: `worker.js`, `tailwind.config.js`, `postcss.config.js` — the exact thing §3 forbids.
6. **[P1] Reconcile `wrangler.toml` with reality.** `main = "worker.js"` and `worker.ts` is a 9-line passthrough to `env.ASSETS.fetch` — the edge deployment can do nothing but serve static files. Remove the Worker or implement real routing.
7. **[P1] Add the missing `LICENSE` file.** `package.json` claims MIT and lists `"LICENSE"` in `files`; the file does not exist at the root.
8. **[P1] Correct the stale benchmark docs.** `docs/BENCHMARKS.md` reports `htmx-bolt.js` at **85.84 KB**; the actual artifact is **107,609 B** — the doc is ~21 KB stale and its gzip figures are unverifiable against the current build.
9. **[P1] Fix the CI size budget.** It sums only `public/htmx-*.js`, silently excluding `appz.js` (35.8 KB) and any non-`htmx-` prefixed engine.
10. **[P2] Pin the toolchain.** CI uses `bun-version: latest` with no `packageManager` field — builds are not reproducible across days.
11. **[P2] Make the CI actually run AGENTS.md §7.** There is no `go build`, no `bun run build:css`, and no `playwright test` in CI, though the maintainer protocol requires all three.

### B. Removals & replacements

12. **[P0] Remove `express@5.2.1` from `dependencies`.** Repo-wide grep for `from "express"` / `require("express")` returns **zero matches** — dead weight installed for every consumer of the package.
13. **[P1] Delete `package-lock.json`.** It coexists with `bun.lock` in a repo whose first rule is zero npm; two lockfiles drift.
14. **[P1] Untrack `dist/`** (172 files, 3.1 MB of committed build output). Build it in CI; publish to Pages or the Worker.
15. **[P1] Remove the 16 MB of binaries and the 137 MB archive** from the working tree (`htmxpro-server.exe` 8.0 MB, `htmxui-server.exe` 8.0 MB, `sun.zip` 137 MB) and route them through `releases/`.
16. **[P1] Delete `postcss.config.js`, `postcss`, and `autoprefixer`.** Nothing invokes PostCSS — `build:css` calls the Tailwind CLI directly. Dead config plus two dead devDependencies.
17. **[P1] Replace `bunx tailwindcss` with a declared `@tailwindcss/cli` devDependency.** Installed `tailwindcss@4.3.3` has **no `bin` field** (`grep '"bin"'` → exit 1), so the CLI is resolved outside the declared, locked dependency set.
18. **[P1] Delete `tailwind.config.js` or wire it properly.** Tailwind v4 ignores `content:` globs unless `@config` is present in the CSS entry; `public/styles/input.css` has no `@config`, so source detection happens incidentally.
19. **[P1] Replace all 13 CDN `<script src="https://unpkg.com/htmx.org@2.0.4">` tags** with a vendored, SRI-pinned local copy. A third-party CDN on the hot path is both a security and an offline blocker.
20. **[P2] Replace the no-op custom elements** (~16 `app-*` tags registered as `class extends HTMLElement {}`) with either real elements exposing lifecycle/attributes, or documented plain elements — today "custom element" implies behaviour the tags do not have.

### C. Improvements & additions

21. **[P0] Minify the engines.** `build:engines` omits `--minify`; `htmx-bolt.js` ships at 107,609 B / 3,031 lines with full bundler helper preamble. Measure the win, then add `--minify` and per-engine budgets.
22. **[P0] Minify the CSS.** `output.css` ships at 196,417 B / 6,266 lines unminified (`--minify` absent). Add a compressed-byte budget as well, since gzip already reduces it to 22,358 B.
23. **[P0] Compress + cache static responses.** `server.ts` sets no `Content-Encoding`, no `Cache-Control`, no `ETag` (checked lines 458–493). Roughly 300 KB uncompressed is re-fetched on every page view. Add gzip/brotli, `immutable` on content-hashed filenames, and `ETag` on HTML.
24. **[P0] Ship a PWA kit:** `manifest.webmanifest`, a service worker with precache + offline shell, `theme-color`, install-prompt handling, and versioned cache invalidation. None of it exists.
25. **[P0] Make offline real.** `htmx-offline.ts` touches `indexedDB` once and never `caches.*` or `serviceWorker` — mutate queues exist, but assets and navigation do not.
26. **[P0] Ship the SSE transport and a live demo.** The wire protocol is documented but unreachable: zero `EventSource` in `src/`, no `hx-ext="sse"` in any view. Add the extension, a vendored copy, and one streaming dashboard page that proves it.
27. **[P1] Ship the WebSocket transport**, same as above — `htmx:wsAfterMessage` is already hooked.
28. **[P1] Add a time-series/chart engine** (sparkline, line, OHLC, axis ticks) — there is no charting primitive, which makes "data-streaming dashboard" aspirational.
29. **[P1] Add stream resilience:** sequence numbers, `Last-Event-ID` resume, delta coalescing, and a reconnect/backoff policy on the client.
30. **[P1] Add a Python integration for real:** a Jinja2 macro library mirroring the 160 components, a FastAPI/ASGI adapter, and a distributable package. Today Python is prose only, while Go gets 54 constructors and a linked library — the parity is asymmetric.
31. **[P1] Add `go:embed` + a `htmxpro` CLI** (`new`, `dev`, `build`, `doctor`) so a Go backend really is one binary, and so the parity/schema/checksum rules become automated checks instead of prose.
32. **[P1] Add markup-level tests for the 160 components.** Only one test file references `views/components`; engines are well tested and the actual product surface is not. Add render snapshots + axe-core a11y assertions (no axe anywhere today).
33. **[P1] Add an auth/session/CSRF reference implementation** for Go and Python. Neither server sets a cookie or a CSRF token; sessions, RBAC enforcement, and headers are entirely absent from the shipped servers.
34. **[P1] Formalize the wire protocol as a machine-readable contract** (JSON Schema + contract tests for `hxMatrixUpdate`, `hxStateUpdate`, OOB, `HX-Trigger`) and add i18n/RTL support (no i18n framework, no `dir="rtl"`, no logical-property audit).

---

## 5. What I could not verify

1. **Runtime truths.** I did not start the servers, so no latency, throughput, or memory claim here is measured. The 100k-row / 60 FPS / 32-node figures come from the repo's own docs and are unverified.
2. **The token-efficiency claim.** "HTMX PRO is more token-efficient for AI" is plausible (fewer degrees of freedom, finite vocabulary) but untested. The metric that would settle it is **tokens-to-first-correct-change including retries**, not tokens per line of source. The retry tax is the real risk: React is heavily overrepresented in model training data, so an agent's first attempt on `hx-*` attributes is more likely to be wrong.
3. **Tailwind's token cost.** Long utility class strings may consume the savings that dropping component boilerplate produces. This is measurable and should be measured before the claim is published.
4. **Two flags worth re-checking:** whether `<app-thermostat>` (used in views) is in the `defineCE` list, and whether the `COMPONENTS` array in `server.ts` matches the 160 files on disk.

---

## 6. Suggested order of attack

1. **Trust & safety first (1 day):** items 1, 2, 3, 12, 13, 7 — typecheck, correct packaging, kill drift and dead deps.
2. **Performance truth (1 day):** items 21, 22, 23, then re-measure and update `docs/BENCHMARKS.md` (item 8).
3. **Prove the platform (1 week):** items 24, 25, 26, 28 — a PWA that works offline and a streaming dashboard that actually consumes SSE are the two demos this project must ship to be more than a library.
4. **Close the parity gap (1 week+):** items 4, 30, 31, 32 — generate the DSL, add Python, add the CLI, test the markup.
