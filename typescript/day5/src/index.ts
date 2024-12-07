import fs from "fs";

interface Rule {
  before: number;
  after: number;
}

interface Input {
  rules: Rule[];
  updates: number[][];
}

export const splitRulesUpdates = (input: string): string[] =>
  input.split("\n\n");

export const parseRules = (input: string): Rule[] =>
  input
    .split("\n")
    .map((line) => line.split("|").map(Number))
    .map(([before, after]) => ({ before, after }));

export const parseUpdates = (input: string): number[][] =>
  input.split("\n").map((line) => line.split(",").map(Number));

export const parseInput = (input: string): Input => {
  const [rules, updates] = splitRulesUpdates(input);
  return {
    rules: parseRules(rules),
    updates: parseUpdates(updates),
  };
};

export const loadFile = (path: string): string => fs.readFileSync(path, "utf8");

export const loadInput = (path: string): Input => parseInput(loadFile(path));

export const isInvalidTuple = (tuple: number[], rules: Rule[]): boolean => {
  const [before, after] = tuple;
  for (const rule of rules) {
    if (rule.before === after && rule.after === before) {
      return true;
    }
  }
  return false;
};

export const isValidUpdate = (update: number[], rules: Rule[]): boolean => {
  for (let i = 1; i < update.length; i++) {
    if (isInvalidTuple([update[i - 1], update[i]], rules)) {
      return false;
    }
  }
  return true;
};

export const removeInvalidUpdates = (
  updates: number[][],
  rules: Rule[]
): number[][][] => {
  const valid: number[][] = [];
  const invalid: number[][] = [];
  for (const update of updates) {
    if (isValidUpdate(update, rules)) {
      valid.push(update);
    } else {
      invalid.push(update);
    }
  }
  return [valid, invalid];
};

export const getMiddleNumber = (update: number[]): number =>
  update[Math.floor(update.length / 2)];

export const totalMiddleNumbers = (updates: number[][]): number =>
  updates.reduce((acc, update) => acc + getMiddleNumber(update), 0);

export const sortIncorrect = (update: number[], rules: Rule[]): number[] => {
  const sorted: number[] = [];
  for (let i = 1; i < update.length; i++) {
    const [before, after] = [update[i - 1], update[i]];
    if (!isInvalidTuple([before, after], rules)) {
      sorted.push(before);
      if (i === update.length - 1) {
        sorted.push(after);
      }
      continue;
    }
    sorted.push(after, before, ...update.slice(i + 1));
    return sortIncorrect(sorted, rules);
  }
  return sorted;
};

const main = () => {
  const input = loadInput("../../data/day5/data.txt");
  const [valid, invalid] = removeInvalidUpdates(input.updates, input.rules);
  const corrected = invalid.map((update) => sortIncorrect(update, input.rules));
  console.log("valid", totalMiddleNumbers(valid));
  console.log("invalid", totalMiddleNumbers(corrected));
};

main();
