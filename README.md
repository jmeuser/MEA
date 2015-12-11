# MEA
MEA calculates the Monthly Expense Allocation of a manager (MMEA) or department (DMEA) given the MEA for each employee position in the basic personnel files stored in .json.

There are three employee positions:
* Developers
* QA Testers
* Managers

Developers and QA Testers report to managers (or form their own department), and  managers report to managers.

Expense is allocated as follows:
* Developers warrant $1000 each
* QA Testers warrant $500 each
* Managers warrant $300 each

The monthly expense allocation warranted a manager and their employees is equal to the amount warranted the manager plus the amount warranted by their employees. Thus, a manager having an employee who is a manager of a developer and a QA tester is warranted $300 + $300 + $1000 + $500 yielding $2100.

A department is a forest of employee trees, and, as such, warrants a monthly expense allocation equal to the sum of the monthly expense allocated to each of its employee trees.
