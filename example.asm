section .data
str_0 db 0x65, 0x78, 0x61, 0x6D, 0x70, 0x6C, 0x65, 0x2E, 0x67, 0x62
str_len_0 equ $ - str_0
section .bss
tmp resq 1000
i resq 1
vars resq 10
section .text
global _start
print:
mov r9, -3689348814741910323
sub rsp, 40
mov BYTE [rsp+31], 10
lea rcx, [rsp+30]
.L2:
mov rax, rdi
lea r8, [rsp+32]
mul r9
mov rax, rdi
sub r8, rcx
shr rdx, 3
lea rsi, [rdx+rdx*4]
add rsi, rsi
sub rax, rsi
add eax, 48
mov BYTE [rcx], al
mov rax, rdi
mov rdi, rdx
mov rdx, rcx
sub rcx, 1
cmp rax, 9
ja  .L2
lea rax, [rsp+32]
mov edi, 1
sub rdx, rax
xor eax, eax
lea rsi, [rsp+32+rdx]
mov rdx, r8
mov rax, 1
syscall
add rsp, 40
ret
_start:
push str_len_0
push str_0
call _func_io_read_file
call _func_std_write
mov rax, 60
xor rdi, rdi
syscall
_func_std_exit:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 60
pop rax
pop rdi
syscall
push rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_std_write:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 1
push 1
pop rax
pop rdi
pop rsi
pop rdx
syscall
push rax
pop rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_std_read:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 1
push 0
pop rax
pop rdi
pop rsi
pop rdx
syscall
push rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_std_time:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 0
push 201
pop rax
pop rdi
syscall
push rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_std_delay:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
call _func_std_time
.addr_0:
pop rax
push rax
push rax
call _func_std_time
pop rax
pop rdi
sub rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
push rax
push rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rdi
push rax
push rdi
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_1
pop rax
pop rdi
push rax
push rdi
jmp .addr_0
.addr_1:
pop rax
pop rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_std_itoa:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
pop rax
push rax
push rax
mov rcx, 0
mov rdx, 1
pop rax
cmp rax, rcx
cmove rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_2
push 1
pop rsi
push rsi
mov rax, 9
xor rdi, rdi
mov rdx, 3
mov r10, 34
 mov r8, -1
xor r9, r9
syscall
mov QWORD [vars+0], rax
pop rax
mov QWORD [vars+8], rax
mov rax, [vars+0]
push rax
pop rax
pop rdi
push rax
push rdi
push 48
pop rax
pop rdi
add rax, rdi
push rax
pop rbx
pop rax
mov BYTE [rax], bl
mov rax, [vars+8]
push rax
mov rax, [vars+0]
push rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
.addr_2:
pop rax
push rax
push rax
pop rax
push rax
push rax
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_3
push 0
pop rax
pop rdi
sub rax, rdi
push rax
.addr_3:
push 0
pop rax
pop rdi
push rax
push rdi
.addr_4:
pop rax
push rax
push rax
pop rax
cmp rax, 0
je .addr_5
push 10
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
xor rdx, rdx
div rdi
push rdx
push rax
pop rax
pop rdi
push rax
push rdi
pop rax
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
jmp .addr_4
.addr_5:
pop rax
pop rax
pop rdi
push rax
push rdi
pop rax
push rax
push rax
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_6
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
.addr_6:
pop rax
pop rdi
push rax
push rdi
pop rax
push rax
push rax
pop rsi
push rsi
mov rax, 9
xor rdi, rdi
mov rdx, 3
mov r10, 34
 mov r8, -1
