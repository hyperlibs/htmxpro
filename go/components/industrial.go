package components

import (
	"fmt"
)

// =============================================================================
// INDUSTRIAL DESIGN BLOCK 1: CIVIL, ARCHITECTURAL & LANDSCAPING STUDIO
// =============================================================================

type BlockCivilArchitectProps struct {
	ProjectID     string
	StructuralSys string
	Scale         string
	TotalArea     float64
	MaxDeflection float64
}

func BlockCivilArchitectStudio(p BlockCivilArchitectProps) Component {
	if p.ProjectID == "" {
		p.ProjectID = "Parkview Tower (Level 04)"
	}
	if p.StructuralSys == "" {
		p.StructuralSys = "Reinforced Post-Tensioned Concrete (C35/45)"
	}
	if p.Scale == "" {
		p.Scale = "1:100"
	}
	if p.TotalArea == 0 {
		p.TotalArea = 842.5
	}
	if p.MaxDeflection == 0 {
		p.MaxDeflection = 4.2
	}

	html := fmt.Sprintf(`
<div class="block-civil-architect rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "activeLayer": "structural", "scale": "%s", "showContours": true, "showLux": false, "totalArea": %.1f, "maxDeflection": %.1f }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-orange-600/20 border border-orange-500/40 flex items-center justify-center text-orange-400 font-black text-sm">
        📐
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-orange-400">Civil & Architectural BIM Studio</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-orange-500/10 text-orange-400 border border-orange-500/20">ISO 19650</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>Project: %s</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">%s</span>
        </h2>
      </div>
    </div>
    <div class="flex items-center gap-2 bg-slate-950 px-3 py-1.5 rounded-2xl border border-slate-800 text-xs">
      <span class="text-slate-400 font-bold">Scale: <span class="text-white font-mono" hx-text="scale">%s</span></span>
      <span class="text-slate-600">•</span>
      <span class="text-orange-400 font-mono font-bold">Axis A1–F8 (UTM-32N)</span>
    </div>
  </div>

  <!-- ZONE 3 & 4: CAD VIEWPORT & STRUCTURAL INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-slate-950 p-5 flex flex-col justify-between overflow-x-auto border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="flex items-center justify-between text-xs font-mono text-orange-400/90 z-10 pointer-events-none">
        <span>LEVEL 04 PLAN • RL +14.200m</span>
        <span>LOAD: 3.0 kN/m² LIVE • 4.5 kN/m² DEAD</span>
      </div>
      <div class="relative w-full max-w-[460px] aspect-[4/3] mx-auto my-auto flex items-center justify-center">
        <svg class="w-full h-full" viewBox="0 0 500 350">
          <rect x="70" y="40" width="360" height="200" rx="8" fill="rgba(30, 41, 59, 0.4)" stroke="#94a3b8" stroke-width="2.5" />
          <rect x="80" y="50" width="16" height="16" fill="#f97316" stroke="#ea580c" />
          <rect x="242" y="50" width="16" height="16" fill="#f97316" stroke="#ea580c" />
          <rect x="404" y="50" width="16" height="16" fill="#f97316" stroke="#ea580c" />
          <line x1="80" y1="140" x2="242" y2="140" stroke="#cbd5e1" stroke-width="2" />
          <circle cx="140" cy="285" r="14" fill="#15803d" opacity="0.7" />
          <circle cx="200" cy="295" r="18" fill="#15803d" opacity="0.7" />
        </svg>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div>
        <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
          <span class="text-[10px] font-black uppercase tracking-widest text-orange-400">Structural & Spatial Metrics</span>
          <span class="text-[10px] font-bold text-emerald-400">Eurocode 2 / IBC</span>
        </div>
        <div class="mt-4 space-y-3">
          <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
            <span class="text-[10px] font-extrabold uppercase text-slate-400 block mb-1">Gross Internal Area (GIA)</span>
            <div class="flex items-baseline justify-between">
              <span class="text-3xl font-black text-white" hx-text="totalArea + ' m²'">%.1f m²</span>
              <span class="text-xs font-bold text-orange-400">86.4%% Efficiency</span>
            </div>
          </div>
          <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
            <span class="text-xs font-bold text-slate-300">Max Mid-Span Deflection</span>
            <span class="text-2xl font-black text-white block mt-1" hx-text="maxDeflection + ' mm'">%.1f mm</span>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-orange-400">BIM Standard: ISO 19650 Level 2</span>
    <span class="text-orange-400 font-mono">@diag: BIM_ENGINE_SYNC</span>
  </div>
</div>`, p.Scale, p.TotalArea, p.MaxDeflection, p.ProjectID, p.StructuralSys, p.Scale, p.TotalArea, p.MaxDeflection)
	return Component{html: html}
}

