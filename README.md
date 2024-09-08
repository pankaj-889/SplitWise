# 💸 SplitEase - Simplify Group Expenses!

Welcome to **SplitEase**! A smart and easy-to-use application for managing group expenses, simplifying payments, and keeping everything settled between friends, family, or colleagues. Whether it's a trip, a shared dinner, or monthly house bills, SplitEase has you covered!

---

## 🎯 Features
- **👥 User-Friendly**: Create users with ease and manage balances automatically.
- **👨‍👩‍👧 Group Management**: Create groups with multiple users to track shared expenses.
- **💰 Smart Expense Tracking**: Add expenses and instantly update everyone's balances based on how the expense is split.
- **📊 Accurate Balance Calculation**: Track what each user owes or is owed.
- **✅ Settlement Simplified**: Settle debts and record payments between users.
- **📜 Payment History**: Keep a clear history of all payments and settlements.

---

## 🏗️ Entities

SplitEase is built around the following key entities that manage all expense-sharing functionality:

1. **User**: 
   - Stores user information and tracks their balance.
   - Fields:
     - `id`: Unique identifier for the user.
     - `name`: Name of the user.
     - `email`: User's email address.
     - `password`: Secure password (hashed).
     - `balance`: Initial balance (set to 0).

2. **Group**: 
   - Represents a collection of users for shared expenses.
   - Fields:
     - `id`: Unique identifier for the group.
     - `name`: Group name (e.g., "Trip to Goa").
     - `created_by`: User ID of the group creator.
     - `created_at`: Timestamp when the group was created.
     - `users`: List of user IDs that belong to the group.

3. **Expense**: 
   - Tracks each expense added to a group.
   - Fields:
     - `expense_id`: Unique identifier for the expense.
     - `group_id`: The ID of the group associated with the expense.
     - `amount`: Total expense amount.
     - `paid_by`: List of user IDs who contributed to the payment.
     - `split_in_between`: List of user IDs among whom the expense is split.

4. **Payment**: 
   - Records payments made between users.
   - Fields:
     - `payment_id`: Unique identifier for the payment.
     - `group_id`: ID of the group where the payment occurred.
     - `payer_id`: User ID of the payer.
     - `payee_id`: User ID of the payee.
     - `amount`: Amount paid.
     - `date_of_payment`: Timestamp of when the payment was made.

5. **Balance**: 
   - Maintains the balance for each user in a group.
   - Fields:
     - `user_id`: User ID for whom the balance is tracked.
     - `group_id`: The group in which the balance is tracked.
     - `amount_owed`: Total amount the user owes others in the group.
     - `amount_due`: Total amount others owe to the user.

6. **Settlement**: 
   - Tracks the settlement of debts between users.
   - Fields:
     - `id`: Unique identifier for the settlement.
     - `from_user`: User ID of the user who is settling their debt.
     - `to_user`: User ID of the user who is being paid.
     - `amount`: Amount being settled.
     - `date`: Timestamp of the settlement.

---

## 🚀 Getting Started

Follow these steps to get the SplitEase app up and running on your machine.

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/splitease.git
