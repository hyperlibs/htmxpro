(() => {
  var __defProp = Object.defineProperty;
  var __getOwnPropNames = Object.getOwnPropertyNames;
  var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
  var __hasOwnProp = Object.prototype.hasOwnProperty;
  var __moduleCache = /* @__PURE__ */ new WeakMap;
  var __toCommonJS = (from) => {
    var entry = __moduleCache.get(from), desc;
    if (entry)
      return entry;
    entry = __defProp({}, "__esModule", { value: true });
    if (from && typeof from === "object" || typeof from === "function")
      __getOwnPropNames(from).map((key) => !__hasOwnProp.call(entry, key) && __defProp(entry, key, {
        get: () => from[key],
        enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable
      }));
    __moduleCache.set(from, entry);
    return entry;
  };
  var __export = (target, all) => {
    for (var name in all)
      __defProp(target, name, {
        get: all[name],
        enumerable: true,
        configurable: true,
        set: (newValue) => all[name] = () => newValue
      });
  };

  // src/appz.ts
  var exports_appz = {};
  __export(exports_appz, {
    runMasterDiagnostics: () => runMasterDiagnostics,
    getAppzChannel: () => getAppzChannel,
    applySegmentLayerStyles: () => applySegmentLayerStyles,
    VIEWPORT_SPECS: () => VIEWPORT_SPECS,
    LayerUX: () => LayerUX,
    LayerSystem: () => LayerSystem,
    LayerSys: () => LayerSys,
    LayerSrv: () => LayerSrv,
    LayerService: () => LayerService,
    LayerPresentation: () => LayerPresentation,
    LayerFirmware: () => LayerFirmware,
    LayerFirm: () => LayerFirm,
    HardwareBridge: () => HardwareBridge,
    AppzNavigator: () => AppzNavigator,
    AppzEngine: () => AppzEngine,
    AppzChan: () => AppzChan,
    Appz: () => Appz
  });
  var VIEWPORT_SPECS = {
    "OS-H": {
      mode: "OS-H",
      formFactor: "OS",
      orientation: "H",
      width: 1024,
      height: 600,
      label: "OS Horizontal",
      description: "Smart Display, In-Vehicle Infotainment (IVI), Edge Console"
    },
    "OS-P": {
      mode: "OS-P",
      formFactor: "OS",
      orientation: "P",
      width: 600,
      height: 1024,
      label: "OS Vertical",
      description: "Smart Mirror, Vertical Kiosk, Doorbell Intercom Station"
    },
    "Tab-H": {
      mode: "Tab-H",
      formFactor: "Tab",
      orientation: "H",
      width: 1180,
      height: 820,
      label: "Tablet Horizontal",
      description: "Dual-Pane Master-Detail, Left Navigation Rail, Telemetry Hub"
    },
    "Tab-P": {
      mode: "Tab-P",
      formFactor: "Tab",
      orientation: "P",
      width: 820,
      height: 1180,
      label: "Tablet Vertical",
      description: "Field Clipboard, POS Register, Equipment Diagnostics"
    },
    "Mob-H": {
      mode: "Mob-H",
      formFactor: "Mob",
      orientation: "H",
      width: 844,
      height: 390,
      label: "Mobile Horizontal",
      description: "Handheld Gaming, Drone Controller, Live Stream HUD"
    },
    "Mob-P": {
      mode: "Mob-P",
      formFactor: "Mob",
      orientation: "P",
      width: 390,
      height: 844,
      label: "Mobile Vertical",
      description: "One-Hand Reach, Dynamic Island, Full-Bleed Card Feed"
    },
    "TV-H": {
      mode: "TV-H",
      formFactor: "TV",
      orientation: "H",
      width: 1280,
      height: 720,
      label: "TV Horizontal",
      description: "16:9 Living Room Dashboard, 4-Zone Command Center"
    },
    "TV-P": {
      mode: "TV-P",
      formFactor: "TV",
      orientation: "P",
      width: 720,
      height: 1280,
      label: "TV Vertical",
      description: "Digital Signage Totem, Flight Board, Retail Directory"
    }
  };

  class LayerFirmware {
    wakeLockActive = false;
    hapticCount = 0;
    haptic(type = "light") {
      this.hapticCount++;
      if (typeof navigator !== "undefined" && "vibrate" in navigator) {
        switch (type) {
          case "light":
            navigator.vibrate(10);
            break;
          case "medium":
            navigator.vibrate(25);
            break;
          case "heavy":
            navigator.vibrate(50);
            break;
          case "selection":
            navigator.vibrate(6);
            break;
          case "success":
            navigator.vibrate([15, 30, 20]);
            break;
          case "warning":
            navigator.vibrate([30, 40, 30]);
            break;
          case "error":
            navigator.vibrate([60, 60, 60]);
            break;
        }
      }
    }
    async wakeLock() {
      if (typeof navigator !== "undefined" && "wakeLock" in navigator) {
        try {
          await navigator.wakeLock.request("screen");
          this.wakeLockActive = true;
          return true;
        } catch (e) {
          return false;
        }
      }
      return false;
    }
    async battery() {
      if (typeof navigator !== "undefined" && "getBattery" in navigator) {
        try {
          const b = await navigator.getBattery();
          return { level: Math.round(b.level * 100), charging: b.charging, voltage: 3.85 };
        } catch (e) {}
      }
      return { level: 94, charging: true, voltage: 4.12 };
    }
    network() {
      if (typeof navigator !== "undefined") {
        const conn = navigator.connection || navigator.mozConnection || navigator.webkitConnection;
        const isOnline = typeof navigator.onLine === "boolean" ? navigator.onLine : true;
        return {
          online: isOnline,
          effectiveType: conn?.effectiveType || "wifi",
          rtt: conn?.rtt || 18,
          rssiBars: isOnline ? 4 : 0
        };
      }
      return { online: true, effectiveType: "wifi", rtt: 12, rssiBars: 4 };
    }
    sensors() {
      return {
        lightLux: 420,
        tempC: 22.4,
        humidity: 48.5,
        cpuTemp: 44.8
      };
    }
    audioBeep(freq = 440, durationMs = 100) {
      if (typeof window !== "undefined" && window.AudioContext) {
        try {
          const ctx = new (window.AudioContext || window.webkitAudioContext);
          const osc = ctx.createOscillator();
          const gain = ctx.createGain();
          osc.connect(gain);
          gain.connect(ctx.destination);
          osc.frequency.value = freq;
          gain.gain.setValueAtTime(0.08, ctx.currentTime);
          gain.gain.exponentialRampToValueAtTime(0.001, ctx.currentTime + durationMs / 1000);
          osc.start();
          osc.stop(ctx.currentTime + durationMs / 1000);
        } catch (e) {}
      }
    }
    diag() {
      const d = {
        layer: "Layer-Firm",
        timestamp: Date.now(),
        battery: { level: 94, charging: true, voltage: 4.12 },
        network: this.network(),
        sensors: this.sensors(),
        hapticSupported: typeof navigator !== "undefined" && "vibrate" in navigator,
        wakeLockActive: this.wakeLockActive,
        hapticTriggerCount: this.hapticCount
      };
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("appz:diag:firm", { detail: d }));
      }
      return d;
    }
  }
  var LayerFirm = new LayerFirmware;

  class AppzChan {
    listeners = new Set;
    send(val) {
      this.listeners.forEach((cb) => {
        try {
          cb(val);
        } catch (e) {
          console.error(`[AppzChan] listener error:`, e);
        }
      });
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("appz:channel", { detail: { payload: val, timestamp: Date.now() } }));
      }
    }
    recv(cb) {
      this.listeners.add(cb);
      return () => this.listeners.delete(cb);
    }
    close() {
      this.listeners.clear();
    }
  }
  var activeChannels = new Map;
  function getAppzChannel(name) {
    if (!activeChannels.has(name)) {
      activeChannels.set(name, new AppzChan);
    }
    return activeChannels.get(name);
  }

  class LayerSystem {
    startTime = Date.now();
    powerState = "awake";
    storageCache = new Map;
    chan(name) {
      return getAppzChannel(name);
    }
    clock() {
      const now = new Date;
      return {
        time: now.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit", hour12: false }),
        date: now.toLocaleDateString([], { weekday: "short", month: "short", day: "numeric" }),
        uptimeSec: Math.floor((Date.now() - this.startTime) / 1000)
      };
    }
    getPowerState() {
      return this.powerState;
    }
    setPowerState(state) {
      this.powerState = state;
      LayerFirm.haptic("selection");
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("appz:power-state", { detail: { state } }));
      }
    }
    storage = {
      get: (key) => {
        if (typeof localStorage !== "undefined") {
          try {
            const val = localStorage.getItem(`appz:${key}`);
            return val ? JSON.parse(val) : null;
          } catch (e) {
            return null;
          }
        }
        return this.storageCache.get(key) || null;
      },
      set: (key, val) => {
        if (typeof localStorage !== "undefined") {
          try {
            localStorage.setItem(`appz:${key}`, JSON.stringify(val));
          } catch (e) {}
        }
        this.storageCache.set(key, val);
      }
    };
    diag() {
      let totalListeners = 0;
      const channelNames = [];
      activeChannels.forEach((ch, name) => {
        channelNames.push(name);
        totalListeners += ch.listeners.size;
      });
      const clk = this.clock();
      const d = {
        layer: "Layer-Sys",
        timestamp: Date.now(),
        uptimeSec: clk.uptimeSec,
        clock: { time: clk.time, date: clk.date },
        powerState: this.powerState,
        channelsCount: activeChannels.size,
        activeChannels: channelNames,
        totalChannelListeners: totalListeners,
        storageKeysCount: typeof localStorage !== "undefined" ? localStorage.length : this.storageCache.size
      };
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("appz:diag:sys", { detail: d }));
      }
      return d;
    }
  }
  var LayerSys = new LayerSystem;

  class LayerService {
    isTelemetryRunning = false;
    alerts = [
      { id: "alt-01", severity: "warning", title: "Solar Inverter Ambient Temp: 58°C", source: "Grid Substation Alpha", timestamp: Date.now() - 120000, acknowledged: false },
      { id: "alt-02", severity: "info", title: "Mesh Node 04 Reconnected (RSSI -42dBm)", source: "IoT Gateway", timestamp: Date.now() - 360000, acknowledged: true }
    ];
    otaState = {
      availableVersion: "v2.4.1",
      currentVersion: "v2.4.0",
      downloading: false,
      progress: 0
    };
    telemetry = {
      start: (intervalMs = 2000) => {
        if (typeof window === "undefined")
          return () => {};
        this.isTelemetryRunning = true;
        const timer = setInterval(() => {
          const powerKw = 4800 + Math.floor(Math.random() * 80) - 40;
          const temp = 22 + parseFloat((Math.random() * 0.8 - 0.4).toFixed(1));
          LayerSys.chan("iot:telemetry").send({ powerKw, temp, timestamp: Date.now() });
        }, intervalMs);
        return () => {
          clearInterval(timer);
          this.isTelemetryRunning = false;
        };
      }
    };
    alarms = {
      list: () => [...this.alerts],
      add: (alert) => {
        const newAlert = {
          ...alert,
          id: `alt-${Date.now().toString(36)}`,
          timestamp: Date.now(),
          acknowledged: false
        };
        this.alerts.unshift(newAlert);
        LayerFirm.haptic("error");
        LayerSys.chan("iot:alerts").send(newAlert);
        return newAlert;
      },
      ack: (id) => {
        const item = this.alerts.find((a) => a.id === id);
        if (item) {
          item.acknowledged = true;
          LayerFirm.haptic("selection");
          LayerSys.chan("iot:alerts-ack").send({ id });
        }
      }
    };
    ota = {
      getStatus: () => ({ ...this.otaState }),
      checkUpdate: async () => ({ updateAvailable: true, version: "v2.4.1" }),
      startInstall: (onProgress) => {
        this.otaState.downloading = true;
        this.otaState.progress = 0;
        let p = 0;
        const interval = setInterval(() => {
          p += 20;
          this.otaState.progress = p;
          onProgress(p);
          if (p >= 100) {
            clearInterval(interval);
            this.otaState.downloading = false;
            this.otaState.currentVersion = this.otaState.availableVersion;
            LayerFirm.haptic("success");
          }
        }, 300);
      }
    };
    diag() {
      const unack = this.alerts.filter((a) => !a.acknowledged);
      const critical = unack.filter((a) => a.severity === "critical");
      const d = {
        layer: "Layer-Srv",
        timestamp: Date.now(),
        telemetryActive: this.isTelemetryRunning,
        totalAlerts: this.alerts.length,
        unacknowledgedAlerts: unack.length,
        criticalAlerts: critical.length,
        otaCurrentVersion: this.otaState.currentVersion,
        otaAvailableVersion: this.otaState.availableVersion,
        otaDownloading: this.otaState.downloading,
        otaProgress: this.otaState.progress
      };
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("appz:diag:srv", { detail: d }));
      }
      return d;
    }
  }
  var LayerSrv = new LayerService;

  class LayerPresentation {
    history = [];
    currentViewMode = "Mob-P";
    push(url, transition = "slide-left") {
      LayerFirm.haptic("selection");
      this.history.push(url);
      if (typeof window !== "undefined" && window.htmx) {
        return window.htmx.ajax("GET", url, {
          target: "app-shellui",
          swap: "innerHTML transition:true"
        });
      }
      return Promise.resolve();
    }
    pop() {
      if (this.history.length <= 1)
        return false;
      LayerFirm.haptic("light");
      this.history.pop();
      const prevUrl = this.history[this.history.length - 1];
      if (typeof window !== "undefined" && window.htmx) {
        window.htmx.ajax("GET", prevUrl, {
          target: "app-shellui",
          swap: "innerHTML transition:true"
        });
      }
      return true;
    }
    showSheet(sheetId, snap = "70%") {
      LayerFirm.haptic("medium");
      if (typeof document === "undefined")
        return;
      const sheet = document.querySelector(sheetId);
      if (!sheet)
        return;
      sheet.classList.remove("hidden");
      sheet.setAttribute("data-state", "open");
      sheet.style.transform = `translateY(0%)`;
    }
    hideSheet(sheetId) {
      LayerFirm.haptic("light");
      if (typeof document === "undefined")
        return;
      const sheet = document.querySelector(sheetId);
      if (!sheet)
        return;
      sheet.setAttribute("data-state", "closed");
      sheet.style.transform = `translateY(100%)`;
      setTimeout(() => sheet.classList.add("hidden"), 250);
    }
    setViewMode(mode) {
      if (!VIEWPORT_SPECS[mode])
        return;
      this.currentViewMode = mode;
      const spec = VIEWPORT_SPECS[mode];
      LayerFirm.haptic("selection");
      if (typeof document !== "undefined") {
        const frame = document.querySelector(".device-viewport-frame");
        if (frame) {
          frame.style.width = `${spec.width}px`;
          frame.style.height = `${spec.height}px`;
          frame.setAttribute("data-view-mode", mode);
          frame.setAttribute("data-form-factor", spec.formFactor);
          frame.setAttribute("data-orientation", spec.orientation);
        }
        const scaler = document.querySelector(".viewport-scaler");
        if (scaler && frame) {
          const availableW = Math.max(320, window.innerWidth - 48);
          const availableH = Math.max(320, window.innerHeight - 150);
          const scaleW = spec.width > availableW ? availableW / spec.width : 1;
          const scaleH = spec.height > availableH ? availableH / spec.height : 1;
          const scale = Math.min(1, Math.min(scaleW, scaleH));
          scaler.style.transform = `scale(${scale})`;
          scaler.style.transformOrigin = "top center";
          scaler.style.height = `${spec.height * scale}px`;
          scaler.style.width = `${spec.width * scale}px`;
        }
        document.querySelectorAll("app-shellui, [app-shellui]").forEach((el) => {
          el.setAttribute("view-mode", mode);
          el.setAttribute("form-factor", spec.formFactor);
          el.setAttribute("orientation", spec.orientation);
        });
        document.querySelectorAll("[data-view-btn]").forEach((btn) => {
          const active = btn.getAttribute("data-view-btn") === mode;
          btn.classList.toggle("bg-indigo-600", active);
          btn.classList.toggle("text-white", active);
          btn.classList.toggle("text-slate-400", !active);
          btn.classList.toggle("bg-slate-800/80", !active);
        });
      }
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("appz:view-changed", { detail: spec }));
      }
    }
    toggleControlCenter() {
      LayerFirm.haptic("medium");
      if (typeof document === "undefined")
        return;
      const cc = document.querySelector("app-control-center, #app-control-center");
      if (!cc)
        return;
      const isOpen = cc.getAttribute("data-open") === "true";
      cc.setAttribute("data-open", (!isOpen).toString());
      if (!isOpen) {
        cc.classList.remove("pointer-events-none", "opacity-0", "-translate-y-full");
        cc.classList.add("pointer-events-auto", "opacity-100", "translate-y-0");
      } else {
        cc.classList.add("pointer-events-none", "opacity-0", "-translate-y-full");
        cc.classList.remove("pointer-events-auto", "opacity-100", "translate-y-0");
      }
    }
    toggleMultitask() {
      LayerFirm.haptic("selection");
      if (typeof document === "undefined")
        return;
      const mt = document.querySelector("app-multitask, #app-multitask");
      if (!mt)
        return;
      const isOpen = mt.getAttribute("data-open") === "true";
      mt.setAttribute("data-open", (!isOpen).toString());
      if (!isOpen) {
        mt.classList.remove("hidden", "opacity-0", "scale-95");
        mt.classList.add("flex", "opacity-100", "scale-100");
      } else {
        mt.classList.add("opacity-0", "scale-95");
        setTimeout(() => {
          mt.classList.remove("flex");
          mt.classList.add("hidden");
        }, 200);
      }
    }
    lock() {
      LayerFirm.haptic("heavy");
      LayerSys.setPowerState("locked");
      if (typeof document === "undefined")
        return;
      const ls = document.querySelector("app-lockscreen, #app-lockscreen");
      if (!ls)
        return;
      ls.classList.remove("hidden", "opacity-0", "pointer-events-none");
      ls.classList.add("opacity-100", "pointer-events-auto");
      ls.setAttribute("data-locked", "true");
    }
    unlock() {
      LayerFirm.haptic("success");
      LayerSys.setPowerState("awake");
      if (typeof document === "undefined")
        return;
      const ls = document.querySelector("app-lockscreen, #app-lockscreen");
      if (!ls)
        return;
      ls.classList.add("opacity-0", "pointer-events-none");
      ls.setAttribute("data-locked", "false");
      setTimeout(() => ls.classList.add("hidden"), 300);
    }
    diag() {
      const spec = VIEWPORT_SPECS[this.currentViewMode];
      let sheetOpen = false;
      let ccOpen = false;
      let mtOpen = false;
      const locked = LayerSys.getPowerState() === "locked";
      if (typeof document !== "undefined") {
        const sheet = document.querySelector('app-bottom-sheet[data-state="open"]');
        sheetOpen = !!sheet;
        const cc = document.querySelector('app-control-center[data-open="true"], #app-control-center[data-open="true"]');
        ccOpen = !!cc;
        const mt = document.querySelector('app-multitask[data-open="true"], #app-multitask[data-open="true"]');
        mtOpen = !!mt;
      }
      const d = {
        layer: "Layer-UX",
        timestamp: Date.now(),
        currentViewMode: this.currentViewMode,
        formFactor: spec.formFactor,
        orientation: spec.orientation,
        viewportDimensions: { width: spec.width, height: spec.height },
        navHistoryDepth: this.history.length,
        activeModals: {
          sheetOpen,
          controlCenterOpen: ccOpen,
          multitaskOpen: mtOpen,
          locked
        },
        registeredCustomElementsCount: 17
      };
      if (typeof window !== "undefined") {
        window.dispatchEvent(new CustomEvent("appz:diag:ux", { detail: d }));
      }
      return d;
    }
  }
  var LayerUX = new LayerPresentation;
  var AppzNavigator = LayerUX;
  var HardwareBridge = LayerFirm;
  function applySegmentLayerStyles(el) {
    const levelAttr = el.getAttribute("level") || el.getAttribute("layer") || "0";
    const level = parseInt(levelAttr, 10);
    const clampedLevel = Math.max(0, Math.min(9, isNaN(level) ? 0 : level));
    el.style.position = el.style.position || "relative";
    el.style.zIndex = String(clampedLevel * 10);
    el.setAttribute("data-resolved-layer", String(clampedLevel));
  }

  class AppzEngine {
    init(root = document) {
      if (typeof document === "undefined")
        return;
      root.querySelectorAll("app-layer, [app-layer], [layer]").forEach((el) => {
        applySegmentLayerStyles(el);
      });
      root.querySelectorAll("[hx-pull-refresh], [app-pull-refresh]").forEach((el) => {
        this.bindPullToRefresh(el);
      });
      root.querySelectorAll("app-bottom-sheet, [app-bottom-sheet]").forEach((el) => {
        this.bindBottomSheet(el);
      });
      root.querySelectorAll("app-joystick, [app-joystick]").forEach((el) => {
        this.bindJoystick(el);
      });
      root.querySelectorAll("app-swipeout, [app-swipeout]").forEach((el) => {
        this.bindSwipeout(el);
      });
      root.querySelectorAll("app-expand-card, [app-expand-card]").forEach((el) => {
        this.bindExpandCard(el);
      });
      root.querySelectorAll("app-gauge, [app-gauge]").forEach((el) => {
        this.bindGauge(el);
      });
      root.querySelectorAll("app-keypad, [app-keypad]").forEach((el) => {
        this.bindKeypad(el);
      });
      root.querySelectorAll("app-statusbar, [app-statusbar]").forEach((el) => {
        this.bindStatusBar(el);
      });
      root.querySelectorAll("app-thermostat, [app-thermostat]").forEach((el) => {
        this.bindThermostat(el);
      });
      root.querySelectorAll("[data-view-btn]").forEach((btn) => {
        btn.addEventListener("click", () => {
          const mode = btn.getAttribute("data-view-btn");
          if (mode)
            LayerUX.setViewMode(mode);
        });
      });
      if (typeof window !== "undefined" && !window.__appzResizeBound) {
        window.__appzResizeBound = true;
        window.addEventListener("resize", () => {
          LayerUX.setViewMode(LayerUX.currentViewMode);
        });
      }
    }
    bindPullToRefresh(el) {
      let startY = 0;
      let currentY = 0;
      const threshold = 65;
      const refreshUrl = el.getAttribute("hx-pull-refresh") || el.getAttribute("app-pull-refresh") || "";
      el.addEventListener("touchstart", (e) => {
        if (el.scrollTop === 0) {
          startY = e.touches[0].clientY;
        }
      }, { passive: true });
      el.addEventListener("touchmove", (e) => {
        if (startY === 0)
          return;
        currentY = e.touches[0].clientY;
        const pullDist = Math.max(0, currentY - startY);
        if (pullDist > 0 && pullDist < 120) {
          el.style.transform = `translateY(${pullDist * 0.4}px)`;
          if (pullDist > threshold)
            LayerFirm.haptic("selection");
        }
      }, { passive: true });
      el.addEventListener("touchend", () => {
        const pullDist = Math.max(0, currentY - startY);
        startY = 0;
        currentY = 0;
        el.style.transform = "";
        if (pullDist >= threshold) {
          LayerFirm.haptic("success");
          if (refreshUrl && typeof window !== "undefined" && window.htmx) {
            window.htmx.ajax("GET", refreshUrl, { target: el, swap: "innerHTML" });
          }
        }
      });
    }
    bindBottomSheet(el) {
      let startY = 0;
      el.style.transition = "transform 0.25s cubic-bezier(0.16, 1, 0.3, 1)";
      const handle = el.querySelector(".drag-handle") || el.querySelector(".sheet-handle") || el;
      handle.addEventListener("touchstart", (e) => {
        startY = e.touches[0].clientY;
      }, { passive: true });
      handle.addEventListener("touchmove", (e) => {
        const deltaY = e.touches[0].clientY - startY;
        if (deltaY > 0) {
          el.style.transform = `translateY(${deltaY}px)`;
        }
      }, { passive: true });
      handle.addEventListener("touchend", (e) => {
        const deltaY = (e.changedTouches ? e.changedTouches[0].clientY : 0) - startY;
        if (deltaY > 90) {
          LayerUX.hideSheet(`#${el.id}`);
        } else {
          el.style.transform = "translateY(0%)";
        }
      });
    }
    bindJoystick(el) {
      const channelName = el.getAttribute("hx-chan-send") || el.getAttribute("channel") || "game:joystick";
      const ch = getAppzChannel(channelName);
      const thumb = el.querySelector(".joystick-thumb") || el.children[0];
      let active = false;
      const handleMove = (clientX, clientY) => {
        const rect = el.getBoundingClientRect();
        const centerX = rect.left + rect.width / 2;
        const centerY = rect.top + rect.height / 2;
        const maxRadius = rect.width / 2;
        let dx = clientX - centerX;
        let dy = clientY - centerY;
        const distance = Math.sqrt(dx * dx + dy * dy);
        if (distance > maxRadius) {
          dx = dx / distance * maxRadius;
          dy = dy / distance * maxRadius;
        }
        if (thumb) {
          thumb.style.transform = `translate(${dx}px, ${dy}px)`;
        }
        const normX = parseFloat((dx / maxRadius).toFixed(2));
        const normY = parseFloat((dy / maxRadius).toFixed(2));
        ch.send({ x: normX, y: normY });
      };
      el.addEventListener("pointerdown", (e) => {
        active = true;
        LayerFirm.haptic("selection");
        handleMove(e.clientX, e.clientY);
      });
      window.addEventListener("pointermove", (e) => {
        if (active)
          handleMove(e.clientX, e.clientY);
      });
      window.addEventListener("pointerup", () => {
        if (active) {
          active = false;
          if (thumb)
            thumb.style.transform = "translate(0px, 0px)";
          ch.send({ x: 0, y: 0 });
        }
      });
    }
    bindSwipeout(el) {
      let startX = 0;
      let currentX = 0;
      const content = el.querySelector(".swipeout-content") || el;
      const actionsRight = el.querySelector(".swipeout-actions-right");
      const actionsLeft = el.querySelector(".swipeout-actions-left");
      content.style.transition = "transform 0.2s cubic-bezier(0.16, 1, 0.3, 1)";
      el.addEventListener("touchstart", (e) => {
        startX = e.touches[0].clientX;
      }, { passive: true });
      el.addEventListener("touchmove", (e) => {
        if (startX === 0)
          return;
        currentX = e.touches[0].clientX;
        const diffX = currentX - startX;
        const damped = diffX * 0.75;
        content.style.transform = `translateX(${damped}px)`;
      }, { passive: true });
      el.addEventListener("touchend", (e) => {
        const diffX = (e.changedTouches ? e.changedTouches[0].clientX : currentX) - startX;
        startX = 0;
        if (diffX < -70 && actionsRight) {
          content.style.transform = `translateX(-120px)`;
          LayerFirm.haptic("medium");
        } else if (diffX > 70 && actionsLeft) {
          content.style.transform = `translateX(120px)`;
          LayerFirm.haptic("medium");
        } else {
          content.style.transform = "translateX(0px)";
        }
      });
    }
    bindExpandCard(el) {
      el.addEventListener("click", () => {
        const isExpanded = el.getAttribute("data-expanded") === "true";
        LayerFirm.haptic("medium");
        if (!isExpanded) {
          el.setAttribute("data-expanded", "true");
          el.classList.add("fixed", "inset-4", "z-50", "shadow-2xl");
        } else {
          el.setAttribute("data-expanded", "false");
          el.classList.remove("fixed", "inset-4", "z-50", "shadow-2xl");
        }
      });
    }
    bindGauge(el) {
      const val = parseFloat(el.getAttribute("value") || "0");
      const min = parseFloat(el.getAttribute("min") || "0");
      const max = parseFloat(el.getAttribute("max") || "100");
      const unit = el.getAttribute("unit") || "";
      const label = el.getAttribute("label") || "";
      const pct = Math.max(0, Math.min(1, (val - min) / (max - min)));
      const radius = 38;
      const circumference = 2 * Math.PI * radius;
      const strokeDashoffset = circumference * (1 - pct * 0.75);
      el.innerHTML = `
      <div class="relative flex flex-col items-center justify-center">
        <svg class="w-24 h-24 transform -rotate-135" viewBox="0 0 100 100">
          <circle cx="50" cy="50" r="${radius}" stroke="currentColor" stroke-width="8" fill="transparent" class="text-slate-800" stroke-dasharray="${circumference * 0.75} ${circumference * 0.25}" />
          <circle cx="50" cy="50" r="${radius}" stroke="url(#gauge-gradient)" stroke-width="8" stroke-linecap="round" fill="transparent"
            stroke-dasharray="${circumference}" stroke-dashoffset="${strokeDashoffset}" class="transition-all duration-700 ease-out" />
          <defs>
            <linearGradient id="gauge-gradient" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stop-color="#38bdf8" />
              <stop offset="50%" stop-color="#6366f1" />
              <stop offset="100%" stop-color="#ec4899" />
            </linearGradient>
          </defs>
        </svg>
        <div class="absolute inset-0 flex flex-col items-center justify-center pointer-events-none">
          <span class="text-lg font-black text-white leading-none">${val}<span class="text-[10px] font-bold text-slate-400 ml-0.5">${unit}</span></span>
          ${label ? `<span class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">${label}</span>` : ""}
        </div>
      </div>
    `;
    }
    bindKeypad(el) {
      const targetInput = el.getAttribute("target") || "#pin-display";
      const maxDigits = parseInt(el.getAttribute("digits") || "4", 10);
      let enteredPin = "";
      const updateDisplay = () => {
        const display = document.querySelector(targetInput);
        if (display) {
          display.innerHTML = Array.from({ length: maxDigits }).map((_, i) => `
          <div class="w-3 h-3 rounded-full border-2 border-indigo-500/60 ${i < enteredPin.length ? "bg-indigo-500 scale-110 shadow-[0_0_8px_rgba(99,102,241,0.8)]" : "bg-transparent"} transition-all"></div>
        `).join("");
        }
        if (enteredPin.length === maxDigits) {
          LayerFirm.haptic("success");
          el.dispatchEvent(new CustomEvent("appz:pin-submit", { detail: { pin: enteredPin } }));
          if (enteredPin === (el.getAttribute("valid-pin") || "1234")) {
            LayerUX.unlock();
            enteredPin = "";
            updateDisplay();
          } else {
            LayerFirm.haptic("error");
            setTimeout(() => {
              enteredPin = "";
              updateDisplay();
            }, 300);
          }
        }
      };
      el.querySelectorAll("[data-key]").forEach((btn) => {
        btn.addEventListener("click", () => {
          const key = btn.getAttribute("data-key");
          LayerFirm.haptic("selection");
          if (key === "clear") {
            enteredPin = "";
          } else if (key === "backspace") {
            enteredPin = enteredPin.slice(0, -1);
          } else if (key && enteredPin.length < maxDigits) {
            enteredPin += key;
          }
          updateDisplay();
        });
      });
      updateDisplay();
    }
    bindStatusBar(el) {
      const timeEl = el.querySelector(".status-time");
      const updateTime = () => {
        if (timeEl) {
          const clock = LayerSys.clock();
          timeEl.textContent = clock.time;
        }
      };
      updateTime();
      setInterval(updateTime, 1000);
      const islandCh = getAppzChannel("iot:island");
      const island = el.querySelector(".dynamic-island");
      if (island) {
        islandCh.recv((event) => {
          if (event && event.text) {
            LayerFirm.haptic("medium");
            island.classList.add("scale-105", "bg-indigo-950", "border-indigo-500/50");
            island.setAttribute("title", event.text);
            setTimeout(() => {
              island.classList.remove("scale-105", "bg-indigo-950", "border-indigo-500/50");
            }, 2500);
          }
        });
      }
    }
    bindThermostat(el) {
      let targetTemp = parseFloat(el.getAttribute("temp") || "21.5");
      const tempDisplay = el.querySelector(".thermostat-temp");
      const updateTemp = (newTemp) => {
        targetTemp = Math.round(newTemp * 10) / 10;
        if (tempDisplay)
          tempDisplay.textContent = targetTemp.toFixed(1);
        LayerFirm.haptic("selection");
        getAppzChannel("iot:thermostat").send({ targetTemp });
      };
      el.querySelector(".thermostat-plus")?.addEventListener("click", () => updateTemp(targetTemp + 0.5));
      el.querySelector(".thermostat-minus")?.addEventListener("click", () => updateTemp(targetTemp - 0.5));
    }
  }
  var Appz = new AppzEngine;
  function runMasterDiagnostics() {
    const report = {
      firm: LayerFirm.diag(),
      sys: LayerSys.diag(),
      srv: LayerSrv.diag(),
      ux: LayerUX.diag(),
      timestamp: Date.now()
    };
    if (typeof window !== "undefined") {
      window.dispatchEvent(new CustomEvent("appz:diag", { detail: report }));
    }
    return report;
  }
  if (typeof window !== "undefined") {
    window.Appz = Appz;
    window.$appz = {
      firm: LayerFirm,
      sys: LayerSys,
      srv: LayerSrv,
      ux: LayerUX,
      diag: runMasterDiagnostics,
      navigator: LayerUX,
      push: LayerUX.push.bind(LayerUX),
      pop: LayerUX.pop.bind(LayerUX),
      sheet: LayerUX.showSheet.bind(LayerUX),
      sheetClose: LayerUX.hideSheet.bind(LayerUX),
      setViewMode: LayerUX.setViewMode.bind(LayerUX),
      toggleControlCenter: LayerUX.toggleControlCenter.bind(LayerUX),
      toggleMultitask: LayerUX.toggleMultitask.bind(LayerUX),
      lock: LayerUX.lock.bind(LayerUX),
      unlock: LayerUX.unlock.bind(LayerUX),
      haptic: LayerFirm.haptic.bind(LayerFirm),
      chan: getAppzChannel,
      send(name, val) {
        getAppzChannel(name).send(val);
      },
      recv(name, cb) {
        return getAppzChannel(name).recv(cb);
      },
      specs: VIEWPORT_SPECS
    };
    const defineCE = (tag) => {
      if (!customElements.get(tag)) {
        customElements.define(tag, class extends HTMLElement {
        });
      }
    };
    [
      "app-shellui",
      "app-bar",
      "app-seg",
      "app-layer",
      "app-navbar",
      "app-bottom-sheet",
      "app-joystick",
      "app-statusbar",
      "app-control-center",
      "app-launcher",
      "app-multitask",
      "app-lockscreen",
      "app-keypad",
      "app-gauge",
      "app-swipeout",
      "app-expand-card",
      "app-thermostat"
    ].forEach(defineCE);
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", () => Appz.init());
    } else {
      Appz.init();
    }
  }
})();
