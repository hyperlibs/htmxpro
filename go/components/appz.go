package components

import (
	"fmt"
	"strings"
)

// =============================================================================
// 1. SHELLUI & 8 DISTINCT VIEWPORT MODES
// =============================================================================

type ShellUIProps struct {
	Title      string
	ViewMode   string // "OS-H" | "OS-P" | "Tab-H" | "Tab-P" | "Mob-H" | "Mob-P" | "TV-H" | "TV-P"
	Mode       string // "mobile" | "game" | "os" | "iot"
	StatusBar  Component
	AppBar     Component
	Segments   []Component
	NavBar     Component
	Overlays   []Component
}

func ShellUI(p ShellUIProps) Component {
	var sb strings.Builder
	viewModeAttr := ""
	if p.ViewMode != "" {
		viewModeAttr = fmt.Sprintf(`view-mode="%s"`, p.ViewMode)
	}
	modeAttr := ""
	if p.Mode != "" {
		modeAttr = fmt.Sprintf(`mode="%s"`, p.Mode)
	}

	sb.WriteString(fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="h-full bg-slate-950 text-slate-100 select-none">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, viewport-fit=cover">
  <title>%s — Appz</title>
  <script src="https://unpkg.com/htmx.org@2.0.4"></script>
  <script src="/htmx-bolt.js"></script>
  <script src="/appz.js"></script>
  <script src="https://cdn.tailwindcss.com"></script>
  <style>
    body { padding-top: env(safe-area-inset-top); padding-bottom: env(safe-area-inset-bottom); }
  </style>
</head>
<body class="h-full flex flex-col antialiased overflow-hidden">
  <app-shellui %s %s class="flex-1 flex flex-col relative overflow-hidden">
    %s
    %s
    <main class="flex-1 overflow-y-auto flex flex-col relative">`, p.Title, viewModeAttr, modeAttr, p.StatusBar.Render(), p.AppBar.Render()))

	for _, seg := range p.Segments {
		sb.WriteString(seg.Render())
	}

	sb.WriteString(fmt.Sprintf(`
    </main>
    %s`, p.NavBar.Render()))

	for _, overlay := range p.Overlays {
		sb.WriteString(overlay.Render())
	}

	sb.WriteString(`
  </app-shellui>
</body>
</html>`)

	return Component{html: sb.String()}
}

// =============================================================================
// 2. LAYER-UX: OS SYSTEM PRIMITIVES
// =============================================================================

// --- Status Bar with Live Clock, Wi-Fi RSSI, Battery & Dynamic Island ---

type StatusBarProps struct {
	ShowIsland   bool
	IslandText   string
	WifiBars     int // 0..4
	BatteryPct   int
	NetworkLabel string
}

func StatusBar(p StatusBarProps) Component {
	wifiIcons := []string{"○", "▂", "▃", "▄", "▅"}
	wifiIcon := "▅"
	if p.WifiBars >= 0 && p.WifiBars < len(wifiIcons) {
		wifiIcon = wifiIcons[p.WifiBars]
	}
	if p.BatteryPct == 0 {
		p.BatteryPct = 94
	}
	if p.NetworkLabel == "" {
		p.NetworkLabel = "5G"
	}

	islandHtml := ""
	if p.ShowIsland {
		islandText := p.IslandText
		if islandText == "" {
			islandText = "IoT Normal"
		}
		islandHtml = fmt.Sprintf(`
    <div class="dynamic-island h-6 px-3 bg-black rounded-full flex items-center gap-1.5 border border-slate-800 text-[10px] font-bold text-slate-300 transition-all cursor-pointer"
         hx-action="$appz.haptic('selection')">
      <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
      <span>%s</span>
    </div>`, islandText)
	}

	html := fmt.Sprintf(`
<app-statusbar class="h-8 bg-slate-950/90 backdrop-blur border-b border-slate-900/80 px-3 flex items-center justify-between shrink-0 z-50 text-[11px] font-semibold text-slate-300">
  <div class="flex items-center gap-1.5">
    <span class="status-time font-bold text-white">12:00</span>
    <span class="text-[9px] text-slate-500 uppercase tracking-wider">%s</span>
  </div>
  %s
  <div class="flex items-center gap-2">
    <span title="Wi-Fi Signal" class="text-xs text-indigo-400 font-mono">%s</span>
    <div class="flex items-center gap-1">
      <span class="text-[10px] font-bold">%d%%</span>
      <span class="text-xs text-emerald-400">🔋</span>
    </div>
  </div>
</app-statusbar>`, p.NetworkLabel, islandHtml, wifiIcon, p.BatteryPct)

	return Component{html: html}
}

// --- Top AppBar ---

type AppBarProps struct {
	Title      string
	BackAction string
	Actions    []Component
}

func AppBar(p AppBarProps) Component {
	backBtn := ""
	if p.BackAction != "" {
		backBtn = fmt.Sprintf(`<button hx-action="%s" class="p-2 -ml-2 rounded-full hover:bg-slate-800 text-slate-300">←</button>`, p.BackAction)
	}
	var actionsSb strings.Builder
	for _, a := range p.Actions {
		actionsSb.WriteString(a.Render())
	}

	html := fmt.Sprintf(`
<app-bar class="h-14 bg-slate-900/90 backdrop-blur border-b border-slate-800/80 px-4 flex items-center justify-between shrink-0 z-40">
  <div class="flex items-center gap-2">
    %s
    <h1 class="font-bold text-base text-white tracking-tight">%s</h1>
  </div>
  <div class="flex items-center gap-2">
    %s
  </div>
</app-bar>`, backBtn, p.Title, actionsSb.String())

	return Component{html: html}
}

// --- Control Center Quick Settings Drawer ---

type QuickTile struct {
	Icon   string
	Label  string
	Active bool
	Action string
}

type ControlCenterProps struct {
	ID    string
	Tiles []QuickTile
}

func ControlCenter(p ControlCenterProps) Component {
	var tilesSb strings.Builder
	for _, t := range p.Tiles {
		activeClass := "bg-slate-800 text-slate-400"
		if t.Active {
			activeClass = "bg-indigo-600 text-white shadow-lg shadow-indigo-600/30"
		}
		tilesSb.WriteString(fmt.Sprintf(`
    <button hx-action="%s" class="p-3 rounded-2xl %s flex flex-col items-center justify-center gap-1 text-xs font-bold transition-all active:scale-95">
      <span class="text-xl">%s</span>
      <span>%s</span>
    </button>`, t.Action, activeClass, t.Icon, t.Label))
	}

	html := fmt.Sprintf(`
<app-control-center id="%s" data-open="false" 
  class="absolute inset-0 z-50 bg-slate-950/80 backdrop-blur-xl p-6 flex flex-col justify-between transition-all duration-300 opacity-0 pointer-events-none -translate-y-full">
  <div>
    <div class="flex items-center justify-between pb-4 border-b border-slate-800/80">
      <span class="font-black text-sm uppercase tracking-wider text-slate-400">Control Center</span>
      <button hx-action="$appz.toggleControlCenter()" class="w-8 h-8 rounded-full bg-slate-800 flex items-center justify-center text-xs text-slate-300">✕</button>
    </div>
    <div class="grid grid-cols-4 gap-3 mt-4">
      %s
    </div>
  </div>
</app-control-center>`, p.ID, tilesSb.String())

	return Component{html: html}
}

// --- Palm/WebOS-inspired Multitask Card Switcher ---

type MultitaskCard struct {
	Title string
	Icon  string
	Badge string
}

type MultitaskProps struct {
	Cards []MultitaskCard
}

func Multitask(p MultitaskProps) Component {
	var cardsSb strings.Builder
	for _, c := range p.Cards {
		cardsSb.WriteString(fmt.Sprintf(`
    <div class="w-56 h-80 rounded-2xl bg-slate-900 border border-slate-700/80 p-4 shadow-2xl flex flex-col justify-between cursor-pointer active:scale-95 transition-all">
      <div class="flex items-center justify-between">
        <span class="text-lg">%s</span>
        <span class="text-[10px] font-bold px-2 py-0.5 rounded-full bg-indigo-500/20 text-indigo-400 border border-indigo-500/30">%s</span>
      </div>
      <p class="font-black text-white text-sm">%s</p>
    </div>`, c.Icon, c.Badge, c.Title))
	}

	html := fmt.Sprintf(`
<app-multitask id="app-multitask" data-open="false" 
  class="absolute inset-0 z-50 bg-slate-950/90 backdrop-blur-2xl p-6 hidden flex-col justify-center items-center gap-6 transition-all duration-300">
  <div class="flex items-center gap-4 overflow-x-auto max-w-full py-4">
    %s
  </div>
  <button hx-action="$appz.toggleMultitask()" class="px-5 py-2 rounded-xl bg-slate-800 text-slate-300 font-bold text-xs">Close</button>
</app-multitask>`, cardsSb.String())

	return Component{html: html}
}

// --- Tactile PIN & Keypad ---

type KeypadProps struct {
	TargetID string
	Digits   int
	ValidPIN string
}

func Keypad(p KeypadProps) Component {
	if p.Digits == 0 {
		p.Digits = 4
	}
	if p.ValidPIN == "" {
		p.ValidPIN = "1234"
	}
	if p.TargetID == "" {
		p.TargetID = "#pin-display"
	}

	keys := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "clear", "0", "backspace"}
	var keysSb strings.Builder
	for _, k := range keys {
		label := k
		if k == "clear" {
			label = "C"
		} else if k == "backspace" {
			label = "⌫"
		}
		keysSb.WriteString(fmt.Sprintf(`
    <button data-key="%s" class="w-16 h-16 rounded-full bg-slate-900/90 hover:bg-slate-800 border border-slate-800 text-lg font-bold text-white flex items-center justify-center active:scale-90 transition-all shadow-md">
      %s
    </button>`, k, label))
	}

	html := fmt.Sprintf(`
<app-keypad target="%s" digits="%d" valid-pin="%s" class="flex flex-col items-center gap-6">
  <div id="%s" class="flex items-center gap-3"></div>
  <div class="grid grid-cols-3 gap-4">
    %s
  </div>
</app-keypad>`, p.TargetID, p.Digits, p.ValidPIN, strings.TrimPrefix(p.TargetID, "#"), keysSb.String())

	return Component{html: html}
}

// =============================================================================
// 3. LAYER-UX: RICH HARDWARE TELEMETRY & CONTROLLER BLOCKS
// =============================================================================

// --- High-Precision Circular SVG Telemetry Gauge ---

type GaugeProps struct {
	Value float64
	Min   float64
	Max   float64
	Unit  string
	Label string
}

func Gauge(p GaugeProps) Component {
	html := fmt.Sprintf(`<app-gauge value="%.1f" min="%.1f" max="%.1f" unit="%s" label="%s" class="inline-block"></app-gauge>`,
		p.Value, p.Min, p.Max, p.Unit, p.Label)
	return Component{html: html}
}

// --- Kinetic Swipeout Action Item ---

type SwipeoutProps struct {
	ID          string
	Content     Component
	LeftAction  Component
	RightAction Component
}

func Swipeout(p SwipeoutProps) Component {
	html := fmt.Sprintf(`
<app-swipeout id="%s" class="relative overflow-hidden rounded-xl bg-slate-900 border border-slate-800 block">
  <div class="swipeout-actions-left absolute inset-y-0 left-0 flex items-center">
    %s
  </div>
  <div class="swipeout-content relative bg-slate-900 p-3 z-10">
    %s
  </div>
  <div class="swipeout-actions-right absolute inset-y-0 right-0 flex items-center">
    %s
  </div>
</app-swipeout>`, p.ID, p.LeftAction.Render(), p.Content.Render(), p.RightAction.Render())
	return Component{html: html}
}

// --- Dual-Dial HVAC Thermostat ---

type ThermostatProps struct {
	TargetTemp float64
	Mode       string // "cool" | "heat" | "auto" | "eco"
}

func Thermostat(p ThermostatProps) Component {
	if p.TargetTemp == 0 {
		p.TargetTemp = 21.5
	}
	html := fmt.Sprintf(`
<app-thermostat temp="%.1f" mode="%s" class="p-6 rounded-3xl bg-slate-900/90 border border-slate-800 flex flex-col items-center gap-4 shadow-xl">
  <span class="text-xs uppercase font-extrabold tracking-widest text-indigo-400">HVAC Climate Control</span>
  <div class="flex items-center gap-6">
    <button class="thermostat-minus w-12 h-12 rounded-full bg-slate-800 hover:bg-slate-700 text-white font-bold text-xl active:scale-90 transition-all">-</button>
    <div class="text-center">
      <span class="thermostat-temp text-4xl font-black text-white">%.1f</span>
      <span class="text-sm font-bold text-slate-400 ml-1">°C</span>
    </div>
    <button class="thermostat-plus w-12 h-12 rounded-full bg-slate-800 hover:bg-slate-700 text-white font-bold text-xl active:scale-90 transition-all">+</button>
  </div>
</app-thermostat>`, p.TargetTemp, p.Mode, p.TargetTemp)
	return Component{html: html}
}

// =============================================================================
// 4. KOTLIN: 0-9 PHOTOSHOP-STYLE SEGMENT & LAYER STACKING
// =============================================================================

type SegmentProps struct {
	ID     string
	Class  string
	Layers []Component
}

func Segment(p SegmentProps) Component {
	var sb strings.Builder
	idAttr := ""
	if p.ID != "" {
		idAttr = fmt.Sprintf(`id="%s"`, p.ID)
	}
	sb.WriteString(fmt.Sprintf(`<app-seg %s class="relative %s">`, idAttr, p.Class))
	for _, l := range p.Layers {
		sb.WriteString(l.Render())
	}
	sb.WriteString(`</app-seg>`)
	return Component{html: sb.String()}
}

func Layer(level int, child Component) Component {
	clamped := level
	if clamped < 0 {
		clamped = 0
	}
	if clamped > 9 {
		clamped = 9
	}
	html := fmt.Sprintf(`<app-layer level="%d" style="z-index: %d;" class="relative">%s</app-layer>`, clamped, clamped*10, child.Render())
	return Component{html: html}
}

// =============================================================================
// 5. NAVIGATION BAR & SPATIAL JOYSTICK
// =============================================================================

type NavBarProps struct {
	Mode     string // "sticky" | "float"
	Behavior string // "hide-on-scroll" | "always-visible"
	Active   string
	Items    []NavItem
}

func NavBar(p NavBarProps) Component {
	var sb strings.Builder
	modeClass := "bg-slate-900 border-t border-slate-800"
	if p.Mode == "float" {
		modeClass = "mx-4 mb-4 rounded-2xl bg-slate-900/90 backdrop-blur-lg border border-slate-700/60 shadow-2xl"
	}

	sb.WriteString(fmt.Sprintf(`<app-navbar mode="%s" behavior="%s" class="h-16 px-4 flex items-center justify-around shrink-0 %s">`, p.Mode, p.Behavior, modeClass))
	for _, it := range p.Items {
		activeClass := "text-slate-400"
		if it.Active || it.Label == p.Active {
			activeClass = "text-indigo-400 font-bold"
		}
		sb.WriteString(fmt.Sprintf(`<a href="%s" class="flex flex-col items-center gap-1 text-[10px] %s"><span>%s</span><span>%s</span></a>`, it.Route, activeClass, it.Icon, it.Label))
	}
	sb.WriteString(`</app-navbar>`)
	return Component{html: sb.String()}
}

type JoystickProps struct {
	ID      string
	Channel string
	Size    int
}

func Joystick(p JoystickProps) Component {
	if p.Size == 0 {
		p.Size = 100
	}
	if p.Channel == "" {
		p.Channel = "game:joystick"
	}
	return Component{html: fmt.Sprintf(`
<app-joystick id="%s" hx-chan-send="%s" style="width: %dpx; height: %dpx;" 
  class="rounded-full bg-slate-800/60 border-2 border-indigo-500/40 backdrop-blur flex items-center justify-center touch-none">
  <div class="joystick-thumb w-10 h-10 rounded-full bg-indigo-500 shadow-lg pointer-events-none transition-transform"></div>
</app-joystick>`, p.ID, p.Channel, p.Size, p.Size)}
}
