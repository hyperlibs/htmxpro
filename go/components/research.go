package components

import (
	"fmt"
)

// =============================================================================
// RESEARCH DESIGN BLOCK 1: PATHOLOGY LAB SIMULATOR
// =============================================================================

type BlockPathologyLabProps struct {
	SpecimenID   string
	Stain        string
	Diagnosis    string
	Mag          string // "4x" | "10x" | "20x" | "40x" | "100x"
	ZStack       float64
	Ki67Pct      float64
	MitoticCount int
}

func BlockPathologyLab(p BlockPathologyLabProps) Component {
	if p.SpecimenID == "" {
		p.SpecimenID = "#BIO-84920"
	}
	if p.Diagnosis == "" {
		p.Diagnosis = "Invasive Ductal Carcinoma (FFPE Core Biopsy)"
	}
	if p.Mag == "" {
		p.Mag = "40x"
	}
	if p.ZStack == 0 {
		p.ZStack = 1.2
	}
	if p.Ki67Pct == 0 {
		p.Ki67Pct = 34.2
	}
	if p.MitoticCount == 0 {
		p.MitoticCount = 18
	}

	html := fmt.Sprintf(`
<div class="block-pathology-lab rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "mag": "%s", "zStack": %.1f, "stainGain": 100, "activeTool": "lasso", "ki67": %.1f, "mitoticCount": %d, "showMasks": true }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-violet-600/20 border border-violet-500/40 flex items-center justify-center text-violet-400 font-black text-sm">
        🔬
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-violet-400">Pathology WSI Scanner</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">CALIBRATED</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>Specimen %s</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">%s</span>
        </h2>
      </div>
    </div>

    <!-- Magnification Objective Turret -->
    <div class="flex items-center gap-1 bg-slate-950 p-1 rounded-2xl border border-slate-800 text-xs font-black">
      <button class="px-2.5 py-1 rounded-xl transition-all"
              hx-class='{ "bg-violet-600 text-white shadow-md shadow-violet-600/30": mag === "4x", "text-slate-400 hover:text-white": mag !== "4x" }'
              hx-action="mag = '4x'; $appz.haptic('selection')">4x</button>
      <button class="px-2.5 py-1 rounded-xl transition-all"
              hx-class='{ "bg-violet-600 text-white shadow-md shadow-violet-600/30": mag === "10x", "text-slate-400 hover:text-white": mag !== "10x" }'
              hx-action="mag = '10x'; $appz.haptic('selection')">10x</button>
      <button class="px-2.5 py-1 rounded-xl transition-all"
              hx-class='{ "bg-violet-600 text-white shadow-md shadow-violet-600/30": mag === "20x", "text-slate-400 hover:text-white": mag !== "20x" }'
              hx-action="mag = '20x'; $appz.haptic('selection')">20x</button>
      <button class="px-2.5 py-1 rounded-xl transition-all"
              hx-class='{ "bg-violet-600 text-white shadow-md shadow-violet-600/30": mag === "40x", "text-slate-400 hover:text-white": mag !== "40x" }'
              hx-action="mag = '40x'; $appz.haptic('selection')">40x</button>
      <button class="px-2.5 py-1 rounded-xl transition-all"
              hx-class='{ "bg-amber-600 text-white shadow-md shadow-amber-600/30": mag === "100x", "text-slate-400 hover:text-white": mag !== "100x" }'
              hx-action="mag = '100x'; $appz.haptic('heavy')">100x Oil</button>
    </div>
  </div>

  <!-- ZONE 2: CALIBRATION & TOOLBAR -->
  <div class="px-5 py-2.5 bg-slate-900/40 border-b border-slate-800/80 flex flex-wrap items-center justify-between gap-3 text-xs">
    <div class="flex items-center gap-1.5">
      <span class="text-[10px] font-black uppercase tracking-wider text-slate-500 mr-1">Tools</span>
      <button class="px-2.5 py-1 rounded-xl font-bold flex items-center gap-1.5 transition-all"
              hx-class='{ "bg-violet-600 text-white": activeTool === "pan", "bg-slate-900 text-slate-400 hover:text-white": activeTool !== "pan" }'
              hx-action="activeTool = 'pan'; $appz.haptic('selection')">
        <span>✋ Pan</span>
      </button>
      <button class="px-2.5 py-1 rounded-xl font-bold flex items-center gap-1.5 transition-all"
              hx-class='{ "bg-violet-600 text-white": activeTool === "lasso", "bg-slate-900 text-slate-400 hover:text-white": activeTool !== "lasso" }'
              hx-action="activeTool = 'lasso'; $appz.haptic('selection')">
        <span>✏️ Tumor Lasso</span>
      </button>
      <button class="px-2.5 py-1 rounded-xl font-bold flex items-center gap-1.5 transition-all"
              hx-class='{ "bg-violet-600 text-white": activeTool === "ruler", "bg-slate-900 text-slate-400 hover:text-white": activeTool !== "ruler" }'
              hx-action="activeTool = 'ruler'; $appz.haptic('selection')">
        <span>📏 µm Caliper</span>
      </button>
      <button class="px-2.5 py-1 rounded-xl font-bold flex items-center gap-1.5 transition-all"
              hx-class='{ "bg-emerald-600 text-white": showMasks, "bg-slate-900 text-slate-400": !showMasks }'
              hx-action="showMasks = !showMasks; $appz.haptic('selection')">
        <span hx-text="showMasks ? '👁️ Masks: ON' : '👁️ Masks: OFF'">👁️ Masks: ON</span>
      </button>
    </div>

    <div class="flex items-center gap-4">
      <div class="flex items-center gap-2">
        <span class="text-[10px] font-bold text-slate-400 uppercase">Z-Stack</span>
        <button class="w-6 h-6 rounded-lg bg-slate-800 text-white font-bold flex items-center justify-center text-xs"
                hx-action="zStack = +(zStack - 0.2).toFixed(1); $appz.haptic('selection')">-</button>
        <span class="font-mono font-bold text-violet-400 text-xs w-12 text-center" hx-text="zStack + ' µm'">%.1f µm</span>
        <button class="w-6 h-6 rounded-lg bg-slate-800 text-white font-bold flex items-center justify-center text-xs"
                hx-action="zStack = +(zStack + 0.2).toFixed(1); $appz.haptic('selection')">+</button>
      </div>

      <div class="flex items-center gap-2">
        <span class="text-[10px] font-bold text-slate-400 uppercase">H&E Gain</span>
        <input type="range" min="50" max="150" value="100" class="w-20 accent-violet-500 cursor-pointer"
               hx-action="$appz.haptic('selection')" />
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: STAGE CANVAS & QUANTITATIVE INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[360px]">
    <div class="lg:col-span-8 relative bg-slate-950 p-4 flex items-center justify-center overflow-hidden border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="relative w-full max-w-[440px] aspect-square rounded-full border-4 border-slate-800 bg-[#160d1f] shadow-inner overflow-hidden flex items-center justify-center">
        <!-- SVG Microscopic Field -->
        <svg class="absolute inset-0 w-full h-full opacity-90" viewBox="0 0 400 400">
          <defs>
            <radialGradient id="heGradGo" cx="50%" cy="50%" r="50%">
              <stop offset="0%" stop-color="#4a154b" stop-opacity="0.85" />
              <stop offset="70%" stop-color="#2a0d33" stop-opacity="0.95" />
              <stop offset="100%" stop-color="#120517" stop-opacity="1" />
            </radialGradient>
          </defs>
          <rect width="400" height="400" fill="url(#heGradGo)" />
          <circle cx="110" cy="90" r="9" fill="#7c2d8a" opacity="0.7" />
          <circle cx="220" cy="180" r="14" fill="#9d174d" opacity="0.8" />
          <circle cx="250" cy="210" r="16" fill="#831843" opacity="0.8" />
        </svg>

        <svg class="absolute inset-0 w-full h-full pointer-events-none" viewBox="0 0 400 400" hx-show="showMasks">
          <ellipse cx="220" cy="180" rx="14" ry="11" fill="rgba(244, 63, 94, 0.25)" stroke="#f43f5e" stroke-width="2" />
          <ellipse cx="250" cy="210" rx="16" ry="13" fill="rgba(244, 63, 94, 0.25)" stroke="#f43f5e" stroke-width="2" />
          <circle cx="235" cy="150" r="5" fill="#f59e0b" stroke="#ffffff" stroke-width="1.5" />
          <text x="245" y="154" fill="#f59e0b" font-size="9" font-weight="900">MITOSIS</text>
        </svg>

        <div class="absolute inset-0 flex items-center justify-center pointer-events-none opacity-40">
          <div class="w-full h-[1px] bg-violet-400"></div>
          <div class="h-full w-[1px] bg-violet-400 absolute"></div>
          <div class="w-20 h-20 rounded-full border border-violet-400 absolute"></div>
        </div>

        <div class="absolute bottom-4 left-6 bg-slate-900/80 backdrop-blur px-2.5 py-1 rounded-xl border border-slate-800 text-[10px] font-bold text-slate-300">
          Obj: <span class="text-violet-400 font-black" hx-text="mag">%s</span> • N.A. 1.40
        </div>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div>
        <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
          <span class="text-[10px] font-black uppercase tracking-widest text-violet-400">Histologic Quantitative Telemetry</span>
          <span class="text-[10px] font-bold text-slate-400">CAP / ASCO</span>
        </div>

        <div class="mt-4 space-y-3">
          <div class="p-3 rounded-2xl bg-slate-950/70 border border-slate-800">
            <div class="flex items-center justify-between">
              <span class="text-xs font-bold text-slate-300">Ki-67 Proliferation Index</span>
              <span class="text-xs font-black text-rose-400">HIGH</span>
            </div>
            <div class="flex items-baseline gap-2 mt-1">
              <span class="text-3xl font-black text-white" hx-text="ki67 + '%%'">%.1f%%</span>
              <span class="text-[10px] text-slate-400">pos. nuclei / 1000 cells</span>
            </div>
            <div class="w-full h-1.5 rounded-full bg-slate-800 mt-2 overflow-hidden">
              <div class="h-full bg-rose-500 rounded-full" style="width: %.1f%%;"></div>
            </div>
          </div>

          <div class="p-3 rounded-2xl bg-slate-950/70 border border-slate-800">
            <span class="text-[10px] font-extrabold uppercase tracking-widest text-slate-400 block mb-1">Nottingham Histologic Score</span>
            <div class="mt-1 text-right">
              <span class="text-xs font-black text-rose-400 bg-rose-500/10 px-2 py-0.5 rounded-md border border-rose-500/20">GRADE III (Poorly Differentiated)</span>
            </div>
          </div>
        </div>
      </div>

      <div class="pt-3 border-t border-slate-800/80">
        <span class="text-[10px] font-extrabold uppercase text-slate-400 block mb-2">Mitotic Figures Counter</span>
        <div class="flex items-center gap-2">
          <button class="flex-1 py-2 rounded-xl bg-slate-800 text-white font-black text-xs"
                  hx-action="mitoticCount = Math.max(0, mitoticCount - 1); $appz.haptic('selection')">- Decrement</button>
          <div class="px-3 py-1.5 rounded-xl bg-slate-950 border border-slate-800 font-mono font-black text-lg text-white"
               hx-text="mitoticCount">%d</div>
          <button class="flex-1 py-2 rounded-xl bg-violet-600 text-white font-black text-xs"
                  hx-action="mitoticCount += 1; $appz.haptic('medium')">+ Tally Mitosis</button>
        </div>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <div class="flex items-center gap-3">
      <div class="flex flex-col items-center">
        <div class="w-20 h-1 bg-white rounded-full"></div>
        <span class="text-[9px] font-mono font-bold text-white mt-0.5">50 µm</span>
      </div>
      <span class="text-slate-600">|</span>
      <span class="text-[10px] font-mono">X: 142.50 µm • Y: 89.12 µm • Z: +%.2f µm</span>
    </div>
    <div class="flex items-center gap-4 text-[10px] font-bold">
      <span class="text-emerald-400">● WSI Sensor Sync: 60.0 FPS</span>
      <span class="text-violet-400">Plan-Apo 1.40 Oil</span>
    </div>
  </div>
</div>`, p.Mag, p.ZStack, p.Ki67Pct, p.MitoticCount, p.SpecimenID, p.Diagnosis, p.ZStack, p.Mag, p.Ki67Pct, p.Ki67Pct, p.MitoticCount, p.ZStack)
	return Component{html: html}
}

