#!/usr/bin/env python3
"""Check statement coverage, merging duplicate blocks from -coverpkg=./...."""
import sys
from pathlib import Path

MODULE = "github.com/ritoon/sdk-pennylane/"
# Keep the generated client separate so it cannot hide gaps in handwritten code.
LIMITS = {"total": 57, "generated": 57, "sdk": 94, "examples": 85,
          "docgen": 93, "spec": 91}


def coverage(profile):
    lines = profile.splitlines()
    if not lines or lines[0] not in {"mode: atomic", "mode: count", "mode: set"}:
        raise ValueError("missing or invalid coverage mode")
    blocks = {}
    for line in lines[1:]:
        location, statements, count = line.split()
        statements, count = int(statements), int(count)
        if statements < 0 or count < 0:
            raise ValueError("negative coverage value")
        if location in blocks and blocks[location][0] != statements:
            raise ValueError("inconsistent duplicate coverage block")
        blocks[location] = (statements, max(count, blocks.get(location, (0, 0))[1]))
    totals = {group: [0, 0] for group in LIMITS}
    for location, (statements, count) in blocks.items():
        filename = location.rsplit(":", 1)[0]
        if not filename.startswith(MODULE):
            raise ValueError("unexpected module: " + filename)
        filename = filename[len(MODULE):]
        if filename == "api.gen.go":
            group = "generated"
        elif filename.startswith("exemple/"):
            group = "examples"
        elif filename.startswith("internal/docgen/"):
            group = "docgen"
        elif filename.startswith("internal/spec/"):
            group = "spec"
        elif "/" not in filename:
            group = "sdk"
        else:
            raise ValueError("add a coverage group for " + filename)
        for key in ("total", group):
            totals[key][0] += statements
            totals[key][1] += statements if count > 0 else 0
    return totals


def check(totals):
    passed = True
    for group, minimum in LIMITS.items():
        total, covered = totals[group]
        percent = 100 * covered / total if total else 0
        ok = total > 0 and percent >= minimum
        print(f"{group}: {percent:.2f}% ({covered}/{total}), minimum {minimum}%: {'OK' if ok else 'FAIL'}")
        passed &= ok
    return passed


if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit("usage: python3 scripts/check-coverage.py coverage.out")
    try:
        sys.exit(0 if check(coverage(Path(sys.argv[1]).read_text())) else 1)
    except (OSError, ValueError) as error:
        sys.exit(str(error))
