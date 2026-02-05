import pytest
import calculations as c

# ========== PASSED ==========
def test_pass_adding():
    # Denne test vil passere
    result = c.Adding(1, 1);
    assert result == 2

def test_invalid_input():
    with pytest.raises(Exception):
        c.Adding("sd", 2)

def test_pass_subtract():
    result = c.Subtracting(2, 1)
    assert result == 1

def test_pass_multiplication():
    result = c.Multiplication(5, 5)
    assert result == 25

def test_pass_dividing():
    result = c.Dividing(10, 5)
    assert result == 2

def test_zero_division():
    with pytest.raises(ZeroDivisionError):
        c.Dividing(0, 1)

# ========== FAILED ==========

def test_fail_adding():
    result = c.Adding(2, 2)
    assert result == 3 

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
