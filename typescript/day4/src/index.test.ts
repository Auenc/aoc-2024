import { describe, expect, it } from "@jest/globals";

import {
  SearchDiagonalLeftDown,
  SearchDiagonalLeftUp,
  SearchDiagonalRightDown,
  SearchDiagonalRightUp,
  SearchDown,
  SearchLeft,
  SearchRight,
  SearchUp,
  searchX,
} from "./index";

describe("SearchRight", () => {
  it("detects true cases", () => {
    expect(SearchRight([["x", "m", "a", "s"]], 0, 0)).toBe(true);
    expect(SearchRight([["a", "b", "x", "m", "a", "s"]], 2, 0)).toBe(true);
  });
  it("detects false cases", () => {
    expect(SearchRight([["m", "a", "s", "x"]], 0, 0)).toBe(false);
    expect(SearchRight([["a", "b", "m", "a", "s", "x"]], 2, 0)).toBe(false);
  });
});

describe("SearchLeft", () => {
  it("detects true cases", () => {
    expect(SearchLeft([["s", "a", "m", "x"]], 3, 0)).toBe(true);
    expect(SearchLeft([["s", "a", "m", "x", "b", "a"]], 3, 0)).toBe(true);
  });
  it("detects false cases", () => {
    expect(SearchLeft([["x", "m", "a", "s"]], 3, 0)).toBe(false);
    expect(SearchLeft([["s", "a", "m", "b", "x", "a"]], 3, 0)).toBe(false);
  });
});

describe("SearchUp", () => {
  it("detects true cases", () => {
    expect(SearchUp([["s"], ["a"], ["m"], ["x"]], 0, 3)).toBe(true);
    expect(SearchUp([["a"], ["b"], ["s"], ["a"], ["m"], ["x"]], 0, 5)).toBe(
      true
    );
  });
  it("detects false cases", () => {
    expect(SearchUp([["m"], ["a"], ["s"], ["x"]], 0, 3)).toBe(false);
    expect(SearchUp([["a"], ["b"], ["m"], ["a"], ["s"], ["x"]], 0, 2)).toBe(
      false
    );
  });
});

describe("SearchDown", () => {
  it("detects true cases", () => {
    expect(SearchDown([["x"], ["m"], ["a"], ["s"]], 0, 0)).toBe(true);
    expect(SearchDown([["a"], ["b"], ["x"], ["m"], ["a"], ["s"]], 0, 2)).toBe(
      true
    );
  });
  it("detects false cases", () => {
    expect(SearchDown([["s"], ["a"], ["m"], ["x"]], 0, 0)).toBe(false);
    expect(SearchDown([["a"], ["b"], ["s"], ["a"], ["m"], ["x"]], 0, 2)).toBe(
      false
    );
  });
});

describe("SearchDiagonalRightDown", () => {
  it("detects true cases", () => {
    expect(
      SearchDiagonalRightDown(
        [
          ["x", ".", ".", "."],
          [".", "m", ".", "."],
          [".", ".", "a", "."],
          [".", ".", ".", "s"],
        ],
        0,
        0
      )
    ).toBe(true);
  });
  it("detects false cases", () => {
    expect(
      SearchDiagonalRightDown(
        [
          ["b", ".", ".", "."],
          [".", "a", ".", "."],
          [".", ".", "r", "."],
          [".", ".", ".", "g"],
        ],
        0,
        0
      )
    ).toBe(false);
  });
});

describe("SearchDiagonalRightUp", () => {
  it("detects true cases", () => {
    expect(
      SearchDiagonalRightUp(
        [
          ["s", ".", ".", "."],
          [".", "a", ".", "."],
          [".", ".", "m", "."],
          [".", ".", ".", "x"],
        ],
        3,
        3
      )
    ).toBe(true);
  });
  it("detects false cases", () => {
    expect(
      SearchDiagonalRightUp(
        [
          ["r", ".", ".", "."],
          [".", "a", ".", "."],
          [".", ".", "l", "."],
          [".", ".", ".", "b"],
        ],
        3,
        0
      )
    ).toBe(false);
  });
});

describe("SearchDiagonalLeftDown", () => {
  it("detects true cases", () => {
    expect(
      SearchDiagonalLeftDown(
        [
          [".", ".", ".", "x"],
          [".", ".", "m", "."],
          [".", "a", ".", "."],
          ["s", ".", ".", "."],
        ],
        3,
        0
      )
    ).toBe(true);
  });
  it("detects false cases", () => {
    expect(
      SearchDiagonalLeftDown(
        [
          [".", ".", ".", "b"],
          [".", ".", "a", "."],
          [".", "m", ".", "."],
          ["s", ".", ".", "."],
        ],
        3,
        0
      )
    ).toBe(false);
  });
});

describe("SearchDiagonalLeftUp", () => {
  it("detects true cases", () => {
    expect(
      SearchDiagonalLeftUp(
        [
          [".", ".", ".", "s"],
          [".", ".", "a", "."],
          [".", "m", ".", "."],
          ["x", ".", ".", "."],
        ],
        0,
        3
      )
    ).toBe(true);
  });
  it("detects false cases", () => {
    expect(
      SearchDiagonalLeftUp(
        [
          [".", ".", ".", "b"],
          [".", ".", "a", "."],
          [".", "m", ".", "."],
          ["x", ".", ".", "."],
        ],
        0,
        3
      )
    ).toBe(false);
  });
});

describe("searchX", () => {
  it("detects true cases", () => {
    expect(
      searchX([
        [".", "m", ".", "s", ".", ".", ".", ".", ".", "."],
        [".", ".", "a", ".", ".", "m", "s", "m", "s", "."],
        [".", "m", ".", "s", ".", "m", "a", "a", ".", "."],
        [".", ".", "a", ".", "a", "s", "m", "s", "m", "."],
        [".", "m", ".", "s", ".", "m", ".", ".", ".", "."],
        [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
        ["s", ".", "s", ".", "s", ".", "s", ".", "s", "."],
        [".", "a", ".", "a", ".", "a", ".", "a", ".", "."],
        ["m", ".", "m", ".", "m", ".", "m", ".", "m", "."],
        [".", ".", ".", ".", ".", ".", ".", ".", ".", "."],
      ])
    ).toBe(9);
  });
});
