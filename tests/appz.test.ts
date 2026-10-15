import { describe, test, expect } from 'bun:test';
import { 
  LayerFirm, 
  LayerSys, 
  LayerSrv, 
  LayerUX, 
  runMasterDiagnostics,
  applySegmentLayerStyles, 
  VIEWPORT_SPECS,
  type AppzViewMode
} from '../src/appz';

describe('Appz 4-Layer OS Stack & @diag Telemetry Suite', () => {

  test('1. VIEWPORT_SPECS contains all 8 required distinct modes', () => {
    const requiredModes: AppzViewMode[] = [
      'OS-H', 'OS-P',
      'Tab-H', 'Tab-P',
      'Mob-H', 'Mob-P',
      'TV-H', 'TV-P'
    ];

    requiredModes.forEach(mode => {
      expect(VIEWPORT_SPECS[mode]).toBeDefined();
      expect(VIEWPORT_SPECS[mode].mode).toBe(mode);
      expect(['OS', 'Tab', 'Mob', 'TV']).toContain(VIEWPORT_SPECS[mode].formFactor);
      expect(['H', 'P']).toContain(VIEWPORT_SPECS[mode].orientation);
      expect(VIEWPORT_SPECS[mode].width).toBeGreaterThan(0);
      expect(VIEWPORT_SPECS[mode].height).toBeGreaterThan(0);
    });
  });

  test('2. Dimensions match ergonomic device standards across all 8 modes', () => {
    expect(VIEWPORT_SPECS['OS-H'].width).toBe(1024);
    expect(VIEWPORT_SPECS['OS-H'].height).toBe(600);
    expect(VIEWPORT_SPECS['OS-P'].width).toBe(600);
    expect(VIEWPORT_SPECS['OS-P'].height).toBe(1024);

    expect(VIEWPORT_SPECS['Tab-H'].width).toBe(1180);
    expect(VIEWPORT_SPECS['Tab-H'].height).toBe(820);
    expect(VIEWPORT_SPECS['Tab-P'].width).toBe(820);
    expect(VIEWPORT_SPECS['Tab-P'].height).toBe(1180);

    expect(VIEWPORT_SPECS['Mob-H'].width).toBe(844);
    expect(VIEWPORT_SPECS['Mob-H'].height).toBe(390);
    expect(VIEWPORT_SPECS['Mob-P'].width).toBe(390);
    expect(VIEWPORT_SPECS['Mob-P'].height).toBe(844);

    expect(VIEWPORT_SPECS['TV-H'].width).toBe(1280);
    expect(VIEWPORT_SPECS['TV-H'].height).toBe(720);
    expect(VIEWPORT_SPECS['TV-P'].width).toBe(720);
    expect(VIEWPORT_SPECS['TV-P'].height).toBe(1280);
  });

  test('3. Layer-Firm: Hardware Abstraction & Layer-Specific @diag', async () => {
    LayerFirm.haptic('selection');
    LayerFirm.haptic('success');

    const batt = await LayerFirm.battery();
    expect(batt.level).toBeGreaterThan(0);
    expect(typeof batt.charging).toBe('boolean');

    const net = LayerFirm.network();
    expect(typeof net.online).toBe('boolean');
    expect(net.rssiBars).toBeGreaterThanOrEqual(0);

    const sensors = LayerFirm.sensors();
    expect(sensors.lightLux).toBeGreaterThan(0);
    expect(sensors.cpuTemp).toBeGreaterThan(0);

    // Layer-Firm specific @diag telemetry
    const diag = LayerFirm.diag();
    expect(diag.layer).toBe('Layer-Firm');
    expect(diag.hapticTriggerCount).toBeGreaterThanOrEqual(2);
    expect(diag.battery).toBeDefined();
    expect(diag.sensors).toBeDefined();
    expect(diag.network).toBeDefined();
  });

  test('4. Layer-Sys: System APIs, IPC Channels & Layer-Specific @diag', () => {
    const ch = LayerSys.chan<{ metric: string; val: number }>('sys:cpu');
    const received: number[] = [];

    const unsub = ch.recv(e => received.push(e.val));
    ch.send({ metric: 'cpu', val: 12 });
    ch.send({ metric: 'cpu', val: 18 });
    expect(received).toEqual([12, 18]);
    unsub();

    LayerSys.setPowerState('ambient');
    expect(LayerSys.getPowerState()).toBe('ambient');
    LayerSys.setPowerState('awake');

    LayerSys.storage.set('brightness', 85);
    expect(LayerSys.storage.get('brightness')).toBe(85);

    // Layer-Sys specific @diag telemetry
    const diag = LayerSys.diag();
    expect(diag.layer).toBe('Layer-Sys');
    expect(diag.channelsCount).toBeGreaterThan(0);
    expect(diag.activeChannels).toContain('sys:cpu');
    expect(diag.clock.time).toBeDefined();
    expect(diag.uptimeSec).toBeGreaterThanOrEqual(0);
  });

  test('5. Layer-Srv: Service Daemons, Alarms & Layer-Specific @diag', () => {
    const alert = LayerSrv.alarms.add({
      severity: 'critical',
      title: 'Emergency Generator Coolant Low',
      source: 'Generator-02'
    });
    expect(alert.id).toBeDefined();
    expect(alert.acknowledged).toBe(false);

    LayerSrv.alarms.ack(alert.id);
    const updated = LayerSrv.alarms.list().find(a => a.id === alert.id);
    expect(updated?.acknowledged).toBe(true);

    const ota = LayerSrv.ota.getStatus();
    expect(ota.currentVersion).toBeDefined();

    // Layer-Srv specific @diag telemetry
    const diag = LayerSrv.diag();
    expect(diag.layer).toBe('Layer-Srv');
    expect(diag.totalAlerts).toBeGreaterThanOrEqual(1);
    expect(diag.otaCurrentVersion).toBe(ota.currentVersion);
  });

  test('6. Layer-UX: Presentation Ergonomics, Layers 0-9 & Layer-Specific @diag', () => {
    LayerUX.setViewMode('OS-H');
    expect(LayerUX.currentViewMode).toBe('OS-H');

    // 0-9 Layer Stacking
    const mockEl = {
      attributes: { level: '7' } as Record<string, string>,
      style: {} as Record<string, string>,
      getAttribute(k: string) { return this.attributes[k]; },
      setAttribute(k: string, v: string) { this.attributes[k] = v; }
    } as any;

    applySegmentLayerStyles(mockEl);
    expect(mockEl.style.zIndex).toBe('70');
    expect(mockEl.getAttribute('data-resolved-layer')).toBe('7');

    // Layer-UX specific @diag telemetry
    const diag = LayerUX.diag();
    expect(diag.layer).toBe('Layer-UX');
    expect(diag.currentViewMode).toBe('OS-H');
    expect(diag.formFactor).toBe('OS');
    expect(diag.orientation).toBe('H');
    expect(diag.viewportDimensions.width).toBe(1024);
    expect(diag.viewportDimensions.height).toBe(600);
    expect(diag.activeModals).toBeDefined();
  });

  test('7. Master @diag across all 4 layers simultaneously', () => {
    const report = runMasterDiagnostics();
    expect(report.firm.layer).toBe('Layer-Firm');
    expect(report.sys.layer).toBe('Layer-Sys');
    expect(report.srv.layer).toBe('Layer-Srv');
    expect(report.ux.layer).toBe('Layer-UX');
    expect(report.timestamp).toBeGreaterThan(0);
  });
});