// =============================================================================
// INDUSTRIAL DESIGN BLOCK 2: MARINE, ECDIS & GPS NAUTICAL NAVIGATION
// =============================================================================

type BlockMarineNavigationProps struct {
	VesselName string
	MMSI       string
	RangeNM    string
	SOGKnots   float64
	COGTrue    int
	DepthM     float64
	CPANM      float64
}

func BlockMarineNavigation(p BlockMarineNavigationProps) Component {
	if p.VesselName == "" {
		p.VesselName = "R/V NEPTUNE EXPLORER"
		p.MMSI = "368492010"
	}
	if p.RangeNM == "" {
		p.RangeNM = "6.0"
	}
	if p.SOGKnots == 0 {
		p.SOGKnots = 14.2
		p.COGTrue = 248
		p.DepthM = 18.4
		p.CPANM = 1.85
	}

	html := fmt.Sprintf(`
<div class="block-marine-navigation rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "chartOrient": "north", "rangeNM": "%s", "sogKnots": %.1f, "cogTrue": %d, "depthM": %.1f, "cpaNM": %.2f, "showRadar": true, "showAis": true }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-blue-600/20 border border-blue-500/40 flex items-center justify-center text-blue-400 font-black text-sm">
        ⚓
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-blue-400">ECDIS Marine Navigation System</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-blue-500/10 text-blue-400 border border-blue-500/20">IMO / SOLAS</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>Vessel: %s</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">MMSI: %s</span>
        </h2>
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: NAUTICAL CHART & TELEMETRY INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-[#07131e] p-4 flex flex-col justify-between overflow-hidden border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="relative w-full max-w-[420px] aspect-square mx-auto flex items-center justify-center">
        <svg class="w-full h-full" viewBox="0 0 400 400">
          <circle cx="200" cy="200" r="180" fill="#0c2538" stroke="#334155" stroke-width="1.5" />
          <circle cx="200" cy="200" r="6" fill="#38bdf8" stroke="#ffffff" stroke-width="1.5" />
          <line x1="200" y1="200" x2="260" y2="168" stroke="#38bdf8" stroke-width="2.5" />
          <polygon points="280,140 286,146 280,152 274,146" fill="#f59e0b" stroke="#ffffff" stroke-width="1" />
        </svg>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div>
        <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
          <span class="text-[10px] font-black uppercase tracking-widest text-blue-400">Navigational Dynamics</span>
          <span class="text-[10px] font-bold text-emerald-400">AIS LIVE</span>
        </div>
        <div class="mt-4 space-y-3">
          <div class="grid grid-cols-2 gap-2">
            <div class="p-3 rounded-2xl bg-slate-950/70 border border-slate-800">
              <span class="text-[10px] font-extrabold uppercase text-slate-400 block">SOG</span>
              <span class="text-3xl font-black text-white" hx-text="sogKnots">%.1f</span>
            </div>
            <div class="p-3 rounded-2xl bg-slate-950/70 border border-slate-800">
              <span class="text-[10px] font-extrabold uppercase text-slate-400 block">COG</span>
              <span class="text-3xl font-black text-white" hx-text="cogTrue + '°'">%d°</span>
            </div>
          </div>
          <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
            <span class="text-xs font-bold text-slate-300">Under-Keel Clearance (UKC)</span>
            <span class="text-2xl font-black text-cyan-400 block mt-1" hx-text="depthM + ' m'">%.1f m</span>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-blue-400">GPS: 01°14'28" N, 103°55'12" E</span>
    <span class="text-blue-400 font-mono">@diag: ECDIS_OK</span>
  </div>
</div>`, p.RangeNM, p.SOGKnots, p.COGTrue, p.DepthM, p.CPANM, p.VesselName, p.MMSI, p.SOGKnots, p.COGTrue, p.DepthM)
	return Component{html: html}
}

// =============================================================================
// INDUSTRIAL DESIGN BLOCK 3: PRECISION AGRICULTURE & AGRONOMY SIMULATION
// =============================================================================

type BlockAgriculturePrecisionProps struct {
	ParcelID   string
	Crop       string
	AreaHa     int
	NDVIMean   float64
	SoilMoist  float64
	PivotAngle float64
}

