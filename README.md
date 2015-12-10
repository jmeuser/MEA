# MEA
MEA calculates the Monthly Expense Allocation of a manager or department given the MEA for each employee position basic personnel files stored in json.

There are three employee positions:
* Developers
* QA Testers
* Managers

Developers and QA Testers report to managers (or form their own department), and  managers report to managers.

Expense is allocated as follows:
* Developers warrent $1000 each
* QA Testers warrent $500 each
* Managers warent $300 each

The monthly expense allocation warrented a manager and their employees is equal to the amount warrented the manager plus the amount warrented by their employees. Thus, a manager having an employee who is a manager of a developer and a QA tester is warrented $300 + $300 + $1000 + $500 yielding $2100.

A department is a forest of employee trees, and, as such, warrents a monthly expense allocation equal to the sum of the monthly expense allocated to each of its employee trees.
