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

// =============================================================================
// DOMAIN S1 EXTENDED: DAILY FORECAST, UV, WIND & AIR QUALITY
// =============================================================================

type ForecastDayItem struct {
	Day  string
	Icon string
	High int
	Low  int
}

type SmartForecastDailyProps struct {
	Days []ForecastDayItem
}

func SmartForecastDaily(p SmartForecastDailyProps) Component {
	if len(p.Days) == 0 {
		p.Days = []ForecastDayItem{
			{Day: "TODAY", Icon: "⛅", High: 22, Low: 14},
			{Day: "FRI", Icon: "☀️", High: 25, Low: 16},
			{Day: "SAT", Icon: "🌧️", High: 18, Low: 12},
			{Day: "SUN", Icon: "⛈️", High: 17, Low: 11},
			{Day: "MON", Icon: "🌤️", High: 21, Low: 13},
		}
	}
	var sb strings.Builder
	for i, d := range p.Days {
		bgClass := "bg-slate-950/60 border-slate-800/80"
		textColor := "text-slate-400"
		if i == 0 {
			bgClass = "bg-indigo-600/20 border-indigo-500/40"
			textColor = "text-indigo-400 font-extrabold"
		}
		sb.WriteString(fmt.Sprintf(`
    <div class="flex-shrink-0 w-20 p-2.5 rounded-2xl %s border flex flex-col items-center gap-1 cursor-pointer active:scale-95 transition-all"
         hx-action="activeDay = %d; $appz.haptic('selection')">
      <span class="text-[10px] %s">%s</span>
      <span class="text-xl">%s</span>
      <span class="text-xs font-black text-white">%d°</span>
      <span class="text-[9px] text-slate-400">%d°</span>
    </div>`, bgClass, i, textColor, d.Day, d.Icon, d.High, d.Low))
	}

	html := fmt.Sprintf(`
<div class="smart-forecast-daily p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-3"
     hx-state='{ "activeDay": 0 }'>
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">7-Day Meteorological Outlook</span>
    <span class="text-[10px] text-slate-400 font-bold">Updated Just Now</span>
  </div>
  <div class="flex items-center gap-2 overflow-x-auto pb-1">
    %s
  </div>
</div>`, sb.String())
	return Component{html: html}
}

type SmartUVIndexDialProps struct {
	UV float64
}

func SmartUVIndexDial(p SmartUVIndexDialProps) Component {
	if p.UV == 0 {
		p.UV = 7.4
	}
	risk := "Moderate"
	if p.UV >= 8 {
		risk = "Very High"
	} else if p.UV >= 6 {
		risk = "High"
	}
	html := fmt.Sprintf(`
<div class="smart-uv-index-dial p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl flex items-center justify-between"
     hx-state='{ "uv": %.1f }'>
  <div>
    <span class="text-[10px] font-black uppercase tracking-widest text-amber-400">Solar Radiation Sensor</span>
    <div class="flex items-baseline gap-2 mt-1">
      <span class="text-4xl font-black text-white" hx-text="uv">%.1f</span>
      <span class="text-xs font-bold text-amber-400 uppercase">%s</span>
    </div>
    <p class="text-[10px] text-slate-400 mt-1">Burn time: ~20 mins • SPF 50+ recommended</p>
  </div>
  <app-gauge value="%.1f" min="0" max="11" unit="" label="UV Risk" class="inline-block"></app-gauge>
</div>`, p.UV, p.UV, risk, p.UV)
	return Component{html: html}
}

type SmartWindCompassProps struct {
	SpeedKmh     float64
	DirectionDeg int
	GustKmh      float64
}

func SmartWindCompass(p SmartWindCompassProps) Component {
	if p.SpeedKmh == 0 {
		p.SpeedKmh = 24.5
		p.DirectionDeg = 45
		p.GustKmh = 38.0
	}
	html := fmt.Sprintf(`
<div class="smart-wind-compass p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl flex items-center justify-between"
     hx-state='{ "speedKmh": %.1f, "directionDeg": %d, "gustKmh": %.1f }'>
  <div>
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Ultrasonic Anemometer</span>
    <div class="flex items-baseline gap-2 mt-1">
      <span class="text-3xl font-black text-white" hx-text="speedKmh">%.1f</span>
      <span class="text-xs font-bold text-slate-400 uppercase">KM/H</span>
    </div>
    <p class="text-[10px] text-slate-400 mt-0.5">Gusting to <span class="font-bold text-white" hx-text="gustKmh">%.1f</span> km/h • Direction NE</p>
  </div>
  <div class="relative w-20 h-20 rounded-full border-2 border-slate-700 bg-slate-950 flex items-center justify-center">
    <span class="absolute top-1 text-[8px] font-black text-slate-400">N</span>
    <span class="absolute bottom-1 text-[8px] font-black text-slate-400">S</span>
    <span class="absolute left-1 text-[8px] font-black text-slate-400">W</span>
    <span class="absolute right-1 text-[8px] font-black text-slate-400">E</span>
    <div class="w-1.5 h-10 bg-indigo-500 rounded-full transform rotate-45 origin-center transition-transform shadow-md"></div>
  </div>
</div>`, p.SpeedKmh, p.DirectionDeg, p.GustKmh, p.SpeedKmh, p.GustKmh)
	return Component{html: html}
}

