import fs from "fs";

export const isSafe = (levels: number[]): boolean => {
  let hasRisen = false;
  let hasFallen = false;

  for (let i = 1; i < levels.length; i++) {
    const previousLevel = levels[i - 1];
    const diff = levels[i] - previousLevel;

    if (diff > 0) {
      hasRisen = true;
    } else if (diff < 0) {
      hasFallen = true;
    } else {
      return false;
    }

    const absDiff = Math.abs(diff);

    if (absDiff > 3) {
      return false;
    }
    if (hasRisen && hasFallen) {
      return false;
    }
  }
  return true;
};

export const isDampenerSafe = (
  levels: number[],
  toRemove: number = -1
): boolean => {
  if (toRemove !== -1) {
    const filtered = levels.filter((_, i) => i !== toRemove);
    const safe = isSafe(filtered);
    if (safe) {
      return true;
    }
    const next = toRemove + 1;
    if (next > levels.length - 1) {
      return false;
    }
    return isDampenerSafe(levels, next);
  }

  const safe = isSafe(levels);
  if (safe) {
    return true;
  }

  return isDampenerSafe(levels, 0);
};

export const countSafe = (levels: number[][]): number =>
  levels.reduce((acc, level) => acc + (isSafe(level) ? 1 : 0), 0);

export const countSafeDampended = (levels: number[][]): number =>
  levels.reduce((acc, level) => acc + (isDampenerSafe(level) ? 1 : 0), 0);

export const loadFile = (path: string): string => fs.readFileSync(path, "utf8");

const parseInput = (input: string): number[][] =>
  input
    .trim()
    .split("\n")
    .map((line) => line.split(" ").map((num) => parseInt(num, 10)));

export const main = () => {
  const contents = loadFile("../../data/day2/data.txt");
  const levels = parseInput(contents);
  const safeLevels = countSafeDampended(levels);
  console.log(safeLevels);
};

main();
