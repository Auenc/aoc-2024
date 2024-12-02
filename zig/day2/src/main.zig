const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const file = std.fs.cwd().openFile("../../data/day2/test.txt", .{}) catch |err| {
        return err;
    };
    defer file.close();

    var reader = file.reader();
    var lineList = std.ArrayList([]u8).init(allocator);
    defer lineList.deinit();

    var buffer: [1024]u8 = undefined;
    while (try reader.readUntilDelimiterOrEof(buffer[0..], '\n')) |line| {
        const owned_line = try allocator.alloc(u8, line.len);
        @memcpy(owned_line, line);

        try lineList.append(owned_line);
    }

    const lines: [][]u8 = try lineList.toOwnedSlice();
    defer for (lines) |line| allocator.free(line);

    const levels = try parse_levels(lines, allocator);
    defer for (levels) |level| allocator.free(level);

    const safe_count = count_safe(levels);
    print("Safe levels: {}\n", .{safe_count});
}

fn parse_levels(lines: [][]u8, allocator: std.mem.Allocator) ![][]i32 {
    var levels = std.ArrayList([]i32).init(allocator);
    defer levels.deinit();

    for (lines) |line| {
        var level = std.ArrayList(i32).init(allocator);
        defer level.deinit();

        for (line) |char| {
            if (char == ' ') {
                continue;
            }

            const digit: i32 = @intCast(char);
            try level.append(digit);
        }

        try levels.append(try level.toOwnedSlice());
    }

    return levels.toOwnedSlice();
}

fn count_safe(levels: [][]i32) u8 {
    var count: u8 = 0;
    for (levels) |level| {
        if (is_safe(level)) {
            count += 1;
        }
    }
    return count;
}

test "count_is_safe" {
    const levels = &[_][]const i32{
        &[_]i32{ 7, 6, 4, 2, 1 },
        &[_]i32{ 1, 2, 7, 8, 9 },
        &[_]i32{ 9, 7, 6, 2, 1 },
        &[_]i32{ 1, 3, 2, 4, 5 },
        &[_]i32{ 8, 6, 4, 4, 1 },
        &[_]i32{ 1, 3, 6, 7, 9 },
    };
    try std.testing.expect(count_safe(levels) == 2);
}

fn is_safe(levels: []const i32) bool {
    var has_risen = false;
    var has_lowered = false;
    for (levels, 0..) |level, i| {
        if (i == 0) {
            continue;
        }
        const prev_level = levels[i - 1];
        const diff = level - prev_level;
        if (diff > 0) {
            has_risen = true;
        } else {
            has_lowered = true;
        }
        if (has_risen and has_lowered) {
            return false;
        }
        if (diff > 3 or diff < -2 or diff == 0) {
            return false;
        }
    }
    return true;
}

test "is_safe" {
    try std.testing.expect(is_safe(&[_]i32{ 7, 6, 4, 2, 1 }));
    try std.testing.expect(!is_safe(&[_]i32{ 1, 2, 7, 8, 9 }));
    try std.testing.expect(!is_safe(&[_]i32{ 9, 7, 6, 2, 1 }));
    try std.testing.expect(!is_safe(&[_]i32{ 1, 3, 2, 4, 5 }));
    try std.testing.expect(!is_safe(&[_]i32{ 8, 6, 4, 4, 1 }));
    try std.testing.expect(is_safe(&[_]i32{ 1, 3, 6, 7, 9 }));
    try std.testing.expect(is_safe(&[_]i32{ 1, 2, 3, 4, 6, 7, 9 }));
}

test "parse_lines" {
    const lines = [_][]u8{
        "7 6 4 2 1\n",
        "1 2 7 8 9\n",
        "9 7 6 2 1\n",
        "1 3 2 4 5\n",
        "8 6 4 4 1\n",
        "1 3 6 7 9\n",
    };
    const allocator = std.heap.page_allocator;
    const levels = try parse_levels(&lines, allocator);
    defer for (levels) |level| allocator.free(level);

    try std.testing.expect(levels.len == 6);
    try std.testing.expect(levels[0] == [_]i32{ 7, 6, 4, 2, 1 });
    try std.testing.expect(levels[1] == [_]i32{ 1, 2, 7, 8, 9 });
    try std.testing.expect(levels[2] == [_]i32{ 9, 7, 6, 2, 1 });
    try std.testing.expect(levels[3] == [_]i32{ 1, 3, 2, 4, 5 });
    try std.testing.expect(levels[4] == [_]i32{ 8, 6, 4, 4, 1 });
    try std.testing.expect(levels[5] == [_]i32{ 1, 3, 6, 7, 9 });
}