// =============================================================================
// RESEARCH DESIGN BLOCK 2: RADIOLOGY DICOM LAB SIMULATOR
// =============================================================================

type BlockRadiologyDICOMProps struct {
	AccessionNo  string
	Modality     string
	Slice        int
	MaxSlice     int
	WindowPreset string
	WindowW      int
	WindowL      int
	ROIMean      float64
	ROIArea      float64
}

func BlockRadiologyDICOM(p BlockRadiologyDICOMProps) Component {
	if p.AccessionNo == "" {
		p.AccessionNo = "#RAD-2026-X99"
	}
	if p.Modality == "" {
		p.Modality = "Chest CT Angiography (0.625mm Thin Slices)"
	}
	if p.Slice == 0 {
		p.Slice = 142
		p.MaxSlice = 320
	}
	if p.WindowW == 0 {
		p.WindowW = 400
		p.WindowL = 40
		p.WindowPreset = "soft"
	}
	if p.ROIMean == 0 {
		p.ROIMean = 42.5
		p.ROIArea = 3.42
	}

	html := fmt.Sprintf(`
<div class="block-radiology-dicom rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "slice": %d, "maxSlice": %d, "windowPreset": "%s", "windowW": %d, "windowL": %d, "activeTool": "roi", "roiMean": %.1f, "roiArea": %.2f, "inverted": false }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-cyan-600/20 border border-cyan-500/40 flex items-center justify-center text-cyan-400 font-black text-sm">
        ☢️
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-cyan-400">Radiology PACS DICOM Studio</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-cyan-500/10 text-cyan-400 border border-cyan-500/20">MPR 16-BIT</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>Accession %s</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">%s</span>
        </h2>
      </div>
    </div>

    <div class="flex items-center gap-2 bg-slate-950 px-3 py-1.5 rounded-2xl border border-slate-800 text-xs">
      <span class="text-slate-400 font-bold">120 kVp</span>
      <span class="text-slate-600">•</span>
      <span class="text-cyan-400 font-mono font-black">CTDIvol: 3.2 mGy</span>
    </div>
  </div>

  <!-- ZONE 2: CALIBRATION & TOOLBAR -->
  <div class="px-5 py-2.5 bg-slate-900/40 border-b border-slate-800/80 flex flex-wrap items-center justify-between gap-3 text-xs">
    <div class="flex items-center gap-1.5">
      <span class="text-[10px] font-black uppercase tracking-wider text-slate-500 mr-1">HU Window</span>
      <button class="px-2.5 py-1 rounded-xl font-bold transition-all"
              hx-class='{ "bg-cyan-600 text-white shadow-md shadow-cyan-600/30": windowPreset === "lung", "bg-slate-900 text-slate-400": windowPreset !== "lung" }'
              hx-action="windowPreset = 'lung'; windowW = 1500; windowL = -600; $appz.haptic('selection')">
        Lung (1500/-600)
      </button>
      <button class="px-2.5 py-1 rounded-xl font-bold transition-all"
              hx-class='{ "bg-cyan-600 text-white shadow-md shadow-cyan-600/30": windowPreset === "soft", "bg-slate-900 text-slate-400": windowPreset !== "soft" }'
              hx-action="windowPreset = 'soft'; windowW = 400; windowL = 40; $appz.haptic('selection')">
        Soft Tissue (400/40)
      </button>
      <button class="px-2.5 py-1 rounded-xl font-bold transition-all"
              hx-class='{ "bg-cyan-600 text-white shadow-md shadow-cyan-600/30": windowPreset === "bone", "bg-slate-900 text-slate-400": windowPreset !== "bone" }'
              hx-action="windowPreset = 'bone'; windowW = 2000; windowL = 300; $appz.haptic('selection')">
        Bone (2000/300)
      </button>
    </div>
  </div>

  <!-- ZONE 3 & 4: MULTI-PLANAR VIEWPORT & INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-black p-4 flex flex-col justify-between overflow-hidden border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="flex items-start justify-between text-[11px] font-mono text-cyan-400/90 pointer-events-none z-10">
        <div>
          <p class="font-black text-white">PATIENT: DOE, JOHN</p>
          <p>MRN: 9028-3419 • THICK: 0.625 mm</p>
        </div>
        <div class="text-right">
          <p>W: <span hx-text="windowW">%d</span> L: <span hx-text="windowL">%d</span></p>
        </div>
      </div>

      <div class="relative w-full max-w-[420px] aspect-square mx-auto flex items-center justify-center">
        <svg class="w-full h-full" viewBox="0 0 400 400">
          <ellipse cx="200" cy="200" rx="170" ry="140" fill="#18202c" stroke="#334155" stroke-width="1.5" />
          <circle cx="200" cy="305" r="24" fill="#cbd5e1" stroke="#ffffff" stroke-width="1" />
          <ellipse cx="200" cy="200" rx="160" ry="130" fill="none" stroke="#94a3b8" stroke-width="3" stroke-dasharray="14 10" />
          <circle cx="135" cy="185" r="14" fill="#475569" stroke="#f43f5e" stroke-width="2" />
        </svg>
      </div>

      <div class="mt-3 flex items-center gap-3 bg-slate-900/80 px-4 py-2 rounded-2xl border border-slate-800">
        <span class="text-xs font-bold text-slate-400">Slice:</span>
        <span class="font-mono font-black text-cyan-400 text-xs w-16" hx-text="slice + ' / ' + maxSlice">%d / %d</span>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div>
        <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
          <span class="text-[10px] font-black uppercase tracking-widest text-cyan-400">Hounsfield Unit Morphometry</span>
          <span class="text-[10px] font-bold text-emerald-400">RECIST 1.1</span>
        </div>

        <div class="mt-4 space-y-3">
          <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
            <span class="text-[10px] font-extrabold uppercase text-slate-400 block mb-1">Target Lesion 1 (LUL)</span>
            <div class="flex items-baseline gap-2">
              <span class="text-3xl font-black text-white" hx-text="roiMean + ' HU'">+%.1f HU</span>
            </div>
            <div class="mt-2 text-right">
              <span class="text-[10px] font-black text-rose-400 bg-rose-500/10 px-2 py-0.5 rounded-md border border-rose-500/20">PROGRESSIVE DISEASE (+29%%)</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-cyan-400">Matrix: 512x512 16-bit</span>
    <span class="text-emerald-400">● 60.0 FPS • Zero-Latency LUT</span>
  </div>
</div>`, p.Slice, p.MaxSlice, p.WindowPreset, p.WindowW, p.WindowL, p.ROIMean, p.ROIArea, p.AccessionNo, p.Modality, p.WindowW, p.WindowL, p.Slice, p.MaxSlice, p.ROIMean)
	return Component{html: html}
}

