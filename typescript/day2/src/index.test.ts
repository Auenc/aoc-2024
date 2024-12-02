import { isSafe, countSafe, isDampenerSafe } from "./index";

import { describe, expect, it } from "@jest/globals";

describe("isSafe", () => {
  it("detects true cases", () => {
    expect(isSafe([7, 6, 4, 2, 1])).toBe(true);
    expect(isSafe([1, 3, 6, 7, 9])).toBe(true);
  });
  it("detects unsafe cases", () => {
    expect(isSafe([1, 2, 7, 8, 9])).toBe(false);
    expect(isSafe([9, 7, 6, 2, 1])).toBe(false);
    expect(isSafe([1, 3, 2, 4, 5])).toBe(false);
    expect(isSafe([8, 6, 4, 4, 1])).toBe(false);
  });
});

describe("countSafe", () => {
  it("counts safe cases", () => {
    expect(
      countSafe([
        [7, 6, 4, 2, 1],
        [1, 3, 6, 7, 9],
        [1, 2, 7, 8, 9],
        [9, 7, 6, 2, 1],
        [1, 3, 2, 4, 5],
        [8, 6, 4, 4, 1],
      ])
    ).toBe(2);
  });
});

describe("isDampenerSafe", () => {
  it("detects true cases", () => {
    expect(isDampenerSafe([7, 6, 4, 2, 1])).toBe(true);
    expect(isDampenerSafe([1, 3, 2, 4, 5])).toBe(true);
    expect(isDampenerSafe([8, 6, 4, 4, 1])).toBe(true);
    expect(isDampenerSafe([1, 3, 6, 7, 9])).toBe(true);
  });
  it("detects unsafe cases", () => {
    expect(isDampenerSafe([1, 2, 7, 8, 9])).toBe(false);
    expect(isDampenerSafe([9, 7, 6, 2, 1])).toBe(false);
  });
});
