import calculations as cf
import sys
import os

def main():
    print("==============================")
    print("========= CALCULATOR =========")
    print("==============================")

    result = 0
    while(True):

        try:
            number = int(input("Enter a number: "))
        except ValueError:
            print("Input needs to be a number")
            continue
    
        if (isinstance(number, (int, float)) != True):
            raise Exception("Input needs to be a number!")

        print("1. Add\n" 
              "2. Minus\n" 
              "3. Multiplication\n"
              "4. Division\n"
              "5. Start over\n"
              "6. Quit\n")

        choice = int(input("Choose a number: "))
        
        if choice == 1:
            result = cf.Adding(result, number)
        elif choice == 2:
            result = cf.Subtracting(result, number)
        elif choice == 3: 
            result = cf.Multiplication(result, number)
        elif choice == 4:
            result = cf.Dividing(result, number)
        elif choice == 5:
            print("Starting over...")
            os.system('cls' if os.name == 'nt' else 'clear')
            main()
        elif choice == 6:
            print("Quitting...")
            sys.exit()
        else:
            print("Invalid choice")
            continue

        print(f"Result: {result}")

if __name__ == "__main__":
    main()