// =============================================================================
// RESEARCH DESIGN BLOCK 3: HUMAN & BIOLOGY CELLULAR ATLAS UI
// =============================================================================

type BlockBiologyAtlasProps struct {
	Organ        string
	CellCount    string
	ActiveGene   string
	ActiveCluster string
	ClusterPct   float64
}

func BlockBiologyAtlas(p BlockBiologyAtlasProps) Component {
	if p.Organ == "" {
		p.Organ = "Left Ventricle Myocardium"
	}
	if p.CellCount == "" {
		p.CellCount = "142,500 Single Cells (10x Chromium 3' v3)"
	}
	if p.ActiveGene == "" {
		p.ActiveGene = "TNNT2"
	}
	if p.ActiveCluster == "" {
		p.ActiveCluster = "Cardiomyocytes"
		p.ClusterPct = 48.2
	}

	html := fmt.Sprintf(`
<div class="block-biology-atlas rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "viewMode": "spatial", "activeGene": "%s", "umiCutoff": 1500, "activeCluster": "%s", "clusterPct": %.1f }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-emerald-600/20 border border-emerald-500/40 flex items-center justify-center text-emerald-400 font-black text-sm">
        🧬
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-emerald-400">Human Biology Cellular Atlas</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">HCA v3.2</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>Organ: %s</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">%s</span>
        </h2>
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: SPATIAL CANVAS & QUANTITATIVE INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-slate-950 p-4 flex flex-col justify-between overflow-hidden border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="flex items-center justify-between text-xs font-bold text-slate-400 z-10 pointer-events-none">
        <span class="flex items-center gap-2">
          <span class="w-2.5 h-2.5 rounded-full bg-emerald-400"></span>
          Expression Gradient (<span class="font-mono text-white" hx-text="activeGene">%s</span>)
        </span>
      </div>

      <div class="relative w-full max-w-[460px] aspect-[4/3] mx-auto my-auto flex items-center justify-center">
        <svg class="w-full h-full" viewBox="0 0 500 350">
          <path d="M 120,70 C 220,30 380,50 420,150 C 450,220 360,310 240,320 C 140,330 70,220 80,140 C 90,90 100,75 120,70 Z"
                fill="#05140f" stroke="#10b981" stroke-width="1.5" stroke-dasharray="4 2" />
          <circle cx="200" cy="120" r="7" fill="#fde047" opacity="0.9" />
          <circle cx="225" cy="115" r="7" fill="#fde047" opacity="0.9" />
          <circle cx="250" cy="130" r="7" fill="#fde047" opacity="0.9" />
        </svg>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div>
        <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
          <span class="text-[10px] font-black uppercase tracking-widest text-emerald-400">Cellular Composition</span>
          <span class="text-[10px] font-bold text-slate-400">Single-Cell scRNA</span>
        </div>

        <div class="mt-4 space-y-3">
          <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
            <span class="text-[10px] font-extrabold uppercase text-slate-400 block mb-1">Selected Cell Population</span>
            <div class="flex items-baseline justify-between">
              <h3 class="text-base font-black text-white" hx-text="activeCluster">%s</h3>
              <span class="text-xl font-black text-emerald-400" hx-text="clusterPct + '%%'">%.1f%%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-emerald-400">Spatial Spot Pitch: 55 µm</span>
    <span class="text-emerald-400 font-mono">@diag: ATLAS_ACTIVE</span>
  </div>
</div>`, p.ActiveGene, p.ActiveCluster, p.ClusterPct, p.Organ, p.CellCount, p.ActiveGene, p.ActiveCluster, p.ClusterPct)
	return Component{html: html}
}

