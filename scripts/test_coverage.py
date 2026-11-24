import contextlib
import importlib.util
import io
from pathlib import Path
import unittest

spec = importlib.util.spec_from_file_location("coverage_check", Path(__file__).with_name("check-coverage.py"))
checker = importlib.util.module_from_spec(spec)
spec.loader.exec_module(checker)


class CoverageTests(unittest.TestCase):
    def test_duplicate_blocks_are_merged(self):
        prefix = checker.MODULE
        totals = checker.coverage(f"mode: atomic\n{prefix}client.go:1.1,2.2 2 0\n{prefix}client.go:1.1,2.2 2 3\n{prefix}client.go:3.1,4.2 1 0\n")
        self.assertEqual(totals["sdk"], [3, 2])
        self.assertEqual(totals["total"], [3, 2])

    def test_thresholds_and_missing_groups_fail_closed(self):
        totals = {group: [100, 100] for group in checker.LIMITS}
        with contextlib.redirect_stdout(io.StringIO()):
            self.assertTrue(checker.check(totals))
            for group, minimum in checker.LIMITS.items():
                with self.subTest(group=group):
                    self.assertFalse(checker.check(totals | {group: [100, minimum - 1]}))
                    self.assertFalse(checker.check(totals | {group: [0, 0]}))

    def test_invalid_profiles_are_rejected(self):
        for text in ["", "mode: unknown", "mode: atomic\nbroken", f"mode: atomic\n{checker.MODULE}new/tool.go:1.1,2.2 1 1"]:
            with self.subTest(text=text), self.assertRaises(ValueError):
                checker.coverage(text)


if __name__ == "__main__":
    unittest.main()
