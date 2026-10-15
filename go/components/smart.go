package components

import (
	"fmt"
	"strings"
)

// =============================================================================
// DOMAIN S1: SMART WEATHER & ATMOSPHERE
// =============================================================================

type SmartWeatherHeroProps struct {
	City      string
	Temp      float64
	Condition string // "Sunny" | "Rain" | "Cloudy" | "Storm"
	High      float64
	Low       float64
	FeelsLike float64
	Humidity  int
	WindKmh   float64
}

func SmartWeatherHero(p SmartWeatherHeroProps) Component {
	if p.City == "" {
		p.City = "San Francisco"
	}
	if p.Condition == "" {
		p.Condition = "Partly Cloudy"
	}
	html := fmt.Sprintf(`
<div class="smart-weather-hero p-5 rounded-3xl bg-gradient-to-br from-indigo-900/60 via-slate-900 to-slate-950 border border-indigo-700/40 shadow-xl flex flex-col justify-between"
     hx-state='{ "temp": %.1f, "condition": "%s" }'>
  <div class="flex items-start justify-between">
    <div>
      <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Atmospheric Telemetry</span>
      <h2 class="text-xl font-black text-white mt-0.5">%s</h2>
      <p class="text-xs text-slate-400">%s</p>
    </div>
    <span class="text-3xl">⛅</span>
  </div>
  <div class="flex items-baseline gap-2 my-3">
    <span class="text-5xl font-black text-white tracking-tighter" hx-text="temp">%.1f</span>
    <span class="text-lg font-bold text-slate-400">°C</span>
  </div>
  <div class="grid grid-cols-3 gap-2 pt-3 border-t border-slate-800/80 text-[10px] text-slate-300">
    <div><span class="text-slate-500 block">Feels</span><span class="font-bold">%.1f°C</span></div>
    <div><span class="text-slate-500 block">Humidity</span><span class="font-bold">%d%%</span></div>
    <div><span class="text-slate-500 block">Wind</span><span class="font-bold">%.1f km/h</span></div>
  </div>
</div>`, p.Temp, p.Condition, p.City, p.Condition, p.Temp, p.FeelsLike, p.Humidity, p.WindKmh)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S2: SMART DIALS & ROTARY CONTROLLERS
// =============================================================================

type SmartThermostatDualZoneProps struct {
	ZoneA string
	TempA float64
	ZoneB string
	TempB float64
	Mode  string
}

func SmartThermostatDualZone(p SmartThermostatDualZoneProps) Component {
	html := fmt.Sprintf(`
<div class="smart-thermostat-dual-zone p-5 rounded-3xl bg-slate-900/90 border border-slate-800 flex flex-col items-center gap-4 shadow-xl"
     hx-state='{ "tempA": %.1f, "tempB": %.1f, "mode": "%s" }'>
  <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Dual-Zone Climate Control</span>
  <div class="grid grid-cols-2 gap-4 w-full">
    <!-- Zone A -->
    <div class="flex flex-col items-center p-3 rounded-2xl bg-slate-950/60 border border-slate-800">
      <span class="text-[10px] font-extrabold text-slate-400">%s</span>
      <div class="flex items-center gap-2 my-2">
        <button class="w-8 h-8 rounded-full bg-slate-800 text-white font-bold text-sm" hx-action="tempA -= 0.5; $appz.haptic('selection')">-</button>
        <span class="text-2xl font-black text-white" hx-text="tempA.toFixed(1)">%.1f</span>
        <button class="w-8 h-8 rounded-full bg-slate-800 text-white font-bold text-sm" hx-action="tempA += 0.5; $appz.haptic('selection')">+</button>
      </div>
      <span class="text-[9px] text-indigo-400 font-bold">● Active Cooling</span>
    </div>
    <!-- Zone B -->
    <div class="flex flex-col items-center p-3 rounded-2xl bg-slate-950/60 border border-slate-800">
      <span class="text-[10px] font-extrabold text-slate-400">%s</span>
      <div class="flex items-center gap-2 my-2">
        <button class="w-8 h-8 rounded-full bg-slate-800 text-white font-bold text-sm" hx-action="tempB -= 0.5; $appz.haptic('selection')">-</button>
        <span class="text-2xl font-black text-white" hx-text="tempB.toFixed(1)">%.1f</span>
        <button class="w-8 h-8 rounded-full bg-slate-800 text-white font-bold text-sm" hx-action="tempB += 0.5; $appz.haptic('selection')">+</button>
      </div>
      <span class="text-[9px] text-emerald-400 font-bold">● Eco Auto</span>
    </div>
  </div>
</div>`, p.TempA, p.TempB, p.Mode, p.ZoneA, p.TempA, p.ZoneB, p.TempB)
	return Component{html: html}
}

type SmartRotaryVolumeProps struct {
	Level    int // 0..100
	Muted    bool
	OutputDev string
}

func SmartRotaryVolume(p SmartRotaryVolumeProps) Component {
	html := fmt.Sprintf(`
<div class="smart-rotary-volume p-5 rounded-3xl bg-slate-900 border border-slate-800 flex flex-col items-center gap-3 shadow-xl"
     hx-state='{ "vol": %d, "muted": %t }'>
  <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Audio Master Level</span>
  <div class="relative w-32 h-32 flex items-center justify-center">
    <div class="w-28 h-28 rounded-full bg-gradient-to-tr from-slate-950 via-slate-800 to-slate-900 border-4 border-indigo-500/40 shadow-inner flex flex-col items-center justify-center cursor-pointer active:scale-95 transition-all"
         hx-action="muted = !muted; $appz.haptic('medium')">
      <span class="text-2xl font-black text-white" hx-text="muted ? 'MUTED' : vol + '%%'">%d%%</span>
      <span class="text-[9px] text-slate-400 font-bold">%s</span>
    </div>
  </div>
  <div class="flex items-center gap-3">
    <button class="px-3 py-1 rounded-xl bg-slate-800 text-xs font-bold text-slate-300" hx-action="vol = Math.max(0, vol - 10); $appz.haptic('selection')">-10</button>
    <button class="px-3 py-1 rounded-xl bg-indigo-600 text-xs font-bold text-white" hx-action="vol = Math.min(100, vol + 10); $appz.haptic('selection')">+10</button>
  </div>
</div>`, p.Level, p.Muted, p.Level, p.OutputDev)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S3: SMART HOME & IOT AUTOMATION
// =============================================================================

type SmartRelaySwitchboardProps struct {
	Relays []RelayItem
}

type RelayItem struct {
	ID     string
	Label  string
	Active bool
	Watts  int
}

func SmartRelaySwitchboard(p SmartRelaySwitchboardProps) Component {
	var relaysSb strings.Builder
	for _, r := range p.Relays {
		activeClass := "bg-slate-800 text-slate-400 border-slate-700/40"
		if r.Active {
			activeClass = "bg-indigo-600 text-white border-indigo-500 shadow-md shadow-indigo-600/20"
		}
		relaysSb.WriteString(fmt.Sprintf(`
    <div class="p-3 rounded-2xl border %s flex items-center justify-between transition-all">
      <div>
        <span class="text-xs font-bold block">%s</span>
        <span class="text-[9px] text-slate-400">%d Watts</span>
      </div>
      <button class="w-8 h-5 rounded-full bg-slate-950 p-0.5 relative transition-all" hx-action="$appz.haptic('selection')">
        <div class="w-4 h-4 rounded-full bg-white shadow-md"></div>
      </button>
    </div>`, activeClass, r.Label, r.Watts))
	}

	html := fmt.Sprintf(`
<div class="smart-relay-switchboard p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-3">
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">8-Channel Power Relays</span>
    <button class="text-[10px] font-extrabold text-rose-400 bg-rose-500/10 px-2 py-0.5 rounded-full border border-rose-500/20" hx-action="$appz.haptic('heavy')">Killswitch</button>
  </div>
  <div class="grid grid-cols-2 gap-2">
    %s
  </div>
</div>`, relaysSb.String())
	return Component{html: html}
}

type SmartSolarPowerFlowProps struct {
	SolarKw   float64
	BatteryKw float64
	HomeKw    float64
	GridKw    float64
}

func SmartSolarPowerFlow(p SmartSolarPowerFlowProps) Component {
	html := fmt.Sprintf(`
<div class="smart-solar-power-flow p-5 rounded-3xl bg-gradient-to-b from-slate-900 to-slate-950 border border-slate-800 shadow-xl space-y-4">
  <div class="flex items-center justify-between">
    <span class="text-[10px] font-black uppercase tracking-widest text-amber-400">Solar Energy Flow Matrix</span>
    <span class="text-[10px] font-bold text-emerald-400">● 100%% Clean</span>
  </div>
  <div class="grid grid-cols-2 gap-3 text-center">
    <div class="p-3 rounded-2xl bg-amber-500/10 border border-amber-500/30">
      <span class="text-xl">☀️</span>
      <span class="text-xs font-bold text-slate-300 block mt-1">Solar PV</span>
      <span class="text-base font-black text-amber-400">%.1f kW</span>
    </div>
    <div class="p-3 rounded-2xl bg-emerald-500/10 border border-emerald-500/30">
      <span class="text-xl">🔋</span>
      <span class="text-xs font-bold text-slate-300 block mt-1">Powerwall</span>
      <span class="text-base font-black text-emerald-400">%.1f kW</span>
    </div>
    <div class="p-3 rounded-2xl bg-indigo-500/10 border border-indigo-500/30">
      <span class="text-xl">🏠</span>
      <span class="text-xs font-bold text-slate-300 block mt-1">Home Load</span>
      <span class="text-base font-black text-indigo-400">%.1f kW</span>
    </div>
    <div class="p-3 rounded-2xl bg-slate-800/60 border border-slate-700/60">
      <span class="text-xl">⚡</span>
      <span class="text-xs font-bold text-slate-300 block mt-1">Grid Export</span>
      <span class="text-base font-black text-slate-300">%.1f kW</span>
    </div>
  </div>
</div>`, p.SolarKw, p.BatteryKw, p.HomeKw, p.GridKw)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S4: AUTOMOTIVE & CAR INFOTAINMENT HUD
// =============================================================================

type SmartCarCockpitHUDProps struct {
	SpeedKmh    int
	BatteryPct  int
	RangeKm     int
	DriveGear   string // "P" | "R" | "N" | "D"
	SpeedLimit  int
}

func SmartCarCockpitHUD(p SmartCarCockpitHUDProps) Component {
	html := fmt.Sprintf(`
<div class="smart-car-cockpit-hud p-6 rounded-3xl bg-slate-950 border border-slate-800 shadow-2xl flex flex-col justify-between"
     hx-state='{ "speed": %d, "gear": "%s" }'>
  <div class="flex items-center justify-between pb-2 border-b border-slate-900">
    <div class="flex items-center gap-2">
      <span class="text-xs font-black text-indigo-400">TESLA-DRIVE</span>
      <span class="text-[10px] font-bold text-slate-500">AUTONOMOUS L4</span>
    </div>
    <div class="flex items-center gap-3 text-xs font-black">
      <span class="text-slate-500">P</span>
      <span class="text-slate-500">R</span>
      <span class="text-slate-500">N</span>
      <span class="text-indigo-400 px-2 py-0.5 rounded-lg bg-indigo-500/20 border border-indigo-500/30">D</span>
    </div>
  </div>
  
  <div class="flex items-center justify-around py-6">
    <!-- Speedometer -->
    <div class="flex flex-col items-center">
      <span class="text-6xl font-black text-white tracking-tight" hx-text="speed">%d</span>
      <span class="text-[11px] font-extrabold uppercase tracking-widest text-slate-400 mt-1">KM / H</span>
    </div>
    <!-- Speed Limit Pill -->
    <div class="w-10 h-10 rounded-full border-4 border-rose-600 bg-white text-black flex items-center justify-center font-black text-xs shadow-md">
      %d
    </div>
  </div>

  <div class="grid grid-cols-2 gap-3 pt-3 border-t border-slate-900 text-xs">
    <div class="flex items-center gap-2">
      <span class="text-emerald-400">🔋</span>
      <span class="font-bold text-white">%d%% Battery</span>
    </div>
    <div class="flex items-center justify-end gap-1 font-bold text-indigo-400">
      <span>%d km</span>
      <span class="text-[10px] text-slate-500 uppercase">Est. Range</span>
    </div>
  </div>
</div>`, p.SpeedKmh, p.DriveGear, p.SpeedKmh, p.SpeedLimit, p.BatteryPct, p.RangeKm)
	return Component{html: html}
}

type SmartTirePressureTPMSProps struct {
	FrontLeft  float64
	FrontRight float64
	RearLeft   float64
	RearRight  float64
}

func SmartTirePressureTPMS(p SmartTirePressureTPMSProps) Component {
	html := fmt.Sprintf(`
<div class="smart-tire-pressure-tpms p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-3">
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Tire Pressure Monitor (TPMS)</span>
    <span class="text-[10px] font-bold text-emerald-400">All Normal</span>
  </div>
  <div class="relative w-48 h-56 mx-auto bg-slate-950 rounded-2xl border border-slate-800 flex items-center justify-center p-3">
    <!-- Vehicle Wireframe Silhouette -->
    <div class="w-20 h-40 border-2 border-slate-700/50 rounded-2xl absolute"></div>
    <!-- 4 Wheels -->
    <div class="absolute top-4 left-2 text-[10px] font-black text-white bg-slate-900 px-1.5 py-0.5 rounded border border-emerald-500/50">FL: %.1f</div>
    <div class="absolute top-4 right-2 text-[10px] font-black text-white bg-slate-900 px-1.5 py-0.5 rounded border border-emerald-500/50">FR: %.1f</div>
    <div class="absolute bottom-4 left-2 text-[10px] font-black text-white bg-slate-900 px-1.5 py-0.5 rounded border border-emerald-500/50">RL: %.1f</div>
    <div class="absolute bottom-4 right-2 text-[10px] font-black text-white bg-slate-900 px-1.5 py-0.5 rounded border border-emerald-500/50">RR: %.1f</div>
    <span class="text-[9px] text-slate-500 font-bold uppercase tracking-widest z-10">BAR</span>
  </div>
</div>`, p.FrontLeft, p.FrontRight, p.RearLeft, p.RearRight)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S5: INDUSTRIAL HMI & FACTORY CONTROLS
// =============================================================================

type SmartConveyorBeltProps struct {
	BeltID    string
	SpeedRPM  int
	Running   bool
	Overload  bool
}

func SmartConveyorBelt(p SmartConveyorBeltProps) Component {
	stateColor := "bg-emerald-500/10 text-emerald-400 border-emerald-500/20"
	if !p.Running {
		stateColor = "bg-slate-800 text-slate-400 border-slate-700"
	}
	html := fmt.Sprintf(`
<div class="smart-conveyor-belt p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-3">
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <div>
      <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Production Line HMI</span>
      <h3 class="text-xs font-bold text-white">%s</h3>
    </div>
    <span class="text-[10px] font-extrabold px-2 py-0.5 rounded-full border %s">%d RPM</span>
  </div>
  <div class="flex items-center gap-3">
    <button class="flex-1 py-2 rounded-xl bg-emerald-600 hover:bg-emerald-500 text-white font-bold text-xs" hx-action="$appz.haptic('success')">Start VFD</button>
    <button class="flex-1 py-2 rounded-xl bg-rose-600 hover:bg-rose-500 text-white font-bold text-xs" hx-action="$appz.haptic('error')">Emergency Stop</button>
  </div>
</div>`, p.BeltID, stateColor, p.SpeedRPM)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S6 & S7: RETAIL POS & HEALTH VITALS
// =============================================================================

type SmartHeartRateProps struct {
	BPM    int
	Status string // "Resting" | "Cardio" | "Peak"
}

func SmartHeartRate(p SmartHeartRateProps) Component {
	html := fmt.Sprintf(`
<div class="smart-heart-rate p-5 rounded-3xl bg-gradient-to-br from-rose-950/40 via-slate-900 to-slate-950 border border-rose-900/40 shadow-xl flex items-center justify-between">
  <div>
    <span class="text-[10px] font-black uppercase tracking-widest text-rose-400">Continuous PPG Sensor</span>
    <div class="flex items-baseline gap-2 mt-1">
      <span class="text-4xl font-black text-white">%d</span>
      <span class="text-xs font-bold text-slate-400 uppercase">BPM</span>
    </div>
    <span class="text-[10px] font-bold text-rose-400 mt-1 block">● %s Zone</span>
  </div>
  <span class="text-4xl animate-pulse">❤️</span>
</div>`, p.BPM, p.Status)
	return Component{html: html}
}
