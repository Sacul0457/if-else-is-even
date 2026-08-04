//! Zig is a general-purpose programming language and toolchain for maintaining
//! robust, optimal and reusable software.
//!
//! Compilation:
//! $ zig build-exe zig.zig
//!
//! Usage:
//! $ ./zig 67 (six seveeeen, six seveeeen, if you don't know the joke, well maybe the joke aint for you)
//!
//! It will output the code into the console.
//! Redirect it into a zig file:
//! $ ./zig 67 > is_even.zig
//!
//! Compile that newly baked is_even program:
//! $ zig build-exe is_even.zig
//!
//! Check the result, is 67 even?:
//! $ ./is_even
//! false
//!
//! Negative numbers work too:
//! $ ./zig -67 > is_even.zig

const std = @import("std");
const process = std.process;
const fmt = std.fmt;
const Io = std.Io;

pub fn main(init: process.Init) !u8 {
    const arena = init.arena;
    const io = init.io;
    const args = try init.minimal.args.toSlice(arena.allocator());

    var stderr_buferr: [1024 * 4]u8 = undefined;
    var stderr_writer = Io.File.stderr().writer(io, &stderr_buferr);
    const stderr = &stderr_writer.interface;

    if (args.len != 2) {
        try stderr.print(
            \\Well all I wanted was 1 argument, but you choose to give me '{d}' instead,
            \\as a punishment you get to eat this error.
            \\
        ,
            .{args.len - 1},
        );

        try stderr.writeAll("Want some free advice? Try ./zig 67\n");
        try stderr.flush();
        return 1;
    }

    const number = fmt.parseInt(i64, args[1], 10) catch |err| {
        switch (err) {
            error.Overflow => try stderr.print(
                \\Well all I wanted was a number that fits in an i64, but you choose to
                \\give me '{s}' instead, as a punishment you get to eat this error.
                \\
            ,
                .{args[1]},
            ),
            error.InvalidCharacter => try stderr.print(
                \\Well all I wanted was a number, but you choose to give me '{s}' instead,
                \\as a punishment you get to eat this error.
                \\
            ,
                .{args[1]},
            ),
        }

        try stderr.writeAll("Want some free advice? Try ./zig 67\n");
        try stderr.flush();
        return 2;
    };

    var stdout_buffer: [1024 * 4]u8 = undefined;
    var stdout_writer = Io.File.stdout().writer(io, &stdout_buffer);
    const stdout = &stdout_writer.interface;

    try stdout.writeAll(
        \\const std = @import("std");
        \\const process = std.process;
        \\const Io = std.Io;
        \\
        \\pub fn main(init: process.Init) !void {
        \\    // If the compiler knew this number at compile time it would resolve all
        \\    // of the stupid if branches below, and none of them would ever execute at runtime.
        \\
    );

    try stdout.print("    var number: i64 = {d};\n", .{number});

    try stdout.writeAll(
        \\    _ = &number;
        \\
        \\    const is_even = blk: {
        \\
    );

    const step: i64 = if (number < 0) -1 else 1;
    var i: i64 = 0;
    while (true) : (i += step) {
        try stdout.print("        if (number == {d}) break :blk {};\n", .{ i, @mod(i, 2) == 0 });
        if (i == number) break;
    }

    try stdout.writeAll(
        \\        unreachable;
        \\    };
        \\
        \\    var stdout_buffer: [8]u8 = undefined;
        \\    var stdout_writer = Io.File.stdout().writer(init.io, &stdout_buffer);
        \\    const stdout = &stdout_writer.interface;
        \\
        \\    try stdout.print("{}\n", .{is_even});
        \\    try stdout.flush();
        \\}
        \\
    );

    try stdout.flush();
    return 0;
}
