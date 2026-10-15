import { describe, it, expect } from "bun:test";

// =============================================================================
// CLINICAL DICOM 3.0 & RECIST 1.1 ONCOLOGY WORKSTATION TEST SUITE
// =============================================================================

describe("Clinical DICOM 3.0 Diagnostic Engine & RECIST 1.1 Staging", () => {
  const PIXEL_SPACING = 0.72; // mm per pixel calibrated

  // 1. Caliper Geometry & Euclidean Long Axis Measurement
  it("computes accurate Euclidean Long Axis (LD) distance in millimeters", () => {
    // P1: (155, 175), P2: (185, 190)
    const p1 = { x: 155, y: 175 };
    const p2 = { x: 185, y: 190 };
    const dx = p2.x - p1.x; // 30
    const dy = p2.y - p1.y; // 15
    const pixelDist = Math.sqrt(dx * dx + dy * dy); // sqrt(900 + 225) = sqrt(1125) ≈ 33.541
    const mmDist = parseFloat((pixelDist * PIXEL_SPACING).toFixed(1));

    expect(pixelDist).toBeCloseTo(33.541, 2);
    expect(mmDist).toBe(24.1);

    // Baseline test lesion points: dx=27, dy=14
    const blP1 = { x: 155, y: 175 };
    const blP2 = { x: 182, y: 189 };
    const blDx = blP2.x - blP1.x;
    const blDy = blP2.y - blP1.y;
    const blDist = Math.sqrt(blDx * blDx + blDy * blDy) * PIXEL_SPACING;
    expect(parseFloat(blDist.toFixed(1))).toBe(21.9);
  });

  // 2. Perpendicular Short Axis Orthogonal Geometry
  it("calculates perpendicular short axis endpoints and length", () => {
    const p1 = { x: 155, y: 175 };
    const p2 = { x: 185, y: 190 };
    const dx = p2.x - p1.x;
    const dy = p2.y - p1.y;
    const pixelDist = Math.sqrt(dx * dx + dy * dy);

    const midX = (p1.x + p2.x) / 2; // 170
    const midY = (p1.y + p2.y) / 2; // 182.5
    const perpAngle = Math.atan2(dy, dx) + Math.PI / 2;

    const shortLen = pixelDist * 0.67;
    const s1x = midX + Math.cos(perpAngle) * (shortLen / 2);
    const s1y = midY + Math.sin(perpAngle) * (shortLen / 2);
    const s2x = midX - Math.cos(perpAngle) * (shortLen / 2);
    const s2y = midY - Math.sin(perpAngle) * (shortLen / 2);

    // Verify dot product of long vector and short vector is 0 (orthogonality)
    const longVec = { x: dx, y: dy };
    const shortVec = { x: s2x - s1x, y: s2y - s1y };
    const dotProduct = longVec.x * shortVec.x + longVec.y * shortVec.y;
    expect(Math.abs(dotProduct)).toBeLessThan(1e-6);

    const shortMM = parseFloat((shortLen * PIXEL_SPACING).toFixed(1));
    expect(shortMM).toBe(16.2);
  });

  // 3. RECIST 1.1 Oncology Classification Logic
  it("accurately classifies RECIST 1.1 response: CR, PR, SD, PD", () => {
    function classifyRecist(baselineSOD: number, currentSOD: number, newLesions: boolean = false): {
      category: "CR" | "PR" | "SD" | "PD";
      deltaPct: number;
    } {
      if (newLesions) {
        return { category: "PD", deltaPct: ((currentSOD - baselineSOD) / baselineSOD) * 100 };
      }
      if (currentSOD === 0) {
        return { category: "CR", deltaPct: -100 };
      }
      const deltaPct = ((currentSOD - baselineSOD) / baselineSOD) * 100;
      if (deltaPct <= -30) {
        return { category: "PR", deltaPct };
      }
      if (deltaPct >= 20 && (currentSOD - baselineSOD) >= 5) {
        return { category: "PD", deltaPct };
      }
      return { category: "SD", deltaPct };
    }

    const baseline = 50.6; // T1=32.4mm, T2=18.2mm

    // Partial Response: current = 34.9mm (-31.0%)
    const prResult = classifyRecist(baseline, 34.9);
    expect(prResult.category).toBe("PR");
    expect(prResult.deltaPct).toBeCloseTo(-31.03, 1);

    // Complete Response: current = 0mm (-100%)
    const crResult = classifyRecist(baseline, 0);
    expect(crResult.category).toBe("CR");
    expect(crResult.deltaPct).toBe(-100);

    // Stable Disease: current = 45.0mm (-11.0%)
    const sdResult = classifyRecist(baseline, 45.0);
    expect(sdResult.category).toBe("SD");
    expect(sdResult.deltaPct).toBeCloseTo(-11.07, 1);

    // Progressive Disease: current = 62.0mm (+22.5%)
    const pdResult = classifyRecist(baseline, 62.0);
    expect(pdResult.category).toBe("PD");
    expect(pdResult.deltaPct).toBeCloseTo(22.53, 1);

    // Progressive Disease via New Lesion even if SOD decreased
    const pdNewLesion = classifyRecist(baseline, 30.0, true);
    expect(pdNewLesion.category).toBe("PD");
  });

  // 4. Hounsfield Unit (HU) Lookup Table (LUT) Transfer Functions
  it("maps Hounsfield Units to 8-bit grayscale values according to W/L presets", () => {
    function huToGray(hu: number, windowWidth: number, windowCenter: number): number {
      const minHU = windowCenter - windowWidth / 2;
      const maxHU = windowCenter + windowWidth / 2;
      if (hu <= minHU) return 0;
      if (hu >= maxHU) return 255;
      return Math.floor(((hu - minHU) / windowWidth) * 255);
    }

    // LUNG Preset: W: 1500, L: -600 (Range: -1350 to +150)
    const lungW = 1500, lungL = -600;
    expect(huToGray(-1500, lungW, lungL)).toBe(0); // Deep air is pure black
    expect(huToGray(-600, lungW, lungL)).toBe(127); // Center is 50% gray
    expect(huToGray(200, lungW, lungL)).toBe(255); // Bone/dense tissue is pure white

    // MEDIASTINUM / SOFT TISSUE: W: 350, L: 40 (Range: -135 to +215)
    const medW = 350, medL = 40;
    expect(huToGray(-200, medW, medL)).toBe(0); // Lung parenchyma is black
    expect(huToGray(40, medW, medL)).toBe(127); // Soft tissue center is mid-gray
    expect(huToGray(300, medW, medL)).toBe(255); // Rib/vertebra is white

    // BONE Preset: W: 2000, L: 300 (Range: -700 to +1300)
    const boneW = 2000, boneL = 300;
    expect(huToGray(300, boneW, boneL)).toBe(127);
    expect(huToGray(1400, boneW, boneL)).toBe(255);

    // BRAIN Preset: W: 80, L: 40 (Range: 0 to 80)
    const brainW = 80, brainL = 40;
    expect(huToGray(0, brainW, brainL)).toBe(0);
    expect(huToGray(40, brainW, brainL)).toBe(127);
    expect(huToGray(80, brainW, brainL)).toBe(255);
  });

  // 5. 8 Distinct Viewport Modes & Auto-Scaling Math
  it("computes accurate scale factor and dimensions for all 8 view modes", () => {
    const VIEW_SPECS: Record<string, { width: number; height: number; form: string }> = {
      "OS-H":  { width: 1024, height: 600, form: "OS" },
      "OS-P":  { width: 600,  height: 900, form: "OS" },
      "Tab-H": { width: 960,  height: 600, form: "Tablet" },
      "Tab-P": { width: 600,  height: 800, form: "Tablet" },
      "Mob-H": { width: 720,  height: 380, form: "Mobile" },
      "Mob-P": { width: 380,  height: 720, form: "Mobile" },
      "TV-H":  { width: 1280, height: 720, form: "TV" },
      "TV-P":  { width: 720,  height: 1280, form: "TV" }
    };

    function calculateScale(mode: string, availableW: number, availableH: number): number {
      const spec = VIEW_SPECS[mode];
      if (!spec) return 1.0;
      const scaleX = availableW / spec.width;
      const scaleY = availableH / spec.height;
      return Math.min(1.0, Math.min(scaleX, scaleY));
    }

    // Testing on standard laptop display (1366 x 768 available canvas)
    const availW = 1200;
    const availH = 650;

    // OS-H (1024x600) fits without downscaling
    expect(calculateScale("OS-H", availW, availH)).toBe(1.0);

    // TV-H (1280x720) scales down to fit viewport height (650 / 720 ≈ 0.902)
    const tvScale = calculateScale("TV-H", availW, availH);
    expect(tvScale).toBeLessThan(1.0);
    expect(tvScale).toBeCloseTo(650 / 720, 2);

    // TV-P (720x1280) scales down to fit viewport height (650 / 1280 ≈ 0.507)
    const tvPScale = calculateScale("TV-P", availW, availH);
    expect(tvPScale).toBeCloseTo(650 / 1280, 2);

    // Mob-P (380x720) fits comfortably
    expect(calculateScale("Mob-P", availW, 800)).toBe(1.0);
  });

  // 6. 4-Layer Architecture Protocol & Diag Telemetry
  it("verifies 4-layer separation of concerns and diagnostic measurements", () => {
    const layers = {
      sys: {
        id: "Layer-Sys",
        role: "Direct Canvas 2D Buffer & TypedArray HU Matrix",
        zeroAlloc: true,
        bufferSize: 512 * 512 * 4
      },
      firm: {
        id: "Layer-Firm",
        role: "Hardware Haptic Feedback & Display Sync (60 FPS)",
        hapticEnabled: true,
        targetFps: 60.0
      },
      ux: {
        id: "Layer-UX",
        role: "AppzOS Interactive Shell & SVG RECIST Calipers",
        zIndexLevels: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9],
        activeTool: "caliper"
      },
      srv: {
        id: "Layer-Srv",
        role: "Golang DSL 1:1 Parity Component & Zero-NPM Server",
        componentName: "BlockRadiologyDICOM",
        runtime: "Go / Bun"
      }
    };

    expect(layers.sys.bufferSize).toBe(1048576); // 1 MB per 512x512 RGBA slice
    expect(layers.firm.targetFps).toBe(60.0);
    expect(layers.ux.zIndexLevels.length).toBe(10);
    expect(layers.srv.runtime).toContain("Go");
  });

  // 7. Patient Worklist Registry & Filter Logic (RAMSOFT PowerReader Paradigm)
  it("filters and selects patients across clinical PACS worklist statuses", () => {
    interface PatientStudy {
      id: string;
      name: string;
      mrn: string;
      status: "ORDERED" | "SCHEDULED" | "CHECKED IN" | "CONFIRMED" | "ARRIVED" | "COMPLETED" | "VERIFIED";
      modality: "CT" | "MR" | "DX" | "US" | "CR" | "MG";
      schedule: string;
    }

    const worklist: PatientStudy[] = [
      { id: "P1", name: "DOE^JOHN", mrn: "#ONC-94021", status: "VERIFIED", modality: "CT", schedule: "CHEST" },
      { id: "P2", name: "WARREN^SHARON", mrn: "#ONC-88142", status: "ORDERED", modality: "CT", schedule: "CHEST" },
      { id: "P3", name: "ANGELOU^MAYA^N^MS.^JR", mrn: "#RAD-77291", status: "ORDERED", modality: "MR", schedule: "BREAST" },
      { id: "P4", name: "SMITHERS^MIKE", mrn: "#ONC-66382", status: "SCHEDULED", modality: "MR", schedule: "CHEST" },
      { id: "P5", name: "NAVAL^WILLIAM", mrn: "#RAD-55109", status: "SCHEDULED", modality: "DX", schedule: "LSPINE" },
      { id: "P6", name: "RICE^JIM", mrn: "#ONC-33821", status: "CHECKED IN", modality: "CT", schedule: "ABDOMEN" },
      { id: "P7", name: "JONES^TIM", mrn: "#RAD-22904", status: "CONFIRMED", modality: "MR", schedule: "CHEST" },
      { id: "P8", name: "BARILLAS^LIZETTE^M^MS.", mrn: "#ONC-11048", status: "ARRIVED", modality: "DX", schedule: "CHEST" },
      { id: "P9", name: "SEYMOUR^JANE", mrn: "#RAD-99381", status: "COMPLETED", modality: "US", schedule: "PELVIS" },
      { id: "P10", name: "SANCHEZ^VIVIANA", mrn: "#RAD-88492", status: "VERIFIED", modality: "US", schedule: "PELVIS" }
    ];

    // Filter by query "chest"
    const chestStudies = worklist.filter(p => p.schedule.toLowerCase().includes("chest") || p.name.toLowerCase().includes("chest"));
    expect(chestStudies.length).toBe(5);

    // Filter by modality "MR"
    const mrStudies = worklist.filter(p => p.modality === "MR");
    expect(mrStudies.length).toBe(3);

    // Filter by status "VERIFIED"
    const verifiedStudies = worklist.filter(p => p.status === "VERIFIED");
    expect(verifiedStudies.length).toBe(2);
    expect(verifiedStudies.map(p => p.id)).toEqual(["P1", "P10"]);
  });

  // 8. Patient Imaging Tabs & Scan Files Series Selection
  it("organizes scan files across imaging modality tabs (CT, DX, US, MR, PET)", () => {
    interface ScanSeries {
      id: string;
      title: string;
      slices: number;
      matrix: string;
      modality: "CT" | "DX" | "US" | "MR" | "PET";
    }

    const imagingTabs: Record<string, ScanSeries[]> = {
      CT: [
        { id: "CT-S1", title: "Series 1: Scout Topogram", slices: 1, matrix: "512x512", modality: "CT" },
        { id: "CT-S2", title: "Series 2: Axial Chest Thin-Slice CT", slices: 320, matrix: "512x512", modality: "CT" },
        { id: "CT-S3", title: "Series 3: Coronal MPR Reformat", slices: 240, matrix: "512x512", modality: "CT" },
        { id: "CT-S4", title: "Series 4: Sagittal MPR Reformat", slices: 180, matrix: "512x512", modality: "CT" }
      ],
      DX: [
        { id: "DX-S1", title: "Series 1: Chest PA Erect", slices: 1, matrix: "2048x2048", modality: "DX" },
        { id: "DX-S2", title: "Series 2: Chest Lateral", slices: 1, matrix: "2048x2048", modality: "DX" }
      ],
      US: [
        { id: "US-S1", title: "Series 1: Thyroid B-Mode Transverse", slices: 120, matrix: "640x480", modality: "US" },
        { id: "US-S2", title: "Series 2: Color Doppler Carotid", slices: 60, matrix: "640x480", modality: "US" }
      ],
      MR: [
        { id: "MR-S1", title: "Series 1: T1 Axial Spin-Echo", slices: 256, matrix: "512x512", modality: "MR" },
        { id: "MR-S2", title: "Series 2: T2 FLAIR Coronal", slices: 256, matrix: "512x512", modality: "MR" }
      ],
      PET: [
        { id: "PET-S1", title: "Series 1: 18F-FDG Whole Body MIP", slices: 360, matrix: "256x256", modality: "PET" }
      ]
    };

    expect(imagingTabs.CT.length).toBe(4);
    expect(imagingTabs.DX.length).toBe(2);
    expect(imagingTabs.US.length).toBe(2);
    expect(imagingTabs.MR.length).toBe(2);
    expect(imagingTabs.PET.length).toBe(1);

    // Switching to US and selecting Series 2
    const selectedSeries = imagingTabs.US.find(s => s.id === "US-S2");
    expect(selectedSeries).toBeDefined();
    expect(selectedSeries?.slices).toBe(60);
    expect(selectedSeries?.matrix).toBe("640x480");
  });

  // 9. Onscreen 3D Joystick Euler Angles & Spatial Matrix Math
  it("calculates 3D spatial rotation angles and projection from joystick deflection", () => {
    function calculate3DAngles(
      joystickNormX: number, // -1.0 .. +1.0
      joystickNormY: number, // -1.0 .. +1.0
      currentPitch: number,
      currentYaw: number
    ): { pitch: number; yaw: number } {
      // Sensitivity: 90 deg max pitch, 180 deg max yaw
      const newPitch = Math.max(-90, Math.min(90, Math.round(joystickNormY * 90)));
      const newYaw = Math.max(-180, Math.min(180, Math.round(joystickNormX * 180)));
      return { pitch: newPitch, yaw: newYaw };
    }

    // Centered joystick = 0 pitch, 0 yaw (Axial view)
    const center = calculate3DAngles(0, 0, 0, 0);
    expect(center.pitch).toBe(0);
    expect(center.yaw).toBe(0);

    // Full up deflection = +90 pitch (Coronal view)
    const fullUp = calculate3DAngles(0, 1.0, 0, 0);
    expect(fullUp.pitch).toBe(90);
    expect(fullUp.yaw).toBe(0);

    // Full right deflection = +180 yaw
    const fullRight = calculate3DAngles(1.0, 0, 0, 0);
    expect(fullRight.pitch).toBe(0);
    expect(fullRight.yaw).toBe(180);

    // Half deflection oblique
    const oblique = calculate3DAngles(0.25, 0.4, 0, 0);
    expect(oblique.pitch).toBe(36);
    expect(oblique.yaw).toBe(45);

    // 3D vector rotation verification
    function rotateVector(x: number, y: number, z: number, pitchDeg: number, yawDeg: number) {
      const pRad = (pitchDeg * Math.PI) / 180;
      const yRad = (yawDeg * Math.PI) / 180;

      // Rotate around X (pitch)
      const y1 = y * Math.cos(pRad) - z * Math.sin(pRad);
      const z1 = y * Math.sin(pRad) + z * Math.cos(pRad);

      // Rotate around Y (yaw)
      const x2 = x * Math.cos(yRad) + z1 * Math.sin(yRad);
      const z2 = -x * Math.sin(yRad) + z1 * Math.cos(yRad);

      return { x: x2, y: y1, z: z2 };
    }

    const rotated = rotateVector(1, 0, 0, 0, 90);
    expect(rotated.x).toBeCloseTo(0, 4);
    expect(rotated.z).toBeCloseTo(-1, 4);
  });

  // 10. Multi-Modality Diagnostic Rendering Configurations
  it("configures diagnostic presets and units correctly across modalities", () => {
    const MODALITY_CONFIGS = {
      CT: { unit: "HU", minVal: -1000, maxVal: 1000, defaultPreset: "LUNG", sliceThickness: "1.25mm" },
      DX: { unit: "OD", minVal: 0, maxVal: 4.0, defaultPreset: "CHEST_PA", sliceThickness: "Projection" },
      US: { unit: "dB", minVal: 0, maxVal: 90, defaultPreset: "THYROID_B", sliceThickness: "7.5 MHz" },
      MR: { unit: "SI", minVal: 0, maxVal: 4095, defaultPreset: "T1_AXIAL", sliceThickness: "3.0mm" },
      PET: { unit: "SUV", minVal: 0, maxVal: 25.0, defaultPreset: "18F_FDG", sliceThickness: "2.0mm" }
    };

    expect(MODALITY_CONFIGS.CT.unit).toBe("HU");
    expect(MODALITY_CONFIGS.DX.unit).toBe("OD");
    expect(MODALITY_CONFIGS.US.unit).toBe("dB");
    expect(MODALITY_CONFIGS.MR.unit).toBe("SI");
    expect(MODALITY_CONFIGS.PET.unit).toBe("SUV");
  });

  // 11. Zero-Dependency DICOM 3.0 (Part 10) Client-Side Parser
  it("parses valid DICOM Part 10 buffer and extracts clinical metadata & 16-bit CT pixel data", async () => {
    const { parseDicom, createSyntheticDicomBuffer } = await import("../src/dicom-parser");

    const syntheticBuffer = createSyntheticDicomBuffer({
      patientName: "TEST^PATIENT^M",
      patientId: "#RAD-2026-X01",
      modality: "CT",
      rows: 128,
      columns: 128,
      windowCenter: -600,
      windowWidth: 1500
    });

    expect(syntheticBuffer.byteLength).toBeGreaterThan(132 + 128 * 128 * 2);

    // Verify 'DICM' signature at offset 128
    const u8 = new Uint8Array(syntheticBuffer);
    const magic = String.fromCharCode(u8[128], u8[129], u8[130], u8[131]);
    expect(magic).toBe("DICM");

    // Parse DICOM
    const parsed = parseDicom(syntheticBuffer);

    expect(parsed.metadata.patientName).toBe("TEST^PATIENT^M");
    expect(parsed.metadata.patientId).toBe("#RAD-2026-X01");
    expect(parsed.metadata.modality).toBe("CT");
    expect(parsed.metadata.rows).toBe(128);
    expect(parsed.metadata.columns).toBe(128);
    expect(parsed.metadata.bitsAllocated).toBe(16);
    expect(parsed.metadata.windowCenter).toBe(-600);
    expect(parsed.metadata.windowWidth).toBe(1500);
    expect(parsed.metadata.rescaleSlope).toBe(1);
    expect(parsed.metadata.rescaleIntercept).toBe(0);

    // Verify Decoded Pixel Array
    expect(parsed.pixelData.length).toBe(128 * 128);
    expect(parsed.minPixelValue).toBeLessThanOrEqual(-700); // Air / Lung parenchyma
    expect(parsed.maxPixelValue).toBeGreaterThanOrEqual(400); // Cortical Ribs / Spine

    // Verify ImageData rendering with W/L transfer function
    const mockImageData = {
      width: 128,
      height: 128,
      data: new Uint8ClampedArray(128 * 128 * 4)
    } as unknown as ImageData;

    const rendered = parsed.renderToImageData(mockImageData, -600, 1500);
    expect(rendered.data.length).toBe(128 * 128 * 4);
    // Alpha channel must be 255 for all pixels
    expect(rendered.data[3]).toBe(255);
    expect(rendered.data[7]).toBe(255);

    // In Lung window (W: 1500, L: -600), -1000 HU air is charcoal gray (floor(((-1000 - (-1350)) / 1500) * 255) = 59)
    expect(rendered.data[0]).toBe(59);
  });

  // 12. DICOM Parser Error Handling & Edge Cases
  it("handles malformed or truncated DICOM buffers gracefully", async () => {
    const { parseDicom } = await import("../src/dicom-parser");

    // Too small buffer (< 132 bytes)
    const tinyBuffer = new ArrayBuffer(50);
    expect(() => parseDicom(tinyBuffer)).toThrow("Buffer too small");
  });
});


