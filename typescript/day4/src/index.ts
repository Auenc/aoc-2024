import fs from "fs";

export const SearchRight = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, 1, 0);

export const SearchLeft = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, -1, 0);

export const SearchUp = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, 0, -1);

export const SearchDown = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, 0, 1);

export const SearchDiagonalRightDown = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, 1, 1);

export const SearchDiagonalRightUp = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, -1, -1);

export const SearchDiagonalLeftDown = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, -1, 1);

export const SearchDiagonalLeftUp = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => searchArray(arr, startingX, startingY, 1, -1);

const searchArray = (
  arr: string[][],
  x: number,
  y: number,
  xDiff: number,
  yDiff: number
): boolean => {
  if (y + yDiff * 3 >= arr.length || y + yDiff * 3 < 0) {
    return false;
  }
  if (x + xDiff * 3 >= arr[y].length || x + xDiff * 3 < 0) {
    return false;
  }
  return (
    arr[y][x] === "x" &&
    arr[y + yDiff][x + xDiff] === "m" &&
    arr[y + yDiff * 2][x + xDiff * 2] === "a" &&
    arr[y + yDiff * 3][x + xDiff * 3] === "s"
  );
};

export const search = (arr: string[][]): number => {
  let count = 0;
  for (let x = 0; x < arr.length; x++) {
    for (let y = 0; y < arr[x].length; y++) {
      const letter = arr[x][y];
      if (letter !== "x") {
        continue;
      }
      if (SearchRight(arr, y, x)) {
        count++;
      }
      if (SearchLeft(arr, y, x)) {
        count++;
      }
      if (SearchUp(arr, y, x)) {
        count++;
      }
      if (SearchDown(arr, y, x)) {
        count++;
      }
      if (SearchDiagonalRightDown(arr, y, x)) {
        count++;
      }
      if (SearchDiagonalRightUp(arr, y, x)) {
        count++;
      }
      if (SearchDiagonalLeftDown(arr, y, x)) {
        count++;
      }
      if (SearchDiagonalLeftUp(arr, y, x)) {
        count++;
      }
    }
  }
  return count;
};

const isX = (
  arr: string[][],
  x: number,
  y: number,
  topLeft: string,
  topRight: string,
  bottomLeft: string,
  bottomRight: string
): boolean => {
  if (y + 2 >= arr.length || x + 2 >= arr[y].length) {
    return false;
  }

  const isTopLeft = arr[y][x] == topLeft;
  const isTopRight = arr[y][x + 2] == topRight;
  const isMiddle = arr[y + 1][x + 1] == "a";
  const isBottomLeft = arr[y + 2][x] == bottomLeft;
  const isBottomRight = arr[y + 2][x + 2] == bottomRight;
  return isTopLeft && isTopRight && isMiddle && isBottomLeft && isBottomRight;
};

const findX0 = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => isX(arr, startingX, startingY, "m", "s", "m", "s");

const findX90 = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => isX(arr, startingX, startingY, "m", "m", "s", "s");

const findX180 = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => isX(arr, startingX, startingY, "s", "m", "s", "m");

const findX270 = (
  arr: string[][],
  startingX: number,
  startingY: number
): boolean => isX(arr, startingX, startingY, "s", "s", "m", "m");

export const searchX = (arr: string[][]): number => {
  let count = 0;
  for (let y = 0; y < arr.length; y++) {
    for (let x = 0; x < arr[y].length; x++) {
      if (findX0(arr, x, y)) {
        count++;
      }
      if (findX90(arr, x, y)) {
        count++;
      }
      if (findX180(arr, x, y)) {
        count++;
      }
      if (findX270(arr, x, y)) {
        count++;
      }
    }
  }
  return count;
};

export const loadFile = (path: string): string => fs.readFileSync(path, "utf8");

export const parseData = (data: string): string[][] =>
  data.split("\n").map((line) => line.split(""));

export const main = () => {
  const contents = loadFile("../../data/day4/data.txt").toLowerCase();
  const data = parseData(contents);
  const countPart1 = search(data);
  console.log("part 1", countPart1);
  const countPart2 = searchX(data);
  console.log("part 2", countPart2);
};

main();