func BlockAgriculturePrecision(p BlockAgriculturePrecisionProps) Component {
	if p.ParcelID == "" {
		p.ParcelID = "Parcel-B4 North Pivot"
		p.Crop = "Zea mays (Corn V6)"
		p.AreaHa = 160
		p.NDVIMean = 0.78
		p.SoilMoist = 28.4
		p.PivotAngle = 142.5
	}

	html := fmt.Sprintf(`
<div class="block-agriculture-precision rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "activeLayer": "ndvi", "pivotAngle": %.1f, "pivotSpeed": 45, "waterMm": 12.5, "ndviMean": %.2f, "soilMoist": %.1f, "pivotRunning": true }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-lime-600/20 border border-lime-500/40 flex items-center justify-center text-lime-400 font-black text-sm">
        🌾
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-lime-400">Precision Agronomy & Telemetry</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-lime-500/10 text-lime-400 border border-lime-500/20">ISOBUS 11783</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>%s (%d Ha)</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">%s</span>
        </h2>
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: PARCEL CANVAS & QUANTITATIVE AGRONOMIC INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-[#0b170c] p-4 flex flex-col justify-between overflow-hidden border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="relative w-full max-w-[420px] aspect-square mx-auto flex items-center justify-center">
        <svg class="w-full h-full" viewBox="0 0 400 400">
          <circle cx="200" cy="200" r="175" fill="#15803d" stroke="#4d7c0f" stroke-width="2" />
          <line x1="200" y1="200" x2="60" y2="305" stroke="#ffffff" stroke-width="3" stroke-linecap="round" />
        </svg>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
        <span class="text-xs font-bold text-slate-300">Mean Parcel NDVI</span>
        <span class="text-3xl font-black text-white block mt-1" hx-text="ndviMean">%.2f</span>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-lime-400">Weather: 26°C • Hum 42%%</span>
    <span class="text-lime-400 font-mono">@diag: AGRI_MESH_OK</span>
  </div>
</div>`, p.PivotAngle, p.NDVIMean, p.SoilMoist, p.ParcelID, p.AreaHa, p.Crop, p.NDVIMean)
	return Component{html: html}
}

// =============================================================================
// INDUSTRIAL DESIGN BLOCK 4: ELECTRICITY, OIL & GAS ENERGY SCADA
// =============================================================================

type BlockEnergySCADAProps struct {
	SubstationID string
	VoltageKv    float64
	FreqHz       float64
	LoadMw       float64
	WellheadPsi  int
}

func BlockEnergySCADA(p BlockEnergySCADAProps) Component {
	if p.SubstationID == "" {
		p.SubstationID = "Substation Delta-4 (230kV / 69kV)"
		p.VoltageKv = 230.4
		p.FreqHz = 60.02
		p.LoadMw = 142.5
		p.WellheadPsi = 2840
	}

	html := fmt.Sprintf(`
<div class="block-energy-scada rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "viewMode": "grid", "voltageKv": %.1f, "freqHz": %.2f, "loadMw": %.1f, "cb101Closed": true, "wellheadPsi": %d, "chokePct": 38 }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-amber-600/20 border border-amber-500/40 flex items-center justify-center text-amber-400 font-black text-sm">
        ⚡
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-amber-400">Energy & SCADA Grid Automation</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">IEC 61850</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>%s</span>
        </h2>
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: SCADA SINGLE-LINE DIAGRAM & POWER INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-slate-950 p-5 flex flex-col justify-between overflow-x-auto border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="relative w-full max-w-[460px] aspect-[4/3] mx-auto my-auto flex items-center justify-center">
        <svg class="w-full h-full" viewBox="0 0 500 350">
          <line x1="40" y1="60" x2="460" y2="60" stroke="#f59e0b" stroke-width="6" stroke-linecap="round" />
          <rect x="125" y="145" width="30" height="30" rx="6" fill="#1e293b" stroke="#10b981" stroke-width="2" />
          <line x1="40" y1="290" x2="460" y2="290" stroke="#06b6d4" stroke-width="5" stroke-linecap="round" />
        </svg>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
        <span class="text-[10px] font-extrabold uppercase text-slate-400 block mb-1">Grid Active Load</span>
        <span class="text-3xl font-black text-white" hx-text="loadMw + ' MW'">%.1f MW</span>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-amber-400">Protocol: IEC 61850 GOOSE</span>
    <span class="text-amber-400 font-mono">@diag: SCADA_GRID_SYNC</span>
  </div>
</div>`, p.VoltageKv, p.FreqHz, p.LoadMw, p.WellheadPsi, p.SubstationID, p.LoadMw)
	return Component{html: html}
}

