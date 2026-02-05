def Adding(num1, num2):
    if (isinstance(num1, (int, float)) != True):
        raise Exception("Input needs to be a number")

    return num1 + num2

def Subtracting(num1, num2):
    return num1 - num2

def Multiplication(num1, num2):
    return num1 * num2

def Dividing(num1, num2):
    if (num1 == 0 or num2 == 0):
        raise ZeroDivisionError("Cant divide by 0!")
    
    return num1 / num2
    
