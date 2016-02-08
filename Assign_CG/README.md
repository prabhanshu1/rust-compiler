###CodeGenerator for Rust language in GO

##Generate executable
```
cd asgn2/
make
```
##Run test cases
```
bin/codegen test/test1.ir	
```
functions:
getreg -> getreg returns 3 values, 1st register allocated, 2nd fresh value which tells if the register we got is fresh or old, or the passed variable is an int, finally the 3rd value gives the old variable which has occupied that register.

Load_and_store => we give the "fresh" old_variable, and the new register. which stores the old variable value into variable and then load the new variable in the register.

Getreg_force => we ask for specific register to this function, other parameters and return value is same as getreg.

Hold_reg => Register edx, eax are reserved when performing divide, this is where hold_reg comes in picture.

preprocess => fuction takes instruction pointer and return the table, associated with each instructions.

parser => intakes the instructions (3 addr code) and calls parse_line converts them for translating in the function translator in the translator.go file.

formats =
trailing comma should have a space after that
[]= is for a[i]=b with samples in test files
