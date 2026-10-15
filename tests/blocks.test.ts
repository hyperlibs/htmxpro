import { describe, test, expect } from "bun:test";
import fs from "fs";
import path from "path";

describe("Research & Industrial Design Blocks Verification", () => {
  const researchBlocks = [
    { file: "views/components/block-pathology-lab.html", rootClass: "block-pathology-lab", domain: "Pathology WSI" },
    { file: "views/components/block-radiology-dicom.html", rootClass: "block-radiology-dicom", domain: "Radiology DICOM" },
    { file: "views/components/block-biology-atlas.html", rootClass: "block-biology-atlas", domain: "Biology Atlas" },
    { file: "views/components/block-crispr-studio.html", rootClass: "block-crispr-studio", domain: "CRISPR Studio" },
    { file: "views/components/block-genomic-map.html", rootClass: "block-genomic-map", domain: "Genomic Map" }
  ];

  const industrialBlocks = [
    { file: "views/components/block-civil-architect.html", rootClass: "block-civil-architect", domain: "Civil/Architecture" },
    { file: "views/components/block-marine-navigation.html", rootClass: "block-marine-navigation", domain: "Marine/ECDIS" },
    { file: "views/components/block-agriculture-precision.html", rootClass: "block-agriculture-precision", domain: "Precision Agri" },
    { file: "views/components/block-energy-scada.html", rootClass: "block-energy-scada", domain: "Energy SCADA" },
    { file: "views/components/block-pcb-semiconductor.html", rootClass: "block-pcb-semiconductor", domain: "PCB & Semiconductor" }
  ];

  const allBlocks = [...researchBlocks, ...industrialBlocks];

  test("All 10 Research & Industrial Design Block files exist on disk", () => {
    for (const b of allBlocks) {
      expect(fs.existsSync(b.file)).toBe(true);
      const stats = fs.statSync(b.file);
      expect(stats.size).toBeGreaterThan(500); // Substantial content
    }
  });

  test("All 10 blocks adhere to strict Design Block encapsulation (no scattered root)", () => {
    for (const b of allBlocks) {
      const content = fs.readFileSync(b.file, "utf-8");
      // Must contain root container class
      expect(content).toContain(b.rootClass);
      // Must contain hx-state for Bolt reactivity
      expect(content).toContain("hx-state=");
      // Must have Zone 1 Header and Zone 5 Footer
      expect(content).toContain("ZONE 1: BLOCK HEADER");
      expect(content).toContain("ZONE 5: AUDIT & @DIAG");
    }
  });

  test("Zero external prohibited framework mentions (anti-copyright check)", () => {
    const prohibited = ["framework7", "framework-7", "f7-", "framework 7"];
    for (const b of allBlocks) {
      const content = fs.readFileSync(b.file, "utf-8").toLowerCase();
      for (const term of prohibited) {
        expect(content.includes(term)).toBe(false);
      }
    }
  });

  test("Go DSL files exist and export matching component constructors", () => {
    expect(fs.existsSync("go/components/research.go")).toBe(true);
    expect(fs.existsSync("go/components/industrial.go")).toBe(true);

    const researchGo = fs.readFileSync("go/components/research.go", "utf-8");
    expect(researchGo).toContain("func BlockPathologyLab");
    expect(researchGo).toContain("func BlockRadiologyDICOM");
    expect(researchGo).toContain("func BlockBiologyAtlas");
    expect(researchGo).toContain("func BlockCrisprStudio");
    expect(researchGo).toContain("func BlockGenomicMap");

    const industrialGo = fs.readFileSync("go/components/industrial.go", "utf-8");
    expect(industrialGo).toContain("func BlockCivilArchitectStudio");
    expect(industrialGo).toContain("func BlockMarineNavigation");
    expect(industrialGo).toContain("func BlockAgriculturePrecision");
    expect(industrialGo).toContain("func BlockEnergySCADA");
    expect(industrialGo).toContain("func BlockPcbSemiconductor");
  });
});