// =============================================================================
// RESEARCH DESIGN BLOCK 4: CRISPR GENE EDITING STUDIO
// =============================================================================

type BlockCrisprStudioProps struct {
	TargetGene     string
	TargetLocus    string
	EditMode       string // "nhej" | "hdr" | "base" | "prime"
	OnTargetScore  int
	OffTargetScore int
	GCContent      float64
}

func BlockCrisprStudio(p BlockCrisprStudioProps) Component {
	if p.TargetGene == "" {
		p.TargetGene = "HBB"
		p.TargetLocus = "chr11:5,227,002"
		p.EditMode = "hdr"
		p.OnTargetScore = 88
		p.OffTargetScore = 94
		p.GCContent = 55.0
	}

	html := fmt.Sprintf(`
<div class="block-crispr-studio rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "targetGene": "%s", "targetLocus": "%s", "editMode": "%s", "onTargetScore": %d, "offTargetScore": %d, "gcContent": %.1f, "activePam": "TGG" }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-purple-600/20 border border-purple-500/40 flex items-center justify-center text-purple-400 font-black text-sm">
        ✂️
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-purple-400">CRISPR-Cas Precision Studio</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-purple-500/10 text-purple-400 border border-purple-500/20">SpCas9 (NGG)</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span hx-text="'Target Gene: ' + targetGene + ' (Exon 1)'">Target Gene: %s (Exon 1)</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">HbS E6V Sickle Mutation (<span class="font-mono text-purple-300" hx-text="targetLocus">%s</span>)</span>
        </h2>
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: DNA CLEAVAGE CANVAS & QUANTITATIVE INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-slate-950 p-6 flex flex-col justify-between overflow-x-auto border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="space-y-4">
        <div class="ml-16 mr-20 bg-purple-950/40 border border-purple-500/40 rounded-xl p-2 text-center">
          <span class="text-[10px] font-extrabold uppercase tracking-widest text-purple-300">20 nt Synthetic Single Guide RNA (sgRNA Protospacer)</span>
        </div>
        <div class="flex items-center gap-1 font-mono font-black text-sm">
          <span class="w-12 text-[10px] font-bold text-slate-500 text-right mr-1">5'—</span>
          <div class="flex items-center gap-1 bg-slate-900/90 p-2 rounded-2xl border border-slate-800 shadow-inner">
            <span class="px-1.5 py-1 rounded bg-slate-950 text-slate-400">A</span>
            <span class="px-1.5 py-1 rounded bg-slate-950 text-slate-400">T</span>
            <span class="px-1.5 py-1 rounded bg-slate-950 text-slate-400">G</span>
            <div class="flex items-center gap-1 px-1.5 py-1 rounded-xl bg-cyan-950/40 border border-cyan-500/50">
              <span class="px-1 py-0.5 rounded bg-cyan-500/20 text-cyan-300">G</span>
              <span class="px-1 py-0.5 rounded bg-cyan-500/20 text-cyan-300">A</span>
              <span class="px-1 py-0.5 rounded bg-cyan-500/20 text-cyan-300">C</span>
            </div>
            <div class="flex items-center gap-1 px-1.5 py-1 rounded-xl bg-amber-950/50 border border-amber-500">
              <span class="px-1 py-0.5 rounded bg-amber-500/30 text-amber-300 font-black">T</span>
              <span class="px-1 py-0.5 rounded bg-amber-500/30 text-amber-300 font-black">G</span>
              <span class="px-1 py-0.5 rounded bg-amber-500/30 text-amber-300 font-black">G</span>
            </div>
          </div>
          <span class="text-[10px] font-bold text-slate-500 ml-1">—3'</span>
        </div>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div>
        <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
          <span class="text-[10px] font-black uppercase tracking-widest text-purple-400">Cutting & Specificity Metrics</span>
        </div>

        <div class="mt-4 space-y-3">
          <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
            <div class="flex items-center justify-between">
              <span class="text-xs font-bold text-slate-300">On-Target Activity Score</span>
              <span class="text-xs font-black text-emerald-400">EXCELLENT</span>
            </div>
            <div class="flex items-baseline gap-2 mt-1">
              <span class="text-3xl font-black text-white" hx-text="onTargetScore + ' / 100'">%d / 100</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-purple-400">Oligo: 5'-CACCGAGTCCTGAGGAGAAGTCTG-3'</span>
    <span class="text-purple-400 font-mono">@diag: CAS9_CALC_OK</span>
  </div>
</div>`, p.TargetGene, p.TargetLocus, p.EditMode, p.OnTargetScore, p.OffTargetScore, p.GCContent, p.TargetGene, p.TargetLocus, p.OnTargetScore)
	return Component{html: html}
}

