import { describe, expect, it } from "@jest/globals";
import {
  isInvalidTuple,
  isValidUpdate,
  loadInput,
  parseRules,
  parseUpdates,
  removeInvalidUpdates,
  sortIncorrect,
  splitRulesUpdates,
  totalMiddleNumbers,
} from ".";

describe("splitRulesUpdates", () => {
  it("splits the input by double newlines", () => {
    const input = `47|54
54|47

75,47,54,23,1
75,47,54,23,1
75,47,54,23,1`;
    expect(splitRulesUpdates(input)).toEqual([
      "47|54\n54|47",
      "75,47,54,23,1\n75,47,54,23,1\n75,47,54,23,1",
    ]);
  });
});

describe("parseRules", () => {
  it("parses the rules", () => {
    const input = `47|54
54|47`;
    expect(parseRules(input)).toEqual([
      { before: 47, after: 54 },
      { before: 54, after: 47 },
    ]);
  });
});

describe("parseUpdates", () => {
  it("parses the updates", () => {
    const input = `75,47,54,23,1
75,47,54,23,1
75,47,54,23,1`;
    expect(parseUpdates(input)).toEqual([
      [75, 47, 54, 23, 1],
      [75, 47, 54, 23, 1],
      [75, 47, 54, 23, 1],
    ]);
  });
});

describe("isInvalidTuple", () => {
  it("returns true if the tuple is invalid", () => {
    expect(isInvalidTuple([47, 54], [{ before: 47, after: 54 }])).toBe(false);
    expect(isInvalidTuple([54, 47], [{ before: 47, after: 54 }])).toBe(true);
    expect(isInvalidTuple([75, 97], [{ before: 97, after: 75 }])).toBe(true);
    expect(isInvalidTuple([13, 29], [{ before: 29, after: 13 }])).toBe(true);
  });
});

describe("isValidateUpdate", () => {
  it("returns true if the update is valid", () => {
    const input = loadInput("../../data/day5/test.txt");
    expect(isValidUpdate([75, 47, 61, 53, 29], input.rules)).toBe(true);
    expect(isValidUpdate([97, 61, 53, 29, 13], input.rules)).toBe(true);
    expect(isValidUpdate([75, 29, 13], input.rules)).toBe(true);
    expect(isValidUpdate([75, 97, 47, 61, 53], input.rules)).toBe(false);
    expect(isValidUpdate([61, 13, 29], input.rules)).toBe(false);
    expect(isValidUpdate([97, 13, 75, 29, 47], input.rules)).toBe(false);
  });
});

describe("removeInvalidTuples", () => {
  it("removes invalid tuples", () => {
    const input = loadInput("../../data/day5/test.txt");
    const expected = [
      [75, 47, 61, 53, 29],
      [97, 61, 53, 29, 13],
      [75, 29, 13],
    ];
    const [valid] = removeInvalidUpdates(input.updates, input.rules);
    expect(valid).toEqual(expected);
  });
});

describe("totalMiddleNumbers", () => {
  it("returns the sum of the middle numbers", () => {
    const input = loadInput("../../data/day5/test.txt");
    const [correctRules] = removeInvalidUpdates(input.updates, input.rules);
    expect(totalMiddleNumbers(correctRules)).toBe(143);
  });
});

describe("sortIncorrect", () => {
  it("sorts the incorrect updates", () => {
    const input = loadInput("../../data/day5/test.txt");
    expect(sortIncorrect([75, 97, 47, 61, 53], input.rules)).toEqual([
      97, 75, 47, 61, 53,
    ]);
  });
});