type SmartAirQualityPillProps struct {
	AQI  int
	PM25 float64
	PM10 float64
	CO2  int
}

func SmartAirQualityPill(p SmartAirQualityPillProps) Component {
	if p.AQI == 0 {
		p.AQI = 32
		p.PM25 = 8.2
		p.PM10 = 14.0
		p.CO2 = 480
	}
	html := fmt.Sprintf(`
<div class="smart-air-quality-pill p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-3"
     hx-state='{ "aqi": %d }'>
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-emerald-400">Air Quality Telemetry</span>
    <span class="text-[10px] font-extrabold px-2 py-0.5 rounded-full bg-emerald-500/20 text-emerald-400 border border-emerald-500/30">● Good (AQI %d)</span>
  </div>
  <div class="grid grid-cols-3 gap-2 text-center text-xs">
    <div class="p-2.5 rounded-2xl bg-slate-950/60 border border-slate-800">
      <span class="text-[9px] text-slate-500 uppercase block font-bold">PM 2.5</span>
      <span class="text-sm font-black text-white mt-0.5 block">%.1f µg</span>
    </div>
    <div class="p-2.5 rounded-2xl bg-slate-950/60 border border-slate-800">
      <span class="text-[9px] text-slate-500 uppercase block font-bold">PM 10</span>
      <span class="text-sm font-black text-white mt-0.5 block">%.1f µg</span>
    </div>
    <div class="p-2.5 rounded-2xl bg-slate-950/60 border border-slate-800">
      <span class="text-[9px] text-slate-500 uppercase block font-bold">CO2</span>
      <span class="text-sm font-black text-white mt-0.5 block">%d ppm</span>
    </div>
  </div>
</div>`, p.AQI, p.AQI, p.PM25, p.PM10, p.CO2)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S2 & S3 EXTENDED: LIGHTING & AUTOMATION
// =============================================================================

type SmartLightingDimmerProps struct {
	Brightness int
}

func SmartLightingDimmer(p SmartLightingDimmerProps) Component {
	if p.Brightness == 0 {
		p.Brightness = 75
	}
	html := fmt.Sprintf(`
<div class="smart-lighting-dimmer p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl flex flex-col items-center gap-3"
     hx-state='{ "brightness": %d }'>
  <span class="text-[10px] font-black uppercase tracking-widest text-amber-400">Atmospheric Lighting Dimmer</span>
  <div class="relative w-32 h-32 flex items-center justify-center">
    <div class="w-28 h-28 rounded-full bg-slate-950 border-4 border-amber-500/40 shadow-lg flex flex-col items-center justify-center cursor-pointer active:scale-95 transition-all"
         hx-action="$appz.haptic('selection')">
      <span class="text-3xl font-black text-white" hx-text="brightness + '%%'">%d%%</span>
      <span class="text-[9px] text-amber-400 font-bold">Warm Glow</span>
    </div>
  </div>
  <div class="flex items-center gap-2 w-full max-w-[200px]">
    <button class="flex-1 py-1.5 rounded-xl bg-slate-800 text-xs font-bold text-slate-300 active:scale-90 transition-all" hx-action="brightness = Math.max(0, brightness - 10); $appz.haptic('selection')">-10%%</button>
    <button class="flex-1 py-1.5 rounded-xl bg-amber-600 text-xs font-bold text-white active:scale-90 transition-all" hx-action="brightness = Math.min(100, brightness + 10); $appz.haptic('selection')">+10%%</button>
  </div>
</div>`, p.Brightness, p.Brightness)
	return Component{html: html}
}

type SmartSmartLockProps struct {
	Locked  bool
	Battery int
	Name    string
}

func SmartSmartLock(p SmartSmartLockProps) Component {
	if p.Battery == 0 {
		p.Battery = 92
	}
	if p.Name == "" {
		p.Name = "Front Entry Deadbolt"
	}
	lockEmoji := "🔒"
	statusText := "Door Locked"
	if !p.Locked {
		lockEmoji = "🔓"
		statusText = "Door Unlocked"
	}
	html := fmt.Sprintf(`
<div class="smart-smart-lock p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-3"
     hx-state='{ "locked": %t }'>
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">%s</span>
    <span class="text-[10px] font-bold text-slate-400">🔋 %d%% Battery</span>
  </div>
  <div class="flex items-center justify-between py-2">
    <div class="flex items-center gap-3">
      <div class="w-12 h-12 rounded-2xl flex items-center justify-center text-xl transition-all"
           hx-class='{ "bg-emerald-500/20 text-emerald-400 border border-emerald-500/30": locked, "bg-amber-500/20 text-amber-400 border border-amber-500/30": !locked }'>
        <span hx-text="locked ? '🔒' : '🔓'">%s</span>
      </div>
      <div>
        <h3 class="text-sm font-black text-white" hx-text="locked ? 'Door Locked' : 'Door Unlocked'">%s</h3>
        <p class="text-[10px] text-slate-500 font-bold">Auto-lock in 30 sec</p>
      </div>
    </div>
    <button class="px-4 py-2 rounded-xl font-bold text-xs transition-all active:scale-95"
            hx-class='{ "bg-slate-800 text-slate-300": locked, "bg-emerald-600 text-white shadow-lg shadow-emerald-600/30": !locked }'
            hx-action="locked = !locked; $appz.haptic(locked ? 'heavy' : 'success')">
      <span hx-text="locked ? 'Tap to Unlock' : 'Tap to Lock'">Tap to Unlock</span>
    </button>
  </div>
</div>`, p.Locked, p.Name, p.Battery, lockEmoji, statusText)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S4 EXTENDED: AUTOMOTIVE & INFOTAINMENT
// =============================================================================

type SmartEVBatteryRangeProps struct {
	Pct      int
	RangeKm  int
	Charging bool
}

func SmartEVBatteryRange(p SmartEVBatteryRangeProps) Component {
	if p.Pct == 0 {
		p.Pct = 84
		p.RangeKm = 438
	}
	html := fmt.Sprintf(`
<div class="smart-ev-battery-range p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-4"
     hx-state='{ "pct": %d, "rangeKm": %d, "charging": %t }'>
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-emerald-400">EV Traction Battery</span>
    <span class="text-[10px] font-bold text-slate-400">400V Architecture</span>
  </div>
  <div class="flex items-center justify-between">
    <div>
      <span class="text-4xl font-black text-white" hx-text="pct + '%%'">%d%%</span>
      <span class="text-[10px] font-bold text-slate-400 block mt-0.5">82.4 kWh Remaining</span>
    </div>
    <div class="text-right">
      <span class="text-2xl font-black text-indigo-400" hx-text="rangeKm + ' km'">%d km</span>
      <span class="text-[9px] font-extrabold uppercase text-slate-500 block">EPA Range</span>
    </div>
  </div>
  <div class="w-full h-3 rounded-full bg-slate-950 p-0.5 border border-slate-800 flex">
    <div class="h-full rounded-full bg-gradient-to-r from-emerald-500 to-indigo-500 transition-all duration-500" style="width: %d%%;"></div>
  </div>
</div>`, p.Pct, p.RangeKm, p.Charging, p.Pct, p.RangeKm, p.Pct)
	return Component{html: html}
}

type SmartDriveModeProps struct {
	Mode string // "chill" | "normal" | "sport" | "track"
}

func SmartDriveMode(p SmartDriveModeProps) Component {
	if p.Mode == "" {
		p.Mode = "sport"
	}
	html := fmt.Sprintf(`
<div class="smart-drive-mode p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-3"
     hx-state='{ "mode": "%s" }'>
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Powertrain Drive Profile</span>
    <span class="text-[10px] font-extrabold uppercase text-amber-400" hx-text="mode.toUpperCase()">%s</span>
  </div>
  <div class="grid grid-cols-4 gap-2">
    <button class="py-2.5 rounded-xl font-black text-xs transition-all active:scale-95 text-center"
            hx-class='{ "bg-emerald-600 text-white shadow-lg shadow-emerald-600/30": mode === "chill", "bg-slate-950 text-slate-400": mode !== "chill" }'
            hx-action="mode = 'chill'; $appz.haptic('selection')">Chill</button>
    <button class="py-2.5 rounded-xl font-black text-xs transition-all active:scale-95 text-center"
            hx-class='{ "bg-indigo-600 text-white shadow-lg shadow-indigo-600/30": mode === "normal", "bg-slate-950 text-slate-400": mode !== "normal" }'
            hx-action="mode = 'normal'; $appz.haptic('selection')">Normal</button>
    <button class="py-2.5 rounded-xl font-black text-xs transition-all active:scale-95 text-center"
            hx-class='{ "bg-amber-600 text-white shadow-lg shadow-amber-600/30": mode === "sport", "bg-slate-950 text-slate-400": mode !== "sport" }'
            hx-action="mode = 'sport'; $appz.haptic('selection')">Sport</button>
    <button class="py-2.5 rounded-xl font-black text-xs transition-all active:scale-95 text-center"
            hx-class='{ "bg-rose-600 text-white shadow-lg shadow-rose-600/30": mode === "track", "bg-slate-950 text-slate-400": mode !== "track" }'
            hx-action="mode = 'track'; $appz.haptic('heavy')">Track</button>
  </div>
</div>`, p.Mode, strings.ToUpper(p.Mode))
	return Component{html: html}
}

type SmartTurnByTurnNavProps struct {
	Distance string
	Street   string
	ETA      string
	TripKm   float64
	TripMin  int
}

func SmartTurnByTurnNav(p SmartTurnByTurnNavProps) Component {
	if p.Distance == "" {
		p.Distance = "450 m"
		p.Street = "Market Street"
		p.ETA = "12:34"
		p.TripKm = 14.2
		p.TripMin = 18
	}
	html := fmt.Sprintf(`
<div class="smart-turn-by-turn-nav p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl flex items-center justify-between"
     hx-state='{ "distance": "%s", "street": "%s", "turn": "left" }'>
  <div class="flex items-center gap-4">
    <div class="w-14 h-14 rounded-2xl bg-indigo-600 flex items-center justify-center text-white text-2xl font-black shadow-lg shadow-indigo-600/30">
      ↰
    </div>
    <div>
      <span class="text-xl font-black text-white" hx-text="'In ' + distance">In %s</span>
      <h3 class="text-xs font-bold text-slate-300" hx-text="street">%s</h3>
      <p class="text-[9px] text-slate-500 font-bold uppercase mt-0.5">Use 2 left lanes</p>
    </div>
  </div>
  <div class="text-right">
    <span class="text-xs font-black text-emerald-400">%s</span>
    <span class="text-[9px] text-slate-500 block font-bold">%.1f km • %d min</span>
  </div>
</div>`, p.Distance, p.Street, p.Distance, p.Street, p.ETA, p.TripKm, p.TripMin)
	return Component{html: html}
}

// =============================================================================
// DOMAIN S6 & S7 & S8: POS, HEALTH RINGS, AMBIENT AOD
// =============================================================================

type SmartPOSCashierProps struct {
	OrderNo  string
	TotalDue float64
	Tendered float64
}

func SmartPOSCashier(p SmartPOSCashierProps) Component {
	if p.OrderNo == "" {
		p.OrderNo = "#4210"
		p.TotalDue = 48.50
		p.Tendered = 50.00
	}
	html := fmt.Sprintf(`
<div class="smart-pos-cashier p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl space-y-4"
     hx-state='{ "totalDue": %.2f, "tendered": %.2f }'>
  <div class="flex items-center justify-between pb-2 border-b border-slate-800/80">
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">POS Checkout Terminal</span>
    <span class="text-[10px] font-bold text-slate-400">Order %s</span>
  </div>
  <div class="flex items-center justify-between">
    <div>
      <span class="text-[10px] text-slate-500 font-bold uppercase block">Total Due</span>
      <span class="text-3xl font-black text-white" hx-text="'$' + totalDue.toFixed(2)">$%.2f</span>
    </div>
    <div class="text-right">
      <span class="text-[10px] text-slate-500 font-bold uppercase block">Change Due</span>
      <span class="text-2xl font-black text-emerald-400" hx-text="'$' + (tendered - totalDue).toFixed(2)">$%.2f</span>
    </div>
  </div>
  <div class="grid grid-cols-3 gap-2 text-xs font-bold">
    <button class="py-2.5 rounded-xl bg-slate-950 border border-slate-800 text-white active:scale-95 transition-all" hx-action="tendered = 50.00; $appz.haptic('selection')">$50 Cash</button>
    <button class="py-2.5 rounded-xl bg-slate-950 border border-slate-800 text-white active:scale-95 transition-all" hx-action="tendered = 60.00; $appz.haptic('selection')">$60 Cash</button>
    <button class="py-2.5 rounded-xl bg-indigo-600 text-white font-extrabold active:scale-95 transition-all shadow-md shadow-indigo-600/30" hx-action="$appz.haptic('success')">NFC Tap 💳</button>
  </div>
</div>`, p.TotalDue, p.Tendered, p.OrderNo, p.TotalDue, p.Tendered-p.TotalDue)
	return Component{html: html}
}

type SmartStepActivityRingsProps struct {
	Move         int
	MoveGoal     int
	Exercise     int
	ExerciseGoal int
	Stand        int
	StandGoal    int
}

func SmartStepActivityRings(p SmartStepActivityRingsProps) Component {
	if p.MoveGoal == 0 {
		p.Move = 620
		p.MoveGoal = 600
		p.Exercise = 45
		p.ExerciseGoal = 30
		p.Stand = 10
		p.StandGoal = 12
	}
	html := fmt.Sprintf(`
<div class="smart-step-activity-rings p-5 rounded-3xl bg-slate-900 border border-slate-800 shadow-xl flex items-center justify-between"
     hx-state='{ "move": %d, "moveGoal": %d, "exercise": %d, "exerciseGoal": %d }'>
  <div>
    <span class="text-[10px] font-black uppercase tracking-widest text-indigo-400">Daily Activity Rings</span>
    <div class="mt-2 space-y-1 text-xs">
      <div class="flex items-center gap-2">
        <span class="w-2.5 h-2.5 rounded-full bg-rose-500"></span>
        <span class="font-bold text-white"><span hx-text="move">%d</span> / %d Cal</span>
      </div>
      <div class="flex items-center gap-2">
        <span class="w-2.5 h-2.5 rounded-full bg-emerald-500"></span>
        <span class="font-bold text-white"><span hx-text="exercise">%d</span> / %d Min</span>
      </div>
      <div class="flex items-center gap-2">
        <span class="w-2.5 h-2.5 rounded-full bg-sky-500"></span>
        <span class="font-bold text-white">%d / %d Hrs</span>
      </div>
    </div>
  </div>
  <div class="relative w-24 h-24 flex items-center justify-center">
    <svg class="w-full h-full transform -rotate-90" viewBox="0 0 100 100">
      <circle cx="50" cy="50" r="40" stroke="rgba(244, 63, 94, 0.2)" stroke-width="6" fill="transparent" />
      <circle cx="50" cy="50" r="40" stroke="#f43f5e" stroke-width="6" stroke-linecap="round" fill="transparent" stroke-dasharray="251" stroke-dashoffset="30" />
      <circle cx="50" cy="50" r="30" stroke="rgba(16, 185, 129, 0.2)" stroke-width="6" fill="transparent" />
      <circle cx="50" cy="50" r="30" stroke="#10b981" stroke-width="6" stroke-linecap="round" fill="transparent" stroke-dasharray="188" stroke-dashoffset="20" />
      <circle cx="50" cy="50" r="20" stroke="rgba(14, 165, 233, 0.2)" stroke-width="6" fill="transparent" />
      <circle cx="50" cy="50" r="20" stroke="#0ea5e9" stroke-width="6" stroke-linecap="round" fill="transparent" stroke-dasharray="125" stroke-dashoffset="35" />
    </svg>
  </div>
</div>`, p.Move, p.MoveGoal, p.Exercise, p.ExerciseGoal, p.Move, p.MoveGoal, p.Exercise, p.ExerciseGoal, p.Stand, p.StandGoal)
	return Component{html: html}
}

type SmartAmbientAODClockProps struct {
	TimeStr string
	DateStr string
	Battery int
	Temp    int
	Alarms  int
}

func SmartAmbientAODClock(p SmartAmbientAODClockProps) Component {
	if p.TimeStr == "" {
		p.TimeStr = "12:04"
		p.DateStr = "Thursday, October 15"
		p.Battery = 94
		p.Temp = 21
		p.Alarms = 2
	}
	html := fmt.Sprintf(`
<div class="smart-ambient-aod-clock p-8 rounded-3xl bg-black border border-slate-900 shadow-2xl flex flex-col items-center justify-center text-center">
  <span class="text-[9px] font-extrabold uppercase tracking-widest text-slate-600 mb-2">Ambient AOD Mode</span>
  <span class="text-6xl font-black text-slate-100 tracking-tight font-mono">%s</span>
  <span class="text-xs font-bold text-slate-500 mt-1">%s</span>
  <div class="flex items-center gap-3 mt-4 text-xs font-bold text-slate-400">
    <span class="flex items-center gap-1">🔋 %d%%</span>
    <span>•</span>
    <span class="flex items-center gap-1">⛅ %d°C</span>
    <span>•</span>
    <span class="flex items-center gap-1">🔔 %d Alarms</span>
  </div>
</div>`, p.TimeStr, p.DateStr, p.Battery, p.Temp, p.Alarms)
	return Component{html: html}
}