// =============================================================================
// RESEARCH DESIGN BLOCK 5: GENOMIC MAP & SYNTENY DASHBOARD
// =============================================================================

type BlockGenomicMapProps struct {
	Chr             string
	Start           int
	End             int
	SelectedVariant string
}

func BlockGenomicMap(p BlockGenomicMapProps) Component {
	if p.Chr == "" {
		p.Chr = "chr11"
		p.Start = 5224000
		p.End = 5230000
		p.SelectedVariant = "rs334"
	}

	html := fmt.Sprintf(`
<div class="block-genomic-map rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "chr": "%s", "start": %d, "end": %d, "zoom": 1, "selectedVariant": "%s", "showCoverage": true, "showVariants": true }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-amber-600/20 border border-amber-500/40 flex items-center justify-center text-amber-400 font-black text-sm">
        🗺️
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-amber-400">High-Resolution Genome Browser</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-amber-500/10 text-amber-400 border border-amber-500/20">GRCh38.p14</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span class="font-mono text-amber-300">%s:%d–%d</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">Cytoband 11p15.4 (HBB Cluster)</span>
        </h2>
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: MULTI-TRACK BROWSER & QUANTITATIVE INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-slate-950 p-5 flex flex-col justify-between overflow-x-auto border-b lg:border-b-0 lg:border-r border-slate-800/80 space-y-4">
      <div class="space-y-1.5 p-3 rounded-2xl bg-slate-900/60 border border-slate-800/80">
        <div class="flex justify-between text-[10px] font-extrabold text-slate-400 uppercase">
          <span>GENCODE v46 Gene Model</span>
          <span class="font-mono text-white">HBB</span>
        </div>
        <div class="relative h-10 w-full flex items-center">
          <div class="w-full h-0.5 bg-slate-700"></div>
          <div class="absolute left-[15%%] w-16 h-7 rounded-lg bg-amber-500 text-slate-950 font-black text-xs flex items-center justify-center">Exon 1</div>
          <div class="absolute left-[45%%] w-24 h-7 rounded-lg bg-amber-500 text-slate-950 font-black text-xs flex items-center justify-center">Exon 2</div>
          <div class="absolute left-[80%%] w-20 h-7 rounded-lg bg-amber-500 text-slate-950 font-black text-xs flex items-center justify-center">Exon 3</div>
        </div>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
        <span class="text-[10px] font-extrabold uppercase text-slate-400 block mb-1">Target Allele</span>
        <h3 class="text-xl font-black text-rose-400 font-mono">%s (A>T)</h3>
        <p class="text-xs text-white font-bold mt-1">p.Glu7Val (Sickle Cell Hemoglobin HbS)</p>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-amber-400">Assembly: GRCh38.p14</span>
    <span class="text-amber-400 font-mono">@diag: GENOME_SYNC_OK</span>
  </div>
</div>`, p.Chr, p.Start, p.End, p.SelectedVariant, p.Chr, p.Start, p.End, p.SelectedVariant)
	return Component{html: html}
}
