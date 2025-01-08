#include<iostream>
using namespace std;




int main(){
    // print
 cout << "_________        .__               .__          __                \n";
    cout << "\\_   ___ \\_____  |  |   ____  __ __|  | _____ _/  |_  ___________ \n";
    cout << "/    \\  \\/\\__  \\ |  | _/ ___\\|  |  \\  | \\__  \\\\   __\\/  _ \\_  __ \\\n";
    cout << "\\     \\____/ __ \\|  |_\\  \\___|  |  /  |__/ __ \\|  | (  <_> )  | \\/\n";
    cout << " \\______  (____  /____/\\___  >____/|____(____  /__|  \\____/|__|   \n";
    cout << "        \\/     \\/          \\/                \\/                    \n";

    // declare int variable
    int num1, num2, operation;
    
    // first number
    std::cout << "Enter your first Number: ";
    std::cin >>num1;

    // second number
    std::cout << "Enter your Second Number: ";
    std::cin >>num2;

    std::cout << "Enter your Number choices [1/2/3/4]: \n";
    std::cout << "1] Addition\n";
    std::cout << "2] Substraction\n";
    std::cout << "3] Multiplication\n";
    std::cout << "4] Division\n";
    std::cin >>operation;


    switch (operation)
    {
    case 1:
        std::cout<< "Answer: " << num1 + num2 << std::endl;
        break;
     case 2:
        std::cout<< "Answer: " << num1 - num2 << std::endl;
        break;
     case 3:
        std::cout<< "Answer: " << num1 * num2 << std::endl;
        break;
     case 4:
        if (num1 == 0)
        {
            std::cout << "first Number can't divide";
        }
        std::cout<< "Answer: " << num1 / num2 << std::endl;
        break;
    
    default:
        std::cout << "Please choose a operation\n";
        break;
    }

