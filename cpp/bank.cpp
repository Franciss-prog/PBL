#include <iostream>
#include <string>
#include <limits>
using namespace std;

// Function to withdraw
int withdraw(int balance) {
    // Step 1: Declare a variable to store the amount the user wants to withdraw.
    int withdrawAmount;

    while (true)
    {
        // Step 2: Prompt the user to enter the withdrawal amount.
        cout << "Enter your desire amount to withdraw: ₱";
        cin >> withdrawAmount;

        // Step 3: Validate the input:
        if (cin.fail() || withdrawAmount <= 0)
        {
        // - Ensure the input is a positive number.
            // clear
            cin.clear();
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            cout << "Invalid input. Please enter a positive integer." << endl;
            continue;
        }
        // Step 4: If the input is invalid or the balance is insufficient:
    
        if (withdrawAmount > balance)
        {
        // - Inform the user and ask them to try again.
              cout << "Insufficient balance. You cannot withdraw more than your current balance." << endl;
              continue;
        }
            
    // Step 5: If the input is valid:
    // - Subtract the withdrawal amount from the balance.
        balance -= withdrawAmount;
    // - Inform the user of the successful withdrawal.
        cout << "Deposit successful. Your new balance is: ₱" << balance << endl;
        break;

    }
     
    // Step 6: Return the updated balance.
    return balance; // Placeholder return
}

// Function to deposit
int deposit(int balance) {
    // Step 1: Declare a variable to store the amount the user wants to deposit.
    int depositAmount;
    
    // Step 2: Prompt the user to enter the deposit amount.
    while (true)
    {
         cout << "Enter your desire amount to deposit" << endl;
         cin  >> depositAmount;
     
    // Step 3: Validate the input:
        if(cin.fail() || depositAmount <= 0) {
        // - Ensure the input is a positive number.
            cin.clear();
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            cout << "Invalid input. Please enter a positive integer." << endl;
            continue;
        }
        // Step 4: If the input is valid:
        // - Add the deposit amount to the balance.
        balance += depositAmount;
        // - Inform the user of the successful deposit.
        cout << "Withdrawal successful. Your new balance is: ₱" << balance << endl;
        break;
    }
    
    // Step 5: Return the updated balance.
    return balance; // Placeholder return
}

// Prompt the user for an action
void prompt(string username, int& balance) {
    int perform = 0;

    cout << "Hello, " << username << "!" << endl;
    cout << "You have a balance of ₱" << balance << endl;

    while (true) {
        cout << "What do you want to do with your balance?\n";
        cout << "1] Withdraw\n";
        cout << "2] Deposit\n";
        cout << "Choose a number [1/2]: ";
        cin >> perform;

        //  
        if (cin.fail() || (perform != 1 && perform != 2)) {
            cin.clear();
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            cout << "Invalid choice. Please choose 1 or 2.\n";
            continue;
        }

       // If guard: Check if the user wants to withdraw
        if (perform == 1) {
            // If guard: Check if the balance is zero
            if (balance == 0) {
                cout << "You don't have enough balance to withdraw.\n";
                continue; // Skip the rest of the loop and prompt again
            }

            // If all checks pass, process the withdrawal
            balance = withdraw(balance);
        }
        // If guard: Check if the user wants to deposit
        if (perform == 2) {
            balance = deposit(balance);
        }
        cout << "Your updated balance is: " << balance << endl;

        // Ask if the user wants to continue
        char choice;
        cout << "Do you want to perform another transaction? (y/n): ";
        cin >> choice;

        if (choice == 'n' || choice == 'N') {
            cout << "Thank you for using the banking system. Goodbye!\n";
            break;
        }
    }
}

int main() {
    cout << "<----------------Welcome to My Simple Banking System-------------->\n";

    // Step 1: Declare variables to store the user's name and balance.
    string username;
    int balance = 0;

    // Step 2: Prompt the user to enter their name.
    cout << "Enter your name: ";
    cin >> username;

    // Step 3: Validate the name input:
    // - Ensure it is not empty.
    while (username.empty()) {
        cout << "Username cannot be empty. Please enter your name: ";
        cin >> username;

    }

    // Step 4: Call the prompt function to handle user transactions.
    prompt(username, balance);

    // Step 5: End the program with a goodbye message.
    cout << "Thank you for using the banking system. Goodbye!\n";
    return 0;
}
