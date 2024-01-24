Sample codes used in training/onboarding.

**To run each sample:**
```bash
cd <folder>/
go build -v

# For Linux and MacOS:
./{binary} [-flag(s)]

# For Windows:
.\{binary}.exe [-flags(s)]
```

**To run `concurrent` folder:**
1. Download the `testcur.csv` file from `Shared drives/Engineering|Main/Temporary/` and store somewhere.
2. Build `concurrent` folder.
```bash
cd concurrent/
go build -v

# For Linux and MacOS:
./concurrent -file $HOME/testcur.csv
# or
./concurrent -file $HOME/testcur.csv -concurrent=true

# For Windows:
.\concurrent.exe -file C:\somefolder\testcur.csv
# or
.\concurrent.exe -file C:\somefolder\testcur.csv -concurrent=true
```