xor r9, r9
syscall
mov QWORD [vars+0], rax
pop rax
mov QWORD [vars+8], rax
push 1
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
sub rax, rdi
push rax
mov rax, [vars+0]
push rax
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
pop rax
push rax
push rax
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_7
mov rax, [vars+0]
push rax
push 45
pop rbx
pop rax
mov BYTE [rax], bl
push 0
pop rax
pop rdi
sub rax, rdi
push rax
.addr_7:
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
push rax
push rdi
.addr_8:
pop rax
push rax
push rax
pop rax
cmp rax, 0
je .addr_9
push 10
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
xor rdx, rdx
div rdi
push rdx
push rax
pop rax
pop rdi
push rax
push rdi
push 0x30
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
push rax
push rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rbx
pop rax
mov BYTE [rax], bl
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
sub rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
jmp .addr_8
.addr_9:
pop rax
pop rax
mov rax, [vars+8]
push rax
mov rax, [vars+0]
push rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_std_atoi:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
pop rax
push rax
push rax
pop rax
xor rbx, rbx
mov BYTE bl, [rax]
push rbx
push 0x2D
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmove rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_10
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
sub rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
push -1
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
jmp .addr_11
.addr_10:
pop rax
push rax
push rax
pop rax
xor rbx, rbx
mov BYTE bl, [rax]
push rbx
push 0x2B
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmove rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_12
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
sub rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
jmp .addr_13
.addr_12:
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
.addr_13:
.addr_11:
push 0
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
push rax
push rdi
.addr_14:
pop rax
push rax
push rax
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovl rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_15
push 1
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
sub rax, rdi
push rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
push 10
pop rax
pop rdi
mul rdi
push rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rdi
push rax
push rdi
pop rax
push rax
push rax
pop rax
xor rbx, rbx
mov BYTE bl, [rax]
push rbx
pop rax
push rax
push rax
push 0x39
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovl rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_16
pop rax
pop rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rax
push 10
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
xor rdx, rdx
div rdi
push rdx
push rax
pop rax
pop rdi
push rax
push rdi
pop rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
.addr_16:
pop rax
push rax
push rax
push 0x30
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
pop rax
cmp rax, 0
je .addr_17
pop rax
pop rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rax
push 10
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
xor rdx, rdx
div rdi
push rdx
push rax
pop rax
pop rdi
push rax
push rdi
pop rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
.addr_17:
push 0x30
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
sub rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
pop rax
pop rdi
push rax
push rdi
push 1
pop rax
pop rdi
add rax, rdi
push rax
pop rax
pop rdi
push rax
push rdi
jmp .addr_14
.addr_15:
pop rax
pop rax
pop rax
pop rdi
mul rdi
push rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_io_open:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 0
push 0
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
push 2
pop rax
pop rdi
pop rsi
pop rdx
syscall
push rax
pop rax
push rax
push rax
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_io_close:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 3
pop rax
pop rdi
syscall
push rax
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_io_read:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 0
pop rax
pop rdi
pop rsi
pop rdx
syscall
push rax
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmovg rcx, rdx
push rcx
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_io_stat:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
push 1
pop rsi
push rsi
mov rax, 9
xor rdi, rdi
mov rdx, 3
mov r10, 34
 mov r8, -1
xor r9, r9
syscall
mov QWORD [vars+16], rax
pop rax
mov QWORD [vars+24], rax
mov rax, [vars+16]
push rax
pop rax
push rax
push rax
pop rax
pop rdi
pop rbx
push rax
push rdi
push rbx
push 4
pop rax
pop rdi
pop rsi
syscall
push rax
pop rax
pop rax
push rax
push rax
pop rax
xor rbx, rbx
mov QWORD rbx, [rax]
push rbx
push 0
mov rcx, 0
mov rdx, 1
pop rax
pop rdi
cmp rax, rdi
cmove rcx, rdx
push rcx
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_io_size:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
call _func_io_stat
pop rax
pop rdi
push rax
push rdi
push 48
pop rax
pop rdi
add rax, rdi
push rax
pop rax
xor rbx, rbx
mov QWORD rbx, [rax]
push rbx
pop rax
pop rdi
push rax
push rdi
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
_func_io_read_file:
pop rbx
mov rcx, [i]
mov [tmp+rcx], rbx
add rcx, 8
mov [i], rcx
pop rax
pop rdi
push rax
push rdi
pop rsi
push rsi
mov rax, 9
xor rdi, rdi
mov rdx, 3
mov r10, 34
 mov r8, -1
xor r9, r9
syscall
mov QWORD [vars+32], rax
pop rax
mov QWORD [vars+40], rax
mov rax, [vars+32]
push rax
pop rax
pop rdi
push rax
push rdi
pop rbx
pop rax
mov QWORD [rax], rbx
mov rax, [vars+32]
xor rbx, rbx
mov QWORD rbx, [rax]
push rbx
call _func_io_open
pop rax
cmp rax, 0
je .addr_18
push 1
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
.addr_18:
push 1
pop rsi
push rsi
mov rax, 9
xor rdi, rdi
mov rdx, 3
mov r10, 34
 mov r8, -1
xor r9, r9
syscall
mov QWORD [vars+48], rax
pop rax
mov QWORD [vars+56], rax
mov rax, [vars+48]
push rax
pop rax
pop rdi
push rax
push rdi
pop rbx
pop rax
mov QWORD [rax], rbx
mov rax, [vars+32]
xor rbx, rbx
mov QWORD rbx, [rax]
push rbx
call _func_io_size
pop rax
cmp rax, 0
je .addr_19
push 1
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
.addr_19:
pop rsi
push rsi
mov rax, 9
xor rdi, rdi
mov rdx, 3
mov r10, 34
 mov r8, -1
xor r9, r9
syscall
mov QWORD [vars+64], rax
pop rax
mov QWORD [vars+72], rax
mov rax, [vars+72]
push rax
mov rax, [vars+64]
push rax
mov rax, [vars+48]
xor rbx, rbx
mov QWORD rbx, [rax]
push rbx
call _func_io_read
pop rax
cmp rax, 0
je .addr_20
push 1
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
.addr_20:
mov rax, [vars+48]
xor rbx, rbx
mov QWORD rbx, [rax]
push rbx
call _func_io_close
pop rax
cmp rax, 0
je .addr_21
push 1
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
.addr_21:
mov rax, [vars+72]
push rax
mov rax, [vars+64]
push rax
mov rcx, [i]
sub rcx, 8
mov [i], rcx
mov rbx, [tmp+rcx]
push rbx
ret
