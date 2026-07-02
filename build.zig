const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    _ = b.addModule(.{
        .name = "megaverse",
        .root_source_file = .{ .path = "src/main.zig" },
    });
}
