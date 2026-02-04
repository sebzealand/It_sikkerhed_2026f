import pytest

# ========== PASSED ==========
def test_pass_adding():
    # Denne test vil passere
    assert 1 + 1 == 2

def test_pass_subtract():
    assert 2 - 1 == 1

def test_pass_multiplication():
    assert 5 * 5 == 25

def test_pass_dividing():
    assert 10 / 5 == 2

# ========== FAILED ==========

def test_fail_adding():
    assert 2 + 2 == 3

def test_fail_subtract():
    assert 2 - 1 == 2

def test_fail_multiplication():
    assert 2 * 2 == 6

def test_fail_dividing():
    assert 6 / 1 == 1

# ========= SKIPPED ==========

@pytest.mark.skip(reason="Springes over med vilje") # Denne test bliver slet ikke kørt
def test_skip():
    assert False # failed test bliver ignoreret
    raise RuntimeError("Test crashede med vilje") # crash bliver også ignoreret

@pytest.mark.skip(reason="Ændrer til skip fremfor failed")
def test_crash():
    # Denne test crasher med en exception
    raise RuntimeError("Test crashede med vilje")

    assert False # failed test bliver ignoreret
