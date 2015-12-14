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

----
Right now main.go generates the random Employee trees that will be used to write the cost_test.go test suite.
Clone down this repo and then

    $ go run main.go

and you should see the following output

    [0 0]
    [0 0 1 0 2]
    [0 0 0 1 3 1 4 5 0 8]
    [0 0 0 0]
    [0 0 1 2 2 0]
    
    [0 0 1 0 2 4]
    {0 0 Manager a}
    {1 0 Manager a}
    {2 1 Manager a}
    {3 0 Developer a}
    {4 2 Manager a}
    {5 4 Developer a}
    
    [0 0 1]
    {0 0 Manager a}
    {1 0 Manager a}
    {2 1 Developer a}
    
    [0]
    {0 0 Manager a}
    
    [0 0 0 2]
    {0 0 Manager a}
    {1 0 QA Tester a}
    {2 0 Manager a}
    {3 2 QA Tester a}
    
    [0 0 0 0 3 1 3]
    {0 0 Manager a}
    {1 0 Manager a}
    {2 0 Developer a}
    {3 0 Manager a}
    {4 3 Developer a}
    {5 1 Developer a}
    {6 3 Developer a}

----

# Next

* func ret in main.go should return

	[]*Employee
* Introduce func saveBPF (Basic Personnel Files) to main.go (saves marhsaled []*Employee to fileName.json).
* Transform main.go into cost_test.go
* Introduce how-to run cost_test.go to README.md