// =============================================================================
// INDUSTRIAL DESIGN BLOCK 5: ELECTRONICS, MULTI-LAYER PCB & SEMICONDUCTOR FAB
// =============================================================================

type BlockPcbSemiconductorProps struct {
	PackageID    string
	LayerCount   int
	ZDiff        float64
	SkewPs       float64
	DieYield     float64
	DrcViolations int
}

func BlockPcbSemiconductor(p BlockPcbSemiconductorProps) Component {
	if p.PackageID == "" {
		p.PackageID = "BGA-1156 SoC Board (Rev D4)"
		p.LayerCount = 12
		p.ZDiff = 90.2
		p.SkewPs = 0.28
		p.DieYield = 94.8
		p.DrcViolations = 0
	}

	html := fmt.Sprintf(`
<div class="block-pcb-semiconductor rounded-3xl bg-slate-950 border border-slate-800/80 shadow-2xl overflow-hidden flex flex-col select-none"
     hx-state='{ "activeLayer": "l1", "zDiff": %.1f, "skewPs": %.2f, "dieYield": %.1f, "drcViolations": %d, "showWafer": false }'>

  <!-- ZONE 1: BLOCK HEADER -->
  <div class="px-5 py-3.5 bg-slate-900/90 border-b border-slate-800 flex flex-wrap items-center justify-between gap-3">
    <div class="flex items-center gap-3">
      <div class="w-8 h-8 rounded-xl bg-teal-600/20 border border-teal-500/40 flex items-center justify-center text-teal-400 font-black text-sm">
        💻
      </div>
      <div>
        <div class="flex items-center gap-2">
          <span class="text-xs font-black uppercase tracking-wider text-teal-400">PCB EDA & Semiconductor Fab</span>
          <span class="px-1.5 py-0.5 rounded text-[9px] font-extrabold bg-teal-500/10 text-teal-400 border border-teal-500/20">IPC CLASS 3</span>
        </div>
        <h2 class="text-sm font-bold text-white flex items-center gap-2">
          <span>%s</span>
          <span class="text-slate-500">•</span>
          <span class="text-xs text-slate-300 font-normal">%d-Layer Megtron 6</span>
        </h2>
      </div>
    </div>
  </div>

  <!-- ZONE 3 & 4: PCB GERBER VIEWPORT & IMPEDANCE INSPECTOR -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-0 min-h-[380px]">
    <div class="lg:col-span-8 relative bg-[#061214] p-5 flex flex-col justify-between overflow-x-auto border-b lg:border-b-0 lg:border-r border-slate-800/80">
      <div class="relative w-full max-w-[460px] aspect-[4/3] mx-auto my-auto flex items-center justify-center">
        <svg class="w-full h-full" viewBox="0 0 500 350">
          <rect x="195" y="115" width="90" height="90" fill="none" stroke="#ffffff" stroke-width="1.2" />
          <path d="M 270,160 L 320,160 L 330,150 L 340,170 L 350,150 L 360,170 L 370,160 L 460,160"
                fill="none" stroke="#2dd4bf" stroke-width="2.5" />
        </svg>
      </div>
    </div>

    <!-- Inspector -->
    <div class="lg:col-span-4 bg-slate-900/50 p-5 flex flex-col justify-between space-y-4">
      <div class="p-3.5 rounded-2xl bg-slate-950/70 border border-slate-800">
        <span class="text-[10px] font-extrabold uppercase text-slate-400 block mb-1">PCIe Gen 5 Differential Z_diff</span>
        <span class="text-3xl font-black text-white font-mono" hx-text="zDiff + ' Ω'">%.1f Ω</span>
      </div>
    </div>
  </div>

  <!-- ZONE 5: AUDIT & @DIAG STATUS FOOTER -->
  <div class="px-5 py-2.5 bg-slate-900/90 border-t border-slate-800 flex flex-wrap items-center justify-between gap-3 text-xs text-slate-400">
    <span class="text-[10px] font-mono text-teal-400">Standard: IPC-A-610 Class 3</span>
    <span class="text-teal-400 font-mono">@diag: EDA_DRC_PASS</span>
  </div>
</div>`, p.ZDiff, p.SkewPs, p.DieYield, p.DrcViolations, p.PackageID, p.LayerCount, p.ZDiff)
	return Component{html: html}
}
