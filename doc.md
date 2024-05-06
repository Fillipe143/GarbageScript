+ -> push(pop() + pop())
- -> push(pop() - pop())
/ -> push(pop() / pop())
* -> push(pop() * pop())

= -> push(pop() == pop())
> -> push(pop() > pop())
< -> push(pop() < pop())
! -> push(pop() == 0)
& -> push(pop() && pop())
| -> push(pop() || pop())
^ -> push(pop() ^  pop())

a -> push(rax)
b -> push(rbx)
c -> push(rcx)
d -> push(rdx)

l -> push(memory.pop())
s -> memory.push(pop())
m -> push(*memory)

xn -> push(n)
y -> duplicate(pop())
z -> swap2()
z3 -> swap3()
t -> push(pop(2))

$n -> syscall(pop() x n)
@ -> print(pop())
. -> pop()

{ block; -> if (pop()) {block}
{ block1] block2; -> if (pop()) {block1} else {block2}
[ cmp { block; -> while(cmp) {block}

i"file.gb" -> import "file.gb"
#name% ( block; -> func name() {block}
%name# -> name()
'name' -> *name

x"Hello world" %std_println#
x0[yx10>{yx2z/.{x" is odd"]x" is even";%std_println#x1+;
